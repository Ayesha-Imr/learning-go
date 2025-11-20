package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentSuccess := float64(productionRate) *(successRate / 100)
    return percentSuccess
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percentSuccessMinute := (float64(productionRate) * (successRate / 100)) / 60
    return int(percentSuccessMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
    left := carsCount % 10
    totalCost := (tens * 95000) + (left * 10000)
    return uint(totalCost)
}
