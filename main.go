package main 

import (
	"fmt"
	"modules-go/utils"
)

func main() {
	fmt.Println(utils.ReverseWords("Hello World"))
	fmt.Println(utils.Round(3.14159, 2))
	fmt.Println(utils.CalculateAverage([]float64{1, 2, 3, 4, 5}))
}
