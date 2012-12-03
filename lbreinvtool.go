package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type NMAPRun struct {
	//XMLName xml.Name `xml:"nmaprun"`
	Scanner string   `xml:"scanner,attr"`
	Args    string   `xml:"args,attr"`
	Start uint64 `xml:"start,attr"`
	Version string `xml:"version,attr"`
	XMLOutputVersion string `xml:"xmloutputversion,attr"`
	ScanInfo ScanInfo `xml:"scaninfo"`
	Host []Host `xml:"host"`
}

type ScanInfo struct {
	//XMLName xml.Name `xml:"scaninfo"`
	Type string `xml:"type,attr"`
	Protocol string `xml:"protocol,attr"`
	NumServices uint64 `xml:"numservices,attr"`
	//Services string `xml:"services,attr"`
}

type Host struct {
	//XMLName xml.Name `xml:"host"`
	StartTime uint64 `xml:"starttime,attr"`
	EndTime uint64 `xml:"endtime,attr"`
	Status Status `xml:"status"`
}

type Status struct{
	//XMLName xml.Name `xml:"status"`
	State string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

func main() {
	//xmlData, err := ioutil.ReadFile("full.xml")
	//xmlData, err := ioutil.ReadFile("multiple.xml")
	xmlData, err := ioutil.ReadFile("single.xml")

	if err != nil {
		fmt.Println("Error opening/reading file:", err)
		os.Exit(1)
	}
	//defer Fp.Close()

	var n NMAPRun
	xml.Unmarshal(xmlData, &n)

	fmt.Printf("%+v\n",n)
}
