package emitters

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
)

type Scoreboard struct {
	FirebaseClient *messaging.Client
}

func (s Scoreboard) Run(RoomCode string) (err error) {
	ctx := context.Background()
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
	s.FirebaseClient.Send(ctx)
	return err
}
