package deployer

import (
	"fmt"
	"os"
	"os/exec"

	shellwords "github.com/mattn/go-shellwords"
	pb "github.com/nlewo/comin/internal/protobuf"

	"github.com/sirupsen/logrus"
)

func runPreDeploymentCommand(command string, d *pb.Deployment) (string, error) {
	args, err := shellwords.Parse(command)
	if err != nil {
		return "", fmt.Errorf("failed to parse command %q: %w", command, err)
	}
	if len(args) == 0 {
		return "", fmt.Errorf("empty command")
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(os.Environ(),
		"COMIN_GIT_SHA="+envGitSha(d),
		"COMIN_GIT_REF="+envGitRef(d),
		"COMIN_GIT_MSG="+envGitMessage(d),
		"COMIN_HOSTNAME="+envCominHostname(d),
		"COMIN_FLAKE_URL="+envCominFlakeUrl(d),
		"COMIN_GENERATION="+envCominGeneration(d),
		"COMIN_STATUS="+envCominStatus(d),
		"COMIN_ERROR_MSG="+envCominErrorMessage(d),
	)

	output, err := cmd.CombinedOutput()
	outputString := string(output)
	if err != nil {
		return outputString, err
	}

	logrus.Debugf("cmd:[%s] output:[%s]", command, outputString)
	return outputString, nil
}
