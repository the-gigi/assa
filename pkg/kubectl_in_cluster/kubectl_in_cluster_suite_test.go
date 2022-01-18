package kubectl_in_cluster_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKubectlInCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubectlInCluster Suite")
}
