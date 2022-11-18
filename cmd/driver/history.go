package driver

import (
	"github.com/yoshiky/zaimapi/cmd/adaptor/controller"
	"github.com/yoshiky/zaimapi/cmd/adaptor/gateway"
	"github.com/yoshiky/zaimapi/cmd/adaptor/presenter"
	"github.com/yoshiky/zaimapi/cmd/usecase/interactor"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	history := controller.History{
		OutputFactory: presenter.NewHistoryOutputPort,
		InputFactory:  interactor.NewHistoryInputPort,
		RepoFactory:   gateway.NewHistoryRepository,
	}

	return history.GetHistoryByDate(r), nil
}
