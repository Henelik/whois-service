package whois

import (
	"strings"
	"testing"
)

func TestWhoisPositive(t *testing.T) {
	// Only compare registry domain ID line to avoid test breaking due to time-dependent variables
	expectedID := "Registry Domain ID: 56fef0bc77624f99aea63b0d0ca1c638-DONUTS"
	got, err := Whois("fishtech.group", "10s")
	if err != nil {
		t.Error(err)
	}

	if strings.Split(got, "\r\n")[1] != expectedID {
		t.Errorf("Expected: %s, got: %s", expectedID, strings.Split(got, "\r\n")[1])
	}
}

func TestWhoisBadTLD(t *testing.T) {
	// Only compare registry domain ID line to avoid test breaking due to time-dependent variables
	expectedErr := "No whois server found for domain fishtech.grou"
	_, err := Whois("fishtech.grou", "10s")
	if err.Error() != expectedErr {
		t.Errorf("Expected: %s, got: %s", expectedErr, err.Error())
	}
}

func TestWhoisBadFormat(t *testing.T) {
	// Only compare registry domain ID line to avoid test breaking due to time-dependent variables
	expectedErr := "fishtech is not a valid domain name"
	_, err := Whois("fishtech", "10s")
	if err.Error() != expectedErr {
		t.Errorf("Expected: %s, got: %s", expectedErr, err.Error())
	}
}

func TestWhoisBadTimeout(t *testing.T) {
	// Only compare registry domain ID line to avoid test breaking due to time-dependent variables
	expectedErr := "time: unknown unit \"o\" in duration \"10o\""
	_, err := Whois("fishtech.group", "10o")
	if err.Error() != expectedErr {
		t.Errorf("Expected: %s, got: %s", expectedErr, err.Error())
	}
}
