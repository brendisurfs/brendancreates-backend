package email

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
)

// SendEmail - sends an email to the desired reciever, carrying frontend form data.
func SendEmail(contactFromAddr, contactSubject, contactMessage string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	myDomain := os.Getenv("MG_DOMAIN")
	privateKey := os.Getenv("MG_PRIVATE_KEY")
	// pubKey := os.Getenv("MG_PUBLIC_KEY")
	myEmail := os.Getenv("MY_EMAIL")

	// mailgun setup
	mg := mailgun.NewMailgun(myDomain, privateKey)

	// mailgun parameters
	sender := contactFromAddr
	subject := contactSubject
	body := contactMessage
	recipient := myEmail

	// message to be sent
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println(message)

	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Fatal("error with sending email ", err)
	}
	log.Printf("ID: %s Resp: %s\n", id, resp)
}
