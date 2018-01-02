## 练习1：循环和函数
```
package main

import (
	"fmt"
	"math"
)
var z float64 = 1
func Sqrt(x float64) float64 {
	var old float64 = 0
	var dis float64 = z - old
	for dis > 0.005 || dis < -0.005 {
		old = z
		z = z - (z*z-x)/2*z
		fmt.Println(z)
		dis = z - old
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
```
