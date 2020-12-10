module github.com/drbyronw/sw-testing-cd-examples/unit-test-functions

go 1.15

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.4
	github.com/shopspring/decimal v1.2.0
	sw-test-4-cd/unit-test-functions v0.0.0-00010101000000-000000000000
)

replace sw-test-4-cd/unit-test-functions => ./
