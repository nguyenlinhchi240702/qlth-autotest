package api_test

import (
	"testing"

	"qlth-autotest/api_test/helper"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateAbsentRequest(t *testing.T) {
	// Gọi hàm login lấy token
	loginResp, err := helper.Login("0355123456", "123123", "PH")
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.AccessToken, "AccessToken không được để trống")

	t.Logf("Token nhận được: %s", loginResp.AccessToken)

	client := resty.New()

	// Gửi yêu cầu tạo đơn nghỉ phép
	createResp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", loginResp.AccessToken).
		SetBody(map[string]interface{}{
			"student_id":  182493,
			"semester_id": 7,
			"class_id":    43702,
			"grade_id":    2580,
			"requests": []map[string]interface{}{
				{
					"absent_date": "24-06-2025", // đổi sang định dạng chuẩn ISO yyyy-mm-dd
					"reason":      "Medical reason",
					"session_id":  1,
					"url":         "",
					"reason_code": "01",
				},
				{
					"absent_date": "22-06-2025", // đổi sang định dạng chuẩn ISO yyyy-mm-dd
					"reason":      "Medical reason",
					"session_id":  2,
					"url":         "",
					"reason_code": "01",
				},
			},
		}).
		Post("https://slldt.lms360.edu.vn/api/qlth/v3/students/absent_requests")

	if err != nil {
		t.Fatalf("Error when sending create absent request: %v", err)
	}

	t.Logf("Response status code: %d", createResp.StatusCode())
	t.Logf("Response body: %s", createResp.String())

	assert.Equal(t, 200, createResp.StatusCode(), "Expected status code 200")

}
