app {
  fetchInterval = 4
}

binance {
  address = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

  assets = [
             "USDT",
              "BTC",
              "BNB",
             "BUSD",
              "ETH",
              "DAI"
  ]

  fiats = [
            "ARS",
            "EUR",
            "USD",
            "AED",
            "AUD",
            "BDT",
            "BHD",
            "BOB",
            "BRL",
            "CAD",
            "CLP",
            "CNY",
            "COP",
            "CRC",
            "CZK",
            "DOP",
            "DZD",
            "EGP",
            "GBP",
            "GEL",
            "GHS",
            "HKD",
            "IDR",
            "INR",
            "JPY",
            "KES",
            "KHR",
            "KRW",
            "KWD",
            "KZT",
            "LAK",
            "LBP",
            "LKR",
            "MAD",
            "MMK",
            "MXN",
            "MYR",
            "NGN",
            "OMR",
            "PAB",
            "PEN",
            "PHP",
            "PKR",
            "PLN",
            "PYG",
            "QAR",
            "RON",
            "RUB",
            "SAR",
            "SDG",
            "SEK",
            "SGD",
            "THB",
            "TND",
            "TRY",
            "TWD",
            "UAH",
            "UGX",
            "UYU",
            "VES",
            "VND",
            "ZAR"
  ]
}

bestchange {
  baseurl = "https://www.bestchange.com/"
  apiurl = "http://api.bestchange.ru/info.zip"
}
