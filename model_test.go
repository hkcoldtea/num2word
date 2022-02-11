package num2word

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	message := fmt.Sprintf("%v != %v", a, b)
	var cnt int
	switch a.(type) {
	case string:
		if a == b {
			return
		}
	case map[string]string:
		switch b.(type) {
		case map[string]string:
			if len(a.(map[string]string)) == len(b.(map[string]string)) {
				for k, v := range a.(map[string]string) {
					if bv, found := b.(map[string]string)[k]; found {
						if bv == v {
							cnt++
						}
					}
				}
			}
		default:
		}
		if cnt == len(a.(map[string]string)) {
			return
		}
	default:
	}
	t.Fatal(message)
}

func Test_tokenize(t *testing.T) {
	assertEqual(t,
		tokenize("0"), map[string]string{
			"Dollars": "0",
			"Cents":   "00",
		})
	assertEqual(t,
		tokenize("4"), map[string]string{
			"Dollars": "4",
			"Cents":   "00",
		})
	assertEqual(t,
		tokenize("16"), map[string]string{
			"Dollars": "16",
			"Cents":   "00",
		})
	assertEqual(t,
		tokenize("247"), map[string]string{
			"Dollars": "247",
			"Cents":   "00",
		})
	assertEqual(t,
		tokenize("1024"), map[string]string{
			"Thousand": "1",
			"Dollars":  "024",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("12345"), map[string]string{
			"Thousand": "12",
			"Dollars":  "345",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("123456"), map[string]string{
			"Thousand": "123",
			"Dollars":  "456",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("1000001"), map[string]string{
			"Million":  "1",
			"Thousand": "000",
			"Dollars":  "001",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("12312312"), map[string]string{
			"Million":  "12",
			"Thousand": "312",
			"Dollars":  "312",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("123456789"), map[string]string{
			"Million":  "123",
			"Thousand": "456",
			"Dollars":  "789",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("1234567890"), map[string]string{
			"Billion":  "1",
			"Million":  "234",
			"Thousand": "567",
			"Dollars":  "890",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("23232323232"), map[string]string{
			"Billion":  "23",
			"Million":  "232",
			"Thousand": "323",
			"Dollars":  "232",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("111222333444"), map[string]string{
			"Billion":  "111",
			"Million":  "222",
			"Thousand": "333",
			"Dollars":  "444",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("1222333444555"), map[string]string{
			"Trillion": "1",
			"Billion":  "222",
			"Million":  "333",
			"Thousand": "444",
			"Dollars":  "555",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("11222333444555"), map[string]string{
			"Trillion": "11",
			"Billion":  "222",
			"Million":  "333",
			"Thousand": "444",
			"Dollars":  "555",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("111222333444555"), map[string]string{
			"Trillion": "111",
			"Billion":  "222",
			"Million":  "333",
			"Thousand": "444",
			"Dollars":  "555",
			"Cents":    "00",
		})
	assertEqual(t,
		tokenize("1234567.89"), map[string]string{
			"Million":  "1",
			"Thousand": "234",
			"Dollars":  "567",
			"Cents":    "89",
		})
}

func Test_translate_three_digits(t *testing.T) {
	assertEqual(t,
		translate_three_digits("123"),
		"One Hundred and Twenty-three")
	assertEqual(t,
		translate_three_digits("404"),
		"Four Hundred and Four")
	assertEqual(t,
		translate_three_digits("520"),
		"Five Hundred and Twenty")
	assertEqual(t,
		translate_three_digits("300"),
		"Three Hundred")
}

func Example_translate_full_amount() {
	fmt.Println(Translate_full_amount("1234567.89"))
	fmt.Println(Translate_full_amount("1000001"))
	fmt.Println(Translate_full_amount("0.99"))
	// Output:
	// One Million Two Hundred and Thirty-four Thousand Five Hundred and Sixty-seven Dollars and Eighty-nine Cents Only
	// One Million One Dollars Only
	// Ninety-nine Cents Only
}
