package observe

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

const (
	BinancePairLTCBTC  = "LTCBTC"
	BinancePairBTCUSDT = "BTCUSDT"
	BinancePairLTCUSDT = "LTCUSDT"
	BinancePairETHUSDT = "ETHUSDT"
	BinancePairETHUSDC = "ETHUSDC"
	BinancePairBTCUSDC = "BTCUSDC"
	BinancePairETHBTC  = "ETHBTC"
	BinancePairDCRBTC  = "DCRBTC"
	BinancePairXLMBTC  = "XLMBTC"
	BinancePairEOSBTC  = "EOSBTC"
	BinancePairZECBTC  = "ZECBTC"
	BinancePairBNBBTC  = "BNBBTC"
	BinancePairADABTC  = "ADABTC"
	BinancePairXTZBTC  = "XTZBTC"
	BinancePairATOMBTC = "ATOMBTC"
	BinancePairZRXBTC  = "ZRXBTC"
)

type BinanceClient struct {
	client *binance.Client
}

func (b *BinanceClient) GetPairs(symbols []string) ([]binance.SymbolPrice, error) {
	prices, err := b.client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	var list []binance.SymbolPrice
	for _, p := range prices {
		for _, s := range symbols {
			if p.Symbol == s {
				list = append(list, *p)
			}
		}
	}
	return list, nil
}

func NewBinanceClient(apiKey string, secret string) (*BinanceClient, error) {
	b := &BinanceClient{}
	clt := binance.NewClient(apiKey, secret)
	if clt == nil {
		return b, fmt.Errorf("unable to create binance client")
	}
	b.client = clt
	return b, nil
}

func getActiveBinancePairs() []string {
	return []string{
		BinancePairLTCBTC,
		BinancePairBTCUSDT,
		BinancePairBTCUSDC,
		BinancePairETHUSDC,
		BinancePairLTCUSDT,
		BinancePairETHUSDT,
		BinancePairETHBTC,
		BinancePairDCRBTC,
		BinancePairXLMBTC,
		BinancePairEOSBTC,
		BinancePairZECBTC,
		BinancePairBNBBTC,
		BinancePairADABTC,
		BinancePairXTZBTC,
		BinancePairATOMBTC,
		BinancePairZRXBTC,
	}
}
