package alarm_ex1

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestService1(t *testing.T) {
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient1()
		service := Service{client}
		err := service.Send("01022334444")
		assert.Nil(t, err)
	})
}

func TestService2(t *testing.T) {
	t.Run("문자 전송에 성공한다.", func(t *testing.T) {
		client := NewMockSMSClient2()
		service := Service{client}

		receiver := "01022334444"
		client.On("Send", newSuccessSMSRequest(receiver)).
			Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.Nil(t, err)
	})
}
