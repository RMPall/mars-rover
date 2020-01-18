package layout

const (
	MaximumCoordinateValue = 50
)

type Coordinate struct {
	X, Y int
}

func NewLayoutCoordinate(x, y int) *Coordinate {
	if x > MaximumCoordinateValue || y > MaximumCoordinateValue {
		return nil
	}

	return &Coordinate{
		X: x,
		Y: y,
	}
}
