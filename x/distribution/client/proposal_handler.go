package client

import (
	"github.com/hschain/hschain/x/distribution/client/cli"
	"github.com/hschain/hschain/x/distribution/client/rest"
	govclient "github.com/hschain/hschain/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
