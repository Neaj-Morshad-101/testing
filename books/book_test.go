package books_test

import (
	"github.com/Neaj-Morshad-101/testing/books"
	"github.com/Neaj-Morshad-101/testing/library"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Books", func() {
	var foxInSocks, lesMis, book *books.Book

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
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		Expect(book.IsValid()).To(BeTrue())
	})

	Describe("Categorizing books", func() {
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

	Describe("Validating Author Name", func() {
		It("can extract the author's last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
		It("interprets a single author name as a last name", func() {
			book.Author = "Hugo"
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})

		It("can extract the author's first name", func() {
			Expect(book.AuthorFirstName()).To(Equal("Victor"))
		})

		It("returns no first name when there is a single author name", func() {
			book.Author = "Hugo"
			Expect(book.AuthorFirstName()).To(BeZero()) //BeZero asserts the value is the zero-value for its type.  In this case: ""
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
