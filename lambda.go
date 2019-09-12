package lambda

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type SNS struct {
	Client snsiface.SNSAPI
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) (string, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := SNS{
		Client: sns.New(sess),
	}

	result, err := svc.PublishMessage(snsEvent.Records[0].SNS.Message)
	if err != nil {
		panic(err.Error())
	}

	return *result.MessageId, nil
}

func (s *SNS) PublishMessage(jsonMessage string) (*sns.PublishOutput, error) {
	var message map[string]interface{}
	err := json.Unmarshal([]byte(jsonMessage), &message)
	if err != nil {
		panic(err.Error())
	}

	message["platform"] = "farmroad"

	payload, err := json.Marshal(message)
	if err != nil {
		panic(err.Error())
	}
	return s.Client.Publish(&sns.PublishInput{
		Message:  aws.String(string(payload)),
		TopicArn: aws.String(os.Getenv("TOPIC_ARN")),
	})
}

func main() {
	lambda.Start(HandleRequest)
}
