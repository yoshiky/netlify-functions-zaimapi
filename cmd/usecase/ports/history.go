package ports

import (
	"github.com/yoshiky/zaimapi/cmd/entity"
	"github.com/aws/aws-lambda-go/events"
)

type HistoryInputPort interface {
	GetHistoryByDate(r events.APIGatewayProxyRequest, startDate string, endDate string) string
}

type HistoryOutputPort interface {
	OutputHistories(*entity.ZaimMoney) string
	OutputError(error) string
}

// Zaimからデータ取得用のポート
type HistoryRepository interface {
	GetHistoryByDate(request events.APIGatewayProxyRequest, startDate string, endDate string) (entity.ZaimMoney, error)
}
