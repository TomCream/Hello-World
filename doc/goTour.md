### 练习1：循环和函数
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

### 练习：slice
```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8,dy)
	for x,_ := range result {
		result[x] = make([]uint8,dx)
		for y,_ := range result[x]{
			//intValue := (x+y)/2
			//intValue := x*y
			//intValue := x^y
			intValue := 2*x+3*y
			result[x][y] = uint8(intValue)
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
```
#### result
![slicePic](../img/slicePic.png)

### 练习：map
```
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	stringList := strings.Fields(s)
	for _,value := range stringList {
		v,flag :=  result[value]
		if flag {
			result[value] = v+1
		} else {
			result[value] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
```
#### result
![mapPic](../img/mapPic.png)

