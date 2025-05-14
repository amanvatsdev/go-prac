package practice02


func DiscountCalculator() func(int) int {
	return func(i int) int {
		return i-50
	}
}