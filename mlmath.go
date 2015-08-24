package mlmath

import (
	"errors"
	"math"
)

type klabel struct {
	values []float64
	label  string
}

func Check_length(p *[]float64, q *[]float64) error {
	var err error = nil

	if len(*p) != len(*q) {
		err = errors.New("mymath: Input slices have not same length")
	}

	return err
}

func Euclide(p []float64, q []float64) float64 {

	res := 0.0

	for i, val := range p {
		res += (val - q[i]) * (val - q[i])
	}
	res = math.Sqrt(res)

	return res
}

func Manhatten(p []float64, q []float64) float64 {

	res := 0.0

	for i, val := range p {
		res += math.Abs(val - q[i])
	}

	return res
}

func KNN(in *[]float64, tdata *[]klabel) {

}
