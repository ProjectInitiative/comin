package deployer

import (
	"context"
	"testing"

	"github.com/nlewo/comin/internal/store"
	"github.com/stretchr/testify/assert"
)

func testDeployerArgs(t *testing.T) *DeployerArgs {
	tmp := t.TempDir()
	s, err := store.New(tmp+"/state.json", tmp+"/gcroots", 1, 1)
	assert.Nil(t, err)
	return &DeployerArgs{
		Store: s,
		DeployFunc: func(ctx context.Context, outPath, operation string) (bool, string, error) {
			return false, "", nil
		},
	}
}


