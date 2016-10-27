package handler

import (
	errorsN "errors"
	"fmt"
	"log"
	"strconv"

	consul "github.com/hashicorp/consul/api"
	"github.com/keighl/postmark"
	"github.com/micro/go-micro/errors"
	proto "github.com/suicidegang/mail-srv/proto/mail"
	"golang.org/x/net/context"
)

var (
	Client *postmark.Client
	Config *consul.KV
)

type Mailing struct{}

func (m *Mailing) SendTemplate(ctx context.Context, req *proto.SendTemplateRequest, res *proto.SendTemplateResponse) error {

	from := recipient(req.Message.FromName, req.Message.From)
	to := recipient(req.Message.ToName, req.Message.To)
	vars := map[string]interface{}{}

	for key, value := range req.Variables {
		vars[key] = value
	}

	// Use consul to retrieve the actual template id
	templateId, err := getTemplateId(req.Template)
	if err != nil {
		return errors.InternalServerError("sg.micro.srv.mail.SendTemplate", err.Error())
	}

	email := postmark.TemplatedEmail{
		TemplateId:    templateId,
		TemplateModel: vars,
		From:          from,
		To:            to,
		TrackOpens:    true,
	}
	log.Printf("[postmark] Sending %+v", email)

	message, err := Client.SendTemplatedEmail(email)
	if err != nil {
		return errors.InternalServerError("sg.micro.srv.mail.SendTemplate", err.Error())
	}

	res.MessageID = message.MessageID
	return nil
}

func recipient(name, email string) string {
	return fmt.Sprintf("\"%s\" <%s>", name, email)
}

func getTemplateId(name string) (int64, error) {

	// Use consul to retrieve list of available templates at runtime
	template, _, err := Config.Get("mail/templates/"+name, nil)
	if err != nil {
		return 0, err
	}

	if template == nil {
		return 0, errorsN.New("Template was not found under mail/templates")
	}

	templateID, err := strconv.Atoi(string(template.Value))
	if err != nil {
		return 0, err
	}

	return int64(templateID), nil
}
