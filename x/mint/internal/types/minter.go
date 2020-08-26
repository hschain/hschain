package types

import (
	"fmt"

	sdk "hschain/types"
	"math"
)

const (
	defaultPeriodCount    = 32
	defaultTotalPerPeriod = 325000000
	defaultTotalPerDay    = 13000000
	defaultDeflation      = 0.91
)

//MintPlan output plan
type MintPlan struct {
	Period         int     `json:"period" yaml:"period"`
	TotalPerPeriod sdk.Int `json:"total_per_period" yaml:"total_per_period"`
	TotalPerDay    sdk.Int `json:"total_per_day" yaml:"total_per_day"`
}

// Minter represents the minting state.
type Minter struct {
	MintPlans []MintPlan `json:"mint_plans" yaml:"mint_plans"` // mint plan
}

//DefaultMintPlans create
func DefaultMintPlans() []MintPlan {
	var plans []MintPlan
	for i := 0; i < defaultPeriodCount; i++ {
		plan := MintPlan{
			Period:         i,
			TotalPerPeriod: sdk.NewInt(int64(defaultTotalPerPeriod)),
			TotalPerDay:    sdk.NewInt(int64(math.Floor(defaultTotalPerDay*math.Pow(defaultDeflation, float64(i)) + 0.5))),
		}
		plans = append(plans, plan)
	}
	return plans
}

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(mintPlans []MintPlan) Minter {
	return Minter{
		MintPlans: mintPlans,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(mintPlans []MintPlan) Minter {
	return NewMinter(
		mintPlans,
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
func DefaultInitialMinter() Minter {
	return InitialMinter(
		DefaultMintPlans(),
	)
}

// validate minter
func ValidateMinter(minter Minter) error {
	if len(minter.MintPlans) <= 0 {
		return fmt.Errorf("mint parameter mintplan length should be greater than 0, is %d", len(minter.MintPlans))
	}
	return nil
}

//当日产量
func (m Minter) CurrentDayProvisions(totalSupply sdk.Int) sdk.Dec {
	total := totalSupply //当前期已产总量
	current := -1        //当前期数

	for i := range m.MintPlans {
		if total.GTE(m.MintPlans[i].TotalPerPeriod) {
			total = total.Sub(m.MintPlans[i].TotalPerPeriod)
		} else {
			current = i
			break
		}
	}

	if current == -1 { //矿已完了
		return sdk.ZeroDec()
	}

	//当前期剩余总量大于日产量
	if m.MintPlans[current].TotalPerPeriod.Sub(total).GTE(m.MintPlans[current].TotalPerDay) {
		return sdk.NewDecFromInt(m.MintPlans[current].TotalPerDay)
	}

	//如果当前期是最后一期的剩余, 全部挖完
	if current == len(m.MintPlans)-1 {
		return sdk.NewDecFromInt(m.MintPlans[current].TotalPerPeriod.Sub(total))
	}

	//可以从下一期借用
	left := m.MintPlans[current].TotalPerPeriod.Sub(total)
	leftRatio := float64(left.Int64()) / float64(m.MintPlans[current].TotalPerDay.Int64())

	return sdk.NewDecFromInt(left.Add(sdk.NewInt(int64(float64(m.MintPlans[current+1].TotalPerDay.Int64()) * (1 - leftRatio)))))

}

// NextPeriodProvisions returns the period provisions based on current total
// supply and mintplans.
//下一次减产后的日产量
func (m Minter) NextPeriodProvisions(totalSupply sdk.Int) sdk.Dec {
	for i := range m.MintPlans {
		if totalSupply.GTE(m.MintPlans[i].TotalPerPeriod) {
			totalSupply = totalSupply.Sub(m.MintPlans[i].TotalPerPeriod)
		} else {
			if i < len(m.MintPlans)-1 {
				return sdk.NewDecFromInt(m.MintPlans[i+1].TotalPerDay)
			}
		}
	}
	return sdk.ZeroDec()
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params, totalSupply sdk.Int) sdk.Coin {
	provisionAmt := m.CurrentDayProvisions(totalSupply).QuoInt(sdk.NewInt(int64(params.BlocksPerDay)))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}
