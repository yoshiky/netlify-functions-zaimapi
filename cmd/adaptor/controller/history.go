package controller

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
)

type History struct {
	OutputFactory func() ports.HistoryOutputPort
	InputFactory  func(o ports.HistoryOutputPort, u ports.HistoryRepository) ports.HistoryInputPort
	RepoFactory   func() ports.HistoryRepository
}

func (h *History) GetHistoryByDate(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	// リクエストパラメタ取得
	sd := request.QueryStringParameters["start_date"]
	ed := request.QueryStringParameters["end_date"]

	outputPort := h.OutputFactory()
	repository := h.RepoFactory()
	inputPort := h.InputFactory(outputPort, repository)
	return inputPort.GetHistoryByDate(request, sd, ed)
}
