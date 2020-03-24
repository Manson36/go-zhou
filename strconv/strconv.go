package strconv

import (
	"fmt"
	"strconv"
)

func main() {
	str := "110000"
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("ParseInt error:", err)
		return
	}

	fmt.Printf("%#v %T\n", ret, int(ret))

	i := 1000
	ret2 := strconv.FormatInt(int64(i), 10)
	fmt.Println(ret2)
	ret3 := strconv.Itoa(i)
	fmt.Println(ret3)
}
