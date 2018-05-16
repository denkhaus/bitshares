//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAssetReserve

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeAssetReserve] =
		sampleDataAssetReserveOperation
}

var sampleDataAssetReserveOperation = `{
  "amount_to_reserve": {
    "amount": 1000000,
    "asset_id": "1.3.472"
  },
  "extensions": [],
  "fee": {
    "amount": 4000000,
    "asset_id": "1.3.0"
  },
  "payer": "1.2.32469"
}`

//end of file
