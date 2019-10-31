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

	for _, f := range files {
		if strings.Contains(f.Name(), ".xml") {
			xmlFile, err := os.Open(f.Name())
			if err != nil {
				fmt.Println("Error reading:", f.Name())
				fmt.Println(err)
			}
			defer xmlFile.Close()
			byteValue, _ := ioutil.ReadAll(xmlFile)
			sCFDI := cfdi.GetShortCFDI(byteValue)
			fmt.Println(sCFDI.ToString())
		}
	}

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
