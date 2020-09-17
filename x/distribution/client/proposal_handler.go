package client

import (
	"hschain/x/distribution/client/cli"
	"hschain/x/distribution/client/rest"
	govclient "hschain/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
