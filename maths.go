package furthermaths

// Adds multiple amount of numbers
func Add(nums ...int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
