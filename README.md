# PayToBit
This is hyperledger hackthon project

## Successfully run on IBM bluemix

### Register

POST /registrar

{
  "enrollId": "user_type1_4",
  "enrollSecret": "014d931a18"
}

### DeploySpec

{
  "jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 1,
    "chaincodeID": {
      "path": "https://github.com/mjqin/PayToBit"
    },
    "ctorMsg": {
      "function": "Init",
      "args": [
        "111","222"
      ]
    },
    "secureContext": "user_type1_4"
  },
  "id": 0
}
