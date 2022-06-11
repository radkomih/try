package try

import "time"

type WaitingStrategy interface {
	Wait()
}

// 	Incremental
// 	Random
// 	Exponential

type ConstantWait struct {
	Interval uint8
}

func (w ConstantWait) Wait() {
	time.Sleep(time.Duration(w.Interval) * time.Second)
}

type Try struct {
	Tries           uint8
	Strategy        WaitingStrategy
	Errors          []error
	FailureCallback func()
}

func (t Try) Execute(callback func()) {
	defer func() {
		if r := recover(); r != nil {
			t.FailureCallback()

			t.Tries--

			if t.Tries > 0 {
				t.Strategy.Wait()
				t.Execute(callback)
			} else {
				callback()
			}
		}
	}()

	callback()
}
