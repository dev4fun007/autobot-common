package common

import (
	"time"
)

type TickerData struct {
	Market       string `json:"market"`
	Change24Hour string `json:"change_24_hour"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Volume       string `json:"volume"`
	LastPrice    string `json:"last_price"`
	Timestamp    int    `json:"timestamp"`
}

type BaseConfig struct {
	Name         string       `json:"name" bson:"name"`
	Market       string       `json:"market" bson:"market"`
	IsActive     bool         `json:"is_active" bson:"is_active"`
	StrategyType StrategyType `json:"strategy_type" bson:"strategy_type"`
}

type RequestMarketOrder struct {
	Quantity   float64
	LastPrice  float64 // For calculation purpose
	Market     string
	Config     BaseConfig
	OrderType  OrderType
	ActionType ActionType
}

type RequestLimitOrder struct {
	Market       string
	Quantity     float64
	PricePerUnit float64
	Config       BaseConfig
	OrderType    OrderType
	ActionType   ActionType
}

type OrderResponse struct {
	Orders []Order `json:"orders" bson:"orders"`
}

type Order struct {
	ID                string    `json:"id" bson:"id"`
	Market            string    `json:"market" bson:"market"`
	OrderType         string    `json:"order_type" bson:"order_type"`
	Side              string    `json:"side" bson:"side"`
	Status            string    `json:"status" bson:"status"`
	FeeAmount         float64   `json:"fee_amount" bson:"fee_amount"`
	Fee               float64   `json:"fee" bson:"fee"`
	TotalQuantity     float64   `json:"total_quantity" bson:"total_quantity"`
	RemainingQuantity float64   `json:"remaining_quantity" bson:"remaining_quantity"`
	AvgPrice          float64   `json:"avg_price" bson:"avg_price"`
	PricePerUnit      float64   `json:"price_per_unit" bson:"price_per_unit"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	ResponseError     string    `json:"response_error" bson:"response_error"`
}

type OrderEvent struct {
	Order        Order        `json:"order" bson:"order"`
	Timestamp    string       `json:"timestamp" bson:"timestamp"`
	BrokerName   string       `json:"broker_name" bson:"broker_name"`
	StrategyName string       `json:"strategy_name" bson:"strategy_name"`
	StrategyType StrategyType `json:"strategy_type" bson:"strategy_type"`
	EventError   string       `json:"event_error" bson:"event_error"`
}
