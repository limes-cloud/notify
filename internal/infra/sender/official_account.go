package sender

import (
	"encoding/json"
	"errors"
)

// NewOfficialAccount 初始化公众号发送
func NewOfficialAccount(ak string, sk string, extra *string) (*Email, error) {
	if extra == nil {
		return nil, errors.New("extra info in empty")
	}
	email := &Email{User: ak, Password: sk}
	if err := json.Unmarshal([]byte(*extra), email); err != nil {
		return nil, err
	}
	return email, nil
}
