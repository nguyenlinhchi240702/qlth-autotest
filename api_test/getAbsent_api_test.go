package api_test

import (
	"fmt"
	"testing"

	"qlth-autotest/api_test/helper"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAbsentRequestHistory(t *testing.T) {
	// Step 1: Login để lấy token
	loginResp, err := helper.Login("0355123456", "123123", "PH")
	//loginResp: là struct hoặc đối tượng chứa dữ liệu trả về sau khi đăng nhập thành công
	//err — giá trị lỗi (nếu có) trong quá trình gọi hàm (ví dụ lỗi mạng, sai mật khẩu, lỗi phân tích dữ liệu phản hồi...)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.AccessToken, "AccessToken không được để trốngg")

	t.Logf("Token nhận được: %s", loginResp.AccessToken)

	client := resty.New()

	// Step 2: Gọi API lấy lịch sử đơn xin nghỉ
	studentID := 182493
	semesterID := 7
	classID := 43702

	url := fmt.Sprintf("https://slldt.lms360.edu.vn/api/qlth/v3/students/%d/absent_requests?semester_id=%d&class_id=%d", studentID, semesterID, classID)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", loginResp.AccessToken). // Không dùng SetAuthToken nếu API không dùng "Bearer"
		Get(url)

	t.Logf("Response status code: %d", resp.StatusCode())
	t.Logf("Response body: %s", resp.String())

	assert.NoError(t, err, "Có lỗi xảy ra khi gửi yêu cầu")
	assert.Equal(t, 200, resp.StatusCode(), "Expected status code 200")
	assert.Contains(t, resp.String(), `"records"`, "Response không chứa trường 'records'")
}
