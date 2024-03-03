package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string { 
	s := make([]string, len(ip))
	
	for i, value := range ip {
		s[i] = strconv.Itoa(int(value))
	}
	return strings.Join(s, ".")
}


func main() { 
	hosts := map[string]IPAddr{
		"loop": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts { 
		fmt.Printf("%v: %v\n", name, ip)
	}
}

