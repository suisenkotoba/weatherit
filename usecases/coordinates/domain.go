package coordinate

type Coordinate struct {
	Lat  float64
	Long float64
}

func CreateCoordinate(point [2]float64) Coordinate {
	return Coordinate{
		Lat:  point[0],
		Long: point[1],
	}
}

func (c *Coordinate) ToPoint() [2]float64 {
	return [2]float64{c.Lat, c.Long}
}
