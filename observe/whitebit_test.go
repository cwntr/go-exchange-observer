package observe

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestWhitebitTicker(t *testing.T) {
	w := NewWhitebitClient(WhiteBitHost)
	resp, err := w.GetTicker()
	fmt.Printf("XSNUSDT: %v\n", resp.Result.XSNUSDT)
	fmt.Printf("BTCUSDT: %v\n", resp.Result.BTCUSDT)
	fmt.Printf("err: %v\n", err)

	askXSNUSDT, err := decimal.NewFromString(resp.Result.XSNUSDT.Ticker.Ask)
	bidXSNUSDT, err := decimal.NewFromString(resp.Result.XSNUSDT.Ticker.Bid)

	askBTCUSDT, err := decimal.NewFromString(resp.Result.BTCUSDT.Ticker.Ask)
	bidBTCUSDT, err := decimal.NewFromString(resp.Result.BTCUSDT.Ticker.Bid)
	if err != nil {
		t.Error()
		return
	}

	satsXSNBTCask := askXSNUSDT.Div(askBTCUSDT).Round(8)
	satsXSNBTCbid := bidXSNUSDT.Div(bidBTCUSDT).Round(8)
	fmt.Printf("XSN PRICE sats sell: %s\n", satsXSNBTCask)
	fmt.Printf("XSN PRICE sats buy: %s\n", satsXSNBTCbid)

}
