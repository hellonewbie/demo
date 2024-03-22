package main

import (
	"GoSqlInject/ch-3/shodan"
	"fmt"
	"log"
	"os"
)

//LneoEmOBDTinC3Fb1C3qu4GxP9QzmRpZ
func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	//apiKey := os.Getenv("SHODAN_API_KEY")
	Client := shodan.New("LneoEmOBDTinC3Fb1C3qu4GxP9QzmRpZ")
	APIinfo, err := Client.APIinfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits:%d\nScan Credits : %d\n\n",
		APIinfo.QueryCredits,
		APIinfo.ScanCredits,
	)
	HostSearch, err := Client.HostSearch(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	for _, host := range HostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IpStr, host.Port)
	}

}
