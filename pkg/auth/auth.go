package auth

import (
	"os"

	"github.com/michimani/gotwi"
)

func Auth() (*gotwi.Client, error) {

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("PHILOGOPHER_BOT_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("PHILOGOPHER_BOT_ACCESS_TOKEN_SECRET"),
	}

	client, err := gotwi.NewClient(in)
	if err != nil {
		return nil, err
	}

	return client, nil
}
