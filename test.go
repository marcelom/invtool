package main

import (
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	//"os"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}
type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}

func main() {
	v := Result{Name: "none", Phone: "none"}

	data := `
	<Person>
		<FullName>Grace R. Emlin</FullName>
		<Company>Example Inc.</Company>
		<Email where='work'>
			<Addr>gre@work.com</Addr>
		</Email>
		<Group>
			<Value>Friends</Value>
			<Value2>Squash</Value2>
		</Group>
		<City>Hanga Roa</City>
		<State>Easter Island</State>
	</Person>`

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}
