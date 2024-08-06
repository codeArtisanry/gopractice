package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received: %s", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"fmt"

// 	"github.com/gorilla/websocket"
// 	"github.com/labstack/echo/v4"
// )

// var (
// 	upgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 	}
// )

// func hello(c echo.Context) error {
// 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading to websocket", err)
// 		return err
// 	}
// 	defer ws.Close()
// 	fmt.Println("====")

// 	for {
// 		// Read message from browser
// 		msgType, msg, err := ws.ReadMessage()
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}

// 		// Print the message to the console
// 		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

// 		// Write message back to browser
// 		if err = ws.WriteMessage(msgType, msg); err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}
// 	}
// }

// func main() {
// 	e := echo.New()
// 	e.Static("/", ".")
// 	e.GET("/ws", hello)

// 	e.Logger.Fatal(e.Start(":8080"))
// }
