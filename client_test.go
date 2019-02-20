package sendpulse

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestResponseErrorData(t *testing.T) {
	e := SendpulseError{http.StatusInternalServerError, "http://test.com", "Something went wrong", "Test message"}
	assert.Equal(t, fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Message), e.Error())
}
