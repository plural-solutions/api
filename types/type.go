package types

import (
	"github.com/lucasgomide/menyoo-api/schema"
)

type Store interface {
	ProductStore
}

type ProductStore interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
	ProductByRestaurantAndID(restaurantID int, productID int) (schema.Product, error)
}

type ProductCmd interface {
	ProductsByRestaurant(restaurantID int) ([]schema.Product, error)
	ProductByRestaurantAndID(restaurantID int, productID int) (schema.Product, error)
}

type OrderCmd interface {
	CreateOrder(order schema.Order) (schema.Order, error)
}
