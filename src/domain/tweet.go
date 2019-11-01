package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

var Tweets []Tweet

func NewTweet(user, text string) *Tweet {
	t := time.Now()

	tweet := Tweet{user, text, &t, 0}

	return &tweet
}

func (this *Tweet) PrintableTweet() string {
	return "@" + this.User + ":" + this.Text
}
