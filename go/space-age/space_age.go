package space

func Age(seconds float64, planet string) float64 {
	var planetSeconds map[string]float64 = map[string]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (planetSeconds[planet] * 365.25 * 24 * 60 * 60)
}
