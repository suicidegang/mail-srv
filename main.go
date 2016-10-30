package main

import (
	"bufio"
	"log"
	"os"

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
			ignoredList := []string{}
			ignored, err := os.Open("./emails.txt")

			if err != nil {
				panic(err)
			}

			defer ignored.Close()
			scanner := bufio.NewScanner(ignored)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				ignoredList = append(ignoredList, scanner.Text())
			}

			handler.Ignored = ignoredList
			log.Printf("[sg.micro.srv.mail] %v ignored domains", len(handler.Ignored))
			return nil
		}),
	)

	service.Init()
	proto.RegisterMailingHandler(service.Server(), new(handler.Mailing))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
