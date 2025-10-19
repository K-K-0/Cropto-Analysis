package model

type Trade struct {
	EventType     string `json:"e"`
	EventTime     int64  `json:"E"`
	Symbol        string `json:"s"`
	TradeID       int64  `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderID  int64  `json:"b"`
	SellerOrderID int64  `json:"a"`
	TradeTime     int64  `json:"T"`
	IsBuyerMaker  bool   `json:"m"`
}
