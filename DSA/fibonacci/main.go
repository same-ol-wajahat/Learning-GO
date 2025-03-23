package main

// fibonacci sequence example
// [0,1,1,2,3,5,8]

func fibOne(position int) int {
	if position < 2 {
		return position
	}

	seq := []int{1, 0}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}
	return seq[len(seq)-1] + seq[len(seq)-2]
}

func fibTwo(position int) int {
	if position < 2 {
		return position
	}
	a, b := fibTwo(position-1), fibTwo(position-2)
	return a + b
}
