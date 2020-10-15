package matching

import (
	"exchange/orders"
	"log"
	"strconv"
	"sync"
	"time"
)

// SettleOrders : continuously wait and attempt to settle orders periodically in an order table
func SettleOrders(ot *orders.OrderTable, wg *sync.WaitGroup, settleFunc func(stl Settlement, fname string), dir string, waitTimeSecs int) {
	defer wg.Done()
	// run indefinitely
	for {
		// wait
		time.Sleep(time.Duration(waitTimeSecs) * time.Second)
		// try to match an order
		settlement, ok := match(ot)
		// write settlement to a JSON if a big and offer are matched
		if ok {
			fname := dir + getSettlementFilename(settlement)
			settleFunc(settlement, fname)
		}
	}
}

// match : attempt to settle orders in an order table instance
func match(ot *orders.OrderTable) (Settlement, bool) {
	// sort the order table
	ot.Sort(orders.BID)   // only sort when bids added (in exchange/orders)
	ot.Sort(orders.OFFER) // only sort when offers added (in exchange/orders)
	// extract first elements from the sorted bids and offers
	if len(ot.Bids) > 0 && len(ot.Offers) > 0 {
		stl, ok := settle(ot.Bids[0], ot.Offers[0])
		if ok {
			ot.Bids = ot.Bids[1:]
			ot.Offers = ot.Offers[1:]
			log.Printf("matched bid %g with offer %g at %g", stl.Bid.Price, stl.Offer.Price, stl.Price)
			return stl, ok
		}
	}
	return Settlement{}, false
}

// settle : settlement logic for a bid and an offer
func settle(b orders.Bid, o orders.Offer) (Settlement, bool) {
	// [TODO]: work with multiple amounts
	if b.Price == o.Price {
		// match at same price if order prices are equal
		return Settlement{Bid: b, Offer: o, Amount: 1, Price: b.Price}, true
	} else if b.Price > o.Price {
		// if order prices aren't equal, match at the mid-price, this way buyer pays slightly less
		// and seller recieves slightly more
		return Settlement{Bid: b, Offer: o, Amount: 1, Price: (b.Price + o.Price) / 2}, true
	}
	return Settlement{}, false
}

// getSettlementFilename : filename representation for a settlement
func getSettlementFilename(stl Settlement) string {
	t := time.Now()
	a := strconv.Itoa(stl.Amount)
	p := strconv.FormatFloat(stl.Price, 'g', 6, 64) // g -> f?
	return t.Format("2006.01.02_15.04.05") + "_" + stl.Bid.Client + "_" + stl.Offer.Client + "_" + a + "_" + p
}
