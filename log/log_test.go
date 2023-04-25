package log

import "testing"

func TestLog(t *testing.T) {
	subject := NewLog[int]()

	if len(subject) != 0 {
		t.Errorf("subject size was not 0 but %d", len(subject))
	}

	subject = subject.Append(1)
	subject = subject.Append(2)

	if len(subject) != 2 {
		t.Errorf("subject size was not 2 but %d", len(subject))
	}

	it, subject := subject.Take()

	if it != 1 {
		t.Errorf("it was not 1 but %d", it)
	}

	if len(subject) != 1 {
		t.Errorf("subject size was not 1 but %d", len(subject))
	}

	it, subject = subject.Take()

	if it != 2 {
		t.Errorf("it was not 2 but %d", it)
	}

	if len(subject) != 0 {
		t.Errorf("subject size was not 0 but %d", len(subject))
	}
}
