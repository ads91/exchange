package matching

import "exchange/orders"

// Settlement : a bid-offer settlement
type Settlement struct {
	Bid    orders.Bid
	Offer  orders.Offer
	Amount int
	Price  float64
}
