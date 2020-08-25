package client

import (
	govclient "hschain/x/gov/client"
	"hschain/x/params/client/cli"
	"hschain/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
