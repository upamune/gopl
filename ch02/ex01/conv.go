package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
