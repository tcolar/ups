// History: Nov 20 13 tcolar Creation

package ups

import (
	"log"
)

// Examples
func main() {
	// You will need to fill in all those with your actual Ups web service data
	ups := Ups{
		UpsUrl:        UPS_API_TEST_URL,
		LicenseNumber: "FCCCAAA44335522DDFF", // example
		AccountNumber: "A5624589",            // example
		Login:         "YourUpsLogin",
		Password:      "YourUpsPassword",
	}

	trackByReference(ups, "replaceWithYourOrderRef")
	trackByNumber(ups, "1Z12345E1512345676")
}

// Looking up some tracking info by Ups tracking number
func trackByNumber(ups Ups, trackingNo string) {
	reply, err := ups.TrackByNumber(trackingNo)
	if err != nil {
		log.Fatal(err)
	}
	Dump(reply)
}

// Looking up some tracking info by reference
func trackByReference(ups Ups, ref string) {
	reply, err := ups.TrackByShipperRef(ref)
	if err != nil {
		log.Fatal(err)
	}
	Dump(reply)
}
