package books_test

import (
	"github.com/Neaj-Morshad-101/testing/books"
	"github.com/Neaj-Morshad-101/testing/library"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Books", func() {
	var foxInSocks, lesMis *books.Book

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

	It("can extract the author's last name", func() {
		book := &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		Expect(book.AuthorLastName()).To(Equal("Hugo"))
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
