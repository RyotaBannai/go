package main

import (
	"fmt"
)

type IPAdder [4]byte

func (v IPAdder) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v[0], v[1], v[2], v[3])
}

func main() {
	hosts := map[string]IPAdder{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
