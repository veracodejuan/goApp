package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	// Example 1: Using the html package from golang.org/x/net
	htmlContent := `<html><body><h1>Hello, Go!</h1></body></html>`
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}
	printNode(doc)

	// Example 2: Using the websocket package
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Received message: %s\n", msg)
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}
	})
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Example 3: Using Prometheus metrics
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "example_counter",
		Help: "An example of a counter in Prometheus",
	})
	prometheus.MustRegister(counter)
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Fatal(http.ListenAndServe(":9090", nil))
	}()

	// Example 4: Using Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Could not create Docker client: %s", err)
	}
	info, err := cli.Info(ctx)
	if err != nil {
		log.Fatalf("Could not retrieve Docker info: %s", err)
	}
	fmt.Printf("Docker info: %v\n", info)
}

func printNode(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("<%s>\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printNode(c)
	}
	if n.Type == html.ElementNode {
		fmt.Printf("</%s>\n", n.Data)
	}
}
