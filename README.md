# crypto-telegram-notifyer
Server that runs and checks the price of configured coins.


## Overview

![Architecture Overview](/docs/beego_telegram_coins-medium.jpg?raw=true "Architecture Overview")


## How to install locally


**Install required dependencies**

**Required:** first install BeeGo and Bee client


```
go get github.com/astaxie/beego
go get -u github.com/superoo7/go-gecko
```


#### Build & Run

To build and run the project, execute:

```
$ bee run
```


#### Server API

These are the Endpoints exposed on this server

##### GET /coins?symbols=<list_of_symbols>

**Example**  

```
curl localhost:8080/coins?symbols=lisk,tron,bitcoin
```

Symbols must be filled with the name of the coins. In future versions this will be replaced with real symbols of the coins.

Example of response:

```
[
  {
    "symbol": "LISK",
    "usd_price": "1.100000",
    "btc_price": "0.000121"
  },
  {
    "symbol": "BITCOIN",
    "usd_price": "9068.419922",
    "btc_price": "1.000000"
  },
  {
    "symbol": "TRON",
    "usd_price": "0.014237",
    "btc_price": "0.000002"
  }
]
```

#### Coingecko API

Check the API Reference: https://www.coingecko.com/es/api#explore-api