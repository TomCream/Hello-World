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

### 练习2：slice
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

### 练习3：map
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

### 练习4 斐波那契闭包
```
package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		result := i+j
		i = j
		j = result
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
#### result
![fibonacci](../img/fibonacci.png)

### 练习5 Stringers
```
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
```
#### resutl
![stringers](../img/stringers.png)
