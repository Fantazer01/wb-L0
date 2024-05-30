package main

import (
	"database/sql"
)

func getOrdersFromDb(connStr string) map[string]Order {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT order_uid, track_number, entry, name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM public.orders")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var cashOrders = make(map[string]Order)

	for rows.Next() {
		order := Order{}
		rows.Scan(&order.Order_uid,
			&order.Track_number,
			&order.Entry,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
			&order.Payment.Transaction,
			&order.Payment.Request_id,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.Payment_dt,
			&order.Payment.Bank,
			&order.Payment.Delivery_cost,
			&order.Payment.Goods_total,
			&order.Payment.Custom_fee,
			&order.Locale,
			&order.Internal_signature,
			&order.Customer_id,
			&order.Delivery_service,
			&order.Shardkey,
			&order.Sm_id,
			&order.Date_created,
			&order.Oof_shard)

		order.Items = getItemsFromDb(db, order.Track_number)

		cashOrders[order.Order_uid] = order
	}
	return cashOrders
}

func getItemsFromDb(db *sql.DB, trackNumber string) []Item {
	itemRows, err := db.Query("SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM public.items WHERE track_number = $1", trackNumber)
	if err != nil {
		panic(err)
	}
	defer itemRows.Close()

	var items = make([]Item, 0)
	for itemRows.Next() {
		item := Item{}
		itemRows.Scan(&item.Chrt_id,
			&item.Track_number,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.Total_price,
			&item.Nm_id,
			&item.Brand,
			&item.Status)
		items = append(items, item)
	}

	return items
}

func insertOrderToDb(connStr string, order Order) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO public.orders(order_uid, track_number, entry, name, phone, zip, city, address, region, email, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)",
		order.Order_uid,
		order.Track_number,
		order.Entry,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
		order.Payment.Transaction,
		order.Payment.Request_id,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.Payment_dt,
		order.Payment.Bank,
		order.Payment.Delivery_cost,
		order.Payment.Goods_total,
		order.Payment.Custom_fee,
		order.Locale,
		order.Internal_signature,
		order.Customer_id,
		order.Delivery_service,
		order.Shardkey,
		order.Sm_id,
		order.Date_created,
		order.Oof_shard)

	if err != nil {
		panic(err)
	}

	insertItemsToDb(db, order.Items)
}

func insertItemsToDb(db *sql.DB, items []Item) {
	for _, item := range items {
		db.Exec("INSERT INTO public.items(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			item.Chrt_id,
			item.Track_number,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.Total_price,
			item.Nm_id,
			item.Brand,
			item.Status)
	}
}
