package models

type RawCurrency struct {
	Id   int
	Name string
}

type RawExchanger struct {
	Id   int
	Name string
}

type RawExchangeRate struct {
	SourceCurrencyId      int
	TargetCurrencyId      int
	ExchangersId          int
	GiveRate              float64
	GetRate               float64
	TargetCurrencyReserve float64
	GoodReviewsCount      int
	BadReviewsCount       int
}

type ExchangeRate struct {
	SourceCurrency        string
	TargetCurrency        string
	ExchangerName         string
	GiveRate              float64
	GetRate               float64
	TargetCurrencyReserve float64
	GoodReviewsCount      int
	BadReviewsCount       int
}
