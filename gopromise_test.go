package gopromise

import (
	"fmt"
	"testing"
	"time"
)

func doSomething() interface{} {
	seconds := 3
	time.Sleep(time.Second * time.Duration(seconds))
	return seconds
}

func funA() interface{} {
	time.Sleep(time.Second * 3)
	return 1
}

func funB() interface{} {
	time.Sleep(time.Second * 2)
	return 3
}

func funC() interface{} {
	time.Sleep(time.Second * 1)
	return 7
}

func TestPromise(t *testing.T) {
	p := Promise(doSomething)
	fmt.Println("Awaiting...")
	res := (<-p).(int)
	res *= 3
	fmt.Println("Done. Result:", res)
}

func TestPromiseAll(t *testing.T) {
	p := All(funA, funB, funC)
	fmt.Println("Awaiting...")
	resA := (<-p).(int)
	resB := (<-p).(int)
	resC := (<-p).(int)
	fmt.Println("Done. Results:", resA, resB, resC)
}

func TestPromiseRace(t *testing.T) {
	p := Race(funA, funB, funC)
	fmt.Println("Awaiting...")
	fastest := (<-p).(int)
	fmt.Println("Done. Fastest:", fastest)
}
