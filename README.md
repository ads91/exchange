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