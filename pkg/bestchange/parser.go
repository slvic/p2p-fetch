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

const (
	endpointTemplate  = `%s-to-%s.html`
	contentTableClass = `content_table`
	tableBodyTag      = `tbody`
)

type Bestchange struct {
	config configs.Bestchange
	body   string
}

func (b Bestchange) GetData(give, get string) error {
	var bestchangeTable models.BestchangeRow

	responceBytes, err := b.sendRequest(give, get)
	if err != nil {
		return fmt.Errorf("could not send request: %s", err.Error())
	}

	document, err := html.Parse(strings.NewReader(string(responceBytes)))
	if err != nil {
		return fmt.Errorf("could not parse html document: %w", err)
	}
	contentTableNode, err := GetNodeByAttrKey(document, "id", contentTableClass)
	if err != nil {
		return fmt.Errorf("could not get node by attribute key: %w", err)
	}
	contentTableBodyNode, err := GetNodeByTag(contentTableNode, tableBodyTag)
	if err != nil {
		return fmt.Errorf("could not get node by tag: %w", err)
	}
	tableRowNodes, err := GetTableRowNodes(contentTableBodyNode)
	if err != nil {
		return fmt.Errorf("could not get table row nodes: %w", err)
	}

	for _, tableRowNode := range tableRowNodes {
		//GetNodeByAttrKey(exchangerAttribute)
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
