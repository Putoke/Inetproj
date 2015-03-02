package models

import (
    "encoding/json"
    "net/http"
)

func Index(w http.ResponseWriter, r * http.Request) {
    trades := GetTrades()
    json.NewEncoder(w).Encode(trades)
}