package api_test

import (
	"testing"

	resty "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"qlth-autotest/api_test/helper" // đường dẫn đúng đến hàm Login()
)

func TestCreateNotification(t *testing.T) {
	// Gọi hàm login lấy token
	loginResp, err := helper.Login("0933554421", "test@123!", "GV")
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.AccessToken, "AccessToken không được để trống")

	t.Logf("Token nhận được: %s", loginResp.AccessToken)

	client := resty.New()

	// Gửi request tạo notification kèm token trong header Authorization
	createResp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", loginResp.AccessToken).
		SetBody(map[string]interface{}{
			"noti_type":       2,
			"title_noti":      "Thông báo kiểm tra auto test",
			"file_url":        "",
			"status":          "",
			"type_received":   3,
			"total_recipient": 1,
			"mode":            2,
			"message":         "Đây là nội dung thông báo thử nghiệm",
			"recipient_ids":   []int{1184813},
			"semester_id":     7,
		}).
		Post("https://slldt.lms360.edu.vn/api/qlth/v3/notifications")

	assert.NoError(t, err)
	assert.Equal(t, 200, createResp.StatusCode())

	t.Log("Response Body:", createResp.String())
}
