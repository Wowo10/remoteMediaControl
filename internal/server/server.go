package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port          int
	clientCounter int
	connections   sync.Map
}

func NewServer() (*http.Server, func()) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	NewServer := &Server{
		port:          port,
		clientCounter: 1,
	}

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("Running on: %s:%d\n", localAddr.IP.String(), port)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, NewServer.Dispose
}

func (s *Server) Dispose() {
	s.connections.Range(func(key, value any) bool {
		conn := key.(*websocket.Conn)
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Server shutting down"))

		// Wait for the connection to close
		time.Sleep(1 * time.Second)

		conn.Close()
		return true
	})
}
