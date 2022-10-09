package rest_test

import (
	"github.com/outtoin/bybit-api/rest"
	"testing"
)

func newByBitAll() *rest.ByBit {
	baseURL := "https://api-testnet.bybit.com/"
	apiKey := "xQAVwCZYVpUhDLirO7"
	secretKey := "4hU8mzLRIErATqrZiGVo1KrLmZcorrHttf3Y"
	b := rest.New(nil, baseURL, apiKey, secretKey, true)
	return b
}

func TestByBit_SwitchIsolated(t *testing.T) {
	b := newByBitAll()
	_, _, err := b.LinearSwitchIsolated("BTCUSDT", true, 10, 10)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestByBit_SetTradingStop(t *testing.T) {
	b := newByBitAll()
	_, _, err := b.LinearSetTradingStop("BTCUSDT", "Sell", 20000.0, 19000.0, 0.0, 0)
	if err != nil {
		t.Error(err)
		return
	}
}
