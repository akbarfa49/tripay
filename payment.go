package tripay

import (
	"github.com/valyala/fasthttp"
)


type BasePayment struct {
	Name string `json:"nama"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Type string `json:"tipe"`
	Fee  int    `json:"biaya_admin"`
}

type Instruction struct {
	// code is the code of payment channel name *the Value must be UPPER*. ex: BRIVA (important)
	PaymentMethod  PaymentChannelCode `json:"code"  query:"code"`
		
	// pay_code is payment code retrieved from tripay
	PayCode    string `json:"pay_code,omitempty"  query:"pay_code,omitempty"`

	// amount used for insert the amount to the instruction
	Amount     int    `json:"amount"  query:"amount,omitzero"`

	
	/* allow_html to set permission of html tag response.
		0 = unallowed
		1 = allowed
		by default is 0
	*/
	Allow_html int    `json:"allow_html"  query:"allow_html"`
}

type InstructionResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data []struct{
		Title string `json:"title"`
		Steps []string `json:"steps"`
	} `json:"data,omitempty"`
}

/*
GetInstruction used for get Step to do the payment and return it in byte

if confused with the response can use InstructionResponse struct to retrieve the value
*/
func (t *Tripay) GetInstruction(instruction Instruction) ([]byte, error){
	b, _ := structToQuery(&instruction)
	uri := []byte(t.Host + "/payment/instruction?" + b)
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
