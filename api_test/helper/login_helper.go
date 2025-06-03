package helper

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

type LoginResponse struct {
	AccessToken string
	UserID      int
	Role        string
}

// Login is a helper function to authenticate and return access token
func Login(username, password, accountType string) (*LoginResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username":     username,
			"password":     password,
			"account_type": accountType,
		}).
		Post("https://slldt.lms360.edu.vn/api/qlth/v3/login")

	if err != nil {
		return nil, fmt.Errorf("error sending login request: %v", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d - %s", resp.StatusCode(), resp.String())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid login response structure")
	}

	token, _ := data["access_token"].(string)
	userIDFloat, _ := data["user_id"].(float64)
	role, _ := data["role"].(string)

	if token == "" {
		return nil, fmt.Errorf("access_token is empty")
	}

	return &LoginResponse{
		AccessToken: token,
		UserID:      int(userIDFloat),
		Role:        role,
	}, nil
}
