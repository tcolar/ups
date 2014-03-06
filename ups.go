// History: Nov 20 13 tcolar Creation

// ups provides access to some UPS XML api's
package ups

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	// Convenience constants for standard ups API url's
	UPS_API_URL       = "https://www.ups.com/ups.app/xml"
	UPS_API_TEST_URL  = "https://wwwcie.ups.com/ups.app/xml" // test server
	UPS_TEST_TRACKING = "1Z12345E1512345676"
)

// Utility to retrieve data from Ups XML API
type Ups struct {
	LicenseNumber, AccountNumber, Login, Password string
	UpsUrl                                        string
}

// Return tracking info for a specific UPS tracking number
func (u Ups) TrackByNumber(trackingNo string) (reply TrackResponse, err error) {
	reqXml := xmlNumberTracking(u, trackingNo)
	content, err := u.PostXml(u.UpsUrl+"/Track", reqXml)
	if err != nil {
		return reply, err
	}
	return u.ParseTrackReply(content)
}

// Return tracking info for a specific shipper reference
// ShipperRef is usually an order ID or other unique identifier
func (u Ups) TrackByShipperRef(shipperRef string) (reply TrackResponse, err error) {
	reqXml := xmlRefTracking(u, shipperRef)
	content, err := u.PostXml(u.UpsUrl+"/Track", reqXml)
	if err != nil {
		return reply, err
	}
	return u.ParseTrackReply(content)
}

// Unmarshal XML response into a TrackReply
func (u Ups) ParseTrackReply(xmlResp []byte) (reply TrackResponse, err error) {
	//log.Printf("%s", xmlResp)
	resp := TrackResponse{}
	err = xml.Unmarshal(xmlResp, &resp)
	return resp, err
}

// Post Xml to UPS API and return response
func (u Ups) PostXml(url string, xml string) (content []byte, err error) {
	resp, err := http.Post(url, "text/xml", strings.NewReader(xml))
	if err != nil {
		return content, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Dump some of the query resuts as an example
func Dump(resp TrackResponse) {
	// Dummy example of using the data
	log.Printf("Successs : %t", !resp.Failed())
	if !resp.Failed() {
		tracking := resp.TrackingNumber()
		log.Printf("Tracking Number: %s", tracking)
		log.Printf("Reference : %s", resp.Shipment.ReferenceNumber)
	} else {
		log.Fatal(resp)
	}
}
