package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gomodule/oauth1/oauth"

	"github.com/yoshiky/zaimapi/cmd/entity"
	"github.com/yoshiky/zaimapi/cmd/usecase/ports"
)

var (
	ConsumerKey       = os.Getenv("CONSUMER_KEY")
	ConsumerSecret    = os.Getenv("CONSUMER_SECRET")
	AccessToken       = os.Getenv("ACCESS_TOKEN")
	AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

	TemporaryCredentialRequestURI = "https://api.zaim.net/v2/auth/request"
	ResourceOwnerAuthorizationURI = "https://auth.zaim.net/users/auth"
	TokenRequestURI               = "https://api.zaim.net/v2/auth/access"
	zaimMoneyHistoryRequestURI    = "https://api.zaim.net/v2/home/money"
)

type historyRepository struct {
	// httpリクエストする関数か？
}

func NewHistoryRepository() ports.HistoryRepository {
	return &historyRepository{}
}

// historyRepositoryにGetHistoryByDateメソッドを生やすと、HistoryRepositoryインターフェースをhistoryRepositoryに実装したことになる
func (h *historyRepository) GetHistoryByDate(request events.APIGatewayProxyRequest, startDate string, endDate string) (entity.ZaimMoney, error) {

	fmt.Printf("ConsumerKey = %s", ConsumerKey)

	// zaimからデータ取得
	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ConsumerKey,
			Secret: ConsumerSecret,
		},
		TemporaryCredentialRequestURI: TemporaryCredentialRequestURI,
		RenewCredentialRequestURI:     ResourceOwnerAuthorizationURI,
		TokenRequestURI:               TokenRequestURI,
	}

	credentials := &oauth.Credentials{
		Token:  AccessToken,
		Secret: AccessTokenSecret,
	}

	// リクエストパラメタ取得
	mapping := request.QueryStringParameters["mapping"]
	start_date := request.QueryStringParameters["start_date"]
	end_date := request.QueryStringParameters["end_date"]

	query := url.Values{
		"mapping":    []string{mapping},
		"start_date": []string{start_date},
		"end_date":   []string{end_date},
	}

	resp, err := oauthClient.Get(nil, credentials, zaimMoneyHistoryRequestURI, query)
	if err != nil {
		// TODO
		return entity.ZaimMoney{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// レスポンスを構造体に変換
	var zaimMoney entity.ZaimMoney
	if err := json.Unmarshal(body, &zaimMoney); err != nil {
		// TODO
		panic(err.Error())
	}

	return zaimMoney, nil
}
