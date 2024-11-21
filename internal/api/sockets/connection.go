package sockets

import (
	"aie/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type UserConnection struct {
	User       string
	Connection *websocket.Conn
}

var ongoingCalls = make(map[string][]UserConnection)

// WebSocket upgrader to upgrade HTTP requests to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func createNewConnection(c *gin.Context, groupId string) *websocket.Conn {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade connection: %v\n", err)
		return nil
	}
	fmt.Println("creating new connection in group :", groupId)
	return conn
}

func Connect(c *gin.Context) {
	// Upgrade the HTTP request to a WebSocket connection
	// socketId := c.Param("socketId")
	groupId := c.Param("groupId")
	userId := utils.GetUserId(c)

	newConnction := createNewConnection(c, groupId)

	userConnectionObject := UserConnection{
		User:       userId,
		Connection: newConnction,
	}

	ongoingCalls[groupId] = append(ongoingCalls[groupId], userConnectionObject)

	defer func() {
		fmt.Println("starting connection close")
		for index, connection := range ongoingCalls[groupId] {
			if connection.User == userId {
				modifiedCalls := append(ongoingCalls[groupId][:index], ongoingCalls[groupId][index+1:]...)
				ongoingCalls[groupId] = modifiedCalls
				newConnction.Close()
				break
			}
		}
	}()

	for {
		// Read message from the client
		messageType, msg, err := newConnction.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		for _, conn := range ongoingCalls[groupId] {
			// Echo the message back to the client
			err = conn.Connection.WriteMessage(messageType, msg)
			if err != nil {
				fmt.Printf("Error writing message: %v\n", err)
				break
			}
		}
	}
}
