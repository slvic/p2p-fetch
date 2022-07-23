package bestchange

import (
	"fmt"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/bestchange/models"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

const endpointTemplate = `%s-to-%s.html`

type Bestchange struct {
	config configs.Bestchange
	body   string
}

func (b Bestchange) GetData(give, get string) error {
	var bestchangeTable models.BestchangeTable

	responceBytes, err := b.sendRequest(give, get)
	if err != nil {
		return fmt.Errorf("could not send request: %s", err.Error())
	}

	parseBestchangeTable(string(responceBytes))
}

func getAttribute(n *html.Node, key string) (string, bool) {

	for _, attr := range n.Attr {

		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func parseBestchangeTable(page string) (models.BestchangeTable, error) {
	tokenizer := html.NewTokenizer(strings.NewReader(page))

	for {
		nextToken := tokenizer.Next()
		switch {
		case nextToken == html.ErrorToken:
			return models.BestchangeTable{}, fmt.Errorf("tokenizer returned error token: %s", nextToken.String())
		case nextToken == html.StartTagToken:
			token := nextToken.String()
			if token
		}
	}
}

func (b Bestchange) sendRequest(give, get string) ([]byte, error) {
	url := fmt.Sprintf(b.config.BaseUrl+endpointTemplate, give, get)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get responce: %s", err.Error())
	}

	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read responce body: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessfull request, status code %d, response body: %s",
			response.StatusCode,
			string(responseBodyBytes))
	}

	return responseBodyBytes, nil
}
