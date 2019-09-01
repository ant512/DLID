package dlidparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const coloradoIssuerID string = "636020"
const connecticutIssuerID string = "636006"
const illinoisIssuerID string = "636035"
const massachusettsIssuerID string = "636002"
const southCarolinaIssuerID string = "636005"
const tennesseeIssuerID string = "636053"

func parseV1(data string, issuer string) (*DLIDLicense, error) {
	start, _, _ := dataRangeV1(data)
	if start == -1 {
		start, _, _ = dataRangeV1(data)
	}
	//yeah this is a hack
	if start > len(data) && strings.Index(data, "DAB") != -1 {
		start = strings.Index(data, "DAB")
	}
	if start > len(data) {
		return nil, fmt.Errorf("couldn't find start of payload")
	}
	end := len(data) - 1

	payload := data[start:end]
	return parseDataV1(payload, issuer)
}

func dataRangeV1(data string) (int, int, error) {

	start, err := strconv.Atoi(data[21:25])

	if err != nil {
		return 0, 0, errors.New("Data contains malformed payload location")
	}

	end, err := strconv.Atoi(data[25:29])

	if err != nil {
		end = len(data) - 1
	} else {
		end += start
	}
	return start, end, nil
}

func parseDataV1(licenceData string, issuer string) (*DLIDLicense, error) {

	// Version 1 of the DLID card spec was published in 2000.  As of 2012, it is
	// the version used in Colorado.

	// We want to strip off the "DL" chunk identifier, but every other state has
	// managed to screw this up too.  Rather than handle this on a
	// state-by-state basis, we'll check to see what's at the target location
	// and handle it appropriately.

	if strings.HasPrefix(licenceData, "DL") {

		// POMG!  They actually got it right!
		licenceData = licenceData[2:]
	} else if strings.HasPrefix(licenceData, "L") {

		// Either the guys in South Carolina can't count or they don't consider
		// the "DL" header part of the licence data.  In either case, their
		// offset is off by one.
		licenceData = licenceData[1:]
	} else {

		// Honestly, the spec really isn't that hard to follow.  I have no idea
		// why just about every implementation gets it wrong.  Massachusetts,
		// Connecticut and Pennsylvania don't include the "DL" chunk header in
		// at least some of their licenses.
		//
		// This else block is here just so I can grumble about badly-implemented
		// specs.
	}

	components := strings.Split(licenceData, "\n")

	license := &DLIDLicense{}

	license.IssuerID = issuer
	license.IssuerName = issuers[issuer]

	// Country is always USA for V1 licenses
	license.Country = "USA"

	for component := range components {

		if len(components[component]) < 3 {
			continue
		}

		identifier := components[component][0:3]
		data := components[component][3:]

		data = strings.Trim(data, " ")

		switch identifier {
		case "DAR":
			license.VehicleClass = data
		case "DAS":
			license.RestrictionCodes = data
		case "DAT":
			license.EndorsementCodes = data
		case "DAA":
			// Early versions of the Colorado implementation screwed up the
			// delimiter - they use a space instead of the specified comma.
			separator := " "
			if strings.Index(data, separator) == -1 {
				separator = ","
			}
			names := strings.Split(data, separator)
			// According to the spec, names are ordered LAST,FIRST,MIDDLE.
			// However, the geniuses in the Colorado and Tennessee DMVs order it
			// FIRST,MIDDLE,LAST.  We'll use the issuer ID number to
			// identify Colorado and adjust appropriately.  Issuer IDs can
			// be found here:
			//
			// http://www.aamva.org/IIN-and-RID/

			if issuer == coloradoIssuerID || issuer == tennesseeIssuerID {
				// Colorado's backwards formatting style...
				license.FirstName = names[0]

				if len(names) > 2 {
					license.MiddleNames = names[1 : len(names)-1]
					license.LastName = names[len(names)-1]
				} else if len(names) > 1 {
					license.LastName = names[1]
				}
			} else {

				// Everyone else, hopefully.
				license.LastName = names[0]

				if len(names) > 1 {
					license.FirstName = names[1]

					if len(names) > 2 {
						license.MiddleNames = names[2:]
					}
				}
			}

		case "DAE":
			license.NameSuffix = data

		case "DAL":
			// Colorado screws up again: they omit the *required* DAG field and
			// substitute the optional DAL field in older licences.
			fallthrough
		case "DAG":
			license.Street = data
		case "DAN":
			// Again, old Colorado licences ignore the spec.
			fallthrough
		case "DAI":
			license.City = data
		case "DAO":
			// Colorado strikes again.  Honestly, what is the point in having a
			// spec if you don't follow it?
			fallthrough
		case "DAJ":
			license.State = data
		case "DAP":
			// More Colorado shenanigans.
			fallthrough
		case "DAK":
			// Colorado uses the 5-digit zip code.  South Carolina uses the
			// 5 digit zip code plus the +4 extension all smooshed together
			// into one long string.  Massachusetts uses the 5 digit zip
			// plus the +4 extension separated by "-".  The zip is
			// apparently never written like that and always uses "+" as a
			// separator.  Who knows what other states managed to
			// accomplish.  At this point your dedicated programmer admits
			// defeat in trying to untangle the incredible mess implemented
			// in this single field; we'll just show the zip as it is
			// stored.
			license.Postal = strings.Trim(data, " ")
		case "DAQ":
			license.CustomerID = data
		case "DBA":
			license.ExpiryDate = parseDateV1(data)
		case "DBB":
			license.DateOfBirth = parseDateV1(data)
		case "DBC":
			// Sex can be stored as M/F if it uses the DLID code.  It could
			// also be stored as 0/1/2/9 if it uses the ANSI D-20 codes,
			// available here:
			//
			// http://www.aamva.org/ANSI-D20-Standard-for-Traffic-Records-Systems/

			switch data {
			case "M":
				fallthrough
			case "1":
				license.Sex = DriverSexMale
			case "F":
				fallthrough
			case "2":
				license.Sex = DriverSexFemale
			default:
				license.Sex = DriverSexNone
			}
		case "DBD":
			license.IssueDate = parseDateV1(data)
		case "DBK":
			// Optional and probably not available
			license.SocialSecurityNumber = data
		case "DAB":
			license.LastName = data
		case "DAC":
			license.FirstName = data
		case "DAD":
			license.MiddleNames = []string{data}
			/*
				default:
					fmt.Printf("Unknown: %v : %v\n", identifier, data)
			*/
		}
	}
	return license, nil
}

func parseDateV1(data string) time.Time {

	year, err := strconv.Atoi(data[:4])

	if err != nil {
		return time.Unix(0, 0)
	}

	month, err := strconv.Atoi(data[4:6])

	if err != nil {
		return time.Unix(0, 0)
	}

	day, err := strconv.Atoi(data[6:8])

	if err != nil {
		return time.Unix(0, 0)
	}

	location, _ := time.LoadLocation("UTC")

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)
}
