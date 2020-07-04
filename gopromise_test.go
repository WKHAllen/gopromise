package gopromise

import (
	"fmt"
	"testing"
	"time"
)

func assert(value bool, t *testing.T, err string) {
	if !value {
		t.Errorf(err)
	}
}

func assertResult(testName string, expected interface{}, observed interface{}, t *testing.T) {
	assert(expected == observed, t, fmt.Sprintf("%v result incorrect; expected %v, got %v", testName, expected, observed))
}

func alertSuccess(testName string) {
	fmt.Println("SUCCESS:", testName)
}

func alertFailure(testName string) {
	fmt.Println("FAILURE:", testName)
}

func alertStatus(testName string, t *testing.T) {
	if !t.Failed() {
		alertSuccess(testName)
	} else {
		alertFailure(testName)
	}
}

func doSomething() interface{} {
	time.Sleep(time.Second)
	return 3
}

func funA() interface{} {
	time.Sleep(time.Millisecond * 300)
	return 1
}

func funB() interface{} {
	time.Sleep(time.Millisecond * 200)
	return 3
}

func funC() interface{} {
	time.Sleep(time.Millisecond * 100)
	return 7
}

func TestPromise(t *testing.T) {
	p := Promise(doSomething)
	res := (<-p).(int)
	assertResult("Promise", 3, res, t)
	alertStatus("Promise", t)
}

func TestPromiseAll(t *testing.T) {
	p := All(funA, funB, funC)
	resA := (<-p).(int)
	resB := (<-p).(int)
	resC := (<-p).(int)
	assertResult("PromiseAll", 1, resA, t)
	assertResult("PromiseAll", 3, resB, t)
	assertResult("PromiseAll", 7, resC, t)
	alertStatus("PromiseAll", t)
}

func TestPromiseRace(t *testing.T) {
	p := Race(funA, funB, funC)
	fastest := (<-p).(int)
	assertResult("PromiseRace", 7, fastest, t)
	alertStatus("PromiseRace", t)
}
