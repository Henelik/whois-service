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
