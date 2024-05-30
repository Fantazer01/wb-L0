package main

import (
	"encoding/json"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

func main() {
	connStr := "user=userwb password=1234 dbname=wb sslmode=disable"

	var cashOrders = getOrdersFromDb(connStr)
	var muCashOrders = &sync.Mutex{}

	go natsStreamingHandler(connStr, cashOrders, muCashOrders)

	serverInit(cashOrders, muCashOrders)
}

func natsStreamingHandler(connStr string, cashOrders map[string]Order, muCashOrders *sync.Mutex) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	var dataChan = make(chan []byte)

	nc.Subscribe("order", func(m *nats.Msg) {
		dataChan <- m.Data
	})

	for !nc.IsClosed() {
		if data, opened := <-dataChan; opened {
			order := &Order{}
			json.Unmarshal(data, order)
			muCashOrders.Lock()
			cashOrders[order.Order_uid] = *order
			muCashOrders.Unlock()
			insertOrderToDb(connStr, *order)
		}
	}
}

func serverInit(cashOrders map[string]Order, muCashOrders *sync.Mutex) {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var id string = r.URL.Query().Get("id")
		if id != "" {
			muCashOrders.Lock()
			val, ok := cashOrders[id]
			muCashOrders.Unlock()

			if !ok {
				http.Error(w, "Order not found by id", http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(val)
		} else {
			http.Error(w, "id parameter is missing", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
