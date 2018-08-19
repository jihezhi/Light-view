package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string {
	return string(v[0]) + "." + string(v[1]) + "." + string(v[2]) + "." + string(v[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.String) //紗霧：在Go中如果不是特殊需求，没必要使用Printf。请尝试使用Print
	}

}
