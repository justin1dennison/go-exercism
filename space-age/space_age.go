package space

type Planet string

const EARTH_SECONDS_IN_YEAR = 31557600

var ORBITAL_PERIODS_COMPARED_TO_EARTH = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61518726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	earthYearAge := seconds / EARTH_SECONDS_IN_YEAR
	return earthYearAge / ORBITAL_PERIODS_COMPARED_TO_EARTH[planet]
}
