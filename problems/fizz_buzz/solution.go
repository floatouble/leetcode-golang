func fizzBuzz(n int) []string {
    result := make([]string, n+1)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i] = "FizzBuzz"
		case i%3 == 0:
			result[i] = "Fizz"
		case i%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(i)
		}
	}
	return result[1:]
}