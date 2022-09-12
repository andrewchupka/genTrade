package integrations

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func handleWebSocketConnection(symbol string, outGoingMessages chan string, token string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	connection, err := openConnection(token)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	// write outgoing message to it
	var body []byte = makeSubscriptionString(symbol)
	fmt.Println(string(body))
	err = connection.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		log.Fatal("handleWebSocketConnection:", err)
		panic(fmt.Sprintf("Error writing on connection: {message: %s}; {error: %s}", body, err))
	}

	messageStream := make(chan string)

	go readMessages(messageStream, connection)

	fmt.Println("Popping messages off the channel")
	for {

		select {
		case message := <-messageStream:
			outGoingMessages <- message
			// log.Println(message)
		case <-interrupt:
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("Error closing connection: {%s}\n", err)
			}
			log.Println("Closed connection to websocket")
			return
		}
	}

}

func openConnection(token string) (*websocket.Conn, error) {
	serverUrl := url.URL{Scheme: "wss", Host: "ws.finnhub.io", RawQuery: fmt.Sprintf("token=%s", token)}

	fmt.Println(serverUrl.String())

	c, responseCode, err := websocket.DefaultDialer.Dial(serverUrl.String(), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s: Error opening connection to %s; {%s}", responseCode, serverUrl.String(), err))
		panic(fmt.Sprintf("%s: Error opening connection to %s; {%s}", responseCode, serverUrl.String(), err))
	}

	return c, nil
}

func readMessages(outGoingMessages chan string, connection *websocket.Conn) {
	for {
		// fmt.Println("Reading messages")
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from connection. Connection may have been closed: {%s}\n", err)
			return
		}
		// fmt.Printf("Sending message on outbound channel %s\n", string(message))
		outGoingMessages <- string(message)
		// fmt.Println("Message sent")
	}
}

func makeSubscriptionString(symbol string) []byte {
	msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": symbol})
	return msg
}
