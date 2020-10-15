package orders

import (
	"log"
	"sort"
)

// Sort : sort both bids and offers in an order table instance
func Sort(ot *OrderTable) {
	ot.Sort(BID)
	ot.Sort(OFFER)
}

// Sort : sort the bids or offers in an order table instance
func (ot *OrderTable) Sort(typ string) {
	log.Print("sorting " + typ)
	if typ == BID {
		ot.Bids = ot.Bids.sort()
	} else if typ == OFFER {
		ot.Offers = ot.Offers.sort()
	}
}

// sort : sort an array of bids
func (b Bids) sort() Bids {
	// order by time, ASCENDING, then price, DESCENDING
	_b := b
	sort.Slice(_b, func(i, j int) bool {
		if _b[i].Price > _b[j].Price {
			return true
		}
		if _b[i].Price < _b[j].Price {
			return false
		}
		return _b[i].Time < _b[j].Time
	})
	return _b
}

// sort : sort an array of offers
func (o Offers) sort() Offers {
	// order by time, ASCENDING, then price, ASCENDING
	_o := o
	sort.Slice(_o, func(i, j int) bool {
		if _o[i].Price < _o[j].Price {
			return true
		}
		if _o[i].Price > _o[j].Price {
			return false
		}
		return _o[i].Time < _o[j].Time
	})
	return _o
}
