package main

// Just some arbitrary interface so we can later use an interface type.
type Thing interface {
	get() int
}

// Two types implementing the interface, with fields of very different types.
type Int struct {
	val int
}

func (s *Int) get() int {
	return s.val
}

type Ptr struct {
	val *int
}

func (s *Ptr) get() int {
	return *s.val
}

// A global variable of interface type, that we will swap back and
// forth between pointing to an `Int` and to a `Ptr`.
var globalVar Thing = &Int{val: 42}

// var rwMutex sync.RWMutex

// Repeatedly invoke the interface method on the global variable.
func repeatGet() {
	for {
		// rwMutex.RLock()
		x := globalVar
		// rwMutex.RUnlock()
		x.get()
	}
}

// Repeatedly change the dynamic type of the global variable.
func repeatSwap() {
	myval := 0
	for {
		// rwMutex.Lock()
		globalVar = &Ptr{val: &myval}
		// rwMutex.Unlock()
		// rwMutex.Lock()
		globalVar = &Int{val: 42}
		// rwMutex.Unlock()
	}
}

func main() {
	go repeatGet()
	repeatSwap()
}
