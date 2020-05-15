package coins

// Bitcoin coinfactory information
var ConcreteCoin = Coin{
	Info: CoinInfo{
		Icon:      "data:image/png;base64,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",
		Tag:       "CCT",
		Name:      "ConcreteCoin (CCT)",
		Segwit:    false,
		Blockbook: "http://114.67.97.142:9130",
		Protocol:  "cct",
		TxVersion: 1,
		TxBuilder: "bitcoinjs",
		HDIndex:   0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18ConcreteCoin Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x0488B21e,
					Private: 0x0488ade4,
				},
				PubKeyHash: 28,
				ScriptHash: 15,
				Wif:        45,
				StakeHash:  63,
			},
		},
	},
}
