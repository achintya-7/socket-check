package main

import (
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	router := gin.New()

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "message", func (s socketio.Conn, msg string) error {
		log.Println(msg)
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Pong\n")
	})
	router.GET("/socket/*any", gin.WrapH(server))
	router.POST("/socket/*any", gin.WrapH(server))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
