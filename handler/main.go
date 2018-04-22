package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dev-drprasad/stephanie-go/mentors"
)

// Response response object
type Response struct {
	Message string `json:"message"`
}

// Handler main excution point
func Handler() ([]mentors.Mentor, error) {
	result := mentors.ScrapeMentors()
	return result, nil
}

func main() {
	lambda.Start(Handler)
}
