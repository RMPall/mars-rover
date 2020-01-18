package positioning

type Orientation int

const (
	N Orientation = iota
	E
	S
	W
)

var StringToOrientation = map[string]Orientation{
	"N": N,
	"E": E,
	"S": S,
	"W": W,
}

func (o Orientation) String() string {
	return [...]string{"N", "E", "S", "W"}[o]
}

// Rotate90DegreesLeft changes the direction of the robot towards left.
func (o Orientation) Rotate90DegreesLeft() Orientation {
	if o == N {
		return W
	}
	return o - 1
}

// Rotate90DegreesLeft changes the direction of the robot towards right.
func (o Orientation) Rotate90DegreesRight() Orientation {
	if o == W {
		return N
	}
	return o + 1
}
