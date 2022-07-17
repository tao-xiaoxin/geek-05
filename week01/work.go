package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for i, v := range arr {
		switch v {
		case "stupid":
			arr[i] = "smart"
		case "weak":
			arr[i] = "strong"
		default:
			break
		}
	}
	fmt.Println(arr)
}
