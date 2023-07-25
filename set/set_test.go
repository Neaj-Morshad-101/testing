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

		Context("When the set contains items", func() {
			It("Should not be empty", func() {
				st := set.NewSet()
				st.Add("red")
				Expect(st.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return an increasing size", func() {
				st := set.NewSet()

				By("Empty set size being 0", func() {
					Expect(st.Size()).To(BeZero())
				})

				By("Adding first item", func() {
					st.Add("Red")

					Expect(st.Size()).To(Equal(1))
				})

				By("Adding a second item", func() {
					st.Add("Blue")

					Expect(st.Size()).To(Equal(2))
				})
			})
		})
	})

})
