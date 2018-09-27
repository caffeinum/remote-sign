var simnet = {
  wif: 0x64,
  messagePrefix: '\x18Bitcoin Signed Message:\n',
  bech32: 'sb',
  bip32: {
    public: 0x0420bd3a,
    private: 0x0420b900
  },
  pubKeyHash: 0x3f,
  scriptHash: 0x7b,
}

module.exports = simnet
