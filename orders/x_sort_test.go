package orders

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	Sort(&OrderTable{})
}

func TestSortBids(t *testing.T) {
	// make some bids
	bids := Bids{
		Bid{"client-04", 1, 150.0, 4},
		Bid{"client-02", 1, 125.0, 1},
		Bid{"client-03", 1, 100.0, 2},
		Bid{"client-01", 1, 100.0, 1},
	}
	// sort
	ot := &OrderTable{Bids: bids, Offers: nil}
	log.Printf("unsorted order table is %v", ot.Bids)
	ot.Sort(BID)
	log.Printf("sorted order table is %v", ot.Bids)
	// expected result
	sortedBids := Bids{
		Bid{"client-04", 1, 150.0, 4},
		Bid{"client-02", 1, 125.0, 1},
		Bid{"client-01", 1, 100.0, 1},
		Bid{"client-03", 1, 100.0, 2},
	}
	// check rows
	for i := 0; i < len(sortedBids); i++ {
		if ot.Bids[i] != sortedBids[i] {
			log.Printf("rows at index %v aren't equal", i)
			t.Fail()
		}
	}
}

func TestSortOffers(t *testing.T) {
	// make some bids
	offers := Offers{
		Offer{"client-04", 1, 150.0, 4},
		Offer{"client-02", 1, 125.0, 1},
		Offer{"client-03", 1, 100.0, 2},
		Offer{"client-01", 1, 100.0, 1},
	}
	// sort
	ot := &OrderTable{Bids: nil, Offers: offers}
	log.Printf("unsorted order table is %v", ot.Offers)
	ot.Sort(OFFER)
	log.Printf("sorted order table is %v", ot.Offers)
	// expected result
	sortedOffers := Offers{
		Offer{"client-01", 1, 100.0, 1},
		Offer{"client-03", 1, 100.0, 2},
		Offer{"client-02", 1, 125.0, 1},
		Offer{"client-04", 1, 150.0, 4},
	}
	// check rows
	for i := 0; i < len(sortedOffers); i++ {
		if ot.Offers[i] != sortedOffers[i] {
			log.Printf("rows at index %v aren't equal", i)
			t.Fail()
		}
	}
}
