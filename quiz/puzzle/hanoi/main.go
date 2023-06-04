package main

import "fmt"

func hanoi(disk int, src, dest, support string) {
	if disk < 1 {
		return
	}

	hanoi(disk-1, src, support, dest)
	fmt.Printf("move %d from %s to %s\n", disk, src, dest)
	hanoi(disk-1, support, dest, src)
}

func getHanoiMovement(disk int, src, dest, support string) [][]string {
	result := make([][]string, 0)

	var _hanoi func(disk int, src, dest, support string)
	_hanoi = func(disk int, src, dest, support string) {
		if disk < 1 {
			return
		}
		_hanoi(disk-1, src, support, dest)
		result = append(result, []string{fmt.Sprint(disk), src, dest})
		_hanoi(disk-1, support, dest, src)
	}
	_hanoi(disk, src, dest, support)
	return result
}

func main() {
	hanoi(3, "A", "C", "B")

	ans := getHanoiMovement(4, "A", "C", "B")
	fmt.Println(ans)
}
