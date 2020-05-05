package websocket

const (
	eventSubscribed   eventType = "subscribed"
	eventRejected     eventType = "rejected"
	eventSnapshot     eventType = "snapshot"
	eventUnsubscribed eventType = "unsubscribed"
	eventUpdate       eventType = "updated"

	actionSubscribe actionType = "subscribe"
	newOrderSingle  actionType = "NewOrderSingle"
	cancelOrder     actionType = "CancelOrderRequest"
	bulkCancel      actionType = "BulkCancelOrderRequest"

	heartbeatChannel channel = "heartbeat"
	symbolsChannel   channel = "symbols"
	l3Channel        channel = "l3"
	l2Channel        channel = "l2"
	pricesChannel    channel = "prices"
	tickerChannel    channel = "ticker"
	tradesChannel    channel = "trades"
	balancesChannel  channel = "balances"
	authChannel      channel = "auth"
	tradingChannel   channel = "trading"

	Granularity60    Granularity = 60
	Granularity300   Granularity = 300
	Granularity900   Granularity = 900
	Granularity3600  Granularity = 3600
	Granularity21600 Granularity = 21600
	Granularity86400 Granularity = 86400

	GTC TimeInForce = "GTC" // Good Till Cancel. The order will rest on the order book until it is cancelled or filled
	GTD TimeInForce = "GTD" // Good Till Date. The order will reset on the order book until it is cancelled, filled, or expired
	FOK TimeInForce = "FOK" // Fill or Kill. The order is either completely filled or cancelled. No Partial Fills are permitted.
	IOC TimeInForce = "IOC" // Immediate or Cancel. The order is either a) completely filled, b) partially filled and the remaining quantity canceled, or c) the order is canceled.

	LIMIT      OrderType = "limit"     // order which has a price limit
	MARKET     OrderType = "market"    // order that will match at any price available in the market, starting from the best prices and filling up to the available balance
	STOP       OrderType = "stop"      // order which has a stop/trigger price, and when that price is reached, it triggers a market order
	STOP_LIMIT OrderType = "stopLimit" // order which has a stop price and limit price, and when the stop price is reached, it triggers a limit order at the limit price

	ORDER_STATUS_PENDING   OrderStatus = "pending"   // is pending acceptance. Only applicable to stop and stop-limit orders
	ORDER_STATUS_OPEN      OrderStatus = "open"      // has been accepted
	ORDER_STATUS_REJECTED  OrderStatus = "rejected"  // has been rejected	Limit and market orders can get rejected if you have no balance to fill the order even partially.
	ORDER_STATUS_CANCELLED OrderStatus = "cancelled" //  has been cancelled	A market order might get in state cancelled if you don’t have enough balance to fill it at market price. Both market orders and limit orders with IOC can have ordStatus ‘cancelled’ if there is no market for them, even without the user requesting the cancellation.
	ORDER_STATUS_FILLED    OrderStatus = "filled"    // has been filled	A limit order get in state cancelled after the user requested a cancellation.
	ORDER_STATUS_PARTIAL   OrderStatus = "partial"   // has been partially filled
	ORDER_STATUS_EXPIRED   OrderStatus = "expired"   // has been expired

	BUY  OrderSide = "buy"
	SELL OrderSide = "sell"

	ALO ExecInst = "ALO"

	ExecutionReport     MsgType = "8"
	OrderCancelRejected MsgType = "9"

	EXEC_TYPE_NEW          ExecType = "0"
	EXEC_TYPE_CANCELLED    ExecType = "4"
	EXEC_TYPE_EXPIRED      ExecType = "C"
	EXEC_TYPE_REJECTED     ExecType = "8"
	EXEC_TYPE_PARTIAL_FILL ExecType = "F"
	EXEC_TYPE_PENDING      ExecType = "A"
	EXEC_TYPE_TRADE_BREAK  ExecType = "H"
	EXEC_TYPE_ORDER_STATUS ExecType = "I"
)
