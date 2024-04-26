package main

import (
	"fmt"
)

/*
Reverse Shuffle Merge

https://www.hackerrank.com/challenges/reverse-shuffle-merge/problem
*/
// Complete the reverseShuffleMerge function below.
func reverseShuffleMerge(s string) string {
	length := len(s)
	count := make(map[byte]int)
	body := []byte(s)
	for i := 0; i < length/2; i = i + 1 {
		j := length - i - 1
		count[body[i]]++
		count[body[j]]++
		t := body[j]
		body[j] = body[i]
		body[i] = t
	}
	need := make(map[byte]int)
	for key, v := range count {
		need[key] = v / 2
	}

	var res []byte
	fmt.Println("reverse string=", string(body))
	for i, v := range body {
		fmt.Println(i, "v=", string(v), "; res=", count[v], need[v], string(res))
		count[v]--
		if need[v] == 0 {
			continue
		}
		need[v]--
		for {
			num := len(res)
			if num > 0 && res[num-1] > v && count[res[num-1]] > need[res[num-1]] {
				fmt.Println("### ", i, "v=", string(v), "; res=", need[res[num-1]], string(res))
				need[res[num-1]]++
				res = res[:num-1]
			} else {
				break
			}
		}
		res = append(res, v)
	}

	//fmt.Println("res=", string(res))
	//fmt.Println("reverse string=", string(body))

	return string(res)
}

func merge() {

	s := "eggegg"

	fmt.Printf("s=%s\n", s)
	result := reverseShuffleMerge(s)
	fmt.Printf("%s\n", result)

}
