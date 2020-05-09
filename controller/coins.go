package controller

import (
	"github.com/eabz/concretepay-backend/models"
	coinfactory "github.com/eabz/concretepay-backend/models/coin-factory"
)

type CoinController struct{}

func (ctrl *CoinController) GetCoins(params models.Params) (interface{}, error) {
	// Here you handle version filtering
	var response models.CoinsResponse
	if params.Version >= 1 {
		response.CoinsAvailable = len(coinfactory.Coins)
		for tag := range coinfactory.Coins {
			response.CoinsTickers = append(response.CoinsTickers, tag)
		}
	}
	return response, nil
}

func (ctrl *CoinController) GetCoinInfo(params models.Params) (interface{}, error) {
	coin, err := coinfactory.GetCoin(params.Tag)
	if err != nil {
		return nil, err
	}
	return coin.Info, nil
}
