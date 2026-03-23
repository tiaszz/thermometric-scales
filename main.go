package main

import "fmt"

func main() {
	metric := choosingMetric()
	convertMetric := choosingCovertedMetric()

	converting(metric, convertMetric)
}

func choosingMetric() int {
	var metric int
	fmt.Println("Choose yours metric.")
	fmt.Print("[1] Celcius\n[2]Fahrenheit\n[3]Kelvin\nChoose your metric: ")
	fmt.Scanf("%d", &metric)
	return metric
}

func choosingCovertedMetric() int {
	var metric int
	fmt.Println("Choose the metric you want to convert.")
	fmt.Print("[1] Celcius\n[2]Fahrenheit\n[3]Kelvin\nChoose your metric: ")
	fmt.Scanf("%d", &metric)
	return metric
}

func converting(metric int, convertedMetric int) {
	var input int
	fmt.Print("Input your degree: ")
	fmt.Scanf("%d", &input)

	switch metric {
	case 1:
		switch convertedMetric {
		case 1:
			panicMessageOwnMetric()
		case 2:
			celciusToFahrenheit(input)
		case 3:
			celciusToKelvin(input)
		default:
			panicMessageInvalid()
		}
	case 2:
		switch convertedMetric {
		case 1:
			fahrenheitToCelcius(input)
		case 2:
			panicMessageOwnMetric()
		case 3:
			fahrenheitToKelvin(input)
		default:
			panicMessageInvalid()
		}
	case 3:
		switch convertedMetric {
		case 1:
			kelvinToCelcius(input)
		case 2:
			kelvinToFahrenheit(input)
		case 3:
			panicMessageOwnMetric()
		default:
			panicMessageInvalid()
		}
	default:
		panicMessageInvalid()
	}
}

func panicMessageInvalid() {
	panic("Choose a valid metric")
}

func panicMessageOwnMetric() {
	panic("Can't convert your degree to your choose one")
}

func celciusToFahrenheit(input int) {
	fahrenheit := input*9/5 + 32
	fmt.Printf("Your degree in fahrenheit is: %d\n", fahrenheit)
}

func celciusToKelvin(input int) {
	kelvin := float64(input) + 273.15
	fmt.Printf("Your degree in kelvin is: %.2f\n", kelvin)
}

func fahrenheitToCelcius(input int) {
	celcius := float64((input - 32)) / 1.8
	fmt.Printf("Your degree in celcius is: %.2f\n", celcius)
}

func fahrenheitToKelvin(input int) {
	kelvin := (float64(input) + 459.67) * 5 / 9
	fmt.Printf("Your degree in kelvin is: %.2f\n", kelvin)
}

func kelvinToCelcius(input int) {
	celcius := float64(input) - 273.15
	fmt.Printf("Your degree in celcius is: %.2f\n", celcius)
}

func kelvinToFahrenheit(input int) {
	fahrenheit := float64(input)*1.8 - 459.67
	fmt.Printf("Your degree is fahrenheit is: %.2f\n", fahrenheit)
}
