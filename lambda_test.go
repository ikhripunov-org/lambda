package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type mockedSnsPublish struct {
	snsiface.SNSAPI
	In   *sns.PublishInput
	Resp sns.PublishOutput
}

func (m mockedSnsPublish) Publish(in *sns.PublishInput) (*sns.PublishOutput, error) {
	*m.In = *in
	return &m.Resp, nil
}

func TestLambdaHandler(t *testing.T) {
	var argument = sns.PublishInput{}
	client := mockedSnsPublish{Resp: sns.PublishOutput{}, In: &argument}
	svc := SNS{
		Client: client,
	}
	svc.PublishMessage("{\"foo\":\"bar\"}")
	if *argument.Message != "{\"foo\":\"bar\",\"platform\":\"farmroad\"}" {
		t.Error("Wrong message")
	}
}
