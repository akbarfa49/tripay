package


type Callback struct{
	Reference string `json:"reference"`
	MerchantRef string `json:"merchant_ref"`
	PaymentMethod string `json:"payment_method_code"`
	PaymentName string `json:"payment_method"`
	CustomerName string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
	CustomerPhone string `json:"customer_phone"`
	CallbackUrl string `json:"callback_url"`
	ReturnUrl string `json:"return_url"`
	Amount int `json:"total_amount"`
	MerchantFee int `json:"fee_merchant"`
	CustomerFee int `json:"fee_customer"`
	TotalFee int `json:"total_fee"`
	AmountReceived int `json:"amount_received"`
	//0 = open payment. 1 = closed payment
	PaymentType PaymentType `json:"is_closed_payment"`
	Status string `json:"status"`
	PaidAt int `json:"paid_at"`
	Note string `json:"note"`
}

func (t *Tripay) CallbackSignature(callbackData Callback)(signature string ){
	h := hmac.New(sha256.New, t.ApiKey)
	b := json.Marshal(&callbackData)
	h.Write(b)
	signature = hex.EncodeToString(h.Sum(nil))
	return signature
}

func (t *Tripay) CompareSignature(Signature1,Signature2 string) bool{
return	hmac.Equal(signature1,signature2)
}