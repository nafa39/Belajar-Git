package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func handlerWebSocket(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	defer ws.Close()

	clients[ws] = true

	fmt.Println("Client connected")

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			break

		}

		fmt.Printf("Message received: %s\n", msg)

		msg = []byte("you said: " + string(msg))

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
				break
			}
		}
	}
	return nil
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.GET("/ws", handlerWebSocket)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
