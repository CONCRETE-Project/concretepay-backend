package main

import (
	"net/http"
	"os"
	"time"

	"github.com/CONCRETE-Project/concretepay-backend/controller"
	"github.com/CONCRETE-Project/concretepay-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	App := GetApp()
	err := App.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// GetApp is used to wrap all the additions to the GIN API.
func GetApp() *gin.Engine {
	App := gin.Default()
	App.Use(cors.Default())
	ApplyRoutes(App)
	return App
}

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		rate := limiter.Rate{
			Period: 1 * time.Hour,
			Limit:  1000,
		}
		store := memory.NewStore()
		limiterMiddleware := mgin.NewMiddleware(limiter.New(store, rate))
		api.Use(limiterMiddleware)

		coinsCtrl := controller.CoinController{}
		stakeCtrl := controller.StakeController{}

		api.GET("coins/:version", func(c *gin.Context) { callWrapper(c, coinsCtrl.GetCoins) })
		api.GET("coin/:tag", func(c *gin.Context) { callWrapper(c, coinsCtrl.GetCoinInfo) })
		api.GET("stake/:tag", func(c *gin.Context) { callWrapper(c, stakeCtrl.GetAddr) })

	}
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

func callWrapper(c *gin.Context, method func(models.Params) (data interface{}, err error)) {
	params := models.Params{
		Tag:     c.Param("tag"),
		Version: c.Param("version"),
	}
	res, err := method(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "status": -1})
		return
	} else {
		c.JSON(200, gin.H{"data": res, "status": 1})
		return
	}
}
