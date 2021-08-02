package twitter

import (
	"context"
	"fmt"
	"github.com/desmos-labs/desmostipbot/tipper"
	"github.com/desmos-labs/desmostipbot/types"
	"github.com/desmos-labs/desmostipbot/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rs/zerolog/log"
)

// Client represents the Twitter API client
type Client struct {
	twitter *twitter.Client
	demux   twitter.Demux
	stream  *twitter.Stream

	tipper *tipper.Tipper
}

// NewClient allows to build a new Client instance
func NewClient(cfg *types.TwitterConfig, tipper *tipper.Tipper) *Client {
	// Build the Oauth 1.0 client that is needed for the streaming APIs
	config := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.AccessToken, cfg.AccessSecret)

	// Create the client
	client := &Client{
		tipper: tipper,
	}

	// Setup the Twitter client
	httpClient := config.Client(context.Background(), token)
	twitterClient := twitter.NewClient(httpClient)
	client.twitter = twitterClient

	// Setup the Demux
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		client.handleMention(tweet)
	}
	client.demux = demux

	return client
}

// StartListening starts the listening for new tweet mentions
func (client *Client) StartListening() error {
	log.Debug().Msg("Starting Twitter stream listening")

	// Start the stream
	stream, err := client.twitter.Streams.Filter(&twitter.StreamFilterParams{
		Track: []string{"@DesmosTipBot"},
	})
	if err != nil {
		return err
	}
	client.stream = stream

	// Receive messages until stopped or stream quits
	go client.demux.HandleChan(client.stream.Messages)

	return nil
}

// handleMention handles the given tweet mention
func (client *Client) handleMention(tweet *twitter.Tweet) {
	if !types.DesmosTipRegEx.MatchString(tweet.Text) {
		return
	}

	amount, user, err := utils.ParseText(tweet.Text)
	if err != nil {
		client.answerWithError(err, tweet)
		return
	}

	// Send the tip
	txHash, err := client.tipper.Tip(amount, user)
	if err != nil {
		client.answerWithError(err, tweet)
	}

	client.answerTweet(types.TipSentMessage(txHash), tweet)
}

// answerWithError answers the tweet having the given tweet id with the specified error
func (client *Client) answerWithError(err error, tweet *twitter.Tweet) {
	client.answerTweet(utils.Capitalize(err.Error()), tweet)
}

// answerTweet answers the tweet having the given id with the provided text
func (client *Client) answerTweet(text string, tweet *twitter.Tweet) {
	// Prepend to the text the username to which answer
	text = fmt.Sprintf("@%[1]s %[2]s", tweet.User.ScreenName, text)

	// Post the answer
	_, _, err := client.twitter.Statuses.Update(text, &twitter.StatusUpdateParams{
		InReplyToStatusID: tweet.ID,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error while sending tweet answer")
	}
}

// Stop stops the client from listening to new tweets
func (client *Client) Stop() {
	if client.stream != nil {
		client.stream.Stop()
	}
}
