package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
			fmt.Println(sCFDI.ToString())
			tax += sCFDI.Impuestos
			total += sCFDI.Total
		}
	}
	fmt.Printf("Total:%.2f\n  Tax: %4.2f", total, tax)

	/*xmlFile, err := os.Open("Amz_3dMaterial.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ReadCorrecctly the file:")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	*/
	//var cfdiData

	//var cfdiData cfdi.Cfdi

	//data := make(map[string]interface{})
	/*xmlError := xml.Unmarshal([]byte(byteValue), &cfdiData)
	if xmlError != nil {
		log.Fatal(xmlError)
	}*/

	/*sCFDI := cfdi.GetShortCFDI(byteValue)

	fmt.Println(sCFDI.ToString())
	*/
	/*
		fmt.Println((*cfdiData.Fecha))
		fmt.Println((*cfdiData.Total))
		fmt.Println((*cfdiData.Subtotal))

		fmt.Println((*cfdiData.CfdiEmisor.RFC))
		fmt.Println((*cfdiData.CfdiReceptor.RFC))
		fmt.Println((*cfdiData.CfdiImpuestos.Totalimpuestostrasladados))
	*/
}
