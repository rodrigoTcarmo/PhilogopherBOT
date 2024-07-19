package hashtag

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/user/userlookup"
	"github.com/michimani/gotwi/user/userlookup/types"
)

func Hashtag(client *gotwi.Client) error {
	fmt.Println("hashtag function has been successfuly called!")

	input := &types.GetMeInput{
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetMe(context.Background(), client, input)
	if err != nil {
		return err
	}
	fmt.Println("aaaaaaa")
	fmt.Println("ID: ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name: ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username: ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt: ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
	return nil
}
