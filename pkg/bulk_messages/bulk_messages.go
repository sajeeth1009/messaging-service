package bulk_messages

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"time"

	"github.com/influenzanet/go-utils/pkg/api_types"
	"github.com/influenzanet/go-utils/pkg/constants"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/dbs/messagedb"
	"github.com/influenzanet/messaging-service/pkg/templates"
	"github.com/influenzanet/messaging-service/pkg/types"
	studyAPI "github.com/influenzanet/study-service/pkg/api"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

const loginTokenLifeTime = 7 * 24 * 60 * 60 // 7 days

func AsyncSendToAllUsers(
	apiClients *types.APIClients,
	messageDBService *messagedb.MessageDBService,
	instanceID string,
	messageTemplate types.EmailTemplate,
) {
	currentWeekday := time.Now().Weekday()
	var filters *umAPI.StreamUsersMsg_Filters

	if messageTemplate.MessageType == constants.EMAIL_TYPE_NEWSLETTER ||
		messageTemplate.MessageType == constants.EMAIL_TYPE_WEEKLY {
		filters = &umAPI.StreamUsersMsg_Filters{
			OnlyConfirmedAccounts:    true,
			UseReminderWeekdayFilter: true,
			ReminderWeekday:          int32(currentWeekday),
		}
	}

	stream, err := apiClients.UserManagementService.StreamUsers(context.Background(),
		&umAPI.StreamUsersMsg{
			InstanceId: instanceID,
			Filters:    filters,
		},
	)
	if err != nil {
		log.Printf("AsyncSendToAllUsers: %v", err)
		return
	}
	startTime := time.Now().Unix()
	counter := 0
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.AsyncSendToAllUsers(_) = _, %v", apiClients.UserManagementService, err)
			break
		}

		outgoingEmail := types.OutgoingEmail{
			MessageType:     messageTemplate.MessageType,
			HeaderOverrides: messageTemplate.HeaderOverrides,
		}
		contentInfos := map[string]string{}

		if user.Account.Type == "email" {
			outgoingEmail.To = []string{user.Account.AccountId}
		}

		if user.Account.AccountConfirmedAt < 1 {
			log.Println("message is not sent, if account not confirmed")
			continue
		}
		if messageTemplate.MessageType == constants.EMAIL_TYPE_NEWSLETTER {
			if !user.ContactPreferences.SubscribedToNewsletter || user.ContactPreferences.ReceiveWeeklyMessageDayOfWeek != int32(currentWeekday) {
				// user does not want to get newsletter
				continue
			}
			outgoingEmail.To = getEmailsByIds(user.ContactInfos, user.ContactPreferences.SendNewsletterTo)

			token, err := getUnsubscribeToken(apiClients.UserManagementService, instanceID, user)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["unsubscribeToken"] = token
		} else if messageTemplate.MessageType == constants.EMAIL_TYPE_WEEKLY {
			if !user.ContactPreferences.SubscribedToWeekly || user.ContactPreferences.ReceiveWeeklyMessageDayOfWeek != int32(currentWeekday) {
				// user does not want to get weekly reminder
				continue
			}
			token, err := getTemploginToken(apiClients.UserManagementService, instanceID, user, messageTemplate.StudyKey, loginTokenLifeTime)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["loginToken"] = token
		}

		subject, content, err := prepareContent(messageTemplate, user.Account.PreferredLanguage, contentInfos)
		if err != nil {
			log.Printf("AsyncSendToAllUsers: %v", err)
			continue
		}

		outgoingEmail.Subject = subject
		outgoingEmail.Content = content

		sendMail(
			apiClients.EmailClientService,
			instanceID,
			messageDBService,
			outgoingEmail,
		)
		counter += 1
	}
	log.Printf("Finished processing %d '%s' messages in %d s.", counter, messageTemplate.MessageType, time.Now().Unix()-startTime)
}

func AsyncSendToStudyParticipants(
	apiClients *types.APIClients,
	messageDBService *messagedb.MessageDBService,
	instanceID string,
	messageTemplate types.EmailTemplate,
	condition *api.ExpressionArg,
) {
	stream, err := apiClients.UserManagementService.StreamUsers(context.Background(), &umAPI.StreamUsersMsg{InstanceId: instanceID})
	if err != nil {
		log.Printf("AsyncSendToStudyParticipants: %v", err)
		return
	}
	currentWeekday := time.Now().Weekday()

	startTime := time.Now().Unix()
	counter := 0
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.AsyncSendToStudyParticipants(_) = _, %v", apiClients.UserManagementService, err)
			break
		}

		if user.Account.AccountConfirmedAt < 1 {
			log.Println("message is not sent, if account not confirmed")
			continue
		}

		profileIDs := make([]string, len(user.Profiles))
		for i, p := range user.Profiles {
			profileIDs[i] = p.Id
		}

		// check if user is in the study with at least one profile
		_, err = apiClients.StudyService.HasParticipantStateWithCondition(context.Background(), &studyAPI.ProfilesWithConditionReq{
			InstanceId: instanceID,
			ProfileIds: profileIDs,
			StudyKey:   messageTemplate.StudyKey,
			Condition:  expressionArgFromMessageToStudyAPI(condition),
		})
		if err != nil {
			log.Println(err)
			continue
		}

		// user profile in study with valid condition
		outgoingEmail := types.OutgoingEmail{
			MessageType:     messageTemplate.MessageType,
			HeaderOverrides: messageTemplate.HeaderOverrides,
		}
		contentInfos := map[string]string{}

		if user.Account.Type == "email" {
			outgoingEmail.To = []string{user.Account.AccountId}
		} else {
			log.Println("AsyncSendToStudyParticipants: account type not supported yet.")
			continue
		}

		if messageTemplate.MessageType == constants.EMAIL_TYPE_NEWSLETTER {
			if !user.ContactPreferences.SubscribedToNewsletter {
				// user does not want to get newsletter
				continue
			}

			outgoingEmail.To = getEmailsByIds(user.ContactInfos, user.ContactPreferences.SendNewsletterTo)

			token, err := getUnsubscribeToken(apiClients.UserManagementService, instanceID, user)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["unsubscribeToken"] = token
		} else if messageTemplate.MessageType == constants.EMAIL_TYPE_WEEKLY {
			if !user.ContactPreferences.SubscribedToWeekly || user.ContactPreferences.ReceiveWeeklyMessageDayOfWeek != int32(currentWeekday) {
				// user does not want to get weekly reminder
				continue
			}
			token, err := getTemploginToken(apiClients.UserManagementService, instanceID, user, messageTemplate.StudyKey, 604800)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["loginToken"] = token
			contentInfos["studyKey"] = messageTemplate.StudyKey
		}

		subject, content, err := prepareContent(messageTemplate, user.Account.PreferredLanguage, contentInfos)
		if err != nil {
			log.Printf("AsyncSendToAllUsers: %v", err)
			continue
		}

		outgoingEmail.Subject = subject
		outgoingEmail.Content = content

		sendMail(
			apiClients.EmailClientService,
			instanceID,
			messageDBService,
			outgoingEmail,
		)
		counter += 1
	}
	log.Printf("Finished processing %d '%s' messages in %d s.", counter, messageTemplate.MessageType, time.Now().Unix()-startTime)
}

