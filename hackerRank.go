package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	var sorted bool
	fmt.Scan(&n)
	lastUnsortedIndex := n - 1
	list := make([][]byte, n)
	for i := 0; i < n; i++ {
		list[i] = make([]byte, 1, 1000000)
		fmt.Scanf("%s", &list[i])
	}
	for !sorted {
		sorted = true
		for i := 0; i < lastUnsortedIndex; i++ {
			if shouldSwap(list[i], list[i+1]) {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		lastUnsortedIndex--
	}
	for i := 0; i < n; i++ {
		fmt.Println(string(list[i]))
	}
}
func shouldSwap(v1 []byte, v2 []byte) bool {
	n1, n2 := len(v1), len(v2)
	if n1 > n2 {
		return true
	} else if n2 > n1 {
		return false
	} else {
		for i := 0; i < n1; i++ {
			if v1[i] > v2[i] {
				return true
			} else if v2[i] > v1[i] {
				return false
			}
		}
		return false
	}
}
