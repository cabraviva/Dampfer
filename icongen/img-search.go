package icongen

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Top 10 results
func internal_getImagesFor(query string) ([]string, error) {
	// Build the Google Images URL.
	// - tbm=isch  → image search mode.
	// - safe=off   → disable SafeSearch.
	// - tbs=iar:s  → filter for square images (1:1 aspect ratio).
	searchURL := "https://www.google.com/search?tbm=isch&q=" + url.QueryEscape(query) +
		"&safe=off&tbs=iar:s"

	// Create an HTTP request with a browser-like User-Agent.
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/114.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Use a regular expression to extract image URLs.
	re := regexp.MustCompile(`https?://[^"]+\.(jpg|jpeg|png|svg|webp|gif)`)
	matches := re.FindAllString(string(body), -1)

	// Remove duplicate URLs and exclude gstatic, wikipedia, wikimedia, and invalid URLs.
	unique := make(map[string]struct{})
	var urls []string
	for _, m := range matches {
		// Filter out gstatic URLs, invalid URLs (like https://www.svg), and wikipedia or wikimedia URLs.
		if !strings.Contains(m, "gstatic.com") &&
			!strings.Contains(m, "wikipedia.org") &&
			!strings.Contains(m, "wikimedia.org") &&
			!strings.Contains(m, "www.svg") &&
			!strings.Contains(m, "www.png") &&
			isValidImageURL(m) {

			if _, ok := unique[m]; !ok {
				unique[m] = struct{}{}
				urls = append(urls, m)
			}
		}
	}

	// Limit the results to the top 10 URLs.
	if len(urls) > 10 {
		urls = urls[:10]
	}

	return urls, nil
}

// GetImagesFor fetches the top 20 image URLs for a given search query.
func GetImagesFor(query string) ([]string, error) {
	r1, err1 := internal_getImagesFor(query + " icon")
	r2, err2 := internal_getImagesFor(query + " logo")

	var err error = nil
	if err1 != nil {
		err = err1
	} else if err2 == nil {
		err = err2
	}
	rFinal := append(r1, r2...)

	return rFinal, err
}

// isValidImageURL checks if a URL looks like a valid image link with an extension like jpg, png, svg, etc.
func isValidImageURL(url string) bool {
	// Check if the URL ends with a valid image extension and is not just a domain name.
	extRegex := `\.(jpg|jpeg|png|svg|)$`
	re := regexp.MustCompile(extRegex)
	return re.MatchString(url)
}
