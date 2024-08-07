package main

func main() {
	s := "leet**cod*e"
	print(removeStars(s))
}
func removeStars(s string) string {
	stack := []rune{}
	for _, v := range s {
		if v != '*' {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}