app {
  fetchInterval = 4
}

binance {
  address = "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"

  assets = [
        "USDT",
        "BTC",
        "BUSD",
        "BNB",
        "ETH",
        "RUB",
        "SHIB"
  ]

  fiats = [
        "USDT",
        "BTC",
        "BUSD",
        "BNB",
        "ETH",
        "RUB",
        "SHIB"
  ]
}

bestchange {
  baseurl = "https://www.bestchange.com/"
  apiurl = "http://api.bestchange.com/info.zip"
}
