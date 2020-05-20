package coins

type CoinNetWorkBip32Info struct {
	Public  int `json:"public"`
	Private int `json:"private"`
}

type CoinNetworkInfo struct {
	MessagePrefix string               `json:"messagePrefix"`
	Bech32        string               `json:"bech32"`
	Bip32         CoinNetWorkBip32Info `json:"bip32"`
	PubKeyHash    int                  `json:"pubKeyHash"`
	StakeHash     int                  `json:"stakeHash"`
	ScriptHash    int                  `json:"scriptHash"`
	Wif           int                  `json:"wif"`
}

type CoinInfo struct {
	Icon         string                     `json:"icon"`
	Tag          string                     `json:"tag"`
	Name         string                     `json:"name"`
	Segwit       bool                       `json:"segwit"`
	Blockbook    string                     `json:"blockbook"`
	BlockbookApi string                     `json:"blockbook_api"`
	Protocol     string                     `json:"protocol"`
	TxVersion    int                        `json:"tx_version"`
	TxBuilder    string                     `json:"tx_builder"`
	HDIndex      int                        `json:"hd_index"`
	Networks     map[string]CoinNetworkInfo `json:"networks"`
}

// Coin is the basic coin structure to get the correct properties for each coin.
type Coin struct {
	Info CoinInfo `json:"info"`
}
