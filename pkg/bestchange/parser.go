package bestchange

import (
	"fmt"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/bestchange/models"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	startTag             = `StartTag`
	endTag               = `EndTag`
	tableBodyTag         = `tbody`
	contentTableId       = `content_table`
	endpointTemplate     = `%s-to-%s.html`
	exchangerAttribute   = `bj`
	giveAttribute        = `fs`
	giveCountryAttribute = `ct`
	giveMinimumAttribute = `fm1`
	giveMaximumAttribute = `fm2`
	getAttribute         = `bi`
	reserveAttribute     = `ar arp`
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

	parseBestchangeTable(string(responceBytes))
}

func parseBestchangeTable(page string) ([]models.BestchangeRow, error) {
	var bestchangeTable []models.BestchangeRow
	contentTableOrrured := false
	var contentTableCounter int
	reader := strings.NewReader(page)
	tokenizer := html.NewTokenizer(reader)
	html.Parse(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return bestchangeTable, nil
			}
			return nil, fmt.Errorf("error token found: %s", tokenizer.Err().Error())
		}
		_, hasAttr := tokenizer.TagName()

		attrKeyBytes, attrValueBytes, moreAttr := tokenizer.TagAttr()
		attrKey := string(attrKeyBytes)
		attrValue := string(attrValueBytes)

		if !contentTableOrrured {
			if attrKey == contentTableId {
				contentTableOrrured = true
			}
			continue
		}

		if attrKey == tableBodyTag {
			switch tt.String() {
			case startTag:
				contentTableCounter += 1
			case endTag:
				contentTableCounter -= 1
			}
		}

		if contentTableCounter == 0 {
			break
		}

		if hasAttr {
			var bestchangeRow models.BestchangeRow
			for {
				switch attrKey {
				case exchangerAttribute:
					bestchangeRow.Exchanger = attrValue
				case giveAttribute:
					bestchangeRow.Give = attrValue
				case giveCountryAttribute:
					bestchangeRow.GiveCountry = attrValue
				case giveMinimumAttribute:
					bestchangeRow.GiveMin = attrValue
				case giveMaximumAttribute:
					bestchangeRow.GiveMax = attrValue
				case getAttribute:
					if _, err := strconv.ParseFloat(attrValue, 64); err != nil {
						bestchangeRow.Give = attrValue
					}
				case reserveAttribute:
					bestchangeRow.Reserve = attrValue
					bestchangeTable = append(bestchangeTable, bestchangeRow)
				}

				if !moreAttr {
					break
				}
			}
		}
	}
	return bestchangeTable, nil
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
