# testing

## Resources: Ginkgo and Gomega
https://onsi.github.io/ginkgo/
https://onsi.github.io/gomega/

Ginkgo is a testing framework for Go designed to help you write expressive tests. It is best paired with the Gomega matcher library. 
When combined, Ginkgo and Gomega provide a rich and expressive DSL (Domain-specific Language) for writing tests.

### Ginkgo: 
Ginkgo is a testing framework for Go designed to help you write expressive tests. 
### Gomega:
Gomega is a matcher/assertion library.



### Installing Ginkgo:
Ginkgo uses go modules. To add Ginkgo to your project, assuming you have a go.mod file setup, just go install it:
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...


#### Useful commands: 
`ginkgo bootstrap` //  add a Ginkgo suite to books pkg 
This will generate a file named books_suite_test.go in the books directory.

`ginkgo generate book` // to test book.go, generate book_test.go