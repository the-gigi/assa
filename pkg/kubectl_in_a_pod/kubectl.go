package kubectl_in_a_pod

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
	"github.com/the-gigi/kugo/kugo"
)

func Deploy() {

}

func Run(args ...string) (combinedOutput string, err error) {
	cmd := strings.Split("exec -it deploy/kubectl -- kubectl", " ")
	cmd = append(cmd, args...)

	return kugo.Run(cmd...)
}
