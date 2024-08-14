package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	// Example HTML content
	htmlContent := `<html><body><h1>Hello, Go!</h1></body></html>`
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}
	printNode(doc)
}

// printNode recursively prints the node tree.
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
