package models

type BestchangeRow struct {
	Exchanger   string
	Give        string
	GiveCountry string
	GiveMin     string
	GiveMax     string
	Get         string
	Reserve     string
	//Reviews    string
}

type ExchangePair struct {
	Give string
	Get  string
}
