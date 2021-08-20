Un-official Tripay Payment Gateway
===============

IMPORTANT: Make sure you read the documentation and understand what these methods are used for!

need atleast Go Version 1.16 and above to use this package
## Instalation
```
go get github.com/akbarfa49/tripay
```

## Example
```GO
tr := tripay.New("your-apikey","your-private-key","your-merchant-code",tripay.Development);

req := tripay.RequestTransaction{
			PaymentMethod: tripay.BNIVA,
			MerchantRef: "INV69",
			Amount: 20000,
			CustomerName: "akbarfa",
			CustomerEmail: "fania@123.com",
			CustomerPhone: "081234567891",
			OrderItems: []tripay.Item{0: {Sku: "duar",Name: "duar",Price: 20000,Quantity: 1}},
		}
		b, err := tr.RequestClosedTransaction(req)
		if err != nil {
			log.Panic(err)
			return
		}
		v:=tripay.ClosedTransactionResponse{}
		if err := json.Unmarshal(b, &v); err != nil{
			log.Panicln(err)
		}
		os.WriteFile("dump/requestclosed.json", b, 0644)
    
```

## Testing

This package is tested using _test package.

Have Request? DM me on instagram.com/akbarfa49
