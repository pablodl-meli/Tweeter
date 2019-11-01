package service

import (
	"fmt"

	"github.com/pablodl-meli/Tweeter/src/domain"
)

var Tweets []domain.Tweet

var LastId int

var TweetsByUser map[string][]domain.Tweet

var contadorDeTweets int

func InitializeService() {
	Tweets = make([]domain.Tweet, 0)
	contadorDeTweets = 0
	TweetsByUser = make(map[string][]domain.Tweet)
}

func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return tweet.Id, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return tweet.Id, fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return tweet.Id, fmt.Errorf("text exceeds 140 characters")
	}
	tweet.Id = LastId + 1
	Tweets = append(Tweets, *tweet)
	LastId++
	TweetsByUser[tweet.User] = append(TweetsByUser[tweet.User], *tweet)
	return tweet.Id, nil
}

func GetTweetById(id int) *domain.Tweet {
	for _, tweet := range Tweets {
		if tweet.Id == id {
			return &tweet
		}
	}
	return nil
}

func CountTweetsByUser(user string) int {
	return len(TweetsByUser[user])
	// for _, tweet := range Tweets {
	// 	if tweet.User == user {
	// 		contadorDeTweets++
	// 	}
	// }
	// return contadorDeTweets
}

func GetTweetsByUser(user string) []domain.Tweet {
	var aux []domain.Tweet
	for _, tweet := range Tweets {
		if tweet.User == user {
			aux = append(aux, tweet)
		}
	}
	return aux
}

func GetLastTweet() *domain.Tweet {
	return &Tweets[len(Tweets)-1]
}

func GetTweets() []domain.Tweet {
	return Tweets
}
