package main

import (
	"github.com/zhangdapeng520/zdpgo_test/libs/assert"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	// 创建断言对象
	assert := assert.New(t)

	// 创建表格
	var tests = []struct {
		input    int // 输入
		expected int // 期望输出
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	// 遍历表格数据
	for _, test := range tests {
		// 断言
		assert.Equal(Calculate(test.input), test.expected)
	}
}
