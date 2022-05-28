# zdpgo_testify
基于testify二次封装的测试框架

## 版本历史
- v0.1.0 2022/04/12  断言功能
- v0.1.1 2022/04/26  优化：代码结构调整
- v0.1.2 2022/05/28  优化：代码结构调整

## 使用案例
### 断言功能
```go
package main

import (
	""
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
```