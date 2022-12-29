package util

import (
	"math"
	"strings"
)

// 输出一个int32位对应的二进制表示
func BinaryFormart(n int64) string {
	sb := strings.Builder{}
	c := int64(math.Pow(2, 31)) //最高位上是1，其他位全是0.这里的int是64位
	for i := 0; i < 32; i++ {
		if n&c != 0 { //判断n的当前位是否为1
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 // “1”往右移动一位
	}

	return sb.String()

}
