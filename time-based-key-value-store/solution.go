package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))
}

type TimeMap struct {
	store map[string][]ValStamp // key : list of [val, timestamp]
}

type ValStamp struct {
	Val  string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]ValStamp)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
    if _, ok := tm.store[key]; !ok {
        tm.store[key] = make([]ValStamp, 0)
    }
    tm.store[key] = append(tm.store[key], ValStamp{value, timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
    var res string
    var values []ValStamp
    if _, ok := tm.store[key]; ok {
        values = tm.store[key]
    }
    l, r := 0, len(values)-1
    for l <= r {
        mid := l + (r-l)/2
        if values[mid].Time <= timestamp {
            res = values[mid].Val
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */