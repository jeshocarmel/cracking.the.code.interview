package ch16

func englishInt(input int) string {

	number := abs(input)
	bigs := []string{"", "thousand ", "million ", "billion "}

	result := ""

	for i := 0; number > 0; i++ {

		chunk := number % 1000
		result = convert(chunk) + bigs[i] + result
		number = number / 1000
	}

	if input < 0 {
		return "negative " + result
	}

	return result

}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func convert(chunk int) string {

	smalls := []string{"", "one ", "two ", "three ", "four ", "five ", "six ", "seven ", "eight ", "nine ", "ten ", "eleven ",
		"twelve ", "thirteen ", "fourteen ", "fifteen ", "sixteen ", "seventeen ", "eighteen ", "nineteen "}
	tens := []string{"", "", "twenty ", "thirty ", "forty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}

	str := ""

	if chunk >= 100 {
		str = str + smalls[chunk/100] + "hundred "
		chunk = chunk % 100
	}

	if chunk > 19 && chunk < 100 {
		str = str + tens[chunk/10]
		chunk = chunk % 10
	}

	if chunk > 0 && chunk <= 19 {
		str = str + smalls[chunk]
	}

	return str

}
