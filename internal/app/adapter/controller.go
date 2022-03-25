package adapter

import (
	"fmt"
	"strconv"

	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/service"
	"github.com/UfiairENE/send-wire-pay/internal/app/application/usecase"
	"github.com/UfiairENE/send-wire-pay/internal/app/domain/valueobject"
	"github.com/gin-gonic/gin"
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	r.POST("/createaccount", ctrl.createaccount)
	r.GET("/getbalance/:user_id", ctrl.getbalance)
	r.GET("/getpaymenturl", ctrl.getPaymentUrl)
	r.PUT("/withdraw/:user_id", ctrl.withdraw)
	r.GET("/verifypayment/:transaction_id", ctrl.verifyPayment)
	return r
}

//SIGNUP USER
func (ctrl Controller) createaccount(c *gin.Context) {
	var User valueobject.NewUser
	c.BindJSON(&User)
	createaccount := usecase.Createaccount(User) // Dependency Injection
	c.JSON(200, createaccount)
}

func (ctrl Controller) createwallet(c *gin.Context) {
	var Wallet valueobject.Wallet
	c.BindJSON(&Wallet)
	createwallet:= usecase.Createwallet(Wallet) // Dependency Injection
	c.JSON(200, createwallet)
}

func (ctrl Controller) getbalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	Wallet := valueobject.Wallet{ID: id}
	getbalance:= usecase.GetBalance(Wallet) // Dependency Injection
	
	c.JSON(200, getbalance)
}

func (ctrl Controller) getPaymentUrl(c *gin.Context) {
	var form model.FundRequest
	c.BindJSON(&form)
	fmt.Println(form)
	resp, err := service.GetPaymentUrlUrl(form)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp)
}

func (ctrl Controller) verifyPayment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("transaction_id"))
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(400, "error getting query parameter user_id")
		return
	}
	resp, err := service.VerifyTransaction(id, userid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
}

func (ctrl Controller) withdraw(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("user_id"))
	var form model.WithdrawalRequest
	c.BindJSON(&form)
	err := service.Withdraw(form, userid)
	if err != nil {
		c.JSON(400, map[string]interface{}{"data": err.Error()})
		return
	}
	c.JSON(200, map[string]interface{}{"data": "withdrawal successful"})
}

