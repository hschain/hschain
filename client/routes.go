package client

import (
	"github.com/gorilla/mux"

	"github.com/hschain/hschain/client/context"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	RegisterRPCRoutes(cliCtx, r)
}
