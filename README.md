# trading-pairs

## How to run 
It requires Go 1.18
1. `go mod download`
2. `go run main.go` 

Server will be started at port 5000. 

## API 

### Binance

Request 
`GET /api/v1/binance/trading-pairs`

Response will be filename as a string containing base currency and quote currency pairs separated by "/" in each line.

eg. `tradingpairs-repo/binance-pairs.txt`

### Coinbase

Request 
`GET /api/v1/coinbase/trading-pairs`

Response will be filename as a string containing base currency and quote currency pairs separated by "/" in each line.

eg. `tradingpairs-repo/coinbase-pairs.txt`

### Mockgen
```bash
mockgen -source=app/repository/trading_pairs.go -destination=app/repository/mocks/trading_pairs_mock.go
```

## Contact 
@amanjain97 or amanjain5221@gmail.com or create an issue on the repo.
