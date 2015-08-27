package mlmath

import (
	"errors"
	"math"
)

type Klabel struct {
	Values []float64
	Label  string
}

func Check_length(p []float64, q []float64) error {
	var err error = nil

	if len(p) != len(q) {
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

func KNN(in []float64, tdata []Klabel, k int) []Klabel {
	k_res := make([]Klabel, k)
	var temp_res Klabel

	for _, tvalues := range tdata {
		temp_val := Manhatten(in, tvalues.Values)
		temp_res.Values = []float64{temp_val}
		temp_res.Label = tvalues.Label
		k_res = append(k_res, temp_res)
	}

	return k_res
}
