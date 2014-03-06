// History: Mar 5 2014 tcolar Creation

package ups

import (
	"fmt"
)

// UPS Xml monkey patching

// Track by Tracking number
func xmlNumberTracking(u Ups, trackingNo string) string {
	return xmlHead(u) +
		fmt.Sprintf(`
<TrackRequest xml:lang="en-US">
 <Request>
   <TransactionReference>
     <CustomerContext>Test</CustomerContext>
     <XpciVersion>1.0</XpciVersion>
   </TransactionReference>
   <RequestAction>Track</RequestAction>
   <RequestOption>activity</RequestOption>
 </Request>
 <TrackingNumber>%s</TrackingNumber>
</TrackRequest>`, trackingNo)
}

// Track by shipper reference number
func xmlRefTracking(u Ups, ref string) string {
	return xmlHead(u) +
		fmt.Sprintf(`
<TrackRequest xml:lang="en-US">
 <Request>
   <TransactionReference>
     <CustomerContext>Test</CustomerContext>
     <XpciVersion>1.0</XpciVersion>
   </TransactionReference>
   <RequestAction>Track</RequestAction>
   <RequestOption>activity</RequestOption>
 </Request>
 <ShipperNumber>%s</ShipperNumber>
 <ReferenceNumber><Value>%s</Value></ReferenceNumber>
</TrackRequest>`, u.AccountNumber, ref)
}

func xmlHead(u Ups) string {
	return fmt.Sprintf(UPS_XML_HEAD,
		u.LicenseNumber, u.Login, u.Password)
}

const (
	UPS_XML_HEAD = `<?xml version="1.0"?>
<AccessRequest xml:lang="en-US">
   <AccessLicenseNumber>%s</AccessLicenseNumber>
   <UserId>%s</UserId>
   <Password>%s</Password>
</AccessRequest>`
)