func expressionArgFromMessageToStudyAPI(arg *api.ExpressionArg) *studyAPI.ExpressionArg {
	if arg == nil {
		return nil
	}
	newArg := &studyAPI.ExpressionArg{
		Dtype: arg.Dtype,
	}
	switch x := arg.Data.(type) {
	case *api.ExpressionArg_Exp:
		newArg.Data = &studyAPI.ExpressionArg_Exp{Exp: expressionFromMessageToStudyAPI(arg.GetExp())}
	case *api.ExpressionArg_Str:
		newArg.Data = &studyAPI.ExpressionArg_Str{Str: arg.GetStr()}
	case *api.ExpressionArg_Num:
		newArg.Data = &studyAPI.ExpressionArg_Num{Num: arg.GetNum()}
	case nil:
		// The field is not set.
	default:
		log.Printf("api.ExpressionArg has unexpected type %T", x)
	}

	return newArg
}

func expressionFromMessageToStudyAPI(arg *api.Expression) *studyAPI.Expression {
	if arg == nil {
		return nil
	}
	newArg := &studyAPI.Expression{
		Name:       arg.Name,
		ReturnType: arg.ReturnType,
	}
	data := make([]*studyAPI.ExpressionArg, len(arg.Data))
	for i, d := range arg.Data {
		data[i] = expressionArgFromMessageToStudyAPI(d)
	}
	newArg.Data = data
	return newArg
}

func getEmailsByIds(contacts []*umAPI.ContactInfo, ids []string) []string {
	emails := []string{}
	for _, c := range contacts {
		if c.Type == "email" {
			for _, id := range ids {
				if c.Id == id {
					emails = append(emails, c.GetEmail())
				}
			}
		}
	}
	return emails
}

func prepareContent(temp types.EmailTemplate, prefLang string, contentInfos map[string]string) (subject string, content string, err error) {
	translation := templates.GetTemplateTranslation(temp, prefLang)
	subject = translation.Subject
	decodedTemplate, err := base64.StdEncoding.DecodeString(translation.TemplateDef)
	if err != nil {
		return "", "", err
	}

	// execute template
	content, err = templates.ResolveTemplate(
		temp.MessageType+prefLang,
		string(decodedTemplate),
		contentInfos,
	)
	return
}

func sendMail(
	emailClient emailAPI.EmailClientServiceApiClient,
	instanceID string,
	messageDBService *messagedb.MessageDBService,
	mail types.OutgoingEmail,
) {
	_, err := emailClient.SendEmail(context.Background(), &emailAPI.SendEmailReq{
		To:              mail.To,
		HeaderOverrides: mail.HeaderOverrides.ToEmailClientAPI(),
		Subject:         mail.Subject,
		Content:         mail.Content,
	})
	if err != nil {
		log.Printf("Error when sending: %v", err)
		_, errS := messageDBService.AddToOutgoingEmails(instanceID, mail)
		if errS != nil {
			log.Printf("Error while saving to outgoing: %v", errS)
		}
		return
	}

	_, err = messageDBService.AddToSentEmails(instanceID, mail)
	if err != nil {
		log.Printf("Error when saving to sent: %v", err)
	}
}

func getTemploginToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
	studyKey string,
	expiresIn int64,
) (token string, err error) {
	resp, err := userClient.GenerateTempToken(context.Background(), &api_types.TempTokenInfo{
		UserId:     user.Id,
		Purpose:    constants.TOKEN_PURPOSE_SURVEY_LOGIN,
		InstanceId: instanceID,
		Expiration: time.Now().Unix() + expiresIn,
		Info: map[string]string{
			"studyKey": studyKey,
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func getUnsubscribeToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
) (token string, err error) {
	resp, err := userClient.GetOrCreateTemptoken(context.Background(), &api_types.TempTokenInfo{
		UserId:     user.Id,
		Purpose:    constants.TOKEN_PURPOSE_UNSUBSCRIBE_NEWSLETTER,
		InstanceId: instanceID,
		Expiration: time.Now().Unix() + 157680000,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
