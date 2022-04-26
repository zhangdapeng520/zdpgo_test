package main

import (
	"github.com/zhangdapeng520/zdpgo_test"
	"testing"
)

func TestCalculate(t *testing.T) {
	test := zdpgo_test.New()
	test.SetTestObj(t)

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
	for _, testData := range tests {
		// 断言
		test.Assert.Equal(Calculate(testData.input), testData.expected)
	}
}
