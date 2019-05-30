package apple

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "http://itunes.apple.com/lookup"
	userAgent      = "go-apple"
)

// Response holds result from apple response
type Response struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

// Result single app result
type Result struct {
	IsGameCenterEnabled              bool          `json:"isGameCenterEnabled"`
	ArtistViewURL                    string        `json:"artistViewUrl"`
	ArtworkURL60                     string        `json:"artworkUrl60"`
	ArtworkURL100                    string        `json:"artworkUrl100"`
	ScreenshotUrls                   []string      `json:"screenshotUrls"`
	IpadScreenshotUrls               []interface{} `json:"ipadScreenshotUrls"`
	AppletvScreenshotUrls            []interface{} `json:"appletvScreenshotUrls"`
	ArtworkURL512                    string        `json:"artworkUrl512"`
	Advisories                       []interface{} `json:"advisories"`
	SupportedDevices                 []string      `json:"supportedDevices"`
	Kind                             string        `json:"kind"`
	Features                         []interface{} `json:"features"`
	TrackCensoredName                string        `json:"trackCensoredName"`
	TrackViewURL                     string        `json:"trackViewUrl"`
	ContentAdvisoryRating            string        `json:"contentAdvisoryRating"`
	LanguageCodesISO2A               []string      `json:"languageCodesISO2A"`
	FileSizeBytes                    string        `json:"fileSizeBytes"`
	SellerURL                        string        `json:"sellerUrl"`
	TrackContentRating               string        `json:"trackContentRating"`
	MinimumOsVersion                 string        `json:"minimumOsVersion"`
	CurrentVersionReleaseDate        time.Time     `json:"currentVersionReleaseDate"`
	GenreIds                         []string      `json:"genreIds"`
	ReleaseDate                      time.Time     `json:"releaseDate"`
	PrimaryGenreName                 string        `json:"primaryGenreName"`
	PrimaryGenreID                   int           `json:"primaryGenreId"`
	TrackID                          int           `json:"trackId"`
	TrackName                        string        `json:"trackName"`
	ReleaseNotes                     string        `json:"releaseNotes"`
	SellerName                       string        `json:"sellerName"`
	WrapperType                      string        `json:"wrapperType"`
	Version                          string        `json:"version"`
	FormattedPrice                   string        `json:"formattedPrice"`
	ArtistID                         int           `json:"artistId"`
	ArtistName                       string        `json:"artistName"`
	Genres                           []string      `json:"genres"`
	Price                            float64       `json:"price"`
	Description                      string        `json:"description"`
	Currency                         string        `json:"currency"`
	BundleID                         string        `json:"bundleId"`
	IsVppDeviceBasedLicensingEnabled bool          `json:"isVppDeviceBasedLicensingEnabled"`
}

// GetAppByID gets app information from apple
func GetAppByID(id int) (*Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?id=%v", defaultBaseURL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result Response

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
