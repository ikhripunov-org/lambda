package lambda

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type mockedSnsPublish struct {
	Client snsiface.SNSAPI
	Resp   sns.PublishOutput
}

func TestLambdaHandler(t *testing.T) {
	var input map[string]interface{}
	json.Unmarshal([]byte("{\"foo\":\"bar\"}"), &input)
	svc := SNS{
		Client: mockedSnsPublish{},
	}
	svc.PublishMessage(input)
}
