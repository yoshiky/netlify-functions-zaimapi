package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/yoshiky/zaimapi/cmd/entity"
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
)

type HistoryPresenter struct {
}

func NewHistoryOutputPort() ports.HistoryOutputPort {
	return &HistoryPresenter{}
}

func (h *HistoryPresenter) OutputHistories(z *entity.ZaimMoney) events.APIGatewayProxyResponse {
	// FromAccountId = 4(クレカ払い)のみ抽出する
	m := []entity.Money{}
	for _, v := range z.Money {
		if v.FromAccountId != 4 {
			continue
		}
		m = append(m, v)
	}

	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       string(res),
		StatusCode: 200,
	}
}

func (h *HistoryPresenter) OutputError(msg string, code int) events.APIGatewayProxyResponse {

	e := entity.ErrorResponse{
		Code:    code,
		Message: msg,
	}
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       string(res),
		StatusCode: code,
	}
}
