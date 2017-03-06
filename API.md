# PayToBit API

Init() // Initial

params:

* cashAddr  // accout for receive money e.g. alipay account
* bitAddr 

response: none

---
Invoke()

---
Query()

## seller 

applyForSell()  // Seller apply for selling coins

params:

response:

{"addr": "bitAddr" }
 
---
bundingCoin()  // Seller bunding coin in a tempory account for selling

params:

* recvAddr
* totalCoin
* transHash

response:

{"recvAddr": recvAddr , "totalCoin": totalCoin , "txID": txID }

---
revokeTx() // Seller revoke his coins

params:

* txID	// tranistion id

response:

{"status":"ok"}


## buyer

getSellingList() // look for coins on selling

params:

* threshold	//	

response:

{"txID1", txID2", ...}

---
getTxByID() // get transaction inforation in detail

params:

* txID

response:

{"recvAddr": recvAddr , "totalCoin": totalCoin , "txID": txID }

---
submitPaymentProof() // buyer paid successfully

params:

* proof 	//
* bitAddr	//

response:

{}
