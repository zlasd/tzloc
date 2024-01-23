package tzloc

//go:generate go run cmd/extract.go -output location.go

func ValidLocation(loc string) bool {
	_, ok := locationMap[loc]
	return ok
}
