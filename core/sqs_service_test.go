package lib

import (
	"testing"

	// "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Must prefix tests with "Test" otherwise will not run

// Unexpected parameter for SQS Delete
func TestNilDelete(t *testing.T) {
	s := &SQSService{} // pointer to SQSService Struct
	m := &Message{}

	r := s.Delete(m)

	assert.Error(t, r, "Unexpected input: SQS message without a receipt handle")
}

// func TestSQSServiceCreate(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// NewMockQueueService(ctrl)

// 	q := getMockSQSClient()
// 	queueURL := "https://queue.amazonaws.com/1111111/myqueue"

// 	fmt.Printf("123")

// 	a := "Hello"
// 	b := "Hello"

// 	assert.Equal(t, a, b, "Should be the same")
// }