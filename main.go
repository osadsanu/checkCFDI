package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/osadsanu/showCFDI/cfdi"
)

func main() {

	folderPath := "."
	//read Files from img folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	var cfdiList []cfdi.ShortCFDI
	var tax float64
	var total float64

	for _, f := range files {
		if strings.Contains(f.Name(), ".xml") {
			//fmt.Println("reading File:", f.Name())
			xmlFile, err := os.Open(f.Name())

			if err != nil {
				fmt.Println("Error reading:", f.Name())
				fmt.Println(err)
				defer xmlFile.Close()
				continue
			}
			defer xmlFile.Close()

			byteValue, err := ioutil.ReadAll(xmlFile)
			if err != nil {
				fmt.Println("cant read: ", f.Name())
				continue
			}
			sCFDI := cfdi.GetShortCFDI(byteValue)
			cfdiList = append(cfdiList, sCFDI)
			//fmt.Println(sCFDI.ToString())
			tax += sCFDI.Impuestos
			total += sCFDI.Total
		}
	}
	//order the list by date, total, and tax as params
	sort.Slice(cfdiList, func(i, j int) bool {
		if len(os.Args) > 1 {
			if os.Args[1] == "d" {
				return cfdiList[i].Fecha < cfdiList[j].Fecha
			} else if os.Args[1] == "total" {
				return cfdiList[i].Total < cfdiList[j].Total
			} else if os.Args[1] == "tax" {
				return cfdiList[i].Impuestos < cfdiList[j].Impuestos
			} else {
				return cfdiList[i].Fecha < cfdiList[j].Fecha
			}
		} else {
			return cfdiList[i].Fecha < cfdiList[j].Fecha
		}
	})

	for _, cfdis := range cfdiList {
		fmt.Println(cfdis.ToString())
	}

	//args
	/*
		for i := 0; i < len(os.Args); i++ {
			fmt.Println("Arg ", i, " = "+os.Args[i])
		}
	*/
	fmt.Printf("Total:%.2f\n  Tax: %4.2f", total, tax)

}
