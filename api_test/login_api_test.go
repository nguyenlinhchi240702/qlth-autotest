package api_test

import (
	"qlth-autotest/api_test/helper"
	"testing"

	// ← chính là import thư mục helpers

	"github.com/stretchr/testify/assert"
)

func TestLoginAPI(t *testing.T) {
	resp, err := helper.Login("0933554421", "test@123!", "GV")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.AccessToken)

	t.Log("Access token:", resp.AccessToken)
}
