# trading-pairs

## How to run 
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

## Contact 
@amanjain97 or amanjain5221@gmaail.com or create an issue on the repo.