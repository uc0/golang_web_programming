package alarm_ex1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	t.Run("TooManyRequest 계속 발생하는 경우, 최대 retry count만큼 재시도한다.", func(t *testing.T) {
		maxRetry := 3
		client := NewMockSMSClient()
		service := Service{client, maxRetry}

		receiver := "01022334444"
		//client.On("Send", newSuccessSMSRequest(receiver)).
		//	Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.ErrorIs(t, err, SMSFailErr)
		client.AssertNumberOfCalls(t, "Send", 3)
	})

	t.Run("TooManyRequest 발생할 때마다 Retry하며, 도중에 성공할 경우 재시도하지 않는다..", func(t *testing.T) {
		maxRetry := 3
		client := NewMockSMSClient()
		service := Service{client, maxRetry}

		receiver := "01022334444"
		//client.On("Send", newSuccessSMSRequest(receiver)).
		//	Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.Nil(t, err)
		client.AssertNumberOfCalls(t, "Send", 2)
	})

	t.Run("client에서 internal server 에러가 발생한 경우, 재시도하지 않는다.", func(t *testing.T) {
		maxRetry := 3
		client := NewMockSMSClient()
		service := Service{client, maxRetry}

		receiver := "01022334444"
		//client.On("Send", newSuccessSMSRequest(receiver)).
		//	Return(SMSResponse{http.StatusOK, "ok"}, nil)

		err := service.Send(receiver)
		assert.ErrorIs(t, err, SMSFailErr)
		client.AssertNumberOfCalls(t, "Send", 1)
	})
}
