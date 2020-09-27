package client

import (
	govclient "github.com/hschain/hschain/x/gov/client"
	"github.com/hschain/hschain/x/params/client/cli"
	"github.com/hschain/hschain/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
