package space

type Planet string

const EARTH_DAYS = 365.25

var ORBITAL_PERIOD = map[Planet]float64{
	"Earth":   EARTH_DAYS,
	"Mercury": EARTH_DAYS * 0.2408467,
	"Venus":   EARTH_DAYS * 0.61519726,
	"Mars":    EARTH_DAYS * 1.8808158,
	"Jupiter": EARTH_DAYS * 11.862615,
	"Saturn":  EARTH_DAYS * 29.447498,
	"Uranus":  EARTH_DAYS * 84.016846,
	"Neptune": EARTH_DAYS * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	if ORBITAL_PERIOD[planet] == 0 {
		return 0
	}
	return seconds / (ORBITAL_PERIOD[planet] * 86400.0)
}
