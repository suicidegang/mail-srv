package main

import (
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/keighl/postmark"
	"github.com/micro/go-micro"
	"github.com/suicidegang/mail-srv/handler"
	proto "github.com/suicidegang/mail-srv/proto/mail"
)

func main() {
	service := micro.NewService(
		micro.Name("sg.micro.srv.mail"),
		micro.Version("0.1"),
		micro.BeforeStart(func() error {
			log.Printf("[sg.micro.srv.mail] Starting service...")

			// We'll need consul internally for some good stuff
			consul, err := api.NewClient(api.DefaultConfig())
			if err != nil {
				return err
			}

			handler.Config = consul.KV()
			serverToken, _, err := handler.Config.Get("mail/server-token", nil)
			if err != nil {
				return err
			}

			handler.Client = postmark.NewClient(string(serverToken.Value), "")
			return nil
		}),
	)

	service.Init()
	proto.RegisterMailingHandler(service.Server(), new(handler.Mailing))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
