package tripay

import (
	"github.com/valyala/fasthttp"
)


type ChannelResponse struct{
	Success bool `json:"success"`
	Message string `json:"Message"`
	Data []struct{
		Group string `json:"group"`
		Code string `json:"code"`
		Name string `json:"name"`
		Type string `json:"type"`
		MerchantFee Fee `json:"fee_merchant"`
		CustomerFee Fee `json:"fee_customer"`
		TotalFee Fee `json:"total_fee"`
		Active bool `json:"active"`
	} `json:"Data"`
}

//due to inconsistent fee data type we can't specify the type of data. the data type may will variant(string, int, float).
type Fee struct{
	Flat interface{} `json:"flat"`
	Percent interface{} `json:"percent"`	
}


/*
GetChannel used for payment channel fetching information from tripay. 

u can retrieve this natively or use ChannelResponse struct

If the parameter is empty, this will return all payment channels information.
*/
func (t *Tripay) GetChannel(code PaymentChannelCode)([]byte, error) {
	uri := []byte(t.Host + "/merchant/payment-channel?code=" + string(code))
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}
type CostResponse struct{
	Success bool `json:"success"`
	Message string `json:"Message"`
	Data []struct{
		Code string `json:"code"`
		Name string `json:"name"`
		Fee struct{
			Flat interface{} `json:"flat"`
			Percent interface{} `json:"percent"`	
			Min interface{} `json:"min"`
			Max interface{} `json:"max"`
		} `json:"Fee"`
		TotalFee struct{
			Merchant int `json:"merchant"`
			Customer int `json:"customer"`
		} `json:"total_fee"`
		Active bool `json:"active"`
	} `json:"Data"`
}
/*
GetCost used for calculating the payment fees. 

u can retrieve this natively or use CostResponse struct.

If the parameter is empty, this will return all payment channels fees information.
*/
func (t *Tripay) GetCost(amount int,code PaymentChannelCode)([]byte, error) {
	var cost struct{
		Amount int `query:"amount"`
		Code PaymentChannelCode `query:"code,omitempty"`
	}
	cost.Amount = amount
	cost.Code = code
	q, _ :=structToQuery(&cost)
	uri := []byte(t.Host + "/merchant/fee-calculator?"+q)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}


type TransactionListResponse struct{
	Success bool `json:"success"`
	Message string `json:"Message"`
	Data []struct{
		Reference string `json:"reference"`
		MerchantRef string `json:"merchant_ref"`
		PaymentType string `json:"payment_selection_type"`

		//same like payment channel code
		PaymentMethod PaymentChannelCode `json:"payment_method"`

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
		PayCode int `json:"pay_code"`
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
		Note string `json:"note"`
		CreatedAt int `json:"created_at"`
		ExpiredAt int `json:"expired_at"`
		PaidAt int `json:"paid_at"`
	} `json:"Data"`

	Pagination struct{
		Sort string `json:"sort"`
		Offset struct{
			From int `json:"from"`
			To int `json:"to"`
		} `json:"offset"`
		CurrentPage int `json:"current_page"`
		//null json data will converted to 0
		PreviousPage int `json:"previous_page"`
		NextPage int `json:"next_page"`
		LastPage int `json:"last_page"`
		PerPage int `json:"per_page"`
		TotalRecords int `json:"total_records"`
	}
}


/*
GetTransaction list used for retrieve payment list

u can retrieve this natively or use TransactionListResponse struct

for param check https://tripay.co.id/developer?tab=merchant-transactions
*/
func (t *Tripay)GetTransactionList(page,perPage int,channelCode PaymentChannelCode, sort, reference,merchantRef, status string)([]byte, error){
	var list struct{
		Page int `query:"page,omitempty"`
		PerPage int `query:"perPage,omitempty"`
		Code PaymentChannelCode `query:"method,omitempty"`
		Sort string `query:"sort,omitempty"`
		Reference string `query:"reference,omitempty"`
		MerchantRef string `query:"merchant_ref,omitempty"`
		Status string `query:"status,omitempty"`
	}
	list.Page = page
	list.PerPage = perPage
	list.Code = channelCode
	list.Sort = sort
	list.Reference = reference
	list.MerchantRef = merchantRef
	list.Status = status

	q, err :=structToQuery(&list)
	if err != nil{
		return nil, err
	}
	uri := []byte(t.Host + "/merchant/transactions?"+q)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURIBytes(uri)
	req.Header.AddBytesV("Authorization", t.ApiKey)
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}