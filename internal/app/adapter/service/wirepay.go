package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func GetPaymentUrlUrl(rq model.FundRequest) (interface{}, error) {
	fmt.Println(rq)
	db := postgres.Connection()
	method := "POST"
	Url, err := url.Parse("https://api.flutterwave.com/v3/payments")
	if err != nil {
		return "", err
	}

	user := model.User{}

	result := db.First(&user, rq.UserID)
	if result.Error != nil {
		panic(result.Error)
	}

	txRef, _ := uuid.NewV4()
	fltBody := model.FlutterRequestBody{
		TxRef:          fmt.Sprintf("%v", txRef),
		Amount:         fmt.Sprintf("%f", float64(rq.Amount)),
		Currency:       strings.ToUpper(rq.Currency),
		RedirectUrl:    "https://trove-assessment.herokuapp.com/loan",
		PaymentOptions: "",
		Customer: model.Customer{
			ID:    int(user.ID),
			Email: user.Email,
			Name:  user.FirstName + " " + user.LastName,
		},
		Customizations: model.Customizations{
			Title:       "Fund Wallet",
			Description: "Funding Wallet",
		},
	}
	fmt.Println("check id", fltBody.Customer.ID)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(fltBody)

	req, _ := http.NewRequest(method, Url.String(), buf)
	req.Header.Add("Authorization", "Bearer FLWSECK_TEST-e034e324f8562b334dc2955f6f3ca3e9-X")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	respModel := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Link string `json:"link"`
		} `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&respModel)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return respModel, nil
}

func VerifyTransaction(transactionID, userID int) (interface{}, error) {
	db := postgres.Connection()
	method := "GET"
	Url, err := url.Parse("https://api.flutterwave.com/v3/transactions/" + strconv.Itoa(transactionID) + "/verify")
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, Url.String(), nil)
	req.Header.Add("Authorization", "Bearer FLWSECK_TEST-e034e324f8562b334dc2955f6f3ca3e9-X")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	respModel := struct {
		Status          string `json:"status"`
		Message         string `json:"message"`
		TransactionInfo struct {
			ID                int     `json:"int"`
			TxRef             string  `json:"tx_ref"`
			FlwRef            string  `json:"flw_ref"`
			DeviceFingerprint string  `json:"device_fingerprint"`
			Amount            float64 `json:"amount"`
			Currency          string  `json:"currency"`
			Charged_amount    float64 `json:"charged_amount"`
			AppFee            float64 `json:"app_fee"`
			MerchantFee       float64 `json:"merchant_fee"`
			ProcessorResponse string  `json:"processor_response"`
			AuthModel         string  `json:"auth_model"`
			IP                string  `json:"ip"`
			Narration         string  `json:"narration"`
			Status            string  `json:"status"`
			PaymentType       string  `json:"payment_type"`
			Created_at        string  `json:"created_at"`
			Account_id        int     `json:"account_id"`
			Card              struct {
				First6digits string `json:"first_6digits"`
				Last4digits  string `json:"last_4digits"`
				Issuer       string `json:"issuer"`
				Country      string `json:"country"`
				Type         string `json:"type"`
				Token        string `json:"token"`
				Expiry       string `json:"expiry"`
			} `json:"card"`
			Meta struct {
				CheckoutInitAddress string `json:"__CheckoutInitAddress"`
			} `json:"meta"`
			AmountSettled float64 `json:"amount_settled"`
			Customer      struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				PhoneNumber string `json:"phone_number"`
				Email       string `json:"email"`
				CreatedAt   string `json:"created_at"`
			} `json:"customer"`
		} `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&respModel)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	transaction := model.Transactions{}

	result := db.Where("transaction_id = ? AND user_id >= ?", transactionID, userID).First(&transaction)

	if respModel.Status == "success" && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		var wallet model.Wallet
		result := db.First(&wallet, "user_id = ?", userID)
		if result.Error != nil {
			panic(result.Error)
		}
		fmt.Println("Wallet", wallet)
		currencySymbol := strings.ToLower(respModel.TransactionInfo.Currency)
		switch currencySymbol {
		case "ngn":
			wallet.NGNAmount += uint(respModel.TransactionInfo.Amount)
		case "ghs":
			wallet.GHSAmount += uint(respModel.TransactionInfo.Amount)
		case "usd":
			wallet.USDAmount += uint(respModel.TransactionInfo.Amount)
		}
		db.Save(&wallet)

		trans := model.Transactions{TransactionID: uint(transactionID), Amount: int(respModel.TransactionInfo.Amount), Trxref: respModel.TransactionInfo.TxRef, UserID: userID}
		db.Create(&trans)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	return respModel, nil
}

func Withdraw(form model.WithdrawalRequest, userID int) error {
	var wallet model.Wallet
	db := postgres.Connection()
	result := db.First(&wallet, "User_id = ?", userID)
	if result.Error != nil {
		panic(result.Error)
	}

	currencySymbol := strings.ToLower(form.Currency)
	switch currencySymbol {
	case "ngn":
		if uint(form.Amount) > wallet.NGNAmount {
			return fmt.Errorf("insufficient Balance")
		}
		wallet.NGNAmount -= uint(form.Amount)
	case "ghs":
		if uint(form.Amount) > wallet.GHSAmount {
			return fmt.Errorf("insufficient Balance")
		}
		wallet.GHSAmount -= uint(form.Amount)
	case "usd":
		if uint(form.Amount) > wallet.USDAmount {
			return fmt.Errorf("insufficient Balance")
		}
		wallet.USDAmount -= uint(form.Amount)
	}
	db.Save(&wallet)
	return nil

}
