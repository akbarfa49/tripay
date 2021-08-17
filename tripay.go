package tripay

import "github.com/valyala/fasthttp"

const (
	Development Environment = iota
	Production
)

type PaymentChannelCode string

const(
	//BRI Virtual Account
	BRIVA PaymentChannelCode =  "BRIVA"
	//Maybank Virtual Account
	MYBVA PaymentChannelCode = "MYBVA"
	//Permata Virtual Account
	PERMATAVA PaymentChannelCode = "PERMATAVA"
	//BNI Virtual Account
	BNIVA PaymentChannelCode = "BNIVA"
	//Mandiri Virtual Account
	MANDIRIVA PaymentChannelCode = "MANDIRIVA"
	//BCA Virtual Account
	BCAVA PaymentChannelCode = "BCAVA"
	//Sinarmas Virtual Account
	SMSVA PaymentChannelCode = "SMSVA"
	//Muamalat Virtual Account
	MUAMALATVA PaymentChannelCode = "MUAMALATVA"
	//CIMB Virtual Account
	CIMBVA PaymentChannelCode = "CIMBVA"
	//BRI Virtual Account (Open Payment)
	BRIVAOP PaymentChannelCode = "BRIVAOP"
	//CIMB Niaga Virtual Account (Open Payment)
	CIMBVAOP PaymentChannelCode = "CIMBVAOP"
	//BCA Virtual Account (Open Payment)
	BCAVAOP PaymentChannelCode = "BCAVAOP"
	//BNI Virtual Account (Open Payment)
	BNIVAOP PaymentChannelCode = "BNIVAOP"
	//Alfamart
	ALFAMART PaymentChannelCode = "ALFAMART"
	//Alfamidi
	ALFAMIDI PaymentChannelCode = "ALFAMIDI"
	//QRIS
	QRIS PaymentChannelCode = "QRIS"
	//QRIS (Customizable)
	QRISC PaymentChannelCode = "QRISC"
	//QRIS (Open Payment)
	QRISOP PaymentChannelCode = "QRISOP"
	//QRIS (Open Payment - Customizable)
	QRISCOP PaymentChannelCode = "QRISCOP"
)

type Tripay struct {
	ApiKey     []byte 
	PrivateKey string 

	/*Environment to decide what this program running on

	use Development var for sandbox
	use Production var for Production node
	*/
	f *fasthttp.Client
	Host string
	MerchantCode string
	MerchantName string
}

type Environment int

func New(ApiKey, PrivateKey, MerchantCode string, environment Environment) *Tripay {
	host := ""
	switch environment{
	case Development:
		host = "https://tripay.co.id/api-sandbox"
	case Production:
		host = "https://tripay.co.id/api"
	}
	return &Tripay{ApiKey: []byte("Bearer "+ApiKey), PrivateKey: PrivateKey, Host: host,f: &fasthttp.Client{}, MerchantCode: MerchantCode}
}

/*SetHttpClient used for set fasthttp.Client
Default fasthttp is used if no set
*/
func (t *Tripay) SetHttpClient(f *fasthttp.Client){
	t.f = f
}

