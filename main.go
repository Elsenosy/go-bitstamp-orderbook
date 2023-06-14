package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-bitstamp-orderbook/models/requests"
	"go-bitstamp-orderbook/models/response"
	"log"
	"nhooyr.io/websocket"
)

const (
	EVENT   = "bts:subscribe"
	CHANNEL = "order_book_btcusd"
	WS_URL  = "wss://ws.bitstamp.net"
)

func main() {
	// Create a background context
	ctx := context.Background()

	// Parsing requests
	requestInJson, err := json.Marshal(prepareRequest(EVENT, CHANNEL))

	if err != nil {
		log.Fatal(err)
	}

	callBitstampWs(ctx, WS_URL, requestInJson)
}

func prepareRequest(event string, channel string) requests.Message {
	// Prepare request message
	var orderBookChannel requests.Channel
	orderBookChannel.Channel = "order_book_btcusd"

	var socketMessage requests.Message
	socketMessage.Event = "bts:subscribe"
	socketMessage.Data = orderBookChannel

	return socketMessage
}

func callBitstampWs(ctx context.Context, ws_url string, requestInJson []byte) {
	// Trigger the connection
	connection, _, err := websocket.Dial(ctx, "wss://ws.bitstamp.net", nil)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close(websocket.StatusInternalError, "")

	err = connection.Write(ctx, websocket.MessageText, requestInJson)

	if err != nil {
		log.Fatal(err)
	}

	for {
		_, message, err := connection.Read(ctx)

		var msgResponse response.BitstampResponse
		json.Unmarshal(message, &msgResponse)

		if err != nil {
			log.Fatal("Something went wrong!, error: ", err)
		}

		fmt.Printf("Message Structure: %v \n", msgResponse.Data)
	}

	if err != nil {
		log.Fatal(err)
	}

	connection.Close(websocket.StatusNormalClosure, "No reason")
}
