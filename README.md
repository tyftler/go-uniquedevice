# go-uniquedevice

The package `uniquedevice` is a simple Go implementation of Wikimedia's approach
to count unique devices accessing a website per day or month.

Other than the common way to set a unique identifying cookie for each device and
to store its access time in a database, we set the access time itself in a cookie.
This package tells your application whether it should count the accessing device.

For further explanation see [blog.wikimedia.org/2016/03/30/unique-devices-dataset](https://blog.wikimedia.org/2016/03/30/unique-devices-dataset/).

For dealing with bots being counted see [wikitech.wikimedia.org/wiki/Analytics/Unique_Devices/Last_access_solution#Nocookie_Offset](https://wikitech.wikimedia.org/wiki/Analytics/Unique_Devices/Last_access_solution#Nocookie_Offset).

## Installation and Docs

Install using `go get github.com/tyftler/go-uniquedevice`.

Full documentation is available at [godoc.org/github.com/tyftler/go-uniquedevice](https://godoc.org/github.com/tyftler/go-uniquedevice).

## Usage

```go
ud := uniquedevice.New(w, r)

if ud.DailyUnique() {
	// count device for this day
}

if ud.MonthlyUnique() {
	// count device for this month
}
```
