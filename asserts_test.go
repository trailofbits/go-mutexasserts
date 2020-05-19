package mutexasserts

import (
	"sync"
	"testing"
)

func init() {
	exit = func(c int) { panic("fatal") }
}

func TestMutexLocked(t *testing.T) {
	m := sync.Mutex{}
	if MutexLocked(&m) {
		t.Errorf("MutexLocked while it shouldn't be")
	}
	m.Lock()
	if !MutexLocked(&m) {
		t.Errorf("MutexLocked not while it should  be")
	}
	m.Unlock()
	if MutexLocked(&m) {
		t.Errorf("MutexLocked while it shouldn't be")
	}
}

func TestRWMutexLocked(t *testing.T) {
	rw := sync.RWMutex{}
	if RWMutexLocked(&rw) || RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.Lock()
	if !RWMutexLocked(&rw) || RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.Unlock()
	if RWMutexLocked(&rw) || RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.RLock()
	if RWMutexLocked(&rw) || !RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.RLock()
	if RWMutexLocked(&rw) || !RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.RUnlock()
	if RWMutexLocked(&rw) || !RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
	rw.RUnlock()
	if RWMutexLocked(&rw) || RWMutexRLocked(&rw) {
		t.Errorf("RWMutex test failed")
	}
}

func TestAssertMutexLocked_MustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m := sync.Mutex{}
	AssertMutexLocked(&m)
}

func TestAssertMutexLocked_MustNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()
	m := sync.Mutex{}
	m.Lock()
	AssertMutexLocked(&m)
}

func TestAssertRWMutexLocked_MustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m := sync.RWMutex{}
	AssertRWMutexLocked(&m)
}

func TestAssertRWMutexLocked_MustNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()
	m := sync.RWMutex{}
	m.Lock()
	AssertRWMutexLocked(&m)
}

func TestAssertRWMutexRLocked_MustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m := sync.RWMutex{}
	AssertRWMutexRLocked(&m)
}

func TestAssertRWMutexRLocked_MustNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked")
		}
	}()
	m := sync.RWMutex{}
	m.RLock()
	AssertRWMutexRLocked(&m)
}
