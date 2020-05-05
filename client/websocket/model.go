package websocket

import "time"

type Symbol string
type Granularity int
type TimeInForce string
type OrderType string
type OrderStatus string
type OrderSide string
type ExecInst string
type MsgType string
type ExecType string

type heartbeatSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
}

type symbolsSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
}

type tradingSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
}

type quoteSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	Symbol  Symbol     `json:"symbol"`
}

type pricesSubscriptionRequest struct {
	Action      actionType  `json:"action"`
	Channel     channel     `json:"channel"`
	Symbol      Symbol      `json:"symbol"`
	Granularity Granularity `json:"granularity"`
}

type authSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	Token   string     `json:"token"`
}

type tickerSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	Symbol  Symbol     `json:"symbol"`
}

type balancesSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
}

type tradesSubscriptionRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	Symbol  Symbol     `json:"symbol"`
}

type HeartbeatMsg struct {
	Event     eventType `json:"event"`
	Channel   channel   `json:"channel"`
	SeqNum    int64     `json:"seqNum"`
	Timestamp time.Time `json:"timestamp"`
}

type msgCommon struct {
	Event   eventType `json:"event"`
	Channel channel   `json:"channel"`
	SeqNum  int64     `json:"seqNum"`
}

type SymbolsSnapshot struct {
	Seqnum    int                  `json:"seqnum"`
	Event     string               `json:"event"`
	Channel   string               `json:"channel"`
	Timestamp time.Time            `json:"timestamp"`
	Symbols   map[Symbol]SymbolMsg `json:"symbols"`
}

type SymbolMsg struct {
	Name                   Symbol  `json:"name"`
	BaseCurrency           string  `json:"base_currency"`
	BaseCurrencyScale      int     `json:"base_currency_scale"`
	CounterCurrency        string  `json:"counter_currency"`
	CounterCurrencyScale   int     `json:"counter_currency_scale"`
	MinPriceIncrement      int     `json:"min_price_increment"`
	MinPriceIncrementScale int     `json:"min_price_increment_scale"`
	MinOrderSize           int     `json:"min_order_size"`
	MinOrderSizeScale      int     `json:"min_order_size_scale"`
	MaxOrderSize           int     `json:"max_order_size"`
	MaxOrderSizeScale      int     `json:"max_order_size_scale"`
	LotSize                int     `json:"lot_size"`
	LotSizeScale           int     `json:"lot_size_scale"`
	Status                 string  `json:"status"`
	ID                     int     `json:"id"`
	AuctionPrice           float64 `json:"auction_price"`
	AuctionSize            float64 `json:"auction_size"`
	AuctionTime            string  `json:"auction_time"`
	Imbalance              float64 `json:"imbalance"`
}

type L3Msg struct {
	Seqnum  int    `json:"seqnum"`
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Symbol  string `json:"symbol"`
	Bids    []struct {
		ID  string  `json:"id"`
		Px  float64 `json:"px"`
		Qty float64 `json:"qty"`
	} `json:"bids"`
	Asks []struct {
		ID  string  `json:"id"`
		Px  float64 `json:"px"`
		Qty float64 `json:"qty"`
	} `json:"asks"`
}

type L2Msg struct {
	Seqnum  int    `json:"seqnum"`
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Symbol  string `json:"symbol"`
	Bids    []struct {
		Px  float64 `json:"px"`
		Qty float64 `json:"qty"`
		Num int     `json:"num"`
	} `json:"bids"`
	Asks []struct {
		Px  float64 `json:"px"`
		Qty float64 `json:"qty"`
		Num int     `json:"num"`
	} `json:"asks"`
}

type PricesMsg struct {
	Seqnum  int       `json:"seqnum"`
	Event   string    `json:"event"`
	Channel string    `json:"channel"`
	Symbol  string    `json:"symbol"`
	Price   []float64 `json:"price"`
}

