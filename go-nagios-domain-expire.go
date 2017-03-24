package main

import (
	"github.com/likexian/whois-parser-go"
	"github.com/likexian/whois-go"
	//"github.com/fatih/color"
	"fmt"
	"flag"
)

var command string

func init() {

	flag.StringVar(&command, "command", command, "" +
		"Commands:\n " +
		"\t\t showConfig - Show configuration file \n" +
		"\t\t check - Run check domains \n")
}




func main() {
	flag.Parse()
	//pinfo := color.New(color.FgYellow,color.BgBlack)
	//panounce := color.New(color.FgHiCyan,color.BgBlack)
	//perror := color.New(color.FgHiRed,color.BgBlack)

	whois_raw, err := whois_parser.ReadFile("./whois.txt")
	//whois_raw := "ios.in.ua"
	result, err := whois_parser.Parser(whois_raw)
	if err == nil {
		// Print the domain status
		fmt.Println(result.Registrar.DomainStatus)

		// Print the domain created date
		fmt.Println(result.Registrar.CreatedDate)

		// Print the domain expiration date
		fmt.Println(result.Registrar.ExpirationDate)

		// Print the registrant name
		fmt.Println(result.Registrant.Name)

		// Print the registrant email address
		fmt.Println(result.Registrant.Email)
	} else {
	fmt.Println(err)
	}

	result1, err1 := whois.Whois("redacid.org.ua")
	if err1 == nil {
		fmt.Println(result1)
	}


}