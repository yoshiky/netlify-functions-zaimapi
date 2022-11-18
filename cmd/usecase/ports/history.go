package ports

import (
	"github.com/yoshiky/zaimapi/cmd/entity"
	"github.com/aws/aws-lambda-go/events"
)

type HistoryInputPort interface {
	GetHistoryByDate(r events.APIGatewayProxyRequest, startDate string, endDate string) events.APIGatewayProxyResponse
}

type HistoryOutputPort interface {
	OutputHistories(*entity.ZaimMoney) events.APIGatewayProxyResponse
	OutputError(msg string, code int) events.APIGatewayProxyResponse
}

// Zaimからデータ取得用のポート
type HistoryRepository interface {
	GetHistoryByDate(request events.APIGatewayProxyRequest, startDate string, endDate string) (entity.ZaimMoney, error)
}
