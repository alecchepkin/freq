package freq_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFreq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Freq Suite")
}
