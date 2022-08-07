package models

type BinanceRequest struct {
	Asset         string   `json:"asset"`
	Fiat          string   `json:"fiat"`
	MerchantCheck bool     `json:"merchantCheck"`
	Page          int32    `json:"page"`
	PayTypes      []string `json:"-"`
	PublisherType *string  `json:"publisherType"`
	Rows          int32    `json:"rows"`
	TradeType     string   `json:"tradeType"`
}

type BinanceResponse struct {
	Code          *string `json:"code"`
	Message       *string `json:"message"`
	MessageDetail *string `json:"messageDetail"`
	Data          []Data  `json:"data"`
}

type Data struct {
	Adv        Adv        `json:"adv"`
	Advertiser Advertiser `json:"advertiser"`
}

type Adv struct {
	AdvNo                           *string        `json:"advNo"`
	Classify                        *string        `json:"classify"`
	TradeType                       *string        `json:"tradeType"`
	Asset                           *string        `json:"asset"`
	FiatUnit                        *string        `json:"fiatUnit"`
	AdvStatus                       *string        `json:"advStatus"`
	PriceType                       *string        `json:"priceType"`
	PriceFloatingRatio              *string        `json:"priceFloatingRatio"`
	RateFloatingRatio               *string        `json:"rateFloatingRatio"`
	CurrencyRate                    *string        `json:"currencyRate"`
	Price                           *string        `json:"price"`
	InitAmount                      *string        `json:"initAmount"`
	SurplusAmount                   *string        `json:"surplusAmount"`
	AmountAfterEditing              *string        `json:"amountAfterEditing"`
	MaxSingleTransAmount            *string        `json:"maxSingleTransAmount"`
	MinSingleTransAmount            *string        `json:"minSingleTransAmount"`
	BuyerKycLimit                   *string        `json:"buyerKycLimit"`
	BuyerRegDaysLimit               *string        `json:"buyerRegDaysLimit"`
	BuyerBtcPositionLimit           *string        `json:"buyerBtcPositionLimit"`
	Remarks                         *string        `json:"remarks"`
	AutoReplyMsg                    *string        `json:"autoReplyMsg"`
	PayTimeLimit                    int32          `json:"payTimeLimit"`
	TradeMethods                    []TradeMethods `json:"tradeMethods"`
	UserTradeCountFilterTime        *string        `json:"userTradeCountFilterTime"`
	UserBuyTradeCountMin            *string        `json:"userBuyTradeCountMin"`
	UserBuyTradeCountMax            *string        `json:"userBuyTradeCountMax"`
	UserSellTradeCountMin           *string        `json:"userSellTradeCountMin"`
	UserSellTradeCountMax           *string        `json:"userSellTradeCountMax"`
	UserAllTradeCountMin            *string        `json:"userAllTradeCountMin"`
	UserAllTradeCountMax            *string        `json:"userAllTradeCountMax"`
	UserTradeCompleteRateFilterTime *string        `json:"userTradeCompleteRateFilterTime"`
	UserTradeCompleteCountMin       *string        `json:"userTradeCompleteCountMin"`
	UserTradeCompleteRateMin        *string        `json:"userTradeCompleteRateMin"`
	UserTradeVolumeFilterTime       *string        `json:"userTradeVolumeFilterTime"`
	UserTradeType                   *string        `json:"userTradeType"`
	UserTradeVolumeMin              *string        `json:"userTradeVolumeMin"`
	UserTradeVolumeMax              *string        `json:"userTradeVolumeMax"`
	UserTradeVolumeAsset            *string        `json:"userTradeVolumeAsset"`
	CreateTime                      *string        `json:"createTime"`
	AdvUpdateTime                   *string        `json:"advUpdateTime"`
	FiatVo                          *string        `json:"fiatVo"`
	AssetVo                         *string        `json:"assetVo"`
	AdvVisibleRet                   *string        `json:"advVisibleRet"`
	AssetLogo                       *string        `json:"assetLogo"`
	AssetScale                      int32          `json:"assetScale"`
	FiatScale                       int32          `json:"fiatScale"`
	PriceScale                      int32          `json:"priceScale"`
	FiatSymbol                      *string        `json:"fiatSymbol"`
	IsTradable                      bool           `json:"isTradable"`
	DynamicMaxSingleTransAmount     *string        `json:"dynamicMaxSingleTransAmount"`
	MinSingleTransQuantity          *string        `json:"minSingleTransQuantity"`
	MaxSingleTransQuantity          *string        `json:"maxSingleTransQuantity"`
	DynamicMaxSingleTransQuantity   *string        `json:"dynamicMaxSingleTransQuantity"`
	TradableQuantity                *string        `json:"tradableQuantity"`
	CommissionRate                  *string        `json:"commissionRate"`
	TradeMethodCommissionRates      []interface{}  `json:"tradeMethodCommissionRates"`
	LaunchCountry                   *string        `json:"launchCountry"`
	AbnormalStatusList              *string        `json:"abnormalStatusList"`
}

type TradeMethods struct {
	PayId                *string `json:"payId"`
	PayMethodId          *string `json:"payMethodId"`
	PayType              *string `json:"payType"`
	PayAccount           *string `json:"payAccount"`
	PayBank              *string `json:"payBank"`
	PaySubBank           *string `json:"paySubBank"`
	Identifier           *string `json:"identifier"`
	IconUrlColor         *string `json:"iconUrlColor"`
	TradeMethodName      *string `json:"tradeMethodName"`
	TradeMethodShortName *string `json:"tradeMethodShortName"`
	TradeMethodBgColor   *string `json:"tradeMethodBgColor"`
}

type Advertiser struct {
	UserNo           *string       `json:"userNo"`
	RealName         *string       `json:"realName"`
	NickName         *string       `json:"nickName"`
	Margin           *string       `json:"margin"`
	MarginUnit       *string       `json:"marginUnit"`
	OrderCount       *string       `json:"orderCount"`
	MonthOrderCount  *int          `json:"monthOrderCount"`
	MonthFinishRate  *float32      `json:"monthFinishRate"`
	AdvConfirmTime   *string       `json:"advConfirmTime"`
	Email            *string       `json:"email"`
	RegistrationTime *string       `json:"registrationTime"`
	Mobile           *string       `json:"mobile"`
	UserType         *string       `json:"userType"`
	TagIconUrls      []interface{} `json:"tagIconUrls"`
	UserGrade        int32         `json:"userGrade"`
	UserIdentity     *string       `json:"userIdentity"`
	ProMerchant      *bool         `json:"proMerchant"`
	IsBlocked        *string       `json:"isBlocked"`
}
