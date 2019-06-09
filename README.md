# Stocks Server
# Port 
    - :8082
# Microservice 
    - http://127.0.0.1:8082/stock/
    - http://127.0.0.1:8082/stock/{symbol}
    - http://127.0.0.1:8082/stock/AAPL?stock_exchange=NYSE
    - http://127.0.0.1:8082/stock/?page=1&limit=50

# Description
   This Stock Server returns the stock prices of the given stock symbol from the exchanges provided in the query
   parameter. If no query parameter is given then return the value of stock from AMEX (default) stock
   exchange. symbol must be a valid stock symbol.
    
    Query Params:
    stock_exchange - Optional. This should accept all valid exchange names. eg.
    stock_exchange=NASDAQ,NYSE
    
    Response:
    {
    "NASDAQ":{
    "symbol":"AAPL",
    "name":"Apple Inc.",
    "price":"154.94",
    "close_yesterday":"154.94",
    "currency":"USD",
    "market_cap":"732835688367",
    "volume":"142022",
    "timezone":"EST",
    "timezone_name":"America/New_York",
    "gmt_offset":"-18000",
    "last_trade_time":"2019-01-16 16:00:01"
    }
    }    

# Prerequisites:

   - Git
   - Golang 
