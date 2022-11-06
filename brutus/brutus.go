package brutus

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/roatt-tilma/CyberSec-subDomainEnumeration/logger"
)

// Brute is a struct to define
// a single job to work on
type Brute struct {
	URL             string
	Word            string
	EnumerationType int
}

// New returns a new Brute object
func New(url string, word string, enumerationtype int) *Brute {
	return &Brute{
		URL:             url,
		Word:            word,
		EnumerationType: enumerationtype,
	}
}

// FormURL forms a URL from a brute object
func (b *Brute) FormURL() string {

	if b.EnumerationType == 0 {
		return fmt.Sprintf("https://%s.%s", b.Word, b.URL)
	} else if b.EnumerationType == 1 {
		return fmt.Sprintf("https://%s/%s", b.URL, b.Word)
	}

	return ""
}

// Try tries to visit a Brute URL and checks the status code
func (b *Brute) Try(success map[string]bool, logs chan logger.Log) {
	url := b.FormURL()

	resp, err := http.Get(url)
	if err != nil {
		//logger.Error("Error occurred while visiting " + url)
		return
	}

	// defer resp.Body.Close()

	statusCode := strconv.Itoa(resp.StatusCode)

	if success[statusCode] {
		logs <- logger.Log{
			Message: fmt.Sprintf("%s [Status code %s]", url, statusCode),
			Func:    logger.Info,
		}
	}
}
