package observe

import (
	"fmt"
	"github.com/cwntr/go-stakenet/explorer"
	"strconv"
	"strings"
	"time"
)

const (
	PairXSNBTC = "XSN_BTC"
	PairXSNLTC = "XSN_LTC"
	PairLTCBTC = "LTC_BTC"

	CurrencyXSN = "XSN"
	CurrencyBTC = "BTC"
	CurrencyLTC = "LTC"

	SourceBinance  = "Binance"
	SourceLivecoin = "Livecoin"
)

var O Observer

type Observer struct {
	*BinanceClient
	*LivecoinClient
	Ticker time.Duration
}

type PricePair struct {
	Pair    string
	Price   float64
	Sources []string
}

type Currency struct {
	Symbol   string
	PriceUSD float64
	PriceBTC float64
}

func InitClients(binanceKey string, binanceSecret string) (err error) {
	o := Observer{}
	o.BinanceClient, err = NewBinanceClient(binanceKey, binanceSecret)
	if err != nil {
		return err
	}
	o.LivecoinClient = NewLivecoinClient(LiveCoinHost)
	O = o
	return nil
}

func GetPrices() ([]PricePair, []Currency, error) {
	var p []PricePair
	var c []Currency
	var xsnBTC float64
	cp, err := O.LivecoinClient.GetMaxBidMinAsk(LiveCoinPairXSNBTC)
	if err != nil {
		//if Livecoin failed, try the CMC based price from XSN explorer
		api := explorer.NewXSNExplorerAPIClient(nil)
		var xsnPrice float64
		var btcPrice float64

		coins := []string{"XSN", "BTC"}
		for _, coin := range coins {
			price, err := api.GetPrices(strings.ToLower(coin))
			if err != nil {
				fmt.Println("unable to get coin price from XSN Explorer API")
				return p, c, err
			}
			if strings.ToUpper(coin) == "XSN" {
				xsnPrice = price.USD
			}
			if strings.ToUpper(coin) == "BTC" {
				btcPrice = price.USD
			}
		}
		xsnBTC = xsnPrice / btcPrice
	} else {
		xsnBTC, err = strconv.ParseFloat(cp.MaxBid, 64)
		if err != nil {
			return p, c, err
		}
	}
	bPairs, err := O.BinanceClient.GetPairs(getActiveBinancePairs())
	if err != nil {
		return p, c, err
	}

	var BTCinUSD float64

	var XSNinUSD float64
	var XSNinBTC float64
	var XSNinLTC float64

	var LTCinUSD float64
	var LTCinBTC float64

	for _, pair := range bPairs {
		if pair.Symbol == BinancePairBTCUSDT {
			BTCinUSD, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairBTCUSDT: %v", err)
				continue
			}
			XSNinUSD = BTCinUSD / (1 / xsnBTC)
			XSNinBTC = xsnBTC
		} else if pair.Symbol == BinancePairLTCBTC {
			LTCinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairLTCBTC: %v", err)
				continue
			}
		} else if pair.Symbol == BinancePairLTCUSDT {
			LTCinUSD, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairLTCUSDT: %v", err)
				continue
			}
		}
	}
	XSNinLTC = xsnBTC / LTCinBTC
	p = append(p, PricePair{Pair: PairXSNBTC, Price: XSNinBTC, Sources: []string{SourceLivecoin}})
	p = append(p, PricePair{Pair: PairXSNLTC, Price: XSNinLTC, Sources: []string{SourceLivecoin, SourceBinance}})
	p = append(p, PricePair{Pair: PairLTCBTC, Price: LTCinBTC, Sources: []string{SourceBinance}})
	c = append(c, Currency{Symbol: CurrencyBTC, PriceUSD: BTCinUSD, PriceBTC: float64(1)})
	c = append(c, Currency{Symbol: CurrencyLTC, PriceUSD: LTCinUSD, PriceBTC: LTCinBTC})
	c = append(c, Currency{Symbol: CurrencyXSN, PriceUSD: XSNinUSD, PriceBTC: XSNinBTC})
	return p, c, nil
}
