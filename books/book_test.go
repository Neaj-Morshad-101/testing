// Need to update book.go before running tests / specs
// Some functionality need to be added. Because they are in the test file.

package books_test

import (
	"github.com/Neaj-Morshad-101/testing/books"
	"github.com/Neaj-Morshad-101/testing/library"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"time"
)

var _ = Describe("Books", func() {
	var foxInSocks, lesMis, book *books.Book

	BeforeEach(func() {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		Expect(book.IsValid()).To(BeTrue())
	})

	Describe("Categorizing books", func() {
		BeforeEach(func() {
			lesMis = &books.Book{
				Title:  "Les Miserables",
				Author: "Victor Hugo",
				Pages:  2783,
			}

			foxInSocks = &books.Book{
				Title:  "Fox In Socks",
				Author: "Dr. Seuss",
				Pages:  24,
			}
			Expect(lesMis.IsValid()).To(BeTrue())
			Expect(foxInSocks.IsValid()).To(BeTrue())
		})
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(books.CategoryNovel))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(books.CategoryShortStory))
			})
		})
	})

	Describe("Extracting the author's first and last name", func() {
		Context("When the author has both names", func() {
			It("can extract the author's last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})
			It("can extract the author's first name", func() {
				Expect(book.AuthorFirstName()).To(Equal("Victor"))
			})
		})
		Context("When the author only has one name", func() {
			BeforeEach(func() {
				book.Author = "Hugo"
			})
			It("interprets a single author name as a last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})
			It("returns no first name when there is a single author name", func() {
				Expect(book.AuthorFirstName()).To(BeZero()) //BeZero asserts the value is the zero-value for its type.  In this case: ""
			})
		})

		Context("When the author has a middle name", func() {
			BeforeEach(func() {
				book.Author = "Victor Marie Hugo"
			})

			It("can extract the author's last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("can extract the author's first name", func() {
				Expect(book.AuthorFirstName()).To(Equal("Victor"))
			})
		})

		Context("When the author has no name", func() {
			It("should not be a valid book and returns empty for first and last name", func() {
				book.Author = ""
				Expect(book.IsValid()).To(BeFalse())
				Expect(book.AuthorLastName()).To(BeZero())
				Expect(book.AuthorFirstName()).To(BeZero())
			})
		})

	})

	Describe("some JSON decoding edge cases", func() {
		var book *books.Book
		var err error
		var json string
		JustBeforeEach(func() {
			book, err = NewBookFromJSON(json)
			Expect(book).To(BeNil())
		})

		When("the JSON fails to parse", func() {
			BeforeEach(func() {
				json = `{
				"title":"Les Miserables",
				"author":"Victor Hugo",
				"pages":2783oops
				}`
			})

			It("errors", func() {
				Expect(err).To(MatchError(books.ErrInvalidJSON))
			})
		})

		When("the JSON is incomplete", func() {
			BeforeEach(func() {
				json = `{
        		"title":"Les Miserables",
       	 		"author":"Victor Hugo",
      			}`
			})

			It("errors", func() {
				Expect(err).To(MatchError(books.ErrIncompleteJSON))
			})
		})
	})

	Describe("Reporting book weight", func() {
		var book *books.Book
		BeforeEach(func() {
			book = &books.Book{
				Title:  "Les Miserables",
				Author: "Victor Hugo",
				Pages:  2783,
				Weight: 500,
			}
			originalWeightUnits := os.Getenv("WEIGHT_UNITS")
			DeferCleanup(func() error {
				return os.Setenv("WEIGHT_UNITS", originalWeightUnits)
			})
		})

		AfterEach(func() {
			err := os.Setenv("WEIGHT_UNITS", originalWeightUnits)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("with no WEIGHT_UNITS environment set", func() {
			BeforeEach(func() {
				err := os.Clearenv("WEIGHT_UNITS")
				Expect(err).NotTo(HaveOccurred())
			})

			It("reports the weight in grams", func() {
				Expect(book.HumanReadableWeight()).To(Equal("500g"))
			})
		})

		Context("when WEIGHT_UNITS is set to oz", func() {
			BeforeEach(func() {
				err := os.Setenv("WEIGHT_UNITS", "oz")
				Expect(err).NotTo(HaveOccurred())
			})

			It("reports the weight in ounces", func() {
				Expect(book.HumanReadableWeight()).To(Equal("17.6oz"))
			})
		})

		Context("when WEIGHT_UNITS is invalid", func() {
			BeforeEach(func() {
				err := os.Setenv("WEIGHT_UNITS", "smoots")
				Expect(err).NotTo(HaveOccurred())
			})

			It("errors", func() {
				weight, err := book.HumanReadableWeight()
				Expect(weight).To(BeZero())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Saving books to a database", func() {
		AfterEach(func() {
			dbClient.Clear() //clear out the database between tests
		})

		JustAfterEach(func() {
			if CurrentSpecReport().Failed() {
				AddReportEntry("db-dump", dbClient.Dump())
			}
		})

		It("saves the book", func() {
			err := dbClient.Save(book)
			Expect(err).NotTo(HaveOccurred())
		})

	})

	It("can fetch a summary of the book from the library service", func(ctx SpecContext) {
		book := &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		summary, err := library.FetchSummary(ctx, book)
		Expect(err).NotTo(HaveOccurred())
		Expect(summary).To(ContainSubstring("Jean Valjean"))
	}, SpecTimeout(time.Second))
})
