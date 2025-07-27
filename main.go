package main

import "fmt"

var sliceA = []int{1, 2, 4, 3, 4, 4, 5}
var sliceB = []int{3, 5, 6, 6, 7, 8}

func main() {
	uniqA, uniqB := task1()

	fmt.Println(uniqA, uniqB)

	uniq2 := task2(uniqA, uniqB)

	fmt.Println(uniq2)

	uniq3 := task3(uniqA, uniqB)

	fmt.Println(uniq3)


}

func task1() ([]int, []int) {
	sliceUA := make([]int, 0)
	sliceUB := make([]int, 0)

	for _, val := range sliceA {
		var same bool
		for _, val1 := range sliceUA {
			if val1 == val {
				same = true
			}
		}

		if same == false {
			sliceUA = append(sliceUA, val)
		}
	}
	

	for _, val2 := range sliceB {
		var same2 bool
		for _, val3 := range sliceUB {
			if val3 == val2 {
				same2 = true
			}
		}

		if same2 == false {
			sliceUB = append(sliceUB, val2)
		}
	}
	

	return sliceUA, sliceUB
}


func task2(uniqA, uniqB []int) []int {
	sliceU2 := make([]int, 0)

	for _, val1 := range uniqA {
		
		for _, val2 := range uniqB {
			if val1 == val2 {
				sliceU2 = append(sliceU2, val1)
			}
		}
	}
	return sliceU2
}

func task3(unicA, unicB []int) []int {
	sliceU3 := make([]int, 0)

	for _, val := range unicA {
		sliceU3 = append(sliceU3, val)
	}

	for _, val1 := range unicB {
		var opsch bool
		for _, val2 := range unicA {
			if val2 == val1 {
				opsch = true
			}
		}

		if opsch == false {
			sliceU3 = append(sliceU3, val1)
		}
	}

	return sliceU3
}
