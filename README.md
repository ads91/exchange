# Exchange

## What is it?

An engine for matching bids and offers. Trades are asynchronougly booked and matched through seperate goroutines.

Trade booking can be achieved via CSV or JSON file trade representations saved to an OS directory or via a JSON HTTP request.

## Usage

Have Go installed and added to your "PATH" OS environment variable. Navigate to where you'd like to clone the git repository then run the following in your terminal/cmd prompt

    git clone https://github.com/ads91/exchange.git

This will clone the exchange git repository. Now build the exchange executable (navigate into the exchange directory)

    go build

Once successfully built, you should now have an executable called exchange in the exchange directory. Run this executable

    ./exchange

## Future enhancements

1. Allow order sizes greater than one "unit". This could be achieved by settling volume that can be matched and re-booking/ordering the outstanding volume back to the order table with the initial orders' timestamps.

2. Writes to an instance of an OrderTable are done concurrently but the data structure itself is not atomic. This needs to change in order to avoid race conditions.

3. Order table sorting is performed concurrently, so an OrderTable instance needs to be locked on sort prior to matching (this would mean reads/writes for orders will be put on hold).

4. Only sort the order tables on addition of new orders, not after some arbitrary elapsed time.