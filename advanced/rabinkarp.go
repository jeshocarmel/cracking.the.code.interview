package advanced

func naiveStringSearch(fullstring, substring string) []int {

	m := len(substring)
	n := len(fullstring)

	found := []int{}

	for i := 0; i+m <= n; i++ {
		if fullstring[i:i+m] == substring {
			found = append(found, i)
		}
	}

	return found
}

func rabinKarpSearch(fullstring, substring string) []int {

	//credits: https://www.geeksforgeeks.org/java-program-for-rabin-karp-algorithm-for-pattern-searching/

	found := []int{}

	d := 256 // number of alphabets in accordance with ascii

	M := len(substring)
	N := len(fullstring)

	i := 0
	j := 0

	p := 0 //hash value for pattern
	t := 0 //hash value for text
	h := 1

	q := 101 // a large prime number

	// The value of h would be "pow(d, M-1)%q"
	for i := 0; i < M-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first
	// window of text

	for i = 0; i < M; i++ {
		p = (d*p + int(rune(substring[i]))) % q
		t = (d*t + int(rune(fullstring[i]))) % q
	}

	//fmt.Println(p, t)

	for i = 0; i <= N-M; i++ {

		if p == t {
			for j = 0; j < M; j++ {
				if fullstring[i+j] != substring[j] {
					break
				}
			}

			if j == M {
				found = append(found, i)
			}
		}

		if i < N-M {

			t = (d*(t-int(rune(fullstring[i]))*h) + int(rune(fullstring[i+M]))) % q

			// We might get negative value of t, converting it
			// to positive
			if t < 0 {
				t = (t + q)
			}
		}
	}
	return found
}
