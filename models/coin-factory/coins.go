package coinfactory

import (
	"errors"
	"strings"

	"github.com/CONCRETE-Project/concretepay-backend/models/coin-factory/coins"
)

// Coins refers to the coins that are being used on the API instance
var Coins = map[string]*coins.Coin{
	"CCT": &coins.ConcreteCoin,
}

// GetCoin is the safe way to check if a coin exists and retrieve the coin data
func GetCoin(tag string) (*coins.Coin, error) {
	coin, ok := Coins[strings.ToUpper(tag)]
	if !ok {
		return nil, errors.New("coin not available")
	}
	coin = &coins.Coin{
		Info: coin.Info,
	}
	return coin, nil
}
