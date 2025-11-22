package deployer_test

import (
	"testing"

	"github.com/nlewo/comin/internal/deployer"
	"github.com/nlewo/comin/internal/protobuf"
	"github.com/nlewo/comin/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestLivelinessCheck(t *testing.T) {
	// Test with a failing liveliness check
	args := deployer.NewTestDeployerArgs(t)
	args.LivelinessCheckCommand = "sh -c 'exit 1'"
	d := deployer.New(args)
	d.Submit(&protobuf.Generation{Uuid: "a"})
	go d.Run()
	deployment := <-d.DeploymentDoneCh
	assert.Equal(t, store.StatusToString(store.Failed), deployment.Status)

	// Test with a succeeding liveliness check
	args = deployer.NewTestDeployerArgs(t)
	args.LivelinessCheckCommand = "sh -c 'exit 0'"
	d = deployer.New(args)
	d.Submit(&protobuf.Generation{Uuid: "b"})
	go d.Run()
	deployment = <-d.DeploymentDoneCh
	assert.Equal(t, store.StatusToString(store.Done), deployment.Status)
}
