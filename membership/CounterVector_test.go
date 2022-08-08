package membership

import (
	"testing"
)

func TestIncrementsAtAnEvenPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.incrementBy1(10)
	counter.incrementBy1(10)

	value := counter.get(10)
	if value != 2 {
		t.Fatalf("Expected value to be 2 but was %v", value)
	}
}

func TestDecrementsAtAnEvenPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.incrementBy1(10)
	counter.decrementBy1(10)

	value := counter.get(10)
	if value != 0 {
		t.Fatalf("Expected value to be 0 but was %v", value)
	}
}

func TestIncrementsAtAnOddPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.incrementBy1(11)
	counter.incrementBy1(11)

	value := counter.get(11)
	if value != 2 {
		t.Fatalf("Expected value to be 2 but was %v", value)
	}
}

func TestDecrementsAtAnOddPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.incrementBy1(11)
	counter.decrementBy1(11)

	value := counter.get(11)
	if value != 0 {
		t.Fatalf("Expected value to be 0 but was %v", value)
	}
}

func TestDecrementsAPresetValueAtAnOddPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter[3] = 240
	counter.decrementBy1(7)

	value := counter.get(7)
	if value != 14 {
		t.Fatalf("Expected value to be 14 but was %v", value)
	}
}

func TestDecrementsAPresetValueAtAnEvenPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter[4] = 14
	counter.decrementBy1(8)

	value := counter.get(8)
	if value != 13 {
		t.Fatalf("Expected value to be 13 but was %v", value)
	}
}

func TestDoesNotDecrementAtAnOddPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.decrementBy1(7)

	value := counter.get(7)
	if value != 0 {
		t.Fatalf("Expected value to be 0 but was %v", value)
	}
}

func TestDoesNotDecrementAtAnEvenPosition(t *testing.T) {
	var counter counterVector = make([]byte, 10)
	counter.decrementBy1(8)

	value := counter.get(8)
	if value != 0 {
		t.Fatalf("Expected value to be 0 but was %v", value)
	}
}
