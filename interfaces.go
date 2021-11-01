package common

import "context"

type Worker interface {
	Close()
	Start(ctx context.Context)
	GetBaseConfig() BaseConfig
	UpdateConfig(config interface{})
	GetDataChannel() chan<- TickerData
}

type BrokerAction interface {
	GetBrokerName() string
	ExecuteMarketOrder(order RequestMarketOrder) (Order, error)
	ExecuteLimitOrder(order RequestLimitOrder) (Order, error)
}

type ApiService interface {
	CreateStrategy(context.Context, []byte, StrategyType) error
	UpdateStrategy(context.Context, string, []byte, StrategyType) error
	GetStrategiesByType(context.Context, StrategyType) ([]byte, error)
	GetStrategyByNameAndType(context.Context, string, StrategyType) ([]byte, error)
	DeleteStrategyByNameAndType(context.Context, string, StrategyType) error
}

type OrderService interface {
	StartOrderService(ctx context.Context)
	ExecuteMarketOrder(order RequestMarketOrder)
	ExecuteLimitOrder(order RequestLimitOrder)
}

type WorkerRegistryService interface {
	GetActiveWorkers() map[string]Worker
	GetRegisteredWorker(configName string, strategyType StrategyType) Worker
	RegisterConfigWorker(worker Worker, strategyType StrategyType)
	UpdateConfigWorkerRegistry(worker Worker, strategyType StrategyType)
	RemoveConfigWorkerFromRegistry(name string, strategyType StrategyType)
}
