package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	// RegisterFailHandler(Fail) is the single line of glue code connecting Ginkgo to Gomega.
	// If we were to avoid dot-imports this would read as gomega.RegisterFailHandler(ginkgo.Fail)
	//- what we're doing here is telling our matcher library (Gomega) which function to call (Ginkgo's Fail) in the event a failure is detected.
	RegisterFailHandler(Fail)

	// RunSpecs() call tells Ginkgo to start the test suite, passing it the *testing.T instance and a description of the suite. You should only ever call RunSpecs once and you can let Ginkgo worry about calling *testing.T for you.

	RunSpecs(t, "Books Suite")
}
