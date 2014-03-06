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
func (u Ups) TrackByNumber(trackingNo string) (reply TrackReply, err error) {
	reqXml := xmlNumberTracking(u, trackingNo)
	content, err := u.PostXml(u.UpsUrl+"/Track", reqXml)
	if err != nil {
		return reply, err
	}
	log.Printf("%s", content)
	return u.ParseTrackReply(content)
}

// Return tracking info for a specific shipper reference
// ShipperRef is usually an order ID or other unique identifier
func (u Ups) TrackByShipperRef(shipperRef string) (reply TrackReply, err error) {
	reqXml := xmlRefTracking(u, shipperRef)
	content, err := u.PostXml(u.UpsUrl+"/Track", reqXml)
	if err != nil {
		return reply, err
	}
	log.Printf("%s", content)
	return u.ParseTrackReply(content)
}

// Unmarshal XML response into a TrackReply
func (u Ups) ParseTrackReply(xmlResp []byte) (reply TrackReply, err error) {
	data := struct {
		Reply TrackReply `xml:"Body>TrackReply"`
	}{}
	err = xml.Unmarshal(xmlResp, &data)
	return data.Reply, err
}

// Post Xml to UPS API and return response
func (u Ups) PostXml(url string, xml string) (content []byte, err error) {
	log.Print(url)
	log.Print(xml)
	resp, err := http.Post(url, "text/xml", strings.NewReader(xml))
	if err != nil {
		return content, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Dump some of the query resuts as an example
func Dump(reply TrackReply) {
	// Dummy example of using the data
	log.Printf("Successs : %t", !reply.Failed())
	if !reply.Failed() {
		tracking := reply.CompletedTrackDetails[0].TrackDetails[0].TrackingNumber
		log.Printf("Tracking Number: %s", tracking)
		log.Print(reply.CompletedTrackDetails[0].TrackDetails[0].ActualDeliveryAddress)
	}
}
