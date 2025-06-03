package api_test

import (
	"fmt"
	"testing"

	"qlth-autotest/api_test/helper"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestRecallAbsentRequest(t *testing.T) {
	loginResp, err := helper.Login("0355123456", "123123", "PH")
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.AccessToken, "AccessToken không được để trống")

	t.Logf("Token nhận được: %s", loginResp.AccessToken)

	client := resty.New()
	absentID := 2054
	url := fmt.Sprintf("https://slldt.lms360.edu.vn/api/qlth/v3/students/absent_requests/recall/%d", absentID)

	createResp, err := client.R().
		SetHeader("Authorization", loginResp.AccessToken).
		Delete(url)

	if err != nil {
		t.Fatalf("Error when sending recall absent request: %v", err)
	}

	t.Logf("Response status code: %d", createResp.StatusCode())
	t.Logf("Response body: %s", createResp.String())

	assert.Equal(t, 200, createResp.StatusCode(), "Expected status code 200")
}
