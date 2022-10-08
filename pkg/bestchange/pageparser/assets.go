package pageparser

import (
	"fmt"
	"io"
	"strings"

	"github.com/slvic/stock-observer/pkg/bestchange/models"
	"golang.org/x/net/html"
)

const (
	linkAttr = `href`

	assetSeparator = `to`
	giveBorder     = `.com/`
	getBorder      = `.html`
)

func ParseBestchangeAssetsRow(page string) (models.ExchangePair, error) {
	var exchangePair models.ExchangePair
	reader := strings.NewReader(page)
	tokenizer := html.NewTokenizer(reader)
	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return models.ExchangePair{}, nil
			}
			return models.ExchangePair{}, fmt.Errorf("error token found: %s", tokenizer.Err().Error())
		}
		_, hasAttr := tokenizer.TagName()

		attrKeyBytes, attrValueBytes, moreAttr := tokenizer.TagAttr()
		attrKey := string(attrKeyBytes)
		attrValue := string(attrValueBytes)

		if hasAttr {
			for {
				if attrKey == linkAttr {
					assetSeparatorIndex := strings.Index(attrValue, assetSeparator)
					giveBorderIndex := strings.Index(attrValue, giveBorder)
					getBorderIndex := strings.Index(attrValue, getBorder)

					give := attrValue[giveBorderIndex+len(giveBorder) : assetSeparatorIndex-1]
					get := attrValue[assetSeparatorIndex+len(assetSeparator)+1 : getBorderIndex]

					exchangePair.Give = give
					exchangePair.Get = get

					return exchangePair, nil
				}

				if !moreAttr {
					break
				}
			}
			return models.ExchangePair{}, fmt.Errorf("could not find link attribute in html element")
		}
		return models.ExchangePair{}, fmt.Errorf("html element has no tag")
	}
}
