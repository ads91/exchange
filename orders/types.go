package orders

// BID : a bid constant
const BID = "bid"

// OFFER : an offer constant
const OFFER = "offer"

// Bids : an array of Bid types
type Bids []Bid

// Offers : an array of Offer types
type Offers []Offer

// Bid : a bid
type Bid struct {
	Client string
	Amount int
	Price  float64
	Time   int64
}

// Offer : an offer
type Offer struct {
	Client string
	Amount int
	Price  float64
	Time   int64
}

// OrderTable : a collection of bids and offers
type OrderTable struct {
	Bids   Bids
	Offers Offers
}

// orderJSON : a JSON-representation of an order
type orderJSON struct {
	Type   string  `json:"Type"`
	Client string  `json:"Client"`
	Amount int     `json:"Amount"`
	Price  float64 `json:"Price"`
}
