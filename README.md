[![Go Report Card](https://goreportcard.com/badge/github.com/suicidegang/mail-srv)](https://goreportcard.com/report/github.com/suicidegang/mail-srv)
[![Build Status](https://travis-ci.org/suicidegang/mail-srv.svg?branch=master)](https://travis-ci.org/suicidegang/mail-srv)

# Mail-srv

Mail-srv is a microservice used to send & track emails using Postmark API. Built for Golang micro.mu

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```
4. Setup config values in Consul KV

	```mail/server-token``` Your postmark server token.
	
	```mail/templates/{template-alias}``` Template aliases "KV" folder
	
4. Download and start the service
	```shell
	go get github.com/suicidegang/mail-srv
	mail-srv
	```

## The API
User server implements the following RPC Methods

Mailing
- SendTemplate

### Mailing.SendTemplate
```shell
$ micro query sg.micro.srv.mail Mailing.SendTemplate '{"template": "user.signup","variables": {"base_url": "https://spartangeek.com"},"message": {"from": "pedidos@spartangeek.com", "fromName": "Lic. Alberto", "to": "fernandez14@outlook.com", "toName": "Nobody"}}'
{
	"messageID": "bdff86f7-75d4-46b7-88f3-93b53a5b8448"
}
```
