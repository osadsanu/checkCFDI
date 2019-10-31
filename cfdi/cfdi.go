// This file was generated from xml Schema using quicktype, do not modify it directly.
// To parse and unparse this xml data, add this code to your project and do:
//
//    cfdi, err := UnmarshalCfdi(bytes)
//    bytes, err = cfdi.Marshal()

//@(.+)(,omitempty)
//$1,attr$2
package cfdi

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

//GetShortCFDI get the short version of the class
func GetShortCFDI(data []byte) ShortCFDI {
	var longData Cfdi
	xmlError := xml.Unmarshal([]byte(data), &longData)
	if xmlError != nil {
		log.Fatal(xmlError)
	}
	auxTotal, _ := strconv.ParseFloat(*longData.Total, 64)
	auxSubtotal, _ := strconv.ParseFloat(*longData.Subtotal, 64)
	auxImpuestos, _ := strconv.ParseFloat(*longData.CfdiImpuestos.Totalimpuestostrasladados, 64)
	sCFDI := ShortCFDI{
		Fecha:       *longData.Fecha,
		Total:       auxTotal,
		SubTotal:    auxSubtotal,
		RfcEmisor:   *longData.CfdiEmisor.RFC,
		RfcReceptor: *longData.CfdiReceptor.RFC,
		Impuestos:   auxImpuestos,
	}
	return sCFDI
}

//ShortCFDI Contains only a few values from CFDI
type ShortCFDI struct {
	Fecha       string
	Total       float64
	SubTotal    float64
	Impuestos   float64
	RfcReceptor string
	RfcEmisor   string
}

//ToString concadenates the ShortCFDI
func (s ShortCFDI) ToString() string {
	return fmt.Sprintf("%s\t%s->%s\t%.2f + %.2f0= %.2f", s.Fecha, s.RfcEmisor, s.RfcReceptor, s.SubTotal, s.Impuestos, s.Total)

}

//Cfdi struct container
type Cfdi struct {
	XMLName            xml.Name                      `xml:"Comprobante"`
	Fecha              *string                       `xml:"Fecha,attr"`
	Certificado        *string                       `xml:"Certificado,attr"`
	Descuento          *string                       `xml:"Descuento,attr,omitempty"`
	Folio              *string                       `xml:"Folio,attr,omitempty"`
	Formapago          *string                       `xml:"Formapago,attr,omitempty"`
	Lugarexpedicion    *string                       `xml:"Lugarexpedicion,attr,omitempty"`
	Metodopago         *string                       `xml:"Metodopago,attr,omitempty"`
	Moneda             *string                       `xml:"Moneda,attr,omitempty"`
	Nocertificado      *string                       `xml:"Nocertificado,attr,omitempty"`
	Sello              *string                       `xml:"Sello,attr,omitempty"`
	Serie              *string                       `xml:"Serie,attr,omitempty"`
	Subtotal           *string                       `xml:"SubTotal,attr,omitempty"`
	Tipodecomprobante  *string                       `xml:"Tipodecomprobante,attr,omitempty"`
	Total              *string                       `xml:"Total,attr,omitempty"`
	Version            *string                       `xml:"Version,attr,omitempty"`
	Xmlns              *string                       `xml:"Xmlns,attr,omitempty"`
	XmlnsAmazonaddenda *string                       `xml:"Xmlns:amazonaddenda,attr,omitempty"`
	XmlnsCfdi          *string                       `xml:"Xmlns:cfdi,attr,omitempty"`
	XmlnsXsi           *string                       `xml:"Xmlns:xsi,attr,omitempty"`
	XsiSchemalocation  *string                       `xml:"Xsi:schemalocation,attr,omitempty"`
	CfdiEmisor         *CfdiEmisor                   `xml:"Emisor,omitempty"`
	CfdiReceptor       *CfdiReceptor                 `xml:"Receptor,omitempty"`
	CfdiImpuestos      *CfdiComprobanteCfdiImpuestos `xml:"Impuestos,omitempty"`
}
type CfdiEmisor struct {
	Nombre        *string `xml:"Nombre,attr,omitempty"`
	Regimenfiscal *string `xml:"RegimenFiscal,attr,omitempty"`
	RFC           *string `xml:"Rfc,attr,omitempty"`
}
type CfdiReceptor struct {
	RFC     *string `xml:"Rfc,attr,omitempty"`
	Usocfdi *string `xml:"UsocCFDI,attr,omitempty"`
}
type CfdiImpuestos struct {
	CfdiTraslados *CfdiTraslados `xml:"cfdi:traslados,omitempty"`
}
type CfdiComprobanteCfdiImpuestos struct {
	Totalimpuestostrasladados *string        `xml:"TotalImpuestosTrasladados,attr,omitempty"`
	CfdiTraslados             *CfdiTraslados `xml:"cfdi:traslados,omitempty"`
}

