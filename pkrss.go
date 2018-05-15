// Package HitBTC is an implementation of the HitBTC API in Golang.
package hitbtc

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type DepositAddress struct {
	Address   string `JSON:"address"`
	PaymentId string `JSON:"paymentId"`
}

type DepositAddress struct {
	Address   string `JSON:"address"`
	PaymentId string `JSON:"paymentId"`
}

type TradeRecord struct {
	Id        uint64    `json:"id"`
	Type      string    `json:"side"`
	Price     float64   `json:"price,string"`
	Quantity  float64   `json:"quantity,string"`
	Timestamp time.Time `json:"timestamp"`
}

func (b *HitBtc) MarketTradeHistory(symbol string) (trades []TradeRecord, err error) {
	payload := make(map[string]string)
	if symbol == "" {
		err = errors.New("symbol parameter is empty")
		return
	}
	r, err := b.client.do("GET", "public/trades/"+symbol, payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &trades)
	return
}

func (b *HitBtc) CancelOrderOne(orderId string) (order Order, err error) {
	payload := make(map[string]string)
	if orderId == "" {
		err = errors.New("order id parameter is empty")
		return
	}
	r, err := b.client.do("DELETE", "order/"+orderId, payload, true)
	if err != nil {
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &order)
	return
}

func (b *HitBtc) DepositAddress(currency string) (address *DepositAddress, err error) {
	payload := make(map[string]string)
	if currency == "" {
		err = errors.New("currency parameter is empty")
		return
	}
	r, err2 := b.client.do("GET", "account/crypto/address/"+currency, payload, true)
	if err2 != nil {
		err = err2
		return
	}
	var response interface{}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(r, &address)
	return
}

func (b *HitBtc) Withdraw(currency string, address string, amount float64) (id string, e error) {
	payload := make(map[string]string)
	payload["currency"] = currency
	payload["address"] = address
	payload["amount"] = strconv.FormatFloat(amount, 10, 8, 64)

	r, err := b.client.do("GET", "account/crypto/withdraw", payload, true)
	if err != nil {
		e = err
		return
	}
	var response map[string]string
	if err = json.Unmarshal(r, &response); err != nil {
		e = err
		return
	}
	if err = handleErr(response); err != nil {
		e = err
		return
	}
	id = response["id"]
	return
}
