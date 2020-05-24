package templates

import (
	"bytes"
	"html/template"

	"github.com/influenzanet/messaging-service/pkg/types"
)

func GetTemplateTranslation(tDef types.EmailTemplate, lang string) types.LocalizedTemplate {
	var defaultTranslation types.LocalizedTemplate
	for _, tr := range tDef.Translations {
		if tr.Lang == lang {
			return tr
		} else if tr.Lang == tDef.DefaultLanguage {
			defaultTranslation = tr
		}
	}
	return defaultTranslation
}

func ResolveTemplate(tempName string, templateDef string, contentInfos map[string]string) (content string, err error) {
	tmpl, err := template.New("").Parse(templateDef)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, contentInfos)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
