package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/derekg/DLID/dlidparser"
)

func main() {
	s, err := dlidparser.Parse("@\n\x1e\rANSI 636000070002DL00410281ZV03190008DLDAQT64235789\nDCSSAMPLE\nDDEN\nDACMICHAEL\nDDFN\nDADJOHN,BOB\nDDGN\nDCUJR\nDCAD\nDCBK\nDCDPH\nDBD06062008\nDBB06061986\nDBA12102012\nDBC1\nDAU068 in\nDAYBRO\nDAG2300 WEST BROAD STREET\nDAIRICHMOND\nDAJVA\nDAK232690000 \nDCF2424244747474786102204\nDCGUSA\nDCK123456789\nDDAM\nDDB06062008\nDDC06062009\nDDD1\rZVZVA01\r")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println()
	fmt.Println(s)
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		foo, err := dlidparser.Parse(string(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(foo)
	}

}
