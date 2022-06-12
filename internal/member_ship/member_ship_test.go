package member_ship

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	/*
		[멤버십 생성]
		요청: 사용자의 이름, 멤버십 타입 (네이버/쿠팡/페이코)
		응답: 멤버십 ID, 멤버십 타입
		- 멤버십 타입은 네이버/쿠팡/페이코만 가능하며 그 외를 입력하면 에러를 반환한다.
		- 사용자의 이름을 입력하지 않은 경우 에러를 반환한다.
		- 이미 존재하는 사용자 이름이라면 에러를 반환한다.
		- 멤버십 ID는 고유해야 한다.
	*/
	t.Run("멤버십을 등록한다.", func(t *testing.T) {
		app := Application{}
		memberShipID, memberShipType, err := app.Create("김희선", "naver")
		assert.Nil(t, err)
		assert.NotNil(t, memberShipID)
		assert.Equal(t, memberShipType, "naver")
	})

	t.Run("멤버십 타입이 맞지 않으면 에러를 반환한다.", func(t *testing.T) {
		app := Application{}
		memberShipID, memberShipType, err := app.Create("김희선", "")
		assert.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("적절하지 않은 타입입니다."), err)
		assert.Empty(t, memberShipID)
		assert.Empty(t, memberShipType)
	})
}
