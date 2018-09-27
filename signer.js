var BigInteger = require('bigi')
var bitcoin = require('bitcoinjs-lib')

var simnet = require('./simnet.js')
var log = require('./log.js')

bitcoin.networks.simnet = simnet

var wif = process.env.LND_WIF || '4NhVdHeWPkKCvoydkM9mkSySkzfeZkrX1YF5LLYRKdtUcJ6QCDq'

var alice_wif = wif
var bob_wif = wif

var alice_keychain_bi = new BigInteger('109193387910854206109060298261105268670715959771736985409204090197612975437563')
var alice_node_bi = new BigInteger('78107253186025767928659915978886941686615109367511192384184146645813729725746')

var keys = {
  alice_master: new bitcoin.ECPair(alice_node_bi),
  bob_master: new bitcoin.ECPair(alice_node_bi),
  alice: new bitcoin.ECPair(alice_keychain_bi),
  bob: new bitcoin.ECPair(alice_keychain_bi)
}

function signMessage (call, callback) {
    var type = call.request.type || "alice_master"
    var privKey = keys[type]

    log('[__________________________________________________________]')

    log('[ADDRESS]', privKey.getAddress())
    log('[PUBLIC ]', privKey.getPublicKeyBuffer())
    log('[PUBLIC X]', privKey.Q.x.toString(10))
    log('[PUBLIC y]', privKey.Q.y.toString(10))
    log('[PRIVATE]', privKey.toWIF())
    log('[PRIVATE D]', privKey.d.toBuffer())
    log('[PRIVATE D]', privKey.d.toString(10))

    log('[__________________________________________________________]')

    var data = call.request.msg
    log('[TO SIGN]', call.request)

    var hash = data
    log('[HASHED]', hash)

    var signed = privKey.sign(hash).toDER()

    log('[SIGNED R]', privKey.sign(hash).r.toString())

    log('[SIGNED S]', privKey.sign(hash).s.toString())

    var signature = signed.toString('base64')
    log('[SIGNATURE]', signature)

    log('[__________________________________________________________]')

    callback(null, { signature })
}

function signCompact (call, callback) {
    log('[__________________________________________________________]')

    var privKey = keys.alice_master

    var data = call.request.msg
    log('[TO SIGN]', call.request)

    var hash

    if (data.length != 32) {
      hash = bitcoin.crypto.sha256(bitcoin.crypto.sha256(data))
    } else {
      hash = data
    }

    log('[HASHED]', hash)

    var buf
    var signed = privKey.sign(hash).toCompact(1, true)
    log('[SIGNED]', signed)

    var signature = signed.toString('base64')
    log('[SIGNATURE]', signature)

    log('[__________________________________________________________]')

    callback(null, { signature })
}

module.exports = {
  signMessage, signCompact, keys
};
