package api_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/drbyronw/accounts/api"
	"github.com/go-chi/chi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	r := chi.NewRouter()
	var wa api.WebApp

	Context("Handlers respond to API requests", func() {
		Context("When we need an connection", func() {
			It("should properly setup router and respond to Home '/' route", func() {
				Expect(api.SetupRoutes(r, &wa)).Should(Succeed())

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(200))

				defer w.Result().Body.Close()
				// Expect(w.Body.String()).To(Equal("BANK Accounts API"))
			})
		})

	})
})
