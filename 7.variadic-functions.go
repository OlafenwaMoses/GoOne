package main
import (
	"fmt"
)

func main(){
	values := []int{10, 21, 82, 45, 21, 34, 12, 44, 44}
	result := varying_input(values...)
	fmt.Println(result)
}

func varying_input(values ...int) int{
	result := 0
	for _, num := range values{
		result += num
	}

	return result

}