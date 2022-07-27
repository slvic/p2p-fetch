package bestchange

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"io"
)

func GetNodeByAttrKey(doc *html.Node, attrKey, attrVal string) (*html.Node, error) {
	var htmlNode *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == attrKey && attr.Val == attrVal {
					htmlNode = node
					return
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if htmlNode != nil {
		return htmlNode, nil
	}
	return nil, errors.New("could not find a node by attribute key")
}

func GetNodeByTag(doc *html.Node, tagName string) (*html.Node, error) {
	var htmlNode *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tagName {
			htmlNode = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if htmlNode != nil {
		return htmlNode, nil
	}
	return nil, errors.New("could not find a node by tag")
}

func GetTableRowNodes(tbody *html.Node) ([]*html.Node, error) {
	var htmlNodes []*html.Node

	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "tr" {
			htmlNodes = append(htmlNodes, node)
		}
		if node.NextSibling == nil {
			return
		}
		crawler(node.NextSibling)
	}
	crawler(tbody.FirstChild)
	if len(htmlNodes) != 0 {
		return htmlNodes, nil
	}
	return nil, errors.New("could not find table rows")
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
