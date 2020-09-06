package whois

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhoisPositive(t *testing.T) {
	// Only compare registry domain name line to avoid test breaking due to time-dependent variables
	expectedDomain := "Domain Name: fishtech.group"
	got, err := Whois("fishtech.group", "10s")

	assert.Nil(t, err)

	assert.Equal(t, expectedDomain, strings.Split(got, "\r\n")[0])
}

func TestWhoisBadTLD(t *testing.T) {
	expectedErr := "No whois server found for domain fishtech.grou"
	_, err := Whois("fishtech.grou", "10s")

	assert.Equal(t, err.Error(), expectedErr)
}

func TestWhoisBadFormat(t *testing.T) {
	expectedErr := "fishtech is not a valid domain name"
	_, err := Whois("fishtech", "10s")

	assert.Equal(t, err.Error(), expectedErr)
}

func TestWhoisBadTimeout(t *testing.T) {
	expectedErr := "time: unknown unit \"o\" in duration \"10o\""
	_, err := Whois("fishtech.group", "10o")

	assert.Equal(t, err.Error(), expectedErr)
}
