package listener

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"stocks/app/services/exchange"
)

//Listens - It will listen the 8080 port for incoming HTTP Requests.
func Listen() {
	//Registering the handler on top of the request path /stock/symbol
	http.HandleFunc("/stock/", func(writer http.ResponseWriter, request *http.Request) {
		//holds the symbol taken from the client request
		symbol := strings.TrimPrefix(request.URL.Path, "/stock/")
		//this call returns the info as stock data and error if any
		exchengName, page, limit := getParams(request)
		info, err := exchange.GetStockInfo(symbol, exchengName, page, limit)
		if err != nil {
			serveResponse(prepareErrorResp(err.Error()), writer)
		}
		//Serves the response to the client
		serveResponse(info, writer)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}

//This will serves the response back to the client
func serveResponse(info []byte, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(info)
	return
}

//reads the exchange from the request,
//no exchange parameter sent from client default will be 'AMEX'
//returns exchange -
//returns page
//returns limit
func getParams(request *http.Request) (string, int, int) {
	//Default Exchane
	stockExchange := "AMEX"
	page := 1
	if len(request.URL.Query()["page"]) > 0 {
		page, _ = strconv.Atoi(request.URL.Query()["page"][0])
	}
	limit := 50
	if len(request.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(request.URL.Query()["limit"][0])
	}
	params := request.URL.Query()["stock_exchange"]
	if len(params) > 0 {
		stockExchange = params[0]
	}
	return stockExchange, page, limit
}

//prepares the error response using the error string
func prepareErrorResp(message string) []byte {
	resp, _ := json.Marshal(map[string]interface{}{"Message": message})
	return resp
}
