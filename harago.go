package harahgo

import (
	"encoding/json"
)

type point struct {
	Lat float64
	Lng float64
}
type polygon []point
type Area struct {
	DistrictID int           `json:"district_id"`
	CityID     int           `json:"city_id"`
	RegionID   int           `json:"region_id"`
	NameAr     string        `json:"name_ar"`
	NameEn     string        `json:"name_en"`
	Boundaries [][][]float64 `json:"boundaries"`
}

func GetHarah(lat float64, lng float64) Area {
	p := point{lat, lng}

	// type Point [int,int]
	byteValue := districtsBlob

	var areas []Area

	json.Unmarshal(byteValue, &areas)

	for _, area := range areas {

		var polygon polygon
		for _, boundaries := range area.Boundaries[0] {
			polygon = append(polygon, point{boundaries[0], boundaries[1]})
		}
		if polygon.pointInside(p) {
			return (area)
		}
	}
	return Area{}
}

func (p polygon) pointInside(point point) bool {

	numCollisions := 0

	for i := 0; i < len(p); i++ {

		iʹ := (i + 1) % (len(p))

		// If, with Lat, the point sits within the edge
		if (p[i].Lat >= point.Lat) != (p[iʹ].Lat >= point.Lat) {

			dLng := p[iʹ].Lng - p[i].Lng
			dLat := point.Lat - p[i].Lat
			dLatʹ := p[iʹ].Lat - p[i].Lat

			if point.Lng <= dLng*dLat/dLatʹ+p[i].Lng {

				// It's a collision, so iterate the numCollisions
				numCollisions++
			}
		}
	}

	// If the number of collisions is ODD, return true (it's inside)
	return ((numCollisions & 1) == 1)
}
