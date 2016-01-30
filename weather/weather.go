package weather

import (
	"net/http"
    "net/url"
    "strings"
)

var getURL string

type UserAgent interface {
    Get(url string) (*http.Response, error)
}

type Weather struct {
    ua UserAgent
}

func New(ua UserAgent) (*Weather) {
    return &Weather{ ua: ua, }
}

func (w Weather) Get(town string, city string) (*http.Response, error) {
    queryParams := url.Values{}
    queryParams.Set("q", strings.Join([]string{town, city},",") )
    queryParams.Set("appid", "44db6a862fba0b067b1930da0d769e98")
    queryParams.Set("units", "metric")

    queryParamsStr := queryParams.Encode()

    apiURL := "http://api.openweathermap.org/data/2.5/weather"

    getURL = strings.Join([]string{ apiURL, "?", queryParamsStr }, "")

    res, err := w.ua.Get(getURL)

    return res, err
}
