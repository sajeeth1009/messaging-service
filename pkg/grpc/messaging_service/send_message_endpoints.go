package messaging_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *messagingServer) SendInstantEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	// TODO: check inputs
	// TODO: find and execute template
	_, err := s.clients.EmailClientService.SendEmail(ctx, &emailAPI.SendEmailReq{
		To: req.To,
		Content: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <title>A Simple Responsive HTML Email</title>
        <style type="text/css">
        body {margin: 0; padding: 0; min-width: 100%!important;}
        .content {width: 100%; max-width: 600px;}
        </style>
    </head>
    <body yahoo bgcolor="#f6f8f1">
        <table width="100%" bgcolor="#f1f923" border="0" cellpadding="0" cellspacing="0">
            <tr>
                <td>
                    <table class="content" align="center" cellpadding="0" cellspacing="0" border="0">
                        <tr>
                            <td>
                                Hello! Your registration was successful!
                            </td>
						</tr>
						<tr>
                            <td>
                                <a href="https://influenzanet.web.app">Open webpage</a>
                            </td>
						</tr>
                    </table>
                </td>
            </tr>
        </table>
    </body>
</html>`,
		FromAddress: "hevesi@flow-one.de",
		FromName:    "InfluenzatNet",
		Subject:     req.MessageType,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ServiceStatus{
		Version: apiVersion,
		Msg:     "message sent",
		Status:  api.ServiceStatus_NORMAL,
	}, nil
}