type CfdiComprobante struct {
	CfdiEmisor      *CfdiEmisor                   `xml:"cfdi:emisor,omitempty"`
	CfdiReceptor    *CfdiReceptor                 `xml:"cfdi:receptor,omitempty"`
	CfdiConceptos   *CfdiConceptos                `xml:"cfdi:conceptos,omitempty"`
	CfdiImpuestos   *CfdiComprobanteCfdiImpuestos `xml:"cfdi:impuestos,omitempty"`
	CfdiComplemento *CfdiComplemento              `xml:"cfdi:complemento,omitempty"`
	CfdiAddenda     *CfdiAddenda                  `xml:"cfdi:addenda,omitempty"`
}

type CfdiAddenda struct {
	AmazonaddendaElementosamazon *AmazonaddendaElementosamazon `xml:"amazonaddenda:elementosamazon,omitempty"`
}

type AmazonaddendaElementosamazon struct {
	AmazonaddendaLosatributos []AmazonaddendaLosatributo `xml:"amazonaddenda:losatributos"`
}

type AmazonaddendaLosatributo struct {
	Identificacionunica *string `xml:"identificacionunica,attr,omitempty"`
	Nombredelatributo   *string `xml:"nombredelatributo,attr,omitempty"`
	Valordelatributo    *string `xml:"valordelatributo,attr,omitempty"`
}

type CfdiComplemento struct {
	TfdTimbrefiscaldigital *TfdTimbrefiscaldigital `xml:"tfd:timbrefiscaldigital,omitempty"`
}

type TfdTimbrefiscaldigital struct {
	Fechatimbrado     *string `xml:"fechatimbrado,attr,omitempty"`
	Nocertificadosat  *string `xml:"nocertificadosat,attr,omitempty"`
	Rfcprovcertif     *string `xml:"rfcprovcertif,attr,omitempty"`
	Sellocfd          *string `xml:"sellocfd,attr,omitempty"`
	Sellosat          *string `xml:"sellosat,attr,omitempty"`
	UUID              *string `xml:"uuid,attr,omitempty"`
	Version           *string `xml:"version,attr,omitempty"`
	XmlnsTfd          *string `xml:"xmlns:tfd,attr,omitempty"`
	XsiSchemalocation *string `xml:"xsi:schemalocation,attr,omitempty"`
}

type CfdiConceptos struct {
	CfdiConcepto *CfdiConcepto `xml:"cfdi:concepto,omitempty"`
}

type CfdiConcepto struct {
	Cantidad         *string `xml:"cantidad,attr,omitempty"`
	Claveprodserv    *string `xml:"claveprodserv,attr,omitempty"`
	Claveunidad      *string `xml:"claveunidad,attr,omitempty"`
	Descripcion      *string `xml:"descripcion,attr,omitempty"`
	Descuento        *string `xml:"descuento,attr,omitempty"`
	Importe          *string `xml:"importe,attr,omitempty"`
	Noidentificacion *string `xml:"noidentificacion,attr,omitempty"`
	Unidad           *string `xml:"unidad,attr,omitempty"`
	Valorunitario    *string `xml:"valorunitario,attr,omitempty"`
}

type CfdiTraslados struct {
	CfdiTraslado *CfdiTraslado `xml:"cfdi:traslado,omitempty"`
}

type CfdiTraslado struct {
	Base       *string `xml:"base,attr,omitempty"`
	Importe    *string `xml:"importe,attr,omitempty"`
	Impuesto   *string `xml:"impuesto,attr,omitempty"`
	Tasaocuota *string `xml:"tasaocuota,attr,omitempty"`
	Tipofactor *string `xml:"tipofactor,attr,omitempty"`
}
