package users_test

import (
	"errors"
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fortigit.com/forticloud/api/helper/helperfakes"
	"fortigit.com/forticloud/api/users"
	"fortigit.com/forticloud/api/users/usersfakes"
)

var _ = Describe("Handler", func() {
	var (
		c *helperfakes.FakeGinContext
		s *usersfakes.FakeService
		h users.UserHandler
	)
	BeforeEach(func() {
		c = &helperfakes.FakeGinContext{}
		s = &usersfakes.FakeService{}

		h = users.UserHandler{
			UserSvc: s,
		}
	})
	Describe("UpdateUser", func() {
		// START OMIT_1
		Context("when user id is not an integer", func() {
			BeforeEach(func() {
				c.ParamReturns("abc") // HLxxx
			})
			It("aborts with bad request", func() {
				h.UpdateUser(c)

				Expect(c.AbortWithStatusCallCount()).To(Equal(1))
				Expect(c.AbortWithStatusArgsForCall(0)).To(Equal(http.StatusBadRequest))
				Expect(c.JSONCallCount()).To(Equal(0))
			})
		})
		// END OMIT_1

		// START OMIT_2
		Context("when user is not found", func() {
			BeforeEach(func() {
				c.ParamReturns("123") // HLxxx
				s.ExistReturns(false)
			})
			It("aborts with 404 not found", func() {
				h.UpdateUser(c)

				Expect(s.ExistArgsForCall(0)).To(Equal(123))
				Expect(c.AbortWithStatusCallCount()).To(Equal(1))
				Expect(c.AbortWithStatusArgsForCall(0)).To(Equal(http.StatusNotFound))
				Expect(c.JSONCallCount()).To(Equal(0))
			})
		})
		// END OMIT_2

		Context("when binding failed", func() {
			BeforeEach(func() {
				c.ParamReturns("124")
				s.ExistReturns(true)
				c.BindJSONReturns(errors.New("binding error"))
			})
			It("returns error", func() {
				h.UpdateUser(c)

				Expect(s.ExistArgsForCall(0)).To(Equal(124))
				Expect(s.UpdateUserCallCount()).To(Equal(0))
				Expect(c.JSONCallCount()).To(Equal(0))
			})
		})

		Context("when saving to db failed", func() {
			BeforeEach(func() {
				c.ParamReturns("124")
				s.ExistReturns(true)
				c.BindJSONReturns(nil)
				s.UpdateUserReturns(nil, errors.New("db error"))
			})
			It("aborts with 500 error", func() {
				h.UpdateUser(c)

				Expect(s.ExistArgsForCall(0)).To(Equal(124))
				Expect(s.UpdateUserCallCount()).To(Equal(1))
				Expect(c.AbortWithStatusCallCount()).To(Equal(1))
				Expect(c.AbortWithStatusArgsForCall(0)).To(Equal(http.StatusInternalServerError))
				Expect(c.JSONCallCount()).To(Equal(0))
			})
		})

		Context("when update successfully", func() {
			BeforeEach(func() {
				user := users.User{
					UserID:    124,
					Email:     "bob@test.com",
					Name:      "Bob",
					AccountID: 11,
				}
				c.ParamReturns("124")
				s.ExistReturns(true)
				c.BindJSONReturns(nil)
				s.UpdateUserReturns(&user, nil)
			})
			It("returns 200 with uer info", func() {
				h.UpdateUser(c)

				Expect(s.ExistArgsForCall(0)).To(Equal(124))
				Expect(s.UpdateUserCallCount()).To(Equal(1))
				Expect(c.AbortWithStatusCallCount()).To(Equal(0))
				Expect(c.JSONCallCount()).To(Equal(1))

				code, u := c.JSONArgsForCall(0)

				resp, ok := u.(*users.User)
				Expect(ok).To(BeTrue())

				fmt.Println(ok)
				fmt.Println(resp)

				Expect(code).To(Equal(200))
				Expect(resp.UserID).To(Equal(uint(124)))
				Expect(resp.Email).To(Equal("bob@test.com"))
				Expect(resp.Name).To(Equal("Bob"))
				Expect(resp.AccountID).To(Equal(uint(11)))
			})
		})
	})
})
