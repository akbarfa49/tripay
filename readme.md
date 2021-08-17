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
t := tripay.New("your-apikey","your-private-key","your-merchant-code",tripay.Development);
```

## Contents available
content method available so far

| Method  | Contents  | Status |
|---|---|---|
| `GetChannel(code PaymentChannelCode)` | `Channel Pembayaran` | OK |
| `GetInstruction(instruction Insruction)` | `Instruksi Pembayaran` | OK |
| `GetTransactionList(page,perPage int,channelCode PaymentChannelCode, sort, reference,merchantRef, status string)` | `Merchant List Pembayaran` | OK |
| `GetCost(amount int, code PaymentChannelCode)` | `Kalkulator Biaya` | OK |
| `RequestClosedTransaction(transaction RequestTransaction)` | `Close Transaksi` | OK |
| `CallbackSignature` | `Callback` | OK |

all response can retrieved by map[string]interface{} or from instance with suffix "Response"

## Testing

This package is tested using _test package.

Have Request? DM me on instagram.com/akbarfa49