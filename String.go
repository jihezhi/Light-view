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

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string {
	fmt.Print(v[0], ".", v[1], ".", v[2], ".", v[3])
	s := string(v[0]) + "." + string(v[1]) + "." + string(v[2]) + "." + string(v[3])
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  IPAddr{127, 0, 0, 1},
		"googleDNS": IPAddr{8, 8, 8, 8},
	}
	//fmt.Println(hosts)
	for _, ip := range hosts {
		fmt.Println(ip.String())
	}

}

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string {
	return fmt.Sprint(v[0], ".", v[1], ".", v[2], ".", v[3]) // 紗霧：很棒！你掌握了更加贴近Go风格的用法。之后你想用Printf也不是不可以，当你真正需要它的时候，比如说如果你需要控制输出格式。
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	//fmt.Print(hosts["googleDNS"])
	for name, ip := range hosts {
		fmt.Print(name, ":", ip, "\n") // 紗霧：完美~
	}

}
