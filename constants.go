package common

const (
	LogComponent = "Component"
)

type ActionType string

const (
	Wait     ActionType = "WAIT"
	Hold     ActionType = "HOLD"
	Buy      ActionType = "BUY"
	Sell     ActionType = "SELL"
	BuyOnce  ActionType = "BUY_ONCE"
	SellOnce ActionType = "SELL_ONCE"
)

type OrderType string

const (
	MarketOrder OrderType = "MARKET_ORDER"
	LimitOrder  OrderType = "LIMIT_ORDER"
)
