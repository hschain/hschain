package types

import (
	"fmt"

	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/params"
)

// Parameter store keys
var (
	KeyMintDenom     = []byte("MintDenom")
	KeyBlocksPerDay  = []byte("BlocksPerDay")
	KeyMintStartTime = []byte("MintStartTime")
)

// mint parameters
type Params struct {
	MintDenom     string `json:"mint_denom" yaml:"mint_denom"`            // type of coin to mint
	BlocksPerDay  uint64 `json:"blocks_per_day" yaml:"blocks_per_day""`   // expected blocks per day
	MintStartTime int64  `json:"mint_start_time" yaml:"mint_start_time""` //mint proces start time
}

// ParamTable for minting module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(mintDenom string, blocksPerDay uint64, mintStartTime int64) Params {

	return Params{
		MintDenom:     mintDenom,
		BlocksPerDay:  blocksPerDay,
		MintStartTime: mintStartTime,
	}
}

// default minting module parameters
func DefaultParams() Params {
	return Params{
		MintDenom:     sdk.DefaultMintDenom,
		BlocksPerDay:  uint64(60 * 60 * 24 / 5), // assuming 5 second block times
		MintStartTime: 0,
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
  MintStartTime:		  %d
`,
		p.MintDenom, p.BlocksPerDay, p.MintStartTime,
	)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyMintDenom, &p.MintDenom},
		{KeyBlocksPerDay, &p.BlocksPerDay},
		{KeyMintStartTime, &p.MintStartTime},
	}
}
