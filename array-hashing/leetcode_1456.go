package main

func maxVowels(s string, k int) int {
	sum := 0
	isVowel := func(s byte) bool {
		return s == 'u' || s == 'e' || s == 'o' || s == 'a' || s == 'i'
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if i >= k && isVowel(s[i-k]) == true {
			count--
		}
		sum = max(sum, count)
	}
	return sum
}
func main() {

}
