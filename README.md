# crypto-telegram-notifyer
Server that runs and checks the price of configured coins.


## Overview

![Architecture Overview](/docs/beego_telegram_coins-medium.jpg?raw=true "Architecture Overview")


## How to install locally


**Install required dependencies**

**Required:** first install BeeGo and Bee client


```
go get github.com/astaxie/beego
go get github.com/go-sql-driver/mysql
go get -u github.com/superoo7/go-gecko
```

## Environment variables

**RECREATE_DB**: indicates if the DB must be recreated when running or code changes
- default value: false
- example: RECREATE_DB=true


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

##### GET /alarms

Get all alarms configured on DB

```
curl localhost:8080/alarms
```

Example of response:

```
[
  {
    "id": 1,
    "name": "bitisi",
    "type": "UP",
    "against": "btc",
    "price": 1232
  },
  {
    "id": 2,
    "name": "Amigo",
    "type": "BELOW",
    "against": "btc",
    "price": 213
  },
  {
    "id": 3,
    "name": "Another test",
    "type": "BELOW",
    "against": "btc",
    "price": 213
  }
]
```


##### POST /alarms

Save a new configured alarm on DB

```
curl -X POST -H 'Content-Type: application/json' -d '{ "name": "lisk", "type": "BELOW", "against": "btc", "price": 213 }' localhost:8080/alarms
```

Example of response:

```
[
  // TODO
]
```


##### POST /actions

Perform all stored alarms checks

```
curl -X POST -H 'Content-Type: application/json' localhost:8080/actions
```


No response content provided. 201 http status code as response.

#### Coingecko API

Check the API Reference: https://www.coingecko.com/es/api#explore-api


## Create DB schema

1. First create the executable to intereact with DB:

```
go build main.go
```

2. Check the generated executable:

```
./main orm syncdb
```


## Generate API documentation

Pending of being implemented! Follow the tutorial: https://beego.me/docs/advantage/docs.md

```
bee run -downdoc=true -gendoc=true
```

Access then to: http://localhost:8080/swagger