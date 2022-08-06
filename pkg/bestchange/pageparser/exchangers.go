package pageparser

import (
	"fmt"
	"github.com/slvic/p2p-fetch/pkg/bestchange/models"
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

const (
	exchangerAttribute   = `bj`
	giveAttribute        = `fs`
	giveCountryAttribute = `ct`
	giveMinimumAttribute = `fm1`
	giveMaximumAttribute = `fm2`
	getAttribute         = `bi`
	reserveAttribute     = `arp`
)

func ParseBestchangeExchangerRow(page string) (models.BestchangeRow, error) {
	var bestchangeRow models.BestchangeRow
	reader := strings.NewReader(page)
	tokenizer := html.NewTokenizer(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return models.BestchangeRow{}, nil
			}
			return models.BestchangeRow{}, fmt.Errorf("error token found: %s", tokenizer.Err().Error())
		}
		_, hasAttr := tokenizer.TagName()

		attrKeyBytes, attrValueBytes, moreAttr := tokenizer.TagAttr()
		attrKey := string(attrKeyBytes)
		attrValue := string(attrValueBytes)

		if hasAttr {
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
				}

				if !moreAttr {
					break
				}
			}
		}
	}
}
