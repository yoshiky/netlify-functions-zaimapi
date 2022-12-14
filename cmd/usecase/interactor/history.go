package interactor

/*
interactorはUseCaseで定義した機能の実装
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．
interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/yoshiky/zaimapi/cmd/internal/codes"
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
)

// TOOD 頭文字小文字にして隠蔽する？
type HistoryInteractor struct {
	OutputPort  ports.HistoryOutputPort
	HistoryRepo ports.HistoryRepository
}

func NewHistoryInputPort(outputPort ports.HistoryOutputPort, historyRepository ports.HistoryRepository) ports.HistoryInputPort {
	return &HistoryInteractor{
		OutputPort:  outputPort,
		HistoryRepo: historyRepository,
	}
}

// usecase(port)HistoryInputPort の実装
// HistoryInteractorにGetHistoryByDate()を実装しているので、HistoryInteractorはHistoryInputPort型でもある(=HistoryInputPortの実装)
// Repoを呼び出し取得した結果をOutputPortに渡す
func (h *HistoryInteractor) GetHistoryByDate(request events.APIGatewayProxyRequest, startDate string, endDate string) events.APIGatewayProxyResponse {

	if startDate == "" || endDate == "" {
		return h.OutputPort.OutputError(string(codes.InvalidParams), 400)
	}

	histories, err := h.HistoryRepo.GetHistoryByDate(request, startDate, endDate)
	if err != nil {
		h.OutputPort.OutputError(err.Error(), 500)
	}
	return h.OutputPort.OutputHistories(&histories)
}
