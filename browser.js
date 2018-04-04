var grpc = require("grpc");
var fs = require("fs");
var bitcoin = require("bitcoinjs-lib");

var host = process.env.BR_HOST || "localhost"
var port = process.env.BR_PORT || "11009"
var url = host+':'+port

var br_descriptor = grpc.load("browser_sign.proto");
var br = br_descriptor.browser;

var Server = new grpc.Server();

function SignMessage () {
}

