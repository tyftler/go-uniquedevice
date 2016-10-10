// Package uniquedevice is a simple Go implementation of Wikimedia's approach
// to count unique devices accessing a website per day or month.
//
// Other than the common way to set a unique identifying cookie for each device and
// to store its access time in a database, we set the access time itself in a cookie.
// This package tells your application whether it should count the accessing device.
//
// For further explanation see
// https://blog.wikimedia.org/2016/03/30/unique-devices-dataset/
//
// For dealing with bots being counted see
// https://wikitech.wikimedia.org/wiki/Analytics/Unique_Devices/Last_access_solution#Nocookie_Offset
package uniquedevice

import (
	"net/http"
	"time"
)

const (
	name    = "last_access"
	layout  = "02-Jan-2006"
	path    = "/"
	expires = 30 * 24 * time.Hour
)

type UniqueDevice struct {
	dailyUnique   bool
	monthlyUnique bool
	noCookie      bool
}

// New returns a struct, containing information about the accessing device.
func New(w http.ResponseWriter, r *http.Request) *UniqueDevice {
	du, mu, nc := check(w, r)

	return &UniqueDevice{du, mu, nc}
}

func check(w http.ResponseWriter, r *http.Request) (bool, bool, bool) {
	t := time.Now()

	// get cookie
	c, err := r.Cookie(name)

	// set/renew cookie
	setCookie(w, t)

	if err != nil {
		if err == http.ErrNoCookie {
			return true, true, true
		}

		// reading cookie failed
		return true, true, false
	}

	// parse last access time
	ct, err2 := time.Parse(layout, c.Value)

	// parsing failed
	if err2 != nil {
		return true, true, false
	}

	if ct.Year() != t.Year() {
		return true, true, false
	}

	if ct.Month() != t.Month() {
		return true, true, false
	}

	if ct.Day() != t.Day() {
		return true, false, false
	}

	return false, false, false
}

func setCookie(w http.ResponseWriter, t time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   t.Format(layout),
		Path:    path,
		Expires: time.Now().Add(expires),
	})
}

// DailyUnique returns whether the device is accessing the website
// the first time a day and thus should be counted.
func (this *UniqueDevice) DailyUnique() bool {
	return this.dailyUnique
}

// MonthlyUnique returns whether the device is accessing the website
// the first time a month and thus should be counted.
func (this *UniqueDevice) MonthlyUnique() bool {
	return this.monthlyUnique
}

// NoCookie returns whether the device provided a cookie. This can
// be used for analyzing bot requests. For further information
// see the package description.
func (this *UniqueDevice) NoCookie() bool {
	return this.noCookie
}
