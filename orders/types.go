package orders

const BID = "bid"
const OFFER = "offer"

type Bids []Bid
type Offers []Offer

type Bid struct {
	Client string
	Amount int
	Price  float64
	Time   int64
}

type Offer struct {
	Client string
	Amount int
	Price  float64
	Time   int64
}

type OrderTable struct {
	Bids   Bids
	Offers Offers
}

type orderJSON struct {
	Type   string  `json:"Type"`
	Client string  `json:"Client"`
	Amount int     `json:"Amount"`
	Price  float64 `json:"Price"`
}
