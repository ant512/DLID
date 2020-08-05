package dlidparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseV2(data string, issuer string) (*DLIDLicense, error) {

	start, end, err := dataRangeV2(data)

	if end >= len(data) {
		return nil, errors.New("Payload location does not exist in data")
	}

	payload := data[start:end]

	if err != nil {
		return nil, err
	}
	return parseDataV2(payload, issuer)
}

func dataRangeV2(data string) (int, int, error) {

	if len(data) < 31 {
		return 0, 0, errors.New("Data is short")
	}
	start, err := strconv.Atoi(data[23:27])

	if err != nil {
		return 0, 0, errors.New("Data contains malformed payload location")
	}

	end, err := strconv.Atoi(data[27:31])

	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("Data contains malformed payload length -%v  %v", start, data[27:31]))
	}

	end += start

	return start, end, nil
}

func parseDataV2(licenceData string, issuer string) (*DLIDLicense, error) {

	// Version 1 of the DLID card spec was published in 2003.

	if !strings.HasPrefix(licenceData, "DL") {
		return nil, errors.New("Missing header in licence data chunk")
	}

	licenceData = licenceData[2:]

	components := strings.Split(licenceData, "\n")

	license := &DLIDLicense{}

	license.IssuerID = issuer
	license.IssuerName = issuers[issuer]

	for component := range components {

		if len(components[component]) < 3 {
			continue
		}

		identifier := components[component][0:3]
		data := components[component][3:]

		data = strings.Trim(data, " ")

		switch identifier {
		case "DCA":
			license.VehicleClass = data

		case "DCB":
			license.RestrictionCodes = data

		case "DCD":
			license.EndorsementCodes = data

		case "DCS":
			license.LastName = data

		case "DCT":

			// This field contains all of the licencee's names except last
			// name.  The V2 spec doc doesn't specify how the names are
			// separated and doesn't provide an example (unlike the 2000
			// doc).  Wisconsin use a space and Virginia use a comma.  This
			// is why standards have to be adequately documented.

			separator := " "

			if strings.Index(data, separator) == -1 {
				separator = ","
			}

			names := strings.Split(data, separator)

			license.FirstName = names[0]

			if len(names) > 1 {
				license.MiddleNames = names[1:]
			}

		case "DAG":
			license.Street = data

		case "DAI":
			license.City = data

		case "DAJ":
			license.State = data

		case "DAK":
			license.Postal = data

		case "DAQ":
			license.CustomerID = data

		case "DBB":
			license.DateOfBirth = parseDateV2(data)

		case "DBC":

			// According to the spec, the standard dropped M/F and two of
			// the ANSI D-20 gender codes in this revision.  The only
			// permissible values are now "1" and "2".

			switch data {
			case "1":
				license.Sex = DriverSexMale
			case "2":
				license.Sex = DriverSexFemale
			default:
				license.Sex = DriverSexNone
			}
		}
	}

	return license, nil
}

func parseDateV2(data string) time.Time {

	// Sooo, let me get this straight.  They switched from a reasonably-standard
	// and universal date format (yyyyMMdd) to the bizarre US lumpy format
	// (MMddyyyy)?  What were they thinking!?
	if len(data) != 8 {
		return time.Unix(0, 0)
	}
	t, err := time.Parse("01022006", data)
	if err != nil {
		return time.Unix(0, 0)
	}
	return t
}
