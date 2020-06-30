package observe

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewLivecoinClient(t *testing.T) {
	lc := NewLivecoinClient(LiveCoinHost)
	pair, err := lc.GetMaxBidMinAsk(LiveCoinPairXSNBTC)
	fmt.Printf("pair: %v", pair)
	if err != nil {
		fmt.Printf("err:%v", err)
		t.Fail()
		return
	}

	xsnBTC, err := strconv.ParseFloat(pair.MaxBid, 64)
	fmt.Printf("xsnBTC: %v", xsnBTC)
	if err != nil {
		fmt.Printf("err:%v", err)
		t.Fail()
		return
	}
}
