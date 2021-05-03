# Exchange

## What is it?

An engine for matching bids and offers. Trades are asynchronougly booked and matched through seperate goroutines.

Trade booking can be achieved via CSV or JSON file trade representations saved to an OS directory or sent via a JSON HTTP request.

## Usage

Have Go installed and added to your "PATH" OS environment variable. Navigate to where you'd like to clone the git repository then run the following in your terminal/cmd prompt

    git clone https://github.com/ads91/exchange.git

This will clone the exchange git repository. Now build the exchange executable (navigate into the exchange directory)

    go build

Once successfully built, you should now have an executable called exchange in the exchange directory. Run this executable

    ./exchange

In the config directory a go file exists that defines application-level configuration items. These are as follows.

- **LOCAL_ORDERS_ENABLED** instructs the exchange to scan for orders saved in a directory local to the application,
- **LOCAL_ORDERS_DIR** the directory for the exchange to scan if local order booking is enabled,
- **LOCAL_ORDERS_SCAN_TIME** the wait time between scanning the local directory for orders,
- **LOCAL_ORDERS_DELETE_ON_READ** true to delete orders on read, false otherwise,
- **HTTP_ORDERS_ENABLED** instructs the exchange to accept orders via HTTP,
- **HTTP_ORDERS_PORT** the port to listen on for orders sent via HTTP,
- **HTTP_ORDERS_END_POINT** the HTTP end point to dispatch POST request orders,
- **SETTLEMENTS_OUTPUT_DIR** a local directory to save settlement JSONs and
- **MATCHING_RATE** the wait time between attempts for the exchange to match orders.

## Matching

Currently, there is one type of matching algorithm implemented, its logic is as follows.

The application is sheduled to match bids and offers within an instance of an order table periodically. An order table is a collection of booked bids and offers. Initially, these bids and offers are sorted - descending for bids and ascending for offers. In this way, a buyer who's prepared to pay the most will be matched with a seller who's prepared to sell at the lowest price.

After sorting by price, the bids and offers are then sorted by time descending. This is to show precedent to a buyer/seller who placed an order of the same price as another market participant but an earlier point in time.

Once the above sorting has been applied, we attempt to match the top of the book orders. If the top of the book bid is equal to the top of the book offer, the two are matched off at that price. If the top of the book bid is higher than the top of the book offer, the two are matched off at the mid-point of their prices. In this way, the buyer's order is filled at slightly less than they were prepared to pay and the seller's order is filled at a price slightly higher than they offered.

## Future enhancements

1. Allow order sizes greater than one "unit". This could be achieved by settling volume that can be matched and re-booking/ordering the outstanding volume back to the order table with the initial orders' timestamps.

2. Writes to an instance of an OrderTable are done concurrently but the data structure itself is not atomic. This needs to change in order to avoid race conditions.

3. Order table sorting is performed concurrently, so an OrderTable instance needs to be locked on sort prior to matching (this would mean reads/writes for orders will be put on hold).

4. Only sort the order tables on addition of new orders, not after some arbitrary elapsed time.
