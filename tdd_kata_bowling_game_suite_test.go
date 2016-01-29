package bowling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTddKataBowlingGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TddKataBowlingGame Suite")
}
