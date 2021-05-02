package config

// local orders config items
var LOCAL_ORDERS_ENABLED = true

var LOCAL_ORDERS_DIR = "/Users/andrewsanderson/Documents/dev/go/src/exchange/data/orders/"

var LOCAL_ORDERS_SCAN_TIME = 2

var LOCAL_ORDERS_DELETE_ON_READ = true

// HTTP orders config items
var HTTP_ORDERS_ENABLED = false

var HTTP_ORDERS_PORT = "8080"

var HTTP_ORDERS_END_POINT = "/order"

// matching & settlements config items
var SETTLEMENTS_OUTPUT_DIR = "/Users/andrewsanderson/Documents/dev/go/src/exchange/data/settlements/"

var MATCHING_RATE = 5
