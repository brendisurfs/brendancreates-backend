package email

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// SendEmail - sends an email to the desired reciever, carrying frontend form data.
func SendEmail(contactFromAddr, contactSubject, contactMessage string) (string, error) {

	myDomain := os.Getenv("MG_DOMAIN")
	privateKey := os.Getenv("MG_PRIVATE_KEY")
	myEmail := os.Getenv("MY_EMAIL")

	// mailgun parameters
	sender := contactFromAddr
	subject := contactSubject
	body := contactMessage
	recipient := myEmail

	// mailgun setup
	mg := mailgun.NewMailgun(myDomain, privateKey)

	// message to be sent
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)
	fmt.Printf("response: %s, id: %v", resp, id)
	return id, err
}
