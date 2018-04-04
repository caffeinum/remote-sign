var grpc = require('grpc');
var fs = require("fs");

var host = process.env.LND_HOST || "localhost"
var port = process.env.LND_PORT || "10009"
var url = host+':'+port

var lndCert = fs.readFileSync("../.lnd/tls.cert");
var credentials = grpc.credentials.createSsl(lndCert);
var lnrpcDescriptor = grpc.load("rpc.proto");
var lnrpc = lnrpcDescriptor.lnrpc;
var lightning = new lnrpc.Lightning(url, credentials);

console.log("[CONNECTING]", url);

lightning.getInfo({}, function(err, response) {
    if ( err ) console.error( err );
    else console.log('GetInfo:', response);
});

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


