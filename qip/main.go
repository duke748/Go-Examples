package main

import (
	"fmt"
	"net"
)

// struct of ipv4 address
type qipIpv4Address struct {
	ipAddress string
	dnsName   string
	ttl       int
}

// Struct of subnet address
type subnetAddress struct {
	subnetAddress string
	subnetName    string
}

// Declare Interface and attach to all the read methods
type qipRecord interface {
	read()
}

func determineQipRecordType(q qipRecord) {
	switch q.(type) {
	case qipIpv4Address:
		fmt.Println("I'm an IPV4 address")
	case subnetAddress:
		fmt.Println("I'm a Subnet address")
	}

}

func main() {
	fmt.Println("Starting QIP API")

	ip1 := qipIpv4Address{
		ipAddress: "192.168.55.256",
		dnsName:   "Router",
		ttl:       300,
	}

	ip2 := qipIpv4Address{
		ipAddress: "192.168.55.100",
		dnsName:   "vcenter.virtualicity.co.uk",
		ttl:       300,
	}

	ip3 := qipIpv4Address{
		dnsName: "vra.virtualicity.co.uk",
		ttl:     300,
	}

	sub1 := subnetAddress{
		subnetAddress: "192.168.55.0",
		subnetName:    "Home",
	}

	ip1.add()
	ip1.delete()
	ip2.read()
	ip3.add()
	ip3.read()
	sub1.read()

	determineQipRecordType(ip1)
	determineQipRecordType(sub1)
}

// Receives qipIpv4Address struct , returns bool on wheter addition was valid or not
func (s qipIpv4Address) add() bool {
	if net.ParseIP(s.ipAddress) == nil {
		fmt.Println("Expected a valid IPV4 address but '", s.ipAddress, "' is not valid")
		return false
	}
	if net.ParseIP(s.ipAddress) != nil {
		fmt.Println("Adding DNS Entry", s.dnsName, "to", s.ipAddress)
		/// code

	}
	return true
}

func (s qipIpv4Address) delete() {
	if net.ParseIP(s.ipAddress) != nil {
		fmt.Println("Deleting DNS Entry", s.dnsName, "to", s.ipAddress)
	}
}

func (s qipIpv4Address) read() {
	if net.ParseIP(s.ipAddress) != nil {
		fmt.Println("Reading IP Address", s.ipAddress)
	}
}

func (s subnetAddress) read() {
	if net.ParseIP(s.subnetAddress) != nil {
		fmt.Println("Reading Subnet Address", s.subnetAddress)
	}
}
