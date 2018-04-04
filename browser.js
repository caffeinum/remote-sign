var grpc = require('grpc')
var fs = require('fs')
var bitcoin = require('bitcoinjs-lib')

var host = process.env.BR_HOST || 'localhost'
var port = process.env.BR_PORT || '11009'
var url = host+':'+port

var br_descriptor = grpc.load('browser_sign.proto')
var br = br_descriptor.browser

var wif = process.env.LND_WIF || '92Qu1Qc6xuJERWre7fPZ5vtjmLz14dKjmVVdTZVZ7MsVGXtPk7E'
var myKey = new bitcoin.ECPair.fromWIF(wif, bitcoin.networks.testnet)

console.log('[ADDRESS]', myKey.getAddress())
console.log('[PUBLIC ]', myKey.getPublicKeyBuffer())
console.log('[PRIVATE]', myKey.toWIF())

var bcrypto = bitcoin.crypto

function signMessage (call, callback) {
    var data = call.request.msg
    console.log('[TO SIGN]', call.request)

    var hash = bitcoin.crypto.sha256(bitcoin.crypto.sha256(data))
    console.log('[HASHED]', hash)

    var signed = myKey.sign(hash).toDER()
    console.log('[SIGNED]', signed)

    var signature = signed.toString('base64')
    console.log('[SIGNATURE]', signature)

    callback(null, { signature })
}

function getServer() {
    var server = new grpc.Server()
    server.addService(br.BrowserSign.service, {
        signMessage: signMessage
    })

    return server
}

if (require.main === module) {
    // If this is run as a script, start a server on an unused port
    var routeServer = getServer()

    routeServer.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
    console.log('[RUNNING]', 'localhost:50051')
    routeServer.start()
}

exports.getServer = getServer
