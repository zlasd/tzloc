// Package tzloc generate the full list of location enumeration
// of IANA timezone database. The data source is from Go standard
// library lib/time/zoneinfo.zip
//
// You can also check the list from wikipedia:
// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//
// Due to the variable database, the Go team seems not likely to
// add the enum to the standard library for keeping backward
// compatibility. You can see the discussion on this issue:
// https://github.com/golang/go/issues/36278
package tzloc

//go:generate go run cmd/extract.go -output location.go
//go:generate gofmt -w location.go

// ValidLocation checks whether the input is a valid location.
func ValidLocation(loc string) bool {
	_, ok := locationMap[loc]
	return ok
}

// GetLocationList returns the full list of timezone locations.
func GetLocationList() []string {
	locations := make([]string, 0, len(locationMap))
	for loc := range locationMap {
		locations = append(locations, loc)
	}
	return locations
}
