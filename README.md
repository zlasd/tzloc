# tzloc
[![Go Reference](https://pkg.go.dev/badge/github.com/zlasd/tzloc.svg)](https://pkg.go.dev/github.com/zlasd/tzloc)

IANA time zone database's location enumeration

This package contains the full list of location enumeration of IANA timezone database generated automatically. The data source is from Go standard library lib/time/zoneinfo.zip

You can also check the list from wikipedia: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

Due to the variable database, the Go team seems not likely to add the enum to the standard library for keeping backward compatibility.
You can see the discussion on this issue: https://github.com/golang/go/issues/36278
