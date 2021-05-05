# Exchange

## What is it?

An engine for matching bids and offers. Orders are asynchronougly booked and matched.

Orders are first matched by price then by time of arrival to the exchange. If the lowest offer price is higher than the highest bid price, the two are matched at the corresponding mid-price.

Order booking is achieved via CSV or JSON trade representations saved to a directory local to the application or consumed via an HTTP request.

## Usage

Have Go installed and added to your "PATH" OS environment variable. Navigate to where you'd like to clone the git repository then run the following in your terminal/cmd prompt

    git clone https://github.com/ads91/exchange.git

This will clone the exchange git repository. Now build the exchange executable (navigate into the exchange directory)

    go build

Once successfully built, you should now have an executable called exchange in the exchange directory. Run this executable

    ./exchange

The application logs the state of the order book with an interval equal to the time specified in the LOCAL_ORDERS_SCAN_TIME configuration item listed below. Initially, the order book is empty and we can add to the order book via approaches.

A JSON order take on the following structure independent of whether it's saved locally to the exchange or dispatched via HTTP.

```json
{
    "Type"   : "bid",
    "Client" : "client001",
    "Amount" : 1,
    "Price"  : 60.0
}
```

Where

- **Type** is the order type, bid or offer,
- **Client** is a unique identifier of the client placing the order,
- **Amount** is the number of units to be matched and
- **Price** is the price at which the client is prepared to buy/sell.

*Note: the exchange currently only fills orders of Amount equal to 1 unit.*

In the config directory a go file exists that defines application-level configuration items. These are as follows.

- **LOCAL_ORDERS_ENABLED** instructs the exchange to scan for orders saved in a local directory, true to scan, false otherwise,
- **LOCAL_ORDERS_DIR** is the directory for the exchange to scan if local order booking is enabled,
- **LOCAL_ORDERS_SCAN_TIME** is the wait time between scanning the local directory for orders,
- **LOCAL_ORDERS_DELETE_ON_READ** can be true to delete orders on read, false otherwise,
- **HTTP_ORDERS_ENABLED** is a flag to instruct the exchange to process HTTP orders, true to accept, false otherwise,
- **HTTP_ORDERS_PORT** is the port to listen on for orders sent via HTTP,
- **HTTP_ORDERS_END_POINT** is the HTTP end point to dispatch POST request orders,
- **SETTLEMENTS_OUTPUT_DIR** is a local directory to save settlement JSONs and
- **MATCHING_RATE** is the wait time between attempts for the exchange to match orders.

There are two ways in which the exchange books orders. The first is via orders represented as CSVs of JSONs being saved to a directory accesible to the exchange. The second method is by dispatching a JSON via HTTP to the exchange. Below we outline the two approaches with an example for each. 

#### Local order booking

In order for the exchange to monitor a directory for orders the LOCAL_ORDERS_ENABLED flag above must be set to true. 

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
