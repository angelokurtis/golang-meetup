package temperatures

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}
