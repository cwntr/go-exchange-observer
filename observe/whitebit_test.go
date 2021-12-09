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
	fmt.Printf("ETHUSDC: %v\n", resp.Result.ETHUSDC)
	fmt.Printf("LTCUSD: %v\n", resp.Result.LTCUSD)
	fmt.Printf("XSN ERR: %v\n", err)

	xsnUSDT, err := decimal.NewFromString(resp.Result.XSNUSDT.Ticker.Ask)
	if err != nil {
		fmt.Printf("XSN ERR: %v\n", err)
		t.Error()
		return
	}
	ethUSDT, err := decimal.NewFromString(resp.Result.ETHUSDC.Ticker.Ask)
	if err != nil {
		fmt.Printf("ETH ERR: %v\n", err)
		t.Error()
		return
	}
	ltcUSD, err := decimal.NewFromString(resp.Result.LTCUSD.Ticker.Ask)

	if err != nil {
		fmt.Printf("LTC ERR: %v\n", err)
		t.Error()
		return
	}
	btcUSD, err := decimal.NewFromString(resp.Result.BTCUSDT.Ticker.Ask)
	if err != nil {
		fmt.Printf("BTC ERR: %v\n", err)
		t.Error()
		return
	}

	for k, r := range []decimal.Decimal{xsnUSDT, ethUSDT, ltcUSD, btcUSD} {
		f, _ := r.Float64()
		fmt.Printf("%d: %.8f \n", k, f)
	}
}
