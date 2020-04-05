package observe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	LiveCoinHost                   = "https://api.livecoin.net/"
	LiveCoinRouteMaxBidMinAsk      = "/exchange/maxbid_minask"
	LiveCoinQueryParamCurrencyPair = "currencyPair"
	LiveCoinPairXSNBTC             = "XSN/BTC"
)

type CurrencyPair struct {
	Symbol string `json:"symbol"`
	MaxBid string `json:"maxBid"`
	MinAsk string `json:"minAsk"`
}

type Pair struct {
	CurrencyPairs []CurrencyPair `json:"currencyPairs"`
}

type LivecoinClient struct {
	ApiHost string
	Client  *http.Client
}

func NewLivecoinClient(apiHost string) *LivecoinClient {
	lc := &LivecoinClient{}
	lc.ApiHost = apiHost
	lc.Client = &http.Client{}
	return lc
}

func (l *LivecoinClient) GetMaxBidMinAsk(pair string) (CurrencyPair, error) {
	fmt.Println(l.ApiHost)
	u, err := url.Parse(l.ApiHost + LiveCoinRouteMaxBidMinAsk)
	if err != nil {
		return CurrencyPair{}, err
	}
	q := u.Query()
	q.Set(LiveCoinQueryParamCurrencyPair, pair)
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
	var p Pair
	err = l.getJson(u.String(), &p)
	if err != nil {
		return CurrencyPair{}, err
	}
	return p.CurrencyPairs[0], nil
}

func (l *LivecoinClient) getJson(url string, target interface{}) error {
	fmt.Println(url)
	r, err := l.Client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
