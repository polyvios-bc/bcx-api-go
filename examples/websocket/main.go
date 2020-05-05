package main

import (
	"bcx-api-go/client/configuration"
	"bcx-api-go/client/websocket"
	"log"
	"time"
)

func main() {
	config := configuration.Configuration{
		Host:    websocket.WsTestEndpoint,
		ApiKey:  "",
		Timeout: 5 * time.Second,
		Env:     configuration.STAGING,
	}

	wsClient := websocket.NewWebSocketClient(config)

	err := wsClient.Start(true)

	if err != nil {
		log.Fatal(err)
	}
	go listenForUpdates(wsClient)

	subscribeToHeartbeats(wsClient)
	subscribeToSymbols(wsClient)
	subscribeToL3(wsClient, websocket.BTCUSD)
	subscribeToL2(wsClient, websocket.ETHBTC)
	subscribeToTrades(wsClient, websocket.BTCUSD)
	subscribeToPrices(wsClient, websocket.BTCUSD, websocket.Granularity60)
	subscribeToTicker(wsClient, websocket.BTCUSD)
	subscribeToBalances(wsClient)
	subscribeToTrading(wsClient)

	sendNewOrder(wsClient)
	wsClient.BulkCancel(nil)
	time.Sleep(5 * time.Second)
	wsClient.Stop()
}

func listenForUpdates(wsClient *websocket.WebSocketClient) {
	for {
		select {
		case balancesMsg := <-wsClient.Balances():
			log.Printf("received balances %+v", balancesMsg)
			break
		case l3Msg := <-wsClient.L3Quotes():
			log.Printf("received l3 quote %+v", l3Msg)
			break
		case l2Msg := <-wsClient.L2Quotes():
			log.Printf("received l2 quote %+v", l2Msg)
			break
		case heartbeatMsg := <-wsClient.Heartbeats():
			log.Printf("received heartbeat %+v", heartbeatMsg)
			break
		case pricesMsg := <-wsClient.Prices():
			log.Printf("received price %+v", pricesMsg)
			break
		case tradesMsg := <-wsClient.Trades():
			log.Printf("received trades %+v", tradesMsg)
			break
		case tickerMsg := <-wsClient.Ticker():
			log.Printf("received ticker %+v", tickerMsg)
			break
		case tradingMsg := <-wsClient.Trading():
			if tradingMsg.IsSnapshot() {
				log.Printf("received trading snapshot %+v", tradingMsg)
			} else if tradingMsg.IsUpdate() {
				log.Printf("received trading update %+v", tradingMsg)
			} else if tradingMsg.IsReject() {
				log.Printf("received trading reject %+v", tradingMsg)
			} else {
				log.Printf("received unknown trading msg %+v", tradingMsg)
			}
			break
		case symbolsMsg := <-wsClient.Symbols():
			log.Printf("received symbols %+v", symbolsMsg)
			break
		}
	}
}

// SUBSCRIPTION EXAMPLES

func subscribeToHeartbeats(wsClient *websocket.WebSocketClient) {
	err := wsClient.SubscribeHeartbeat()

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToSymbols(wsClient *websocket.WebSocketClient) {
	err := wsClient.SubscribeToSymbols()
	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToL3(wsClient *websocket.WebSocketClient, symbol websocket.Symbol) {
	err := wsClient.SubscribeToL3(symbol)

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToL2(wsClient *websocket.WebSocketClient, symbol websocket.Symbol) {
	err := wsClient.SubscribeToL2(symbol)

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToPrices(wsClient *websocket.WebSocketClient, symbol websocket.Symbol, granularity websocket.Granularity) {
	err := wsClient.SubscribeToPrices(symbol, granularity)

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToTicker(wsClient *websocket.WebSocketClient, symbol websocket.Symbol) {
	err := wsClient.SubscribeToTicker(symbol)

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToTrades(wsClient *websocket.WebSocketClient, symbol websocket.Symbol) {
	err := wsClient.SubscribeToTrades(symbol)

	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToBalances(wsClient *websocket.WebSocketClient) {
	err := wsClient.SubscribeToBalances()

	if err != nil {
		log.Fatal(err)
	}
}

func authenticate(wsClient *websocket.WebSocketClient, token string) {
	err := wsClient.Authenticate(token)
	if err != nil {
		log.Fatal(err)
	}
}

func subscribeToTrading(wsClient *websocket.WebSocketClient) {
	err := wsClient.SubscribeToTrading()

	if err != nil {
		log.Fatal(err)
	}
}

// Trading Examples

func sendNewOrder(wsClient *websocket.WebSocketClient) {
	wsClient.NewOrderSingleMessage(websocket.NewOrderSingleMsg{
		ClOrdID:     "polyviosTest10",
		Symbol:      websocket.BTCUSD,
		OrdType:     websocket.LIMIT,
		TimeInForce: websocket.GTC,
		Side:        websocket.BUY,
		OrderQty:    10.0,
		Price:       3400.0,
		ExecInst:    websocket.ALO,
	})
}
