package main

import "fmt"

// func CanJump(arr []uint) bool {
//     if len(arr) == 0 {
//         return false
//     }
    
//     maxReach := 0
//     for i := 0; i <= maxReach && i < len(arr); i++ {
//         if i+int(arr[i]) > maxReach {
//             maxReach = i + int(arr[i])
//         }
//         if maxReach >= len(arr)-1 {
//             return true
//         }
//     }
//     return false
// }

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}

	maxReach := 0
	for i := 0; i <= maxReach && i < len(arr); i++ {
		if int(arr[i]) + i > maxReach {
			maxReach = int(arr[i]) + i
		}
		if maxReach >= len(arr)-1 {
			return true
		}
	}
	return false
}

func main() {
	input1 := []uint{0, 0, 0, 0}
	fmt.Println(CanJump(input1))

	input2 := []uint{5, 0, 0, 0, 0}
	fmt.Println(CanJump(input2))

	input3 := []uint{1, 1, 0}
	fmt.Println(CanJump(input3))
}