package tripay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/valyala/fasthttp"
)


const (
	OpenPayment PaymentType = iota
	ClosedPayment 
)

type PaymentType int

type RequestTransaction struct {
	MerchantRef string `json:"merchant_ref,omitempty"`

	Method PaymentChannelCode `json:"method,omitempty"`

	CustomerName  string `json:"customer_name,omitempty"`
	CustomerEmail string `json:"customer_email,omitempty"`
	CustomerPhone string `json:"customer_phone,omitempty"`
	CallbackUrl   string `json:"callback_url,omitempty"`
	ReturnUrl     string `json:"return_url,omitempty"`
	Amount        int    `json:"amount,omitempty"`
	CheckoutUrl   string `json:"checkout_url,omitempty"`
	OrderItems   []Item  `json:"order_items,omitempty"`

	//unixtimestamp
	ExpiredTime int `json:"expired_time,omitempty"`

	//hmac-sha256
	Signature   string `json:"signature"`
}

type ClosedTransactionResponse struct{
		Success bool `json:"success"`
		Message string `json:"Message"`
		Data struct{
			Reference string `json:"reference"`
			MerchantRef string `json:"merchant_ref"`
			PaymentType string `json:"payment_selection_type"`
	
			//same like payment channel code
			PaymentMethod string `json:"payment_method"`
			PaymentName string `json:"payment_name"`
			CustomerName string `json:"customer_name"`
			CustomerEmail string `json:"customer_email"`
			CustomerPhone string `json:"customer_phone"`
			CallbackUrl string `json:"callback_url"`
			ReturnUrl string `json:"return_url"`
			Amount int `json:"amount"`
			MerchantFee interface{} `json:"fee_merchant"`
			CustomerFee interface{} `json:"fee_customer"`
			TotalFee interface{} `json:"total_fee"`
			AmountReceived int `json:"amount_received"`
			PayCode string `json:"pay_code"`
			PayUrl string `json:"pay_url"`
			CheckoutUrl string `json:"checkout_url"`
			OrderItems []struct{
				Sku string `json:"sku"`
				Name string `json:"name"`
				Price int `json:"price"`
				Quantity int `json:"quantity"`
				Subtotal int `json:"Subtotal"`
			}`json:"order_items"`
			Status string `json:"status"`
			PaidTime string `json:"paid_time,omitempty"`
			ExpiredTime int `json:"expired_time"`
			Instructions []struct{
				Title string `json:"title"`
				Steps []string `json:"steps"`
			} `json:"instructions"`
			/*Reference string `json:"reference,omitempty"`
			MerchantRef string `json:"merchant_ref,omitempty"`
			PaymentType string `json:"payment_selection_type,omitempty"`
	
			//same like payment channel code
			PaymentMethod string `json:"payment_method,omitempty"`
			PaymentName string `json:"payment_name,omitempty"`
			CustomerName string `json:"customer_name,omitempty"`
			CustomerEmail string `json:"customer_email,omitempty"`
			CustomerPhone string `json:"customer_phone,omitempty"`
			CallbackUrl string `json:"callback_url,omitempty"`
			ReturnUrl string `json:"return_url,omitempty"`
			Amount int `json:"amount,omitempty"`
			MerchantFee interface{} `json:"fee_merchant,omitempty"`
			CustomerFee interface{} `json:"fee_customer,omitempty"`
			TotalFee interface{} `json:"total_fee,omitempty"`
			AmountReceived int `json:"amount_received,omitempty"`
			PayCode int `json:"pay_code,omitempty"`
			PayUrl string `json:"pay_url,omitempty"`
			CheckoutUrl string `json:"checkout_url,omitempty"`
			OrderItems []struct{
				Sku string `json:"sku,omitempty"`
				Name string `json:"name,omitempty"`
				Price int `json:"price,omitempty"`
				Quantity int `json:"quantity,omitempty"`
				Subtotal int `json:"Subtotal,omitempty"`
			}`json:"order_items,omitempty"`
			Status string `json:"status,omitempty"`
			PaidTime string `json:"paid_time,omitempty"`
			ExpiredTime int `json:"expired_time,omitempty"`
			Instructions []struct{
				Title string `json:"title,omitempty"`
				Steps string `json:"steps,omitempty"`
			} `json:"instructions,omitempty"`*/
		} `json:"data"`	
	}

type Item struct{
		Sku      string `json:"sku,omitempty"`
		Name     string `json:"name,omitempty"`
		Price    int    `json:"price,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
	
}


func (t *Tripay) RequestClosedTransaction(transaction RequestTransaction) ([]byte,error){
	h := hmac.New(sha256.New, []byte(t.PrivateKey))
	h.Write([]byte(t.MerchantCode+transaction.MerchantRef+strconv.Itoa(transaction.Amount)))
	transaction.Signature = hex.EncodeToString(h.Sum(nil))
	b, err :=json.Marshal(&transaction)
	if err != nil{
		return nil, err
	}
	uri := []byte(t.Host + "/transaction/create")
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.AppendBody(b)
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := t.f.Do(req, res); err != nil {
		return nil,err
	}
	return res.Body(), nil
}


func (t *Tripay) ClosedTransactionDetails(reference string) ([]byte, error){
	uri := []byte(t.Host+"/transaction/detail?reference="+reference)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := t.f.Do(req, res); err != nil {
		return nil,err
	}
	return res.Body(), nil
}

//this will be continued once i get production API
func (t *Tripay) RequestOpenTransaction(transaction RequestTransaction) ([]byte,error){
	h := hmac.New(sha256.New, []byte(t.PrivateKey))
	h.Write([]byte(t.MerchantCode+transaction.MerchantRef+strconv.Itoa(transaction.Amount)))
	transaction.Signature = hex.EncodeToString(h.Sum(nil))
	b, err :=json.Marshal(&transaction)
	if err != nil{
		return nil, err
	}
	uri := []byte(t.Host+"/open-payment/create")
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.AppendBody(b)
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := t.f.Do(req, res); err != nil {
		return nil,err
	}
	return res.Body(), nil
}

