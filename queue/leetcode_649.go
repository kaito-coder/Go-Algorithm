package main

func main() {

}
func predictPartyVictory(senate string) string {
	radiantQueue := []int{}
	direQueue := []int{}

	for i, s := range senate {
		if s == 'R' {
			radiantQueue = append(radiantQueue, i)
		} else {
			direQueue = append(direQueue, i)
		}
	}
	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		r := radiantQueue[0]
		d := radiantQueue[0]
		radiantQueue = radiantQueue[1:]
		direQueue = radiantQueue[1:]
		if r < d {
			radiantQueue = append(radiantQueue, r+len(senate))
		} else {
			direQueue = append(direQueue, d+len(senate))
		}
	}
	if len(radiantQueue) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}

}
