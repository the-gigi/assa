package kubectl_in_cluster

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/the-gigi/kugo"
	"os"
	"strings"
	"text/template"
)

var (
	supportedVersions = map[string]bool{
		"1.21": true,
		"1.22": true,
		"1.23": true,
	}
)

func prepareDeploymentManifest(version string, namespace, serviceAccount string) (filename string, err error) {
	if !supportedVersions[version] {
		err = errors.Errorf("Unsupported version: %s. Supported version are: %v", version, supportedVersions)
		return
	}

	filename = "deployment-" + version + ".yaml"
	file, err := os.Create(filename)
	if err != nil {
		return
	}

	writer := bufio.NewWriter(file)
	t := template.New("")
	_, err = t.Parse(deploymentTemplate)
	if err != nil {
		return
	}
	err = t.Execute(writer, struct {
		Name           string
		Version        string
		Namespace      string
		ServiceAccount string
	}{
		"kubectl-" + version,
		version,
		namespace,
		serviceAccount,
	})
	if err != nil {
		return
	}

	return
}

func Deploy(version string, namespace, serviceAccount string) (err error) {
	filename, err := prepareDeploymentManifest(version, namespace, serviceAccount)
	if err != nil {
		return
	}

	cmd := "apply -f " + filename
	_, err = kugo.Run(cmd)
	return
}

func Run(args ...string) (combinedOutput string, err error) {
	cmd := strings.Split("exec -it deploy/kubectl -- kubectl", " ")
	cmd = append(cmd, args...)

	return kugo.Run(cmd...)
}
