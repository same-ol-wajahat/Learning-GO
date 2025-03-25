package main

import "testing"

func TestGetAtRandomPosition(t *testing.T) {

	c := NewChessMatch()

	m1 := "g3"
	m2 := "Nf5"
	m3 := "d4"
	m4 := "f2"

	c.insertAt(0, m1)
	c.insertAt(1, m2)
	c.insertAt(2, m3)
	c.insertAt(3, m4)

	result := c.getAt(1)

	if result == nil || result.data != m2 {
		t.Errorf("expected %s but got %s", m2, result.data)
	}
}

func TestInsertAtRandomPosition(t *testing.T) {
	testCases := []struct {
		name     string
		index    int
		data     string
		expected string
	}{
		{
			name:     "first move",
			index:    0,
			data:     "g3",
			expected: "g3",
		},
		{
			name:     "Seocnd move",
			index:    1,
			data:     "f3",
			expected: "f3",
		},
		{
			name:     "Third move",
			index:    2,
			data:     "Nf2",
			expected: "Nf2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewChessMatch()
			c.insertAt(tc.index, tc.data)

			result := c.getAt(0)
			if tc.expected != result.data {
				t.Errorf("extpected: %s got: %s", tc.expected, result.data)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	testCases := []struct {
		name     string
		move     *Move
		expected string
	}{
		{
			name:     "first move",
			move:     NewMove("g3", nil),
			expected: "g3",
		},
		{
			name:     "Seocnd move",
			move:     NewMove("f3", nil),
			expected: "f3",
		},
		{
			name:     "Third move",
			move:     NewMove("Nf2", nil),
			expected: "Nf2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewChessMatch()
			c.prepend(tc.move)

			result := c.getAt(0)
			if tc.expected != result.data {
				t.Errorf("extpected: %s got: %s", tc.expected, result.data)
			}
		})
	}
}
