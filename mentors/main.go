package mentors

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
var consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
var accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
var accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

func getMentorWebsite(twitterURL string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	resp, _ := client.Head(twitterURL)
	website := resp.Header.Get("Location")
	return website
}

func getMentorsFromTwitterResult(result anaconda.SearchResponse, sourceTweetID string) []Mentor {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})

	mentors := make([]Mentor, 0)
	if err == nil {
		svc := dynamodb.New(sess)

		for _, tweet := range result.Statuses {
			if tweet.QuotedStatus != nil {

				mentor := Mentor{
					UserID:        tweet.User.IdStr,
					FullName:      tweet.User.Name,
					UserName:      tweet.User.ScreenName,
					Bio:           tweet.User.Description,
					Tweet:         tweet.FullText,
					TweetID:       tweet.IdStr,
					ProfileImage:  strings.Replace(tweet.User.ProfileImageUrlHttps, "_normal", "", 1),
					SourceTweetID: sourceTweetID,
				}
				if len(tweet.User.URL) > 0 {
					mentorWebsite := getMentorWebsite(tweet.User.URL)
					if len(mentorWebsite) > 0 {
						mentor.Website = mentorWebsite
					}
				}
				mentors = append(mentors, mentor)
				if !dryRun {
					av, _ := dynamodbattribute.MarshalMap(mentor)

					input := &dynamodb.PutItemInput{
						Item:      av,
						TableName: aws.String("mentors"),
					}
					_, err = svc.PutItem(input)
					if err != nil {
						fmt.Println("Got error calling PutItem:")
						fmt.Println(err.Error())
						os.Exit(1)
					}
				}
				fmt.Println(mentor.Tweet)
			}
		}
	}
	return mentors
}

var dryRun = false

// ScrapeMentors hit twitter API and get twitters mentors and store them in dynamoDB
func ScrapeMentors() []Mentor {
	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)

	qp := url.Values{}
	qp.Set("tweet_mode", "extended")
	qp.Set("count", "100")
	// qp.Set("since_id", "987788865212768256")

	result := make([]Mentor, 0)

	sourceTweetIds := []string{"993733247161925632", "889004724669661184"}

	for _, tweetID := range sourceTweetIds {
		searchResult, _ := api.GetSearch("https://twitter.com/sehurlburt/status/"+tweetID, qp)
		fmt.Println(searchResult.Metadata.MaxId)
		mentors := getMentorsFromTwitterResult(searchResult, tweetID)
		result = append(result, mentors...)
	}

	return result

}

// https://twitter.com/nnja/status/984197138371391489

// https://twitter.com/sehurlburt/status/993733247161925632
