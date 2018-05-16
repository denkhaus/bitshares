//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeVestingBalanceWithdraw

package samples

import (
	"github.com/denkhaus/bitshares/gen/data"
	"github.com/denkhaus/bitshares/types"
)

func init() {
	data.OpSampleMap[types.OperationTypeVestingBalanceWithdraw] =
		sampleDataVestingBalanceWithdrawOperation
}

var sampleDataVestingBalanceWithdrawOperation = `{
  "amount": {
    "amount": 406050000,
    "asset_id": "1.3.0"
  },
  "fee": {
    "amount": 200000,
    "asset_id": "1.3.0"
  },
  "owner": "1.2.23043",
  "vesting_balance": "1.13.48"
}`

//end of file
