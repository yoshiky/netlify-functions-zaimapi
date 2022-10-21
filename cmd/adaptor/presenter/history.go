package presenter

import (
	"encoding/json"
	"fmt"

	"github.com/yoshiky/zaimapi/cmd/entity"
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
)

type HistoryPresenter struct {
}

func NewHistoryOutputPort() ports.HistoryOutputPort {
	return &HistoryPresenter{}
}

func (h *HistoryPresenter) OutputHistories(z *entity.ZaimMoney) string {
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
	return string(res)
}

func (h *HistoryPresenter) OutputError(err error) string {
	// return events.APIGatewayProxyResponse{
	// 	Body:       err.Error(),
	// 	StatusCode: 500,
	// }, nil
	return err.Error()
}
