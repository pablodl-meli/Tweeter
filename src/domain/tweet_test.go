package domain_test

import (
	"testing"

	"github.com/pablodl-meli/Tweeter/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTweet("grupoesfera", " This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, expectedText, text, "The expected text is "+expectedText)
	// if text != expectedText {
	// 	t.Errorf("The expected text is %s but was %s", expectedText, text)
	// }

}
