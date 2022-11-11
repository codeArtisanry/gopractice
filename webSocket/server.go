package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket", err)
		return err
	}
	defer ws.Close()
	fmt.Println("====")

	for {
		// Read message from browser
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = ws.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return nil
		}
	}
}

func main() {
	e := echo.New()
	e.Static("/", ".")
	e.GET("/ws", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
