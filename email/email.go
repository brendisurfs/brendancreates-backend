package email

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var myDomain string = "brendancreates.dev"
var privateKey string = "my-private-key"
var myEmail string = "brendan.prednis@pm.me"

// SendEmail - sends an email to the desired reciever, carrying frontend form data.
func SendEmail(contactFromAddr, contactSubject, contactMessage string) {

	mg := mailgun.NewMailgun(myDomain, privateKey)

	sender := contactFromAddr
	recipient := myEmail
	subject := contactSubject
	body := contactMessage

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
