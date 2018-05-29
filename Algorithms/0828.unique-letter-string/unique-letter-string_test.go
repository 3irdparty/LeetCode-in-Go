package problem0828

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	ans int
}{

	{
		"ABC",
		10,
	},

	{
		"ABA",
		8,
	},

	// 可以有多个 testcase
}

func Test_uniqueLetterString(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, uniqueLetterString(tc.S), "输入:%v", tc)
	}
}

func Benchmark_uniqueLetterString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniqueLetterString(tc.S)
		}
	}
}
