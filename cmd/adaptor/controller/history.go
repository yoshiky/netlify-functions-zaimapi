package controller

import (
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
	"github.com/aws/aws-lambda-go/events"
)

type History struct {
	OutputFactory func() ports.HistoryOutputPort
	InputFactory  func(o ports.HistoryOutputPort, u ports.HistoryRepository) ports.HistoryInputPort
	RepoFactory   func() ports.HistoryRepository
}

func (h *History) GetHistoryByDate(request events.APIGatewayProxyRequest) string {
	// リクエストパラメタ取得
	// mapping := request.QueryStringParameters["mapping"]
	sd := request.QueryStringParameters["start_date"]
	ed := request.QueryStringParameters["end_date"]

	outputPort := h.OutputFactory()
	repository := h.RepoFactory()
	inputPort := h.InputFactory(outputPort, repository)
	return inputPort.GetHistoryByDate(request, sd, ed)
}
