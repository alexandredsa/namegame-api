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
}
