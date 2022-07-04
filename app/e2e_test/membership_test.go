package e2e_test

import (
	"comento_git_practice/app"
	"comento_git_practice/app/membership"
	"fmt"
	"github.com/gavv/httpexpect"
	"log"
	"net/http"
	"testing"
)

func TestTossRecreate(t *testing.T) {
	e := newHttpExpectConfig(t)

	t.Run("토스 멤버십을 신청한 후 삭제했다면, 다시 신청할 수 없다.", func(t *testing.T) {
		// given: 토스 멤버십을 신청한다.
		membershipCreateRequest := e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// when: 토스 멤버십을 삭제한다.
		e.DELETE(fmt.Sprintf("/memberships/%s", membershipCreateRequest.Value("ID").Raw())).
			Expect().
			Status(http.StatusOK)

		// then: 토스 멤버십을 다시 신청할 수 없다. 멤버십의 상태가 "탈퇴한 회원"이다.
		e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusBadRequest).
			JSON().Object().
			Value("message").Equal("재가입할 수 없습니다.")
	})
}

func TestGetMembershipByOwner(t *testing.T) {
	e := newHttpExpectConfig(t)

	t.Run("시나리오(1) 멤버십의 주인만 멤버십을 조회할 수 있다.", func(t *testing.T) {
		// given: 멤버십을 생성한다.
		membershipCreateResponse := e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "andy",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// when: 멤버십을 생성한 사용자가 로그인한다.
		loginResponse := e.POST("/login").
			WithFormField("name", "andy").
			WithFormField("password", "andy").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// then: 사용자의 멤버십 단건조회를 할 수 있다.
		e.GET(fmt.Sprintf("/memberships/%s", membershipCreateResponse.Value("id").Raw())).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Value("token").Raw())).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("id").Equal(membershipCreateResponse.Value("id").Raw())
	})
}

func TestGetMembershipByAdmin(t *testing.T) {
	e := newHttpExpectConfig(t)
	log.Println(e)

	t.Run("시나리오(2) Admin 사용자는 멤버십 전체 조회를 할 수 있다.", func(t *testing.T) {
		// given: 생성된 멤버십이 존재한다.
		e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "1",
				MembershipType: "toss",
			}).
			Expect().
			Status(http.StatusCreated)
		e.POST("/memberships").
			WithJSON(membership.CreateRequest{
				UserName:       "2",
				MembershipType: "naver",
			}).
			Expect().
			Status(http.StatusCreated)

		// when: Admin 사용자가 로그인한다.
		loginResponse := e.POST("/login").
			WithFormField("name", "admin").
			WithFormField("password", "admin").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		// then: 멤버십 전체 조회를 할 수 있다.
		memberships := e.GET("/memberships").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Value("token").Raw())).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			Value("memberships").Array()

		memberships.Length().Equal(2)

	})
}

func newHttpExpectConfig(t *testing.T) *httpexpect.Expect {
	handler := app.NewEcho(*app.DefaultConfig())

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e
}
