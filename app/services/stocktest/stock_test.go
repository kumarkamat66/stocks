package stocktest

import (
	"stocks/app/services/exchange"
	"stocks/app/services/listener"
	"testing"
)

func TestGetStockInfo(t *testing.T) {
	//Starting the Http server on Port 8082
	go listener.Listen()

	//Case -1
	//Test case for getting the Stock details with symbol
	statusCode, resp, _ := exchange.DoHTTPRequest("http://127.0.0.1:8082/stock/AAPL")
	if statusCode != 200 || len(resp) == 0 {
		t.Error("Expected result for getting AAPL stock info failed")
	}

	//Case -2
	//Test case for getting the Stock details with symbol and stock exchange
	statusCode, resp, _ = exchange.DoHTTPRequest("http://127.0.0.1:8082/stock/AAPL?stock_exchange=NYSE")
	if statusCode != 200 || len(resp) == 0 {
		t.Error("Expected result for getting AAPL stock info failed")
	}

	//Case -3
	//Test case for getting the list of symbols under the AMEX by default
	statusCode, resp, _ = exchange.DoHTTPRequest("http://127.0.0.1:8082/stock/")
	if statusCode != 200 || len(resp) == 0 {
		t.Error("Expected result for getting symbols under AMEX stock exchange failed")
	}

	//Case -4
	//Test case for getting the list of symbols under the AMEX by default, with page and limit
	statusCode, resp, _ = exchange.DoHTTPRequest("http://127.0.0.1:8082/stock/?page=1&limit=50")
	if statusCode != 200 || len(resp) == 0 {
		t.Error("Expected result for getting lsit if symbols with paging failed")
	}

}
