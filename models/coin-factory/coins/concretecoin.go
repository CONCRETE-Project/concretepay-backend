package coins

// Bitcoin coinfactory information
var ConcreteCoin = Coin{
	Info: CoinInfo{
		Icon:      "data:image/png;base64,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",
		Tag:       "CCT",
		Name:      "ConcreteCoin",
		Segwit:    false,
		Blockbook: "",
		Protocol:  "cct",
		TxVersion: 1,
		TxBuilder: "bitcoinjs",
		HDIndex:   0,
		Networks: map[string]CoinNetworkInfo{
			"P2PKH": {
				MessagePrefix: "\x18 DarkNet Signed Message:\n",
				Bech32:        "",
				Bip32: CoinNetWorkBip32Info{
					Public:  0x042e2635,
					Private: 0x041e342d,
				},
				PubKeyHash: 28,
				StakeHash:  63,
				ScriptHash: 15,
				Wif:        45,
			},
		},
	},
}