package storage

import "go-kit/orderservice"

var OrderStore map[string][entity.Order] = make([]orderservice.Order, 5, 100)

const OrderStoreMaxSize int = 100
