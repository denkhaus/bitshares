//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeTransferFromBlind

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeTransferFromBlind] =
		sampleDataTransferFromBlindOperation
}

var sampleDataTransferFromBlindOperation = `{
  "amount": {
    "amount": 100000,
    "asset_id": "1.3.0"
  },
  "blinding_factor": "30806481146974b091f415b857aef06c4e7d4468ea95d43c44cee62c5216eb04",
  "fee": {
    "amount": 2000000,
    "asset_id": "1.3.0"
  },
  "inputs": [
    {
      "commitment": "02fee663b188f809d8f476eb17757ac7e8f9bf28c694f9007dde1b90d6d3e72afb",
      "owner": {
        "account_auths": [],
        "address_auths": [],
        "key_auths": [],
        "weight_threshold": 0
      }
    }
  ],
  "to": "1.2.35298"
}`

//end of file
