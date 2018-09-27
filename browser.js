var grpc = require('grpc')
var fs = require('fs')

var log = require('./log.js')
var signer = require('./signer.js')

var br = grpc.load('browser_sign.proto').browser

function getServer() {
    var server = new grpc.Server()
    server.addService(br.BrowserSign.service, {
        signMessage: signer.signMessage,
        signCompact: signer.signCompact,
    })

    return server
}

if (require.main === module) {
    // If this is run as a script, start a server on an unused port
    var routeServer = getServer()

    routeServer.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
    log('[RUNNING]', 'localhost:50051')
    routeServer.start()
}

exports.getServer = getServer
