package types

import (
	"fmt"

	sdk "hschain/types"
	"hschain/x/params"
)

// Parameter store keys
var (
	KeyMintDenom    = []byte("MintDenom")
	KeyBlocksPerDay = []byte("BlocksPerDay")
)

// mint parameters
type Params struct {
	MintDenom    string `json:"mint_denom" yaml:"mint_denom"`          // type of coin to mint
	BlocksPerDay uint64 `json:"blocks_per_day" yaml:"blocks_per_day""` // expected blocks per day
}

// ParamTable for minting module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(mintDenom string, blocksPerDay uint64) Params {

	return Params{
		MintDenom:    mintDenom,
		BlocksPerDay: blocksPerDay,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:    sdk.DefaultBondDenom,
		BlocksPerDay: uint64(60 * 60 * 24 / 5), // assuming 5 second block times
	}
}

// validate params
func ValidateParams(params Params) error {
	if params.MintDenom == "" {
		return fmt.Errorf("mint parameter MintDenom can't be an empty string")
	}
	return nil
}

func (p Params) String() string {
	return fmt.Sprintf(`Minting Params:
  Mint Denom:             %s
  BlocksPerDay:			  %d
`,
		p.MintDenom, p.BlocksPerDay,
	)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyMintDenom, &p.MintDenom},
		{KeyBlocksPerDay, &p.BlocksPerDay},
	}
}
