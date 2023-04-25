package heap

import "testing"

func TestHeap(t *testing.T) {
	subject := NewHeap[int]()
	subject = subject.Push(1)
	subject = subject.Push(2)

	if len(subject) != 2 {
		t.Errorf("len of subject was not 2 but %d", len(subject))
	}

	result, subject := subject.Pop()

	if result != 2 {
		t.Errorf("result was not 2 but %d", result)
	}

	if len(subject) != 1 {
		t.Errorf("len of subject was not 1 but %d", len(subject))
	}

	result, subject = subject.Pop()

	if result != 1 {
		t.Errorf("result was not 1 but %d", result)
	}

	if len(subject) != 0 {
		t.Errorf("len of subject was not 0 but %d", len(subject))
	}
}

func TestAddToEmptyHeap(t *testing.T) {
	subject := NewHeap[int]()
	subject = subject.Push(1)

	_, subject = subject.Pop()

	subject = subject.Push(2)

	if len(subject) != 1 {
		t.Errorf("len of subject was not 1 but %d", len(subject))
	}

	var result int
	result, subject = subject.Pop()

	if result != 2 {
		t.Errorf("result was not 2 but %d", result)
	}

	if len(subject) != 0 {
		t.Errorf("len of subject was not 0 but %d", len(subject))
	}
}

func TestPopInALoop(t *testing.T) {
	subject := NewHeap[int]()
	subject = subject.Push(1)
	subject = subject.Push(2)

	var it int
	for {
		it, subject = subject.Pop()

		if it == 0 {
			t.Error("it was 0")
		}

		if len(subject) == 0 {
			break
		}
	}

	if len(subject) != 0 {
		t.Errorf("len of subject was not 0 but %d", len(subject))
	}
}
