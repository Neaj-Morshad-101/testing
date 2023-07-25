package set_test

import (
	"github.com/Neaj-Morshad-101/testing/set"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Describe("Emptiness", func() {
		Context("When the set does not contain items", func() {
			It("Should be empty", func() {
				st := set.NewSet()
				Expect(st.IsEmpty()).To(BeTrue())
			})
		})
	})

})
