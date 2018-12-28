package domain

import "time"

type Order struct {
	Model
	OrderTime   *time.Time `json:"order_time"`
	ReceiveTime *time.Time `json:"receive_time"`
	Status      *string    `json:"status"`
	ShopID      UUID       `json:"shop_id"`
	DetailID    UUID       `json:"detail_id"`
	AccountID   UUID       `json:"account_id"`
}

type OrderDate struct {
	Date  int
	Month int
	Year  int
}
