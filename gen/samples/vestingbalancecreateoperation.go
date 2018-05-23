//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeVestingBalanceCreate

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeVestingBalanceCreate] =
		sampleDataVestingBalanceCreateOperation
}

var sampleDataVestingBalanceCreateOperation = `{
  "amount": {
    "amount": 1,
    "asset_id": "1.3.1564"
  },
  "creator": "1.2.374566",
  "fee": {
    "amount": "3034918500",
    "asset_id": "1.3.1564"
  },
  "owner": "1.2.374566",
  "policy": [
    0,
    {
      "begin_timestamp": "2017-11-24T15:44:54",
      "vesting_cliff_seconds": 300,
      "vesting_duration_seconds": 600
    }
  ]
}`

//end of file
