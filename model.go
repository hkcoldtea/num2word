package num2word

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/* A function to format numbers into human readable forms */
func format_num(amount string) string {
	f64, err := ParseFloat(amount)
	if err != nil {
		log.Println("Invalid number format given")
	}
	full_amount := fmt.Sprintf("%.2f", f64)
	return full_amount
}

/* For the cents parts */
func tokenize_cents(amount string) map[string]string {
	var cents string
	strsplit := strings.Split(amount, ".")
	cents = strsplit[1]
	var cents_tokens = map[string]string{"Cents": cents}
	return cents_tokens
}

/* For the dollars parts */
func tokenize_integer(amount string) map[string]string {
	var integer_tokens = make(map[string]string, 0)
	strsplit := strings.Split(amount, ".")
	integer_parts := strings.Split(strsplit[0], ",")
	number, err := strconv.ParseInt(strings.Join(integer_parts, ""), 10, 64)
	if err != nil {
		return integer_tokens
	}
	for _, name := range NUM_NAME {
		grpNumber := number % 1000
		number /= 1000
		ia := int(grpNumber)
		if number > 0 {
			integer_tokens[name] = fmt.Sprintf("%03d", ia)
		} else {
			integer_tokens[name] = strconv.Itoa(ia)
		}
		if number <= 0 {
			break
		}
	}
	return integer_tokens
}

/* Combine dollars parts and cents parts */
func tokenize(amount string) map[string]string {
	full_amount := format_num(amount)
	tokens := tokenize_integer(full_amount)
	for key, value := range tokenize_cents(full_amount) {
		tokens[key] = value
	}
	return tokens
}

func translate_three_digits(three_digits string) string {
	intHundreds_digit, _ := strconv.Atoi(string(three_digits[0]))
	tens_digit := string(three_digits[1])
	ones_digit := string(three_digits[2])

	hundreds_word := ONE_DIGIT[intHundreds_digit]

	if tens_digit == "0" && ones_digit == "0" {
		return hundreds_word + " Hundred"
	} else if tens_digit == "0" && ones_digit != "0" {
		ones_word := translate_one_digit(ones_digit)
		return hundreds_word + " " + HUNDRED + " and " + ones_word
	} else {
		tens_and_ones := tens_digit + ones_digit
		tens_and_ones_word := translate_two_digits(tens_and_ones)
		return hundreds_word + " " + HUNDRED + " and " + tens_and_ones_word
	}
}

func translate_two_digits(two_digits string) string {
	int2digit, _ := strconv.Atoi(two_digits)
	if _, found := TWO_DIGITS[int2digit]; found {
		return TWO_DIGITS[int2digit]
	} else {
		tens_part, _ := strconv.Atoi(string(two_digits[0]))
		tens_part *= 10
		ones_part, _ := strconv.Atoi(string(two_digits[1]))
		tens_word := TWO_DIGITS[tens_part]
		ones_word := strings.ToLower(ONE_DIGIT[ones_part])
		return tens_word + "-" + ones_word
	}

	return ""
}

func translate_one_digit(one_digit string) string {
	int1digit, _ := strconv.Atoi(one_digit)
	return ONE_DIGIT[int1digit]
}

func translate_digits(strdigits string) string {
	intdigits, _ := strconv.Atoi(strdigits)
	digits := strconv.Itoa(intdigits)
	if len(digits) == 3 {
		return translate_three_digits(digits)
	} else if len(digits) == 2 {
		return translate_two_digits(digits)
	} else if len(digits) == 1 {
		return translate_one_digit(digits)
	}
	return ""
}

/* A function that prints given number in words */
func Translate_full_amount(full_amount string) string {
	tokens := tokenize(full_amount)
	var tokens_with_rm_zero map[string]string
	tokens_with_rm_zero = make(map[string]string, 0)
	for key, value := range tokens {
		i64, err := strconv.ParseInt(value, 10, 64)
		if err == nil && i64 != 0 {
			tokens_with_rm_zero[key] = tokens[key]
		}
	}
	var tokens_words map[string]string
	tokens_words = make(map[string]string, 0)
	for key, _ := range tokens_with_rm_zero {
		tokens_words[key] = translate_digits(tokens_with_rm_zero[key])
	}
	integer_part_words := ""
	cents_part_words := ""

	/* Base cases */
	if len(tokens_words) == 1 {
		if _, found := tokens_words["Cents"]; found {
			cents_part_words := tokens_words["Cents"] + " Cents Only"
			return cents_part_words
		}
	} else if len(tokens_words) == 0 {
		return "Zero"
	}
	for i := len(NUM_NAME)-1; i >= 0; i-- {
		for key, value := range tokens_words {
			if key == NUM_NAME[i] {
				integer_part_words += value + " " + key + " "
			}
			if i == 0 && key == "Cents" {
				cents_part_words += value + " " + key + " "
			}
		}
	}

	var full_words string
	if len(cents_part_words) == 0 {
		full_words = integer_part_words + "Only"
	} else {
		full_words = integer_part_words + "and " + cents_part_words + "Only"
	}

	return full_words
}
