package mlmath

import (
	"errors"
	"math"
)

type Klabel struct {
	Values []float64
	Label  string
}

// Bubblesort for Klabel and repetitive sorting for small k
func KlabelBsort(a []Klabel) []Klabel {
	a_len := len(a)
	a_temp := make([]Klabel, a_len)

	copy(a_temp, a)

	for i, _ := range a {
		if i+1 != a_len {
			if a[i+1].Values[0] < a[i].Values[0] {
				a[i+1].Values[0], a[i].Values[0] = a[i].Values[0], a[i+1].Values[0]
			}
			copy(a_temp, a)

			for i, _ := range a_temp {
				if a_temp[i].Values[0] != a[i].Values[0] {
					KlabelBsort(a)
				}
			}
		}
	}

	return a
}

func Check_length(p []float64, q []float64) error {
	var err error = nil

	if len(p) != len(q) {
		err = errors.New("mymath: Input slices have not same length")
	}

	return err
}

// Euclidian distance, prone to runaways
func Euclide(p []float64, q []float64) float64 {

	res := 0.0

	for i, val := range p {
		res += (val - q[i]) * (val - q[i])
	}
	res = math.Sqrt(res)

	return res
}

//Manhatten metric
func Manhatten(p []float64, q []float64) float64 {

	res := 0.0

	for i, val := range p {
		res += math.Abs(val - q[i])
	}

	return res
}

//kNN
//calculate distance and store results in slice of size k
//if new value is smaller than biggest value in slice, replace
func KNN(in []float64, tdata []Klabel, k int) []Klabel {
	k_res := make([]Klabel, 0)
	var temp_res Klabel

	for _, tvalues := range tdata {
		temp_val := Manhatten(in, tvalues.Values)
		temp_res.Values = []float64{temp_val}
		temp_res.Label = tvalues.Label
		if len(k_res) < k {
			k_res = append(k_res, temp_res)
		} else {
			k_res := KlabelBsort(k_res)
			if temp_res.Values[0] < k_res[2].Values[0] {
				k_res[2] = temp_res
			}
		}
	}

	return k_res
}
