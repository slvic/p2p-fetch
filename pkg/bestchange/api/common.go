package api

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/slvic/p2p-fetch/pkg/bestchange/models"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	bcApiZipFileName = `bestChange.zip`
	bcApiFolder      = `bestChange`
)

const (
	currenciesFile       = `bestChange/bm_cy.dat`
	exchangerOfficesFile = `bestChange/bm_exch.dat`
	exchangeRatesFile    = `bestChange/bm_rates.dat`

	dataSeparator = `;`
)

func getBcApiFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not get bc api file: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong responce status code: %d", resp.StatusCode)
	}

	out, err := os.Create(bcApiZipFileName)
	if err != nil {
		return fmt.Errorf("could not create zip file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("could not write responce body to a zip file: %w", err)
	}

	return nil
}

func unzipSource(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return fmt.Errorf("could not open zip reader: %w", err)
	}
	defer reader.Close()

	absPath, err := filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("could not get absolute path to a destination folder: %w", err)
	}

	for _, file := range reader.File {
		err = unzipFile(file, absPath)
		if err != nil {
			return fmt.Errorf("could not unzip file %s: %w", file.Name, err)
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return fmt.Errorf("could not create a directory: %w", err)
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("clould not create a file: %w", err)
	}

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return fmt.Errorf("could not open a file: %w", err)
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return fmt.Errorf("could not open a file: %w", err)
	}
	defer zippedFile.Close()

	if _, err = io.Copy(destinationFile, zippedFile); err != nil {
		return fmt.Errorf("could not copy zipped file to a destination file: %w", err)
	}

	return nil
}

func getExchangeRates(
	rawExchangeRates []models.RawExchangeRate,
	exchangers map[int]string,
	currencies map[int]string,
) []models.ExchangeRate {

	var exchangeRates []models.ExchangeRate

	for _, rawExchangeRate := range rawExchangeRates {
		var exchangeRate models.ExchangeRate

		exchangeRate.SourceCurrency = currencies[rawExchangeRate.SourceCurrencyId]
		exchangeRate.TargetCurrency = currencies[rawExchangeRate.TargetCurrencyId]
		exchangeRate.ExchangerName = exchangers[rawExchangeRate.ExchangersId]
		exchangeRate.GiveRate = rawExchangeRate.GiveRate
		exchangeRate.GetRate = rawExchangeRate.GetRate
		exchangeRate.TargetCurrencyReserve = rawExchangeRate.TargetCurrencyReserve
		exchangeRate.GoodReviewsCount = rawExchangeRate.GoodReviewsCount
		exchangeRate.BadReviewsCount = rawExchangeRate.BadReviewsCount

		exchangeRates = append(exchangeRates, exchangeRate)
	}

	return exchangeRates
}

func getRawCurrencies(fileName string) (map[int]string, error) {
	currencies := make(map[int]string)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open a file: %w", err)
	}
	defer file.Close()
	encodedReader := transform.NewReader(file, charmap.Windows1251.NewDecoder())

	scanner := bufio.NewScanner(encodedReader)
	for scanner.Scan() {
		var currency models.RawCurrency

		currencyData := strings.Split(scanner.Text(), dataSeparator)

		currency.Id, err = strconv.Atoi(currencyData[0])
		if err != nil {
			return nil, fmt.Errorf("could not convert currency id string to integer: %w", err)
		}
		currency.Name = currencyData[2]

		currencies[currency.Id] = currency.Name
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("a scanner error: %w", err)
	}

	return currencies, nil
}

func getRawExchangers(fileName string) (map[int]string, error) {
	exchangers := make(map[int]string)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open a file: %w", err)
	}
	defer file.Close()
	encodedReader := transform.NewReader(file, charmap.Windows1251.NewDecoder())

	scanner := bufio.NewScanner(encodedReader)

	for scanner.Scan() {
		var exchanger models.RawExchanger

		currencyData := strings.Split(scanner.Text(), dataSeparator)

		exchanger.Id, err = strconv.Atoi(currencyData[0])
		if err != nil {
			return nil, fmt.Errorf("could not convert exchanger id string to integer: %w", err)
		}
		exchanger.Name = currencyData[1]

		exchangers[exchanger.Id] = exchanger.Name
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("a scanner error: %w", err)
	}

	return exchangers, nil
}

func getRawExchangeRates(fileName string) ([]models.RawExchangeRate, error) {
	var exchangeRates []models.RawExchangeRate

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open a file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var exchangeRate models.RawExchangeRate

		currencyData := strings.Split(scanner.Text(), dataSeparator)

		exchangeRate.SourceCurrencyId, err = strconv.Atoi(currencyData[0])
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange rate source currency id string to integer: %w", err)
		}

		exchangeRate.TargetCurrencyId, err = strconv.Atoi(currencyData[1])
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange rate target source id string to integer: %w", err)
		}

		exchangeRate.ExchangersId, err = strconv.Atoi(currencyData[2])
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange rate exchanger's id string to integer: %w", err)
		}

		exchangeRate.GiveRate, err = strconv.ParseFloat(currencyData[3], 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange give rate string to integer: %w", err)
		}

		exchangeRate.GetRate, err = strconv.ParseFloat(currencyData[4], 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange get rate string to integer: %w", err)
		}

		exchangeRate.TargetCurrencyReserve, err = strconv.ParseFloat(currencyData[5], 64)
		if err != nil {
			return nil, fmt.Errorf("could not convert exchange rate target currency reserve string to integer: %w", err)
		}

		reviews := strings.Split(currencyData[6], ".")
		switch len(reviews) {
		case 1:
			exchangeRate.GoodReviewsCount, err = strconv.Atoi(reviews[0])
			if err != nil {
				return nil, fmt.Errorf("could not convert exchange rate revievs count string to integer: %w", err)
			}
		case 2:
			exchangeRate.GoodReviewsCount, err = strconv.Atoi(reviews[1])
			if err != nil {
				return nil, fmt.Errorf("could not convert exchange rate revievs count string to integer: %w", err)
			}
			exchangeRate.BadReviewsCount, err = strconv.Atoi(reviews[0])
			if err != nil {
				return nil, fmt.Errorf("could not convert exchange rate revievs count string to integer: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported reviews count format, there are %d review types", len(reviews))
		}

		exchangeRates = append(exchangeRates, exchangeRate)
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("a scanner error: %w", err)
	}

	return exchangeRates, nil
}
