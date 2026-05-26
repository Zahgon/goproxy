package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/coder/websocket"
)

const (
	echoAddr  = ":12345"
	proxyAddr = ":54321"
	proxyURL  = "http://localhost:54321"
	echoURL   = "wss://localhost:12345"
)

func echo(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func startEchoServer() { _ = "STUB: not implemented"; return }

func startProxy() { _ = "STUB: not implemented"; return }

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	startEchoServer()
	startProxy()

	parsedProxy, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatal("unable to parse proxy URL:", err)
	}

	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, echoURL, &websocket.DialOptions{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				Proxy: http.ProxyURL(parsedProxy),
			},
		},
		Subprotocols: []string{"p1"},
	})
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close(websocket.StatusNormalClosure, "")

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.Read(ctx)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C: // Message send
			// Write current time to the websocket client every 1 second
			if err := c.Write(ctx, websocket.MessageText, []byte(t.String())); err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt: // Server shutdown
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.Close(websocket.StatusNormalClosure, "")
			if err != nil {
				log.Println("write close:", err)
				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
