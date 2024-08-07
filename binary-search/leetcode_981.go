package main

func main() {

}

type TimeMap struct {
	store map[string][]struct {
		value     string
		timestamp int
	}
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]struct {
			value     string
			timestamp int
		}),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key],
		struct {
			value     string
			timestamp int
		}{
			value,
			timestamp,
		},
	)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.store[key]; !ok {
		return ""
	} else {
		values := this.store[key]
		res := ""
		l, r := 0, len(values)-1
		for l <= r {
			mid := l + (r-l)/2
			if values[mid].timestamp <= timestamp {
				res = values[mid].value
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return res
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
