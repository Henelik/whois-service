package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

// Whois finds the whois data for a given domain name
// waitDuration must be in a parseable time format
func Whois(domain string, waitDuration string) (result string, err error) {
	// split the domain name and ensure valid format
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		err = fmt.Errorf("%s is not a valid domain name", domain)
		return
	}

	// look for this top level domain's whois server in the list
	server, ok := server_list[parts[len(parts)-1]]
	if !ok {
		err = fmt.Errorf("No whois server found for domain %s", domain)
		return
	}

	// parse timeout string
	timeout, err := time.ParseDuration(waitDuration)
	if err != nil {
		return
	}

	// connect to the server
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), timeout)
	if err != nil {
		return
	}

	defer connection.Close()

	// send the domain name to the server
	connection.Write([]byte(domain + "\r\n"))

	// read the response
	buffer, err := ioutil.ReadAll(connection)
	if err != nil {
		return
	}

	result = string(buffer[:])

	return
}
