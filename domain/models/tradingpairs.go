package models

type TradingPairs struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

type BinanceTradingPairs struct {
	Symbols []struct {
		Symbol string `json:"symbol"`
	} `json:"symbols"`
}

type CoinbaseTradingPair struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}
