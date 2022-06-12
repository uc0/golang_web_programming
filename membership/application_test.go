package membership

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMembership(t *testing.T) {
	t.Run("멤버십을 생성한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		req := CreateRequest{"jenny", "naver"}
		res, err := app.Create(req)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("이미 등록된 사용자 이름이 존재할 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})

		req := CreateRequest{"jenny", "toss"}
		_, err := app.Create(req)

		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"", "toss"}
		_, err := app.Create(req)

		assert.NotNil(t, err)
	})

	t.Run("멤버십 타입을 입력하지 않은 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", ""}
		_, err := app.Create(req)

		assert.NotNil(t, err)
	})

	t.Run("naver/toss/payco 이외의 타입을 입력한 경우 실패한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))

		req := CreateRequest{"jenny", "foo"}
		_, err := app.Create(req)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("membership 정보를 갱신한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{createRes.ID, "jenny", "toss"}
		res, err := app.Update(req)

		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID)
		assert.NotEmpty(t, res.UserName)
		assert.Equal(t, req.MembershipType, res.MembershipType)
	})

	t.Run("수정하려는 사용자의 이름이 이미 존재하는 사용자 이름이라면 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{"foo", "jenny", "toss"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("멤버십 아이디를 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		_, _ = app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{"", "jenny", "naver"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("사용자 이름을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{createRes.ID, "", "naver"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("멤버쉽 타입을 입력하지 않은 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{createRes.ID, "jenny", ""}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})

	t.Run("주어진 멤버쉽 타입이 아닌 경우, 예외 처리한다.", func(t *testing.T) {
		app := NewApplication(*NewRepository(map[string]Membership{}))
		createRes, _ := app.Create(CreateRequest{"jenny", "naver"})

		req := UpdateRequest{createRes.ID, "jenny", "foo"}
		_, err := app.Update(req)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("멤버십을 삭제한다.", func(t *testing.T) {

	})

	t.Run("id를 입력하지 않았을 때 예외 처리한다.", func(t *testing.T) {

	})

	t.Run("입력한 id가 존재하지 않을 때 예외 처리한다.", func(t *testing.T) {

	})
}
