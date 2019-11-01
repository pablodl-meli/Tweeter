package service

import (
	"testing"

	"github.com/pablodl-meli/Tweeter/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetLastTweet()

	assert.Equal(t, user, publishedTweet.User, "User should be 'grupoesfera'")
	assert.Equal(t, text, publishedTweet.Text, "The text should be 'this is my first tweet'")
	assert.NotNil(t, publishedTweet.Date, "Time should not be nil")

	// if publishedTweet.User != user &&
	// 	publishedTweet.Text != text {
	// 	t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
	// 		user, text, publishedTweet.User, publishedTweet.Text)
	// }
	// if publishedTweet.Date == nil {
	// 	t.Error("Expected date can't be nil")
	// }

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.Equal(t, "user is required", err.Error(), "error should say 'user is required")
	// if err != nil && err.Error() != "user is required" {
	// 	t.Error("Expected error is user is required")
	// }
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	tweetManager := NewTweetManager()
	var firstTweet, secondTweet *domain.Tweet // Fill the tweets with data
	firstTweet = domain.NewTweet("pablo", "tweet de pablo")
	secondTweet = domain.NewTweet("paulo", "tweet de paulo")

	// Operation
	tweetManager.PublishTweet(firstTweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	assert.Equal(t, 2, len(publishedTweets), "length should be 2")

	firstPublishedTweet := publishedTweets[0]

	secondPublishedTweet := publishedTweets[1]

	assert.Equal(t, firstPublishedTweet.User, firstTweet.User, "user should be pablo")
	assert.Equal(t, firstPublishedTweet.Text, firstTweet.Text, "text should be 'tweet de pablo'")
	assert.Equal(t, secondPublishedTweet.User, secondTweet.User, "users should be paulo")
	assert.Equal(t, secondPublishedTweet.Text, secondTweet.Text, "text should be 'tweet de paulo'")

	// if !isValidTweet(t, firstPublishedTweet, user, text) {
	//     return
	// }

	// Same for secondPublishedTweet
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()
	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	assert.Equal(t, publishedTweet.User, tweet.User, "user should be grupoesfera")
	assert.Equal(t, publishedTweet.User, tweet.User, "user should be grupoesfera")
	assert.NotNil(t, publishedTweet.Id, tweet.Id, "id should not be nil")

	// isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization

	tweetManager := NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	tweetManager.PublishTweet(tweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	assert.Equal(t, 2, count, "El usuario grupoesfera debe tener dos tweets")

	// // Validation
	// if count != 2 {
	// 	t.Errorf("Expected count is 2 but was %d", count)
	// }
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	// publish the 3 tweets
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	assert.Equal(t, 2, len(tweets), "Los tweets de grupo esfera son 2")
	assert.Equal(t, firstPublishedTweet.User, user, "El usuario debe ser grupoesfera")
	assert.Equal(t, firstPublishedTweet.Text, text, "El texto debe ser This is my first tweet")
	assert.Equal(t, secondPublishedTweet.User, user, "El usuario debe ser grupoesfera")
	assert.Equal(t, secondPublishedTweet.Text, secondText, "El texto debe ser This is my second tweet")

	// // Validation
	// if len(tweets) != 2 { /* handle error */
	// }
	// firstPublishedTweet := tweets[0]
	// secondPublishedTweet := tweets[1]
	// // check if isValidTweet for firstPublishedTweet and secondPublishedTweet
}
