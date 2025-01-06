package sender

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestEmail_Send(t *testing.T) {
	email, err := NewEmail("860808187@qq.com", "fyudafdzqmhwbfbd", proto.String(`{"name":"上海","host":"smtp.qq.com","port":25}`))
	assert.NoError(t, err)
	err = email.Send("1280291001@qq.com", "测试", "测试内容")
	assert.NoError(t, err)
}
