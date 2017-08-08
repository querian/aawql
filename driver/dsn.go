package driver

import (
	"database/sql/driver"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// Dsn represents a data source name.
type Dsn struct {
	AdwordsID, APIVersion,
	DeveloperToken, AccessToken,
	ClientID, ClientSecret,
	RefreshToken string
	SkipColumnHeader,
	SupportsZeroImpressions,
	UseRawEnumValues bool
}

// NewDsn returns a new instance of Dsn.
func NewDsn(id string) *Dsn {
	return &Dsn{AdwordsID: id}
}

// String outputs the data source name as string.
// Output:
// 123-456-7890:v201607:true:false:false|dEve1op3er7okeN|1234567890-c1i3n7iD.com|c1ien753cr37|1/R3Fr35h-70k3n
func (d *Dsn) String() (n string) {
	if d.AdwordsID == "" {
		return
	}

	n = d.AdwordsID
	n += DsnOptSep + d.APIVersion
	n += DsnOptSep + strconv.FormatBool(d.SupportsZeroImpressions)
	n += DsnOptSep + strconv.FormatBool(d.SkipColumnHeader)
	n += DsnOptSep + strconv.FormatBool(d.UseRawEnumValues)

	if d.DeveloperToken != "" {
		n += DsnSep + d.DeveloperToken
	}
	if d.AccessToken != "" {
		n += DsnSep + d.AccessToken
	}
	if d.ClientID != "" {
		n += DsnSep + d.ClientID
	}
	if d.ClientSecret != "" {
		n += DsnSep + d.ClientSecret
	}
	if d.RefreshToken != "" {
		n += DsnSep + d.RefreshToken
	}

	return
}

// ParseDSN extract the main component of the dsn and prepare for the next analysis
func ParseDSN(dsn string) (values url.Values, err error) {

	if dsn[0:7] != "adwords" {
		err = driver.ErrBadConn
		return
	}

	var idx int
	if idx = strings.Index(dsn, "?"); idx == -1 {
		err = errors.New("query option are mandatory")
		return
	}

	return url.ParseQuery(dsn[idx+1:])
}
