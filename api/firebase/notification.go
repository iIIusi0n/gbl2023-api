package firebase

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	fbConfig "gbl-api/api/firebase/config"
	"gbl-api/config"
	"log"
)

func SendNotification(title, body string) {
	app := fbConfig.GetFirebaseApp()

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: config.RegistriationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
