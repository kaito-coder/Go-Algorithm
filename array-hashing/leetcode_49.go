package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

}
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	res := make([][]string, 0)
	count := make(map[[26]byte]int)
	for _, str := range strs {
		var key [26]byte
		for _, s := range str {
			key[s-'a']++
		}
		if _, ok := count[key]; ok {
			res[count[key]] = append(res[count[key]], str)
		} else {
			count[key] = len(res)
			res = append(res, []string{str})
		}
	}
	return res

}
