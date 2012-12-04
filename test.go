package main

import (
	"encoding/xml"
	"fmt"
)

//UC is Upper Case, LC is Lower Case

type EmailUC struct {
	Where string `xml:"Where,attr"`
	Addr  string
}

type PersonUC struct {
	XMLName xml.Name `xml:"Person"`
	Name    string
	Email   []EmailUC
}

type EmailLC struct {
	Where string `xml:"where,attr"`
	Addr  string `xml:"addr"`
}

type PersonLC struct {
	XMLName xml.Name  `xml:"person"`
	Name    string    `xml:"name"`
	Email   []EmailLC `xml:"email"`
}

func main() {
	data1UC := `<Person><Name>John Doe</Name>
                   <Email Where="home"><Addr>jd@home.com</Addr></Email>
                   <Email Where='work'><Addr>jd@work.com</Addr></Email>
                   </Person>`
	data2UC := `<Person><Name>John Doe</Name>
                   <Email Where="home"><Addr>jd@home.com</Addr></Email>
                   </Person>`
	data1LC := `<person><name>Fulano de Tal</name>
                   <email where="home"><addr>fdt@home.com</addr></email>
                   <email where='work'><addr>fdt@work.com</addr></email>
                   </person>`
	data2LC := `<person><name>Fulano de Tal</name>
                   <email where="home"><addr>fdt@home.com</addr></email>
                   </person>`

	var Person1UC, Person2UC PersonUC
	var Person1LC, Person2LC PersonLC

	err := xml.Unmarshal([]byte(data1UC), &Person1UC)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Upper Case with multiple Email: %+v\n", Person1UC)

	err = xml.Unmarshal([]byte(data2UC), &Person2UC)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Upper Case with single Email: %+v\n", Person2UC)

	err = xml.Unmarshal([]byte(data1LC), &Person1LC)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Lower Case with multiple Email: %+v\n", Person1LC)

	err = xml.Unmarshal([]byte(data2LC), &Person2LC)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("Lower Case with single Email: %+v\n", Person2LC)

}
