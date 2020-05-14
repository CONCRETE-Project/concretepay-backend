package coins

// Bitcoin coinfactory information
var Beetlecoin = Coin{
	Info: CoinInfo{
		Icon:      "data:image/png;base64,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",
		Tag:       "BEET",
		Name:      "BeetleCoin (BEET)",
		Segwit:    false,
		Blockbook: "https://111.229.169.195:9130",
		Protocol:  "beet",
		TxVersion: 1,
		TxBuilder: "bitcoinjs",
		HDIndex:   0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18Beetlecoin Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x04b24746,
					Private: 0x04b2430c,
				},
				PubKeyHash: 75,
				ScriptHash: 85,
				Wif:        127,
				StakeHash:  63,
			},
		},
	},
}
