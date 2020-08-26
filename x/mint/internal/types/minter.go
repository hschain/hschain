package types

import (
	"fmt"

	sdk "hschain/types"
)

//MintPlan output plan
type MintPlan struct {
	Period         uint   `json:"period" yaml:"period"`
	TotalPerPeriod uint64 `json:"total_per_period" yaml:"total_per_period"`
	TotalPerDay    uint64 `json:"total_per_day" yaml:"total_per_day"`
}

// Minter represents the minting state.
type Minter struct {
	Inflation        sdk.Dec    `json:"inflation" yaml:"inflation"`                 // current annual inflation rate
	AnnualProvisions sdk.Dec    `json:"annual_provisions" yaml:"annual_provisions"` // current annual expected provisions
	MintPlans        []MintPlan `json:"mint_plans" yaml:"mint_plans"`               // mint plan
}

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions sdk.Dec, mintPlans []MintPlan) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
		MintPlans:        mintPlans,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation sdk.Dec, mintPlans []MintPlan) Minter {
	return NewMinter(
		inflation,
		sdk.NewDec(0),
		mintPlans,
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		sdk.NewDecWithPrec(13, 2),
		[]MintPlan{
			{0, 325000000, 1300000},
			{1, 325000000, 1300000 * 0.9},
			{2, 325000000, 1300000 * 0.9 * 0.9},
		},
	)
}

// validate minter
func ValidateMinter(minter Minter) error {
	if minter.Inflation.LT(sdk.ZeroDec()) {
		return fmt.Errorf("mint parameter Inflation should be positive, is %s",
			minter.Inflation.String())
	}
	return nil
}

// NextInflationRate returns the new inflation rate for the next hour.
func (m Minter) NextInflationRate(params Params, bondedRatio sdk.Dec) sdk.Dec {
	// The target annual inflation rate is recalculated for each previsions cycle. The
	// inflation is also subject to a rate change (positive or negative) depending on
	// the distance from the desired ratio (67%). The maximum rate change possible is
	// defined to be 13% per year, however the annual inflation is capped as between
	// 7% and 20%.

	// (1 - bondedRatio/GoalBonded) * InflationRateChange
	inflationRateChangePerYear := sdk.OneDec().
		Sub(bondedRatio.Quo(params.GoalBonded)).
		Mul(params.InflationRateChange)
	inflationRateChange := inflationRateChangePerYear.Quo(sdk.NewDec(int64(params.BlocksPerYear)))

	// adjust the new annual inflation for this next cycle
	inflation := m.Inflation.Add(inflationRateChange) // note inflationRateChange may be negative
	if inflation.GT(params.InflationMax) {
		inflation = params.InflationMax
	}
	if inflation.LT(params.InflationMin) {
		inflation = params.InflationMin
	}

	return inflation
}

// NextAnnualProvisions returns the annual provisions based on current total
// supply and inflation rate.
func (m Minter) NextAnnualProvisions(_ Params, totalSupply sdk.Int) sdk.Dec {
	return m.Inflation.MulInt(totalSupply)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}
