package zdpgo_test

import (
	"github.com/zhangdapeng520/zdpgo_test/libs/assert"
	"testing"
)

type Test struct {
	Assert *assert.Assertions
	Config *Config
}

// New 新建测试对象
func New() *Test {
	return NewWirthConfig(Config{})
}

// NewWirthConfig 根据配置新建测试对象
func NewWirthConfig(config Config) *Test {
	test := Test{}

	test.Config = &config
	test.Assert = assert.New(config.TestObj)

	return &test
}

// SetTestObj 测试测试对象
func (t *Test) SetTestObj(t1 *testing.T) {
	t.Config.TestObj = t1
	t.Assert = assert.New(t1)
}
