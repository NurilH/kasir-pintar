package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"test_kasir_pintar/models"
	response "test_kasir_pintar/responses"

	"github.com/labstack/echo/v4"
)

var t = time.Now()
var time_stamp = t.Format("20060102150405")
var merchant_key = "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="
var i_mid = "IONPAYTEST"
var ref_no = "ORD" + time_stamp

func PostRegistrationController(c echo.Context) error {
	post_body := models.VirtuanAcount{}
	c.Bind(&post_body)

	// build merchant data
	merchantData := time_stamp + i_mid + ref_no + post_body.Amt + merchant_key
	// hash sha256 merchantData
	var mer_tok = sha256.Sum256([]byte(merchantData))

	var token = fmt.Sprintf("%x", string(mer_tok[:]))

	post_body.TimeStamp = time_stamp
	post_body.ReferenceNo = ref_no
	post_body.IMid = i_mid
	post_body.MerchantToken = token

	post, _ := json.Marshal(post_body)
	ReqBody := bytes.NewBuffer(post)
	log.Println("log request", ReqBody)
	fmt.Println("")
	api_url := "https://dev.nicepay.co.id/nicepay/direct/v2/registration"

	resp, err := http.Post(api_url, "application/json", ReqBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	ResBody := models.ResVirtual{}

	Res_Body := bytes.NewBuffer(body)
	log.Println("log respons", Res_Body)
	fmt.Println("")
	json.Unmarshal(body, &ResBody)

	if ResBody.ResultCd != "0000" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", ResBody))

}
