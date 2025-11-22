package deployer

import (
	"fmt"
	"strings"

	pb "github.com/nlewo/comin/internal/protobuf"
)

func envGitSha(d *pb.Deployment) string {
	return d.Generation.SelectedCommitId
}

func envGitRef(d *pb.Deployment) string {
	return fmt.Sprintf("%s/%s", d.Generation.SelectedRemoteName, d.Generation.SelectedBranchName)
}

func envGitMessage(d *pb.Deployment) string {
	return strings.Trim(d.Generation.SelectedCommitMsg, "\n")
}

func envCominGeneration(d *pb.Deployment) string {
	return d.Generation.Uuid
}

func envCominHostname(d *pb.Deployment) string {
	return d.Generation.Hostname
}

func envCominStatus(d *pb.Deployment) string {
	return d.Status
}

func envCominErrorMessage(d *pb.Deployment) string {
	return d.ErrorMsg
}

func envCominFlakeUrl(d *pb.Deployment) string {
	return d.Generation.FlakeUrl
}

