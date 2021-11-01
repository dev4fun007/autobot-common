package common

const (
	LogComponent = "Component"
)

type ActionType string

const (
	Buy      ActionType = "BUY"
	Sell     ActionType = "SELL"
	BuyOnce  ActionType = "BUY_ONCE"
	SellOnce ActionType = "SELL_ONCE"

	Wait ActionType = "WAIT"
	Hold ActionType = "HOLD"
)

type StrategyType string

const (
	InvalidStrategy         StrategyType = "INVALID_STRATEGY"
	SimpleSupportResistance StrategyType = "SIMPLE_SUPPORT_RESISTANCE"
)

var StrategyTypeList = [...]StrategyType{SimpleSupportResistance}

type OrderType string

const (
	MarketOrder OrderType = "MARKET_ORDER"
	LimitOrder  OrderType = "LIMIT_ORDER"
)
