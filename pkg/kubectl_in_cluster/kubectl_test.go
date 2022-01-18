package kubectl_in_cluster

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kubectl", func() {
	var err error

	Context("PreDeploy Tests", func() {
		It("Should successfully preapre manifest for a valid version", func() {
			filename, err := prepareDeploymentManifest("1.22", "namespace", "service-account")
			Ω(err).Should(BeNil())
			Ω(filename).ShouldNot(BeEmpty())
		})
		It("Should fail to preapre a manifeat for invalid version", func() {
			_, err = prepareDeploymentManifest("1.13", "namespace", "service-account")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("Deploy Tests", func() {
		It("Should successfully deploy a valid version", func() {
			err = Deploy("1.22", "namespace", "service-account")
			Ω(err).Should(BeNil())
		})
		It("Should fail deploy an invalid version", func() {
			err = Deploy("1.13", "namespace", "service-account")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("Run Tests", func() {

	})
})
