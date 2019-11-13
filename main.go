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

type ArgsFlags struct {
	OrderByDate     bool
	OrderByTotal    bool
	OrderByTax      bool
	RecursiveSearch bool
}

var cfdiList []cfdi.ShortCFDI
var tax float64
var total float64
var flags ArgsFlags

func main() {
	flags = checkArgs(os.Args)
	fmt.Println(flags)

	folderPath := "."
	//read Files from img folder
	checkForCFDI(folderPath)

	//order the list by date, total, and tax as params
	//TODO optimize: put the if before the sort function.
	sort.Slice(cfdiList, func(i, j int) bool {
		if flags.OrderByDate {
			return cfdiList[i].Fecha < cfdiList[j].Fecha
		}
		if flags.OrderByTotal {
			return cfdiList[i].Total < cfdiList[j].Total
		}
		if flags.OrderByTax {
			return cfdiList[i].Impuestos < cfdiList[j].Impuestos
		}
		return cfdiList[i].Fecha < cfdiList[j].Fecha

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
func checkArgs(args []string) ArgsFlags {
	flags := ArgsFlags{
		OrderByDate:     false,
		OrderByTotal:    false,
		OrderByTax:      false,
		RecursiveSearch: false,
	}
	orderFlagsCount := 0
	for _, flag := range args {
		if flag == "-r" {
			flags.RecursiveSearch = true
		}
		if flag == "-d" {
			flags.OrderByDate = true
			orderFlagsCount++
		}
		if flag == "-total" {
			flags.OrderByTotal = true
			orderFlagsCount++
		}
		if flag == "-tax" {
			flags.OrderByTax = true
			orderFlagsCount++
		}
	}
	if orderFlagsCount > 1 {
		fmt.Println("please only use 1 order flag per use \n-tax \t order by tax\n-d\t order by date\n-total\t order by total\n")
	}

	return flags
}

func checkForCFDI(fileName string) {
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fullpath := fmt.Sprintf("%s/%s", fileName, f.Name())
		//fmt.Println(fullpath)
		if strings.Contains(f.Name(), ".xml") { ///TODO: change the if to (the file has to end with .xml not just containt it)
			//fmt.Println("reading File:", f.Name())
			xmlFile, err := os.Open(fullpath)

			if err != nil {
				fmt.Println("Error reading:", fullpath)
				fmt.Println(err)
				defer xmlFile.Close()
				return
			}
			defer xmlFile.Close()

			byteValue, err := ioutil.ReadAll(xmlFile)
			if err != nil {
				fmt.Println("cant read: ", fullpath)
				return
			}
			sCFDI := cfdi.GetShortCFDI(byteValue)
			cfdiList = append(cfdiList, sCFDI)
			//fmt.Println(sCFDI.ToString())
			tax += sCFDI.Impuestos
			total += sCFDI.Total
		}
		if flags.RecursiveSearch {
			if f.IsDir() {
				aux := fmt.Sprintf("%s/%s", fileName, fullpath)
				fmt.Println("\t recursive path: ", aux)
				checkForCFDI(aux)
				/*fmt.Println("is dir! ", f.Name())
				subFiles, err := ioutil.ReadDir(f.Name())
				if err != nil {
					log.Fatal(err)
				}
				for _, f2 := range subFiles {
					aux := fmt.Sprintf("%s/%s", f.Name(), f2.Name())
					fmt.Println(aux)
					checkForCFDI(aux)

				}*/
			}
		}
	}

}
