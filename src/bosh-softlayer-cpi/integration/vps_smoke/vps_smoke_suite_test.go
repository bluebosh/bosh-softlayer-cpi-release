package vps_smoke_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestVpsSmoke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VpsSmoke Suite")
}

var _ = BeforeSuite(func() {
	VPSExec := "test_assets/vps"
	postgresURL, err := GetPostgresURL()
	Expect(err).ToNot(HaveOccurred())

	command := exec.Command(VPSExec, "--scheme=https", "--tls-host=0.0.0.0", "--tls-port=1443", "--tls-certificate=./../test_assets/server.pem", "--tls-key=./../test_assets/server.key", "--databaseDriver", "postgres", "--databaseConnectionString", postgresURL)
	_, err = Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	//KillAndWait()
})

func GetPostgresURL() (string, error) {
	TargetURL := os.Getenv("POSTGRES_URL")
	if TargetURL == "" {
		return "", errors.New("POSTGRES_URL environment must be set")
	}

	return TargetURL, nil
}