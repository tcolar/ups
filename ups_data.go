// History: Mar 5 2014 tcolar Creation
package ups

// Structures to unmarshall the Ups Xml answer into
// See specs here: https://www.ups.com/upsdeveloperkit/downloadresource?loc=en_US

// TrackResponse : Track reply (root) (xml:"TrackResponse")
type TrackResponse struct {
	Response Response
	Shipment Shipment
}

// Failed is a shortcut to check if the response is failed or not
func (r TrackResponse) Failed() bool {
	return r.Response.ResponseStatusCode != "1"
}

// TrackingNumber is a shortcut to the tracking number
func (r TrackResponse) TrackingNumber() string {
	return r.Shipment.Package.TrackingNumber
}

type Response struct {
	TransactionReference      TransactionReference
	ResponseStatusCode        string
	ResponseStatusDescription string
	Error                     Error
}

type TransactionReference struct {
	CustomerContext       string
	TransactionIdentifier string
	XpciVersion           string
}

type Shipment struct {
	InquiryNumber                CodeDescVal
	ShipmentIdentificationNumber string
	ShipmentType                 CodeDesc
	CandidateBookmark            string
	Shipper                      Shipper
	ShipTo                       ShipTo
	ShipmentWeight               ShipmentWeight
	Service                      CodeDesc
	ReferenceNumber              CodeDescVal
	CurrentStatus                CodeDesc
	PickupDate                   string
	DeliveryDetails              DeliveryDetails
	Volume                       CodeDescVal
	BillToName                   string
	// TODO ..... more
	Package Package
}

type Package struct {
	// TODO : more
	TrackingNumber string
}

type DeliveryDetails struct {
	DeliveryDate           string
	Date                   string
	Time                   string
	ServiceCenter          string
	City                   string
	StateProvinceCode      string
	DeliveryDateTime       []DeliveryDateTime
	PickUpServiceCenter    Address
	NumberOfPieces         string
	NumberOfPallets        string
	ShipmentServiceOptions ShipmentServiceOptions
}

type ShipmentServiceOptions struct {
	// TODO : more
}

type DeliveryDateTime struct {
	Type CodeDesc
	Date string
	Time string
}

type ShipmentWeight struct {
	UnitOfMeasurement UnitOfMeasurement
}

type UnitOfMeasurement struct {
	Code        string
	Description string
	Weight      string
}

type ShipTo struct {
	Address Address
}

type Shipper struct {
	ShipperNumber string
	Address       Address
}

type Address struct {
	AddressLine1      string
	AddressLine2      string
	AddressLine3      string
	City              string
	StateProvinceCode string
	PostalCode        string
	CountryCode       string
}

type CodeDesc struct {
	Code        string
	Description string
}

type CodeDescVal struct {
	Code        string
	Description string
	Value       string
}

type Error struct {
	ErrorSeverity       string
	ErrorCode           string
	ErrorDescription    string
	MinimumRetrySeconds string
	ErrorLocation       []ErrorLocation
	ErrorDigest         []string
}

type ErrorLocation struct {
	ErrorLocationElementName   string
	ErrorLocationAttributeName string
}
