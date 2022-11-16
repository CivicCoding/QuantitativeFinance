package wallet

import (
	"QuantitativeFinance/setting"
	"testing"
)

func Test_GetFundingAsset(t *testing.T) {
	setting.SetUp("app")
	GetFundingAsset("BUSD")
}
