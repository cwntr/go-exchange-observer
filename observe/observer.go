package observe

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cwntr/go-stakenet/explorer"
)

const (
	PairXSNBTC  = "XSN_BTC"
	PairXSNLTC  = "XSN_LTC"
	PairLTCBTC  = "LTC_BTC"
	PairETHBTC  = "ETH_BTC"
	PairDCRBTC  = "DCR_BTC"
	PairXLMBTC  = "XLM_BTC"
	PairEOSBTC  = "EOS_BTC"
	PairZECBTC  = "ZEC_BTC"
	PairBNBBTC  = "BNB_BTC"
	PairADABTC  = "ADA_BTC"
	PairXTZBTC  = "XTZ_BTC"
	PairATOMBTC = "ATOM_BTC"
	PairZRXBTC  = "ZRX_BTC"

	CurrencyXSN  = "XSN"
	CurrencyBTC  = "BTC"
	CurrencyLTC  = "LTC"
	CurrencyETH  = "ETH"
	CurrencyDCR  = "DCR"
	CurrencyXLM  = "XLM"
	CurrencyEOS  = "EOS"
	CurrencyZEC  = "ZEC"
	CurrencyBNB  = "BNB"
	CurrencyADA  = "ADA"
	CurrencyXTZ  = "XTZ"
	CurrencyATOM = "ATOM"
	CurrencyZRX  = "ZRX"

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

	var LTCinUSD, LTCinBTC float64
	var ETHinUSD, ETHinBTC float64
	var DCRinUSD, DCRinBTC float64
	var XLMinUSD, XLMinBTC float64
	var EOSinUSD, EOSinBTC float64
	var ZECinUSD, ZECinBTC float64
	var BNBinUSD, BNBinBTC float64
	var ADAinUSD, ADAinBTC float64
	var XTZinUSD, XTZinBTC float64
	var ATOMinUSD, ATOMinBTC float64
	var ZRXinUSD, ZRXinBTC float64

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
		} else if pair.Symbol == BinancePairETHBTC {
			ETHinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairETHBTC: %v", err)
				continue
			}
			ETHinUSD = BTCinUSD / (1 / ETHinBTC)
		} else if pair.Symbol == BinancePairDCRBTC {
			DCRinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairDCRBTC: %v", err)
				continue
			}
			DCRinUSD = BTCinUSD / (1 / DCRinBTC)
		} else if pair.Symbol == BinancePairXLMBTC {
			XLMinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairXLMBTC: %v", err)
				continue
			}
			XLMinUSD = BTCinUSD / (1 / XLMinBTC)
		} else if pair.Symbol == BinancePairEOSBTC {
			EOSinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairEOSBTC: %v", err)
				continue
			}
			EOSinUSD = BTCinUSD / (1 / EOSinBTC)
		} else if pair.Symbol == BinancePairZECBTC {
			ZECinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairZECBTC: %v", err)
				continue
			}
			ZECinUSD = BTCinUSD / (1 / ZECinBTC)
		} else if pair.Symbol == BinancePairBNBBTC {
			BNBinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairBNBBTC: %v", err)
				continue
			}
			BNBinUSD = BTCinUSD / (1 / BNBinBTC)
		} else if pair.Symbol == BinancePairADABTC {
			ADAinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairADATBTC: %v", err)
				continue
			}
			ADAinUSD = BTCinUSD / (1 / ADAinBTC)
		} else if pair.Symbol == BinancePairZRXBTC {
			ZRXinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairZRXBTC: %v", err)
				continue
			}
			ZRXinUSD = BTCinUSD / (1 / ZRXinBTC)
		} else if pair.Symbol == BinancePairATOMBTC {
			ATOMinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairATOMBTC: %v", err)
				continue
			}
			ATOMinUSD = BTCinUSD / (1 / ATOMinBTC)
		} else if pair.Symbol == BinancePairXTZBTC {
			XTZinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairXTZTBTC: %v", err)
				continue
			}
			XTZinUSD = BTCinUSD / (1 / XTZinBTC)
		}
	}

	XSNinLTC = xsnBTC / LTCinBTC
	p = append(p, PricePair{Pair: PairXSNBTC, Price: XSNinBTC, Sources: []string{SourceLivecoin}})
	p = append(p, PricePair{Pair: PairXSNLTC, Price: XSNinLTC, Sources: []string{SourceLivecoin, SourceBinance}})
	p = append(p, PricePair{Pair: PairLTCBTC, Price: LTCinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairETHBTC, Price: ETHinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairDCRBTC, Price: DCRinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairXLMBTC, Price: XLMinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairEOSBTC, Price: EOSinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairZECBTC, Price: ZECinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairBNBBTC, Price: BNBinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairADABTC, Price: ADAinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairXTZBTC, Price: XTZinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairATOMBTC, Price: ATOMinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairZRXBTC, Price: ZRXinBTC, Sources: []string{SourceBinance}})

	c = append(c, Currency{Symbol: CurrencyBTC, PriceUSD: BTCinUSD, PriceBTC: float64(1)})
	c = append(c, Currency{Symbol: CurrencyLTC, PriceUSD: LTCinUSD, PriceBTC: LTCinBTC})
	c = append(c, Currency{Symbol: CurrencyXSN, PriceUSD: XSNinUSD, PriceBTC: XSNinBTC})
	c = append(c, Currency{Symbol: CurrencyETH, PriceUSD: ETHinUSD, PriceBTC: ETHinBTC})
	c = append(c, Currency{Symbol: CurrencyDCR, PriceUSD: DCRinUSD, PriceBTC: DCRinBTC})
	c = append(c, Currency{Symbol: CurrencyXLM, PriceUSD: XLMinUSD, PriceBTC: XLMinBTC})
	c = append(c, Currency{Symbol: CurrencyEOS, PriceUSD: EOSinUSD, PriceBTC: EOSinBTC})
	c = append(c, Currency{Symbol: CurrencyZEC, PriceUSD: ZECinUSD, PriceBTC: ZECinBTC})
	c = append(c, Currency{Symbol: CurrencyBNB, PriceUSD: BNBinUSD, PriceBTC: BNBinBTC})
	c = append(c, Currency{Symbol: CurrencyADA, PriceUSD: ADAinUSD, PriceBTC: ADAinBTC})
	c = append(c, Currency{Symbol: CurrencyXTZ, PriceUSD: XTZinUSD, PriceBTC: XTZinBTC})
	c = append(c, Currency{Symbol: CurrencyATOM, PriceUSD: ATOMinUSD, PriceBTC: ATOMinBTC})
	c = append(c, Currency{Symbol: CurrencyZRX, PriceUSD: ZRXinUSD, PriceBTC: ZRXinBTC})
	return p, c, nil
}