type TickerMsg struct {
	Seqnum         int     `json:"seqnum"`
	Event          string  `json:"event"`
	Channel        string  `json:"channel"`
	Symbol         string  `json:"symbol"`
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type TradesMsg struct {
	Seqnum    int       `json:"seqnum"`
	Event     string    `json:"event"`
	Channel   string    `json:"channel"`
	Symbol    string    `json:"symbol"`
	Timestamp time.Time `json:"timestamp"`
	Side      string    `json:"side"`
	Qty       float64   `json:"qty"`
	Price     float64   `json:"price"`
	TradeID   string    `json:"trade_id"`
}

type BalancesSnapshot struct {
	Seqnum              int          `json:"seqnum"`
	Event               string       `json:"event"`
	Channel             string       `json:"channel"`
	Balances            []BalanceMsg `json:"balances"`
	TotalAvailableLocal float64      `json:"total_available_local"`
	TotalBalanceLocal   float64      `json:"total_balance_local"`
}

type BalanceMsg struct {
	Currency       string  `json:"currency"`
	Balance        float64 `json:"balance"`
	Available      float64 `json:"available"`
	BalanceLocal   float64 `json:"balance_local"`
	AvailableLocal float64 `json:"available_local"`
	Rate           float64 `json:"rate"`
}

type RejectMsg struct {
	Seqnum  int    `json:"seqnum"`
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type TradingMsg interface {
	IsSnapshot() bool
	IsReject() bool
	IsUpdate() bool
}

type TradingSnapshot struct {
	Seqnum  int     `json:"seqnum"`
	Event   string  `json:"event"`
	Channel string  `json:"channel"`
	Orders  []Order `json:"orders"`
}

func (t *TradingSnapshot) IsSnapshot() bool { return true }
func (t *TradingSnapshot) IsReject() bool   { return false }
func (t *TradingSnapshot) IsUpdate() bool   { return false }

type Order struct {
	OrderID      string    `json:"orderID"`
	ClOrdID      string    `json:"clOrdID"`
	Symbol       string    `json:"symbol"`
	Side         string    `json:"side"`
	OrdType      string    `json:"ordType"`
	OrderQty     float64   `json:"orderQty"`
	LeavesQty    float64   `json:"leavesQty"`
	CumQty       float64   `json:"cumQty"`
	AvgPx        float64   `json:"avgPx"`
	OrdStatus    string    `json:"ordStatus"`
	TimeInForce  string    `json:"timeInForce"`
	Text         string    `json:"text"`
	ExecType     string    `json:"execType"`
	ExecID       string    `json:"execID"`
	TransactTime time.Time `json:"transactTime"`
	MsgType      int       `json:"msgType"`
	LastPx       float64   `json:"lastPx"`
	LastShares   float64   `json:"lastShares"`
	TradeID      string    `json:"tradeId"`
	Price        float64   `json:"price"`
}

type TradingUpdated struct {
	Seqnum       int       `json:"seqnum"`
	Event        string    `json:"event"`
	Channel      string    `json:"channel"`
	OrderID      string    `json:"orderID"`
	ClOrdID      string    `json:"clOrdID"`
	Symbol       string    `json:"symbol"`
	Side         string    `json:"side"`
	OrdType      string    `json:"ordType"`
	OrderQty     float64   `json:"orderQty"`
	LeavesQty    float64   `json:"leavesQty"`
	CumQty       float64   `json:"cumQty"`
	AvgPx        float64   `json:"avgPx"`
	OrdStatus    string    `json:"ordStatus"`
	TimeInForce  string    `json:"timeInForce"`
	Text         string    `json:"text"`
	ExecType     string    `json:"execType"`
	ExecID       string    `json:"execID"`
	TransactTime time.Time `json:"transactTime"`
	MsgType      int       `json:"msgType"`
	LastPx       float64   `json:"lastPx"`
	LastShares   float64   `json:"lastShares"`
	TradeID      string    `json:"tradeId"`
	Price        float64   `json:"price"`
}

func (t *TradingUpdated) IsSnapshot() bool { return false }
func (t *TradingUpdated) IsReject() bool   { return false }
func (t *TradingUpdated) IsUpdate() bool   { return true }

type TradingReject struct {
	Seqnum    int    `json:"seqnum"`
	Event     string `json:"event"`
	Channel   string `json:"channel"`
	Text      string `json:"text"`
	ClOrdID   string `json:"clOrdID"`
	OrdStatus string `json:"ordStatus"`
	Action    string `json:"action"`
}

func (t *TradingReject) IsSnapshot() bool { return false }
func (t *TradingReject) IsReject() bool   { return true }
func (t *TradingReject) IsUpdate() bool   { return false }

type newOrderSingleRequest struct {
	Action      actionType  `json:"action"`
	Channel     channel     `json:"channel"`
	ClOrdID     string      `json:"clOrdID"`
	Symbol      Symbol      `json:"symbol"`
	OrdType     OrderType   `json:"ordType"`
	TimeInForce TimeInForce `json:"timeInForce"`
	Side        OrderSide   `json:"side"`
	OrderQty    float64     `json:"orderQty"`
	Price       float64     `json:"price"`
	ExecInst    ExecInst    `json:"execInst"`
}

type NewOrderSingleMsg struct {
	ClOrdID     string      `json:"clOrdID"`
	Symbol      Symbol      `json:"symbol"`
	OrdType     OrderType   `json:"ordType"`
	TimeInForce TimeInForce `json:"timeInForce"`
	Side        OrderSide   `json:"side"`
	OrderQty    float64     `json:"orderQty"`
	Price       float64     `json:"price"`
	ExecInst    ExecInst    `json:"execInst"`
}

type cancelOrderRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	OrderID string     `json:"orderID"`
}

type bulkCancelOrderRequest struct {
	Action  actionType `json:"action"`
	Channel channel    `json:"channel"`
	Symbol  Symbol     `json:"symbol,omitempty"`
}
