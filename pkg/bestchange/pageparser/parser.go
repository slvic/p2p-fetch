package pageparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slvic/stock-observer/internal/configs"
	"github.com/slvic/stock-observer/pkg/bestchange/models"
	"golang.org/x/net/html"
)

const (
	endpointTemplate  = `%s-to-%s.html`
	exchangersTableId = `content_table`
	assetsTableId     = `curr_tab_c`
	assetsLinkClass   = `lc`
	tableBodyTag      = `tbody`
)

type BestchangePageParser struct {
	config configs.Bestchange
}

func NewBestchangeParser(cfg configs.Bestchange) *BestchangePageParser {
	return &BestchangePageParser{config: cfg}
}

func (b BestchangePageParser) GetAssets() ([]models.ExchangePair, error) {
	var exchangePairs []models.ExchangePair

	rawPage, err := b.getRawPage()
	if err != nil {
		return nil, fmt.Errorf("could not get raw page: %s", err.Error())
	}

	assetsTableNode, err := GetNodeByAttrKey(rawPage, "id", assetsTableId)
	if err != nil {
		return nil, fmt.Errorf("could not get assets table node by attribute key: %w", err)
	}
	assetsTableBodyNode, err := GetNodeByTag(assetsTableNode, tableBodyTag)
	if err != nil {
		return nil, fmt.Errorf("could not get assets table body node by tag: %w", err)
	}

	assetsTableRowNodes, err := GetTableRowNodes(assetsTableBodyNode)
	if err != nil {
		return nil, fmt.Errorf("could not get assets table row nodes: %w", err)
	}

	for _, tableNodeRow := range assetsTableRowNodes {
		linkElement, err := GetNodeByAttrKey(tableNodeRow, "class", assetsLinkClass)
		if err != nil {
			return nil, fmt.Errorf("could not get assent's link: %w", err)
		}
		exchangePair, err := ParseBestchangeAssetsRow(RenderNode(linkElement))
		if err != nil {
			log.Printf("could not parse bestchange assets row %s \n\r error: %s", RenderNode(linkElement), err.Error())
			continue
		}
		exchangePairs = append(exchangePairs, exchangePair)
	}

	return exchangePairs, nil
}

func (b BestchangePageParser) GetExchangers(exchanges []models.ExchangePair) error {
	for _, exchange := range exchanges {
		err := b.getExchangersByPair(exchange)
		if err != nil {
			log.Printf(
				"could not get exchange by pair give: %s, get: %s \n\n Error: %s",
				exchange.Give,
				exchange.Get,
				err.Error(),
			)
			continue
		}
	}

	return nil
}

func (b BestchangePageParser) getExchangersByPair(exchange models.ExchangePair) error {
	var bestchangeTable []models.BestchangeRow

	rawExchangers, err := b.getRawExchangers(exchange)
	if err != nil {
		return fmt.Errorf("could not raw exchangers: %s", err.Error())
	}

	exchangersTableNode, err := GetNodeByAttrKey(rawExchangers, "id", exchangersTableId)
	if err != nil {
		return fmt.Errorf("could not get exchangers table node by attribute key: %w", err)
	}
	exchangersTableBodyNode, err := GetNodeByTag(exchangersTableNode, tableBodyTag)
	if err != nil {
		return fmt.Errorf("could not get exchangers table body node by tag: %w", err)
	}
	exchangersTableRowNodes, err := GetTableRowNodes(exchangersTableBodyNode)
	if err != nil {
		return fmt.Errorf("could not get exchangers table row nodes: %w", err)
	}

	for _, tableRowNode := range exchangersTableRowNodes {
		row, err := ParseBestchangeExchangerRow(RenderNode(tableRowNode))
		if err != nil {
			log.Printf("could not parse bestchange exchanger row %s \n\r error: %s", RenderNode(tableRowNode), err.Error())
			continue
		}
		bestchangeTable = append(bestchangeTable, row)
	}
	return nil
}

func (b BestchangePageParser) getRawExchangers(exchange models.ExchangePair) (*html.Node, error) {
	url := fmt.Sprintf(b.config.BaseUrl+endpointTemplate, exchange.Give, exchange.Get)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get responce: %s", err.Error())
	}

	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read responce body: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unsuccessfull request, status code %d, response body: %s",
			response.StatusCode,
			string(responseBodyBytes))
	}

	document, err := html.Parse(strings.NewReader(string(responseBodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("could not parse raw exchangers: %w", err)
	}

	return document, nil
}

func (b BestchangePageParser) getRawPage() (*html.Node, error) {
	response, err := http.Get(b.config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("could not get responce: %s", err)
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

	document, err := html.Parse(strings.NewReader(string(responseBodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("could not parse raw page: %w", err)
	}

	return document, nil
}
