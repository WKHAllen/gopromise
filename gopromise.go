package gopromise

// Promise represents a pending task
func Promise(callback func() interface{}) <-chan interface{} {
	p := make(chan interface{})

	go func() {
		defer close(p)
		p <- callback()
	}()

	return p
}

// All will run multiple pending tasks
func All(callbacks ...func() interface{}) <-chan interface{} {
	p := make(chan interface{})

	go func() {
		defer close(p)

		remaining := len(callbacks)
		retVals := make([]interface{}, len(callbacks))

		for idx, callback := range callbacks {
			go func(index int, callbackFunc func() interface{}) {
				retVals[index] = callbackFunc()
				remaining--
			}(idx, callback)
		}

		for remaining > 0 {}

		for _, retVal := range retVals {
			p <- retVal
		}
	}()

	return p
}

// Race will run multiple tasks and return the first one completed
func Race(callbacks ...func() interface{}) <-chan interface{} {
	switch len(callbacks) {
		case 0: return nil
		case 1: return Promise(callbacks[0])
		default:
			return Promise(func() interface{} {
				var result interface{}
	
				select {
					case result = <-Promise(callbacks[0]):
					case result = <-Race(callbacks[1:]...):
				}
	
				return result
			})
	}
}
