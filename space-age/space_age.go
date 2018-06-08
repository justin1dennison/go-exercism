package space

type Planet string

const earthSecondsInYear = 31557600

var orbitalPeriodsComparedToEarth = map[Planet]float64{
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
	earthYearAge := seconds / earthSecondsInYear
	return earthYearAge / orbitalPeriodsComparedToEarth[planet]
}
