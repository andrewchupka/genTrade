package integrations

import (
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
	var body string = makeSubscriptionString(symbol)
	err = connection.WriteMessage(websocket.TextMessage, []byte(body))
	if err != nil {
		log.Fatal("handleWebSocketConnection:", err)
		panic(fmt.Sprintf("Error writing on connection: {message: %s}; {error: %s}", body, err))
	}

	messageStream := make(chan string)

	go readMessages(messageStream, connection)

	for {
		select {
		case message := <- messageStream:
			outGoingMessages <- message
			log.Println(message)
			return
		case <- interrupt:
			err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("Error closing connection: {%s}\n", err)
			}
			return 
		}
	}

}

func openConnection(token string) (*websocket.Conn, error) {
	serverUrl := url.URL{Scheme: "ws", Host: "ws.finnhub.io", RawQuery: fmt.Sprintf("token=%s", token) }

	fmt.Println(serverUrl.String())

	c, _, err := websocket.DefaultDialer.Dial(serverUrl.String(), nil)
	if err != nil {
		log.Fatal("openWebSocketConnection: ", err)
		panic(fmt.Sprintf("Error opening connection to %s; {%s}", serverUrl.String(), err))
	}

	return c, nil
}

func readMessages(outGoingMessages chan string, connection *websocket.Conn) {
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from connection: {%s}\n", err)
			return
		}
		outGoingMessages <- string(message)
	}

}

func makeSubscriptionString(symbol string) string {
	return fmt.Sprintf("{'type': 'subscribe-news', 'symbol': '%s'", symbol)
}

