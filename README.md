Ups
=====

Some UPS (United Parcel Services) API support for GoLang (ATM mostly for tracking)

I did not bother dealing with all the details and only created what I needed so far.

I might add more over time but for now it provides:
- Retrieving Tracking info by either:
  Tracking number, PO number, or shipper reference number (~order ID)
  The data is unmarshalled from XML into Go structures for more practical usage.

See [ups_example.go](ups_example.go) for usage examples

Note that you will need an API key and Password as well as Accont and Meter numbers from Ups.
See: https://www.ups.com/upsdeveloperkit/


