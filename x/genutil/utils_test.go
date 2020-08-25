package genutil

import (
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"hschain/tests"

	"github.com/stretchr/testify/require"
)

func TestExportGenesisFileWithTime(t *testing.T) {
	t.Parallel()
	dir, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()

	fname := filepath.Join(dir, "genesis.json")
	require.NoError(t, ExportGenesisFileWithTime(fname, "test", nil, json.RawMessage(""), time.Now()))
}
