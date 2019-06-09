package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//@param API Key for authentication
//@param URL to fetch the trade info
const (
	Apikey    = "7qPl7fYevNrFsIDBJo8fObZ0C3hCRoIFqsXD4YbccH9SQ37ETKKKHiTZdGaL"
	URL       = "https://www.worldtradingdata.com/api/v1/stock?stock_exchange=%s&api_token=%s&symbol=%s"
	SearchURL = "https://api.worldtradingdata.com/api/v1/stock_search?stock_exchange=%s&api_token=%s&page=%d&limit=%d"
)

//GetStockInfo - this will take symbol, exchangeName as arguemnts
//returns - response as []byte and
// return error if any
func GetStockInfo(symbol string, exchangeName string, page int, limit int) ([]byte, error) {
	//handling the optional parameter symbol, if it is empty then it will ignore the param
	// else we will use the search Url to fetch the data
	urlWithParams := ""
	if len(symbol) > 0 {
		urlWithParams = fmt.Sprintf(URL, exchangeName, Apikey, symbol)
	} else {
		urlWithParams = fmt.Sprintf(SearchURL, exchangeName, Apikey, page, limit)
	}
	//Requesting the data through API call
	statusCode, data, err := DoHTTPRequest(urlWithParams)
	if err != nil {
		log.Println("Error while requesting the data from Exchange with response code ", statusCode, err)
		return nil, err
	}
	// converting data in to byte to serve response back
	responseData, err := json.Marshal(stringToMap(data))
	if err != nil {
		log.Println("Error while parsing the response from stock exchange ", err.Error())
	}
	return responseData, err
}

//doHTTPRequest - It requests the resource from the specified endpoint URL
func DoHTTPRequest(url string) (int, string, error) {
	client := &http.Client{}
	dataInBytes := bytes.NewBufferString("")
	req, err := http.NewRequest("GET", url, dataInBytes)
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), err
}

//converts data from string to Map
func stringToMap(data string) map[string]interface{} {
	obj := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(data), &obj)
	return obj
}
