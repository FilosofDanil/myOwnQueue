package internal

import (
	"myOwnQueue/pkg"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	// Arrange
	var q = pkg.NewQueue()

	//Act
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	expectedLength := 3

	//Assert
	if q.Length() != expectedLength {
		t.Errorf("Expected 3 but was %d", q.Length())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	// Arrange
	var q = pkg.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	//Act
	value, err := q.Dequeue()
	if err != nil {
		t.Errorf("Error in Dequeuing process")
	}
	expectedValue := 1

	//Assert
	if value != expectedValue {
		t.Errorf("Expected 4 but was %d", value)
	}
}

func TestQueue_Error(t *testing.T) {
	// Arrange
	var q = pkg.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	//Act
	for {
		if _, err := q.Dequeue(); err == nil {
			continue
		} else {
			break
		}
	}
	//Assert
	if val, err := q.Dequeue(); err == nil || val != 0 {
		t.Errorf("Expected error handling and value equal to zero!")
	}
}
