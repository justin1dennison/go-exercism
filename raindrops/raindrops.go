package raindrops

import "fmt"

type Factor struct {
	label string
	value int
}
var factors = []Factor{
	Factor{ label: "Pling", value: 3}, 
	Factor{ label: "Plang", value: 5}, 
	Factor{ label: "Plong", value: 7},
}


func Convert(num int) string {
	result := ""
	for _, factor := range factors {
		if num % factor.value == 0 {
			result += factor.label
		}	
	}
	if result == ""{
		return fmt.Sprintf("%d", num)
	}
	return result
	
}


