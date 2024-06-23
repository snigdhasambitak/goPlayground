package main

import "fmt"

type IPAddress [4]byte

func (ip IPAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {

	iprange := map[string]IPAddress{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for k, v := range iprange {
		fmt.Printf("%v: %v\n", k, v)
	}
}
