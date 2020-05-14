package controller

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/CONCRETE-Project/concretepay-backend/models"
)

type StakeController struct{}

var stakeAddr = []string{
	"SZ2xSuY4U4uPXzohHMpgpoVKCJDYsmgacU",
	"SfZcVnLkjnPwga9JjgwDv57AaMvMvwh63f",
	"SVxyTfNdcKDVAe9aLsp2UTakvfizH2oKkm",
	"SQVsnQR2PuAp2GL4E266Dw21euRXmi18Yd",
	"SiDADEodbRsGsAviMycidcv3gyhb763SA1",
	"SP3onsa9ncnhwouutpHARQEmQt2fsc4Pjr",
}

func (ctrl *StakeController) GetAddr(params models.Params) (interface{}, error) {
	// Here you handle version filtering
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 4
	fmt.Println(rand.Intn(max-min+1) + min)
	return (stakeAddr[rand.Intn(max-min+1)+min]), nil
}
