package observe

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
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
	PairETHUSDT = "ETH_USDT"
	PairBTCUSDT = "BTC_USDT"
	PairETHUSDC = "ETH_USDC"
	PairBTCUSDC = "BTC_USDC"
	PairBTCETH  = "BTC_ETH"
	PairWBTCETH = "WBTC_ETH"

	CurrencyXSN  = "XSN"
	CurrencyWBTC = "WBTC"
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
	CurrencyUSDT = "USDT"
	CurrencyUSDC = "USDC"

	SourceBinance  = "Binance"
	SourceWhitebit = "Whitebit"
)

var O Observer

type Observer struct {
	*BinanceClient
	*WhitebitClient
	Ticker time.Duration
}

type PricePair struct {
	Pair    string
	Price   float64
	Bid     float64
	Ask     float64
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
		fmt.Printf("err while initializing o.BinanceClient: %v", err)
		return err
	}
	o.WhitebitClient = NewWhitebitClient(WhiteBitHost)
	O = o
	return nil
}

func GetPrices() ([]PricePair, []Currency, error) {
	var p []PricePair
	var c []Currency

	wbTicker, err := O.WhitebitClient.GetTicker()
	if err != nil {
		fmt.Printf("err while O.WhitebitClient.GetTicker(): %v", err)
		return p, c, err
	}

	askXSNUSDT, err := decimal.NewFromString(wbTicker.Result.XSNUSDT.Ticker.Ask)
	bidXSNUSDT, err := decimal.NewFromString(wbTicker.Result.XSNUSDT.Ticker.Bid)

	askBTCUSDT, err := decimal.NewFromString(wbTicker.Result.BTCUSDT.Ticker.Ask)
	bidBTCUSDT, err := decimal.NewFromString(wbTicker.Result.BTCUSDT.Ticker.Bid)
	if err != nil {
		fmt.Printf("err decimal.NewFromString: %v", err)
		return p, c, err
	}

	//Binance
	bPairs, err := O.BinanceClient.GetPairs(getActiveBinancePairs())
	if err != nil {
		fmt.Printf("err while O.BinanceClient.GetPairs: %v", err)
		return p, c, err
	}

	var BTCinUSD float64

	var XSNinUSD float64
	var XSNinBTC float64
	var XSNAskInBTC, _ = askXSNUSDT.Div(askBTCUSDT).Round(8).Float64()
	var XSNBidInBTC, _ = bidXSNUSDT.Div(bidBTCUSDT).Round(8).Float64()

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
	var WBTCinETH float64

	// set BTC USD first, since its mandatory for converting other coins
	for _, pair := range bPairs {
		if pair.Symbol == BinancePairBTCUSDT {
			BTCinUSD, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairBTCUSDT: %v", err)
				continue
			}
		}

	}

	//ask is market price
	XSNinUSD = BTCinUSD / (1 / XSNAskInBTC)
	XSNinBTC = XSNAskInBTC

	for _, pair := range bPairs {
		if pair.Symbol == BinancePairLTCBTC {
			LTCinBTC, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairLTCBTC: %v", err)
				continue
			}
		} else if pair.Symbol == BinancePairWBTCETH {
			WBTCinETH, err = strconv.ParseFloat(pair.Price, 64)
			if err != nil {
				fmt.Printf("err parsing BinancePairWBTCETH: %v", err)
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

	p = append(p, PricePair{Pair: PairXSNBTC, Price: XSNinBTC, Ask: XSNAskInBTC, Bid: XSNBidInBTC, Sources: []string{SourceWhitebit}})
	p = append(p, PricePair{Pair: PairLTCBTC, Price: LTCinBTC, Ask: LTCinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairETHBTC, Price: ETHinBTC, Ask: ETHinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairETHUSDT, Price: ETHinUSD, Ask: ETHinUSD, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairBTCUSDT, Price: BTCinUSD, Ask: BTCinUSD, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairETHUSDC, Price: ETHinUSD, Ask: ETHinUSD, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairBTCUSDC, Price: BTCinUSD, Ask: BTCinUSD, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairDCRBTC, Price: DCRinBTC, Ask: DCRinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairXLMBTC, Price: XLMinBTC, Ask: XLMinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairEOSBTC, Price: EOSinBTC, Ask: EOSinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairZECBTC, Price: ZECinBTC, Ask: ZECinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairBNBBTC, Price: BNBinBTC, Ask: BNBinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairADABTC, Price: ADAinBTC, Ask: ADAinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairXTZBTC, Price: XTZinBTC, Ask: XTZinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairATOMBTC, Price: ATOMinBTC, Ask: ATOMinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairZRXBTC, Price: ZRXinBTC, Ask: ZRXinBTC, Sources: []string{SourceBinance}})
	p = append(p, PricePair{Pair: PairWBTCETH, Price: WBTCinETH, Ask: WBTCinETH, Sources: []string{SourceBinance}})

	btcEthPrice := BTCinUSD / ETHinUSD
	p = append(p, PricePair{Pair: PairBTCETH, Price: btcEthPrice, Ask: btcEthPrice, Sources: []string{SourceBinance}})

	c = append(c, Currency{Symbol: CurrencyBTC, PriceUSD: BTCinUSD, PriceBTC: float64(1)})
	c = append(c, Currency{Symbol: CurrencyUSDT, PriceUSD: float64(1), PriceBTC: BTCinUSD})
	c = append(c, Currency{Symbol: CurrencyUSDC, PriceUSD: float64(1), PriceBTC: BTCinUSD})
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
