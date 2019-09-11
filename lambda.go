package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type SNS struct {
	Client snsiface.SNSAPI
}

func HandleRequest(ctx context.Context, input map[string]interface{}) (string, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := SNS{
		Client: sns.New(sess),
	}
	result, err := svc.PublishMessage(input)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(*result.MessageId)

	return fmt.Sprintf("Hello %s!", ""), nil
}

func (s *SNS) PublishMessage(input map[string]interface{}) (*sns.PublishOutput, error) {
	input["platform"] = "farmroad"
	message, err := json.Marshal(input)
	if err != nil {
		panic(err.Error())
	}

	return s.Client.Publish(&sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(os.Getenv("TOPIC_ARN")),
	})
}

func main() {
	lambda.Start(HandleRequest)
}
