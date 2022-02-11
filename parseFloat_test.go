package num2word

import (
	"fmt"
	"testing"
)

func TestParseFloat(t *testing.T) {
	Must := func(v float64, err error) float64 {
		if err != nil {
			t.Fatal(err)
		}
		return v
	}

	MustEqual := func(v1, v2 float64) {
		if v1 != v2 {
			t.Fatal(fmt.Sprintf("number is not equal, v1=%f, v2=%f", v1, v2))
		}
	}

	MustEqual(Must(ParseFloat("1E2")), 100)
	MustEqual(Must(ParseFloat("1E-5")), 0.00001)
	MustEqual(Must(ParseFloat("1.6543E2")), 165.43)
	MustEqual(Must(ParseFloat("0.89E2")), 89)
	MustEqual(Must(ParseFloat("1.6543E-2")), 0.016543)
	MustEqual(Must(ParseFloat("156,819,129")), 156819129)
	MustEqual(Must(ParseFloat("156819129")), 156819129)
	MustEqual(Must(ParseFloat(".1E0")), 0.1)
	MustEqual(Must(ParseFloat(".1E1")), 1)
	MustEqual(Must(ParseFloat("0E1")), 0)
}
