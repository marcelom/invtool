package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type NMAPRun struct {
	//XMLName xml.Name `xml:"nmaprun"`
	Scanner          string   `xml:"scanner,attr"`
	Args             string   `xml:"args,attr"`
	Start            uint64   `xml:"start,attr"`
	Version          string   `xml:"version,attr"`
	XMLOutputVersion string   `xml:"xmloutputversion,attr"`
	ScanInfo         ScanInfo `xml:"scaninfo"`
	Host             []Host   `xml:"host"`
	RunStats         RunStats `xml:"runstats"`
}

type ScanInfo struct {
	//XMLName xml.Name `xml:"scaninfo"`
	Type        string `xml:"type,attr"`
	Protocol    string `xml:"protocol,attr"`
	NumServices uint64 `xml:"numservices,attr"`
	//Services string `xml:"services,attr"`
}

type Host struct {
	//XMLName xml.Name `xml:"host"`
	StartTime  uint64       `xml:"starttime,attr"`
	EndTime    uint64       `xml:"endtime,attr"`
	Status     Status       `xml:"status"`
	HostName   []HostName   `xml:"hostnames>hostname"`
	Address    []Address    `xml:"address"`
	Port       []Port       `xml:"ports>port"`
	OSClass    []OSClass    `xml:"os>osclass"`
	OSMatch    []OSMatch    `xml:"os>osmatch"`
	HostScript []HostScript `xml:"hostscript>script"`
}

type Status struct {
	//XMLName xml.Name `xml:"status"`
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

type HostName struct {
	//XMLName xml.Name `xml:"status"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type Address struct {
	//XMLName xml.Name `xml:"address"`
	Addr   string `xml:"addr,attr"`
	Type   string `xml:"addrtype,attr"`
	Vendor string `xml:"vendor,attr"`
}

type Port struct {
	//XMLName xml.Name `xml:"port"`
	Protocol string  `xml:"protocol,attr"`
	PortID   uint32  `xml:"portid,attr"`
	State    State   `xml:"state"`
	Service  Service `xml:"service"`
}

type State struct {
	//XMLName xml.Name `xml:"state"`
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

type Service struct {
	//XMLName xml.Name `xml:"service"`
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Ostype  string `xml:"ostype,attr"`
	Method  string `xml:"method,attr"`
	Conf    uint32 `xml:"conf,attr"`
}

type OSClass struct {
	//XMLName xml.Name `xml:"osclass"`
	Type     string `xml:"type,attr"`
	Vendor   string `xml:"vendor,attr"`
	OSFamily string `xml:"osfamily,attr"`
	OSGen    string `xml:"osgen,attr"`
	Accuracy uint8  `xml:"accuracy,attr"`
}

type OSMatch struct {
	//XMLName xml.Name `xml:"osmat"`
	Name     string `xml:"name,attr"`
	Accuracy uint8  `xml:"accuracy,attr"`
	Line     uint64 `xml:"line,attr"`
}

type HostScript struct {
	//XMLName xml.Name `xml:"hostscript"`
	ID     string `xml:"id,attr"`
	Output string `xml:"output,attr"`
}

type RunStats struct {
	//XMLName xml.Name `xml:"hostscript"`
	Finished   Finished   `xml:"finished"`
	HostsTotal HostsTotal `xml:"hosts"`
}

type Finished struct {
	EndTime uint64 `xml:"time,attr"`
}
type HostsTotal struct {
	HostsUp    uint32 `xml:"up,attr"`
	HostsDown  uint32 `xml:"down,attr"`
	Hoststotal uint32 `xml:"total,attr"`
}

func main() {
	xmlData, err := ioutil.ReadFile("full.xml")
	//xmlData, err := ioutil.ReadFile("multiple.xml")
	//xmlData, err := ioutil.ReadFile("single.xml")

	if err != nil {
		fmt.Println("Error opening/reading file:", err)
		os.Exit(1)
	}
	//defer Fp.Close()

	var n NMAPRun
	xml.Unmarshal(xmlData, &n)

	fmt.Printf("%+v\n", n)
}
