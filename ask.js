var grpc = require('grpc');
var fs = require("fs");

var host = "localhost"
var port = "50051"
var url = host+':'+port

//var lndCert = fs.readFileSync("../.lnd/tls.cert");
var credentials = grpc.credentials.createInsecure();
var descriptor = grpc.load("browser_sign.proto");
var service = new descriptor.browser.BrowserSign(url, credentials);

console.log("[CONNECTING]", url);

var msg = "31702890165552608143474968043917386623439832025424163227238141880065588587353"
var connectionData = "00005ac4ff0202a280179f0dc77233b7fbace08134520d7098f79dc1774d8c1d9c7ba4817668f33399ff30326132383031373966306463373732333362370000000000000000000000000000"

service.signMessage({ msg: connectionData }, function(err, response) {
    if ( err ) console.error( err );
    else console.log('GetInfo:', response);
});

/*
var call = lightning.subscribeInvoices({});

call.on('data', function(invoice) {
    console.log(invoice);
})
.on('end', function() {
  // The server has finished sending
})
.on('status', function(status) {
  // Process status
  console.log("Current status",  status);
});

*/
