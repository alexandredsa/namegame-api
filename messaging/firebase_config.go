package messaging

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type FirebaseConfig struct {
}

func (fc FirebaseConfig) CreateClient() (*messaging.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	return app.Messaging(ctx)

	// registrationToken := "dkQJJ-3BLbg:APA91bFtqfdcku1WMPX2CkuCtJO9EomUucrh-aFs3X3mMJj636MPR7jbkRY-i0OtzqE1fDLs7bKrPbgN7wCVKoJkWECHjMz6uvyhFP8BXRTWadxVUicJdvg7zt2W1buJ0A7wk37ILho6"

	// message := &messaging.Message{
	// 	Data: map[string]string{
	// 		"score": "850",
	// 		"time":  "2:45",
	// 	},
	// 	Token: registrationToken,
	// }

	// response, err := client.Send(ctx, message)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("Successfully sent message:", response)

}
