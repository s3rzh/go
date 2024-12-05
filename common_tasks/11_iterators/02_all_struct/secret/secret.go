package secret

type Int struct {
	Val     int
	Visible bool
}

type Ints []Int

func (ii Ints) All(yield func(Int) bool) {
	for _, n := range ii {
		if !yield(n) {
			return
		}
	}
}
