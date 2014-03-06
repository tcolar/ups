// History: Mar 5 2014 tcolar Creation
package ups

// Structures to unmarshall the Ups Xml answer into
// See specs here: https://www.ups.com/upsdeveloperkit/downloadresource?loc=en_US

// TrackResponse : Track reply (root) (xml:"TrackResponse")
type TrackResponse struct {
	Response Response
	Shipment []Shipment
}

// Failed is a shortcut to check if the response is failed or not
func (r TrackResponse) Failed() bool {
	return r.Response.ResponseStatusCode != "1"
}

// TrackingNumber is a shortcut to the tracking number
func (r TrackResponse) TrackingNumber() string {
	return r.Shipment[0].Package.TrackingNumber
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
	ShipmentType                 CodeDescVal
	CandidateBookmark            string
	Shipper                      Shipper
	ShipTo                       ShipTo
	ShipmentWeight               Weight
	Service                      CodeDescVal
	ReferenceNumber              CodeDescVal
	CurrentStatus                CodeDescVal
	PickupDate                   string
	DeliveryDetails              DeliveryDetails
	Volume                       CodeDescVal
	BillToName                   string
	COD                          COD
	EstimatedDeliveryDetails     DeliveryDetails
	SignedForByName              string
	Activity                     []ShipmentActivity
	DescriptionOfGoods           string
	CargoReady                   DateTime
	Manifest                     DateTime
	CarrierActivityInformation   []CarrierActivityInformation
	ScheduledDeliveryDate        string
	ScheduledDeliveryTime        string
	FileNumber                   string
	Appointment                  Appointment
	DeliveryDateUnavailable      TypeDesc
	Package                      Package
}

type Package struct {
	TrackingNumber          string
	RescheduledDeliveryDate string
	RescheduledDeliveryTime string
	Redirect                Redirect
	Reroute                 AddressHolder
	ReturnTo                AddressHolder
	PackageServiceOptions   PackageServiceOptions
	Activity                []PkgActivity
	AlternateTrackingInfo   []TypeDescVal
	Accessorial             []CodeDescVal
}

type PkgActivity struct {
	AlternateTrackingInfo   []TypeDescVal
	ActivityLocation        PkgActivityLocation
	Status                  Status
	Date                    string
	Time                    string
	NextScheduleActivity    DateTime
	Message                 []CodeDescVal
	PackageWeight           Weight
	ReferenceNumber         []CodeDescVal
	ProductType             CodeDescVal
	LocationAssured         string
	AlternateTrackingNumber []string
}

type Status struct {
	StatusType CodeDescVal
	StatusCode CodeHolder
}

type CodeHolder struct {
	Code string
}

type PkgActivityLocation struct {
	Address               Address
	AddressArtifactFormat AddressArtifact
	TransportFacility     TypeCode
	Code                  string
	Description           string
	SignedForByName       string
	PODLetter             PODLetter
}

type PODLetter struct {
	HTMLImage                      string
	ElectronicDeliveryNotification NameHolder
}

type NameHolder struct {
	Name string
}

type AddressArtifact struct {
	StreetNumberLow    string
	StreetPrefix       string
	StreetName         string
	StreetSuffix       string
	StreetType         string
	PoliticalDivision1 string
	PoliticalDivision2 string
	PostcodePrimaryLow string
	CountryCode        string
}

type PackageServiceOptions struct {
	SignatureRequired        CodeDescVal
	ImportControl            string
	CommercialInvoiceRemoval string
	UPScarbonneutral         string
	USPSPICNumber            string
	ExchangeBased            string
	PackAndCollect           string
}

type AddressHolder struct {
	Address Address
}

type Redirect struct {
	CompanyName  string
	LocationID   string
	PickupDate   string
	UPSAPAddress Address
}

type TypeDesc struct {
	Type        string
	Description string
}

type Appointment struct {
	Made      DateTime
	Requested DateTime
	BeginTime string
	EndTime   string
}

type CarrierActivityInformation struct {
	CarrierId       string
	Description     string
	Status          string
	Arrival         DateTime
	Departure       DateTime
	OriginPort      string
	DestinationPort string
}

type ShipmentActivity struct {
	ActivityLocation       []ActivityLocation
	Description            string
	Date                   string
	Time                   string
	Trailer                string
	OriginPortDetails      PortDetails
	DestinationPortDetails PortDetails
}

type ActivityLocation struct {
	Address []Address
}

type PortDetails struct {
	OriginPort         string
	EstimatedDeparture DateTime
}

type COD struct {
	CODAmount Ammount
}

type Ammount struct {
	CurrencyCode  string
	MonetaryValue string
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

type DateTime struct {
	Date string
	Time string
}

type DeliveryDateTime struct {
	Type CodeDescVal
	Date string
	Time string
}

type Weight struct {
	UnitOfMeasurement UnitOfMeasurement
	Weight            string
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

type TypeCode struct {
	Code string
	Type string
}

type CodeDescVal struct {
	Code        string
	Description string
	Value       string
}

type TypeDescVal struct {
	Type        string
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
