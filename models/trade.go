package models
import (
    "log"
)


type Trade struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price string `json:"price"`
    Amount string `json:"amount"`
    Dt string `json:"dt"`
    Buyer string `json:"buyer"`
    Seller string `json:"seller"`
}

func GetTrades()  []*Trade {

    rows, err  := db.Query("SELECT * FROM trades")
    if err != nil {
        log.Fatal(err)
    }


    trades := make([]*Trade, 0 , 10)

    var id, name, price, amount, dt, buyer, seller string

    for rows.Next() {
        err = rows.Scan(&id, &name, &price, &amount, &dt, &buyer, &seller)
        if err != nil {
            log.Fatal(err)
        }
        trades = append(trades, &Trade{id, name, price, amount, dt, buyer, seller})
    }

    return trades

}