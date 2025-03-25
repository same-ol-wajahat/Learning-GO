package main

type Move struct {
	data string
	next *Move
}

func NewMove(data string, next *Move) *Move {
	return &Move{
		data: data,
		next: next,
	}
}

type ChessMatch struct {
	head *Move
}

func NewChessMatch() *ChessMatch {
	return &ChessMatch{}
}

func (c *ChessMatch) getAt(index int) *Move {
	crt := c.head
	pos := 0

	if pos < index && crt != nil {
		crt = crt.next
		pos++
	}
	return crt
}

func (c *ChessMatch) insertAt(index int, data string) {
	if c.head == nil {
		c.head = NewMove(data, nil)
		return
	}

	prev := c.getAt(index - 1)
	prev.next = NewMove(data, prev.next)
}

func (c *ChessMatch) prepend(m *Move) {
	second := c.head
	c.head = m
	c.head.next = second
}
