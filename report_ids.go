package api

// TODO: make truly distinct values?

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// TODO: The ID "system" here will have to be updated to allow for custom reports, at some point...

type ID interface {
	// Valid must test the value against a list of known good values for the type of ID
	Valid() bool
	// String must return the contained int value as as string
	String() string
	// Title must return the title of the report / section
	Title() string
}

const unknownID = -1

type ReportID int

const (
	ReportIDUnknown            ReportID = unknownID
	ReportIDClients            ReportID = 1
	ReportIDNetwork            ReportID = 2
	ReportIDApplications       ReportID = 3
	ReportIDWLANs              ReportID = 4
	ReportIDAirtimeUtilization ReportID = 5
	ReportIDSessionsSummary    ReportID = 6
	ReportIDOverview           ReportID = 7
	ReportIDAPsReboot          ReportID = 8
	ReportIDAccessPoints       ReportID = 9
)

// AsReportID will attempt to take the input and translate it into a known ReportID
func AsReportID(in interface{}) (ReportID, error) {
	if i, err := toInt(in); err != nil {
		return ReportIDUnknown, err
	} else if rid := ReportID(i); rid.Valid() {
		return rid, nil
	} else {
		return ReportIDUnknown, fmt.Errorf("\"%d\" is not a valid ReportID", i)
	}
}

func (id ReportID) String() string {
	return strconv.Itoa(int(id))
}

func (id ReportID) Valid() bool {
	return ReportIDClients <= id && id <= ReportIDAccessPoints
}

func (id ReportID) Title() string {
	switch id {
	case ReportIDClients:
		return "Clients Report"

	}
}

type ClientsSectionID int

const (
	ClientSectionIDUnknown                           ClientsSectionID = unknownID
	ClientsSectionIDOverview                         ClientsSectionID = 12
	ClientsSectionIDTop10UniqueClientsByTraffic      ClientsSectionID = 13
	ClientsSectionIDDetails                          ClientsSectionID = 14
	ClientsSectionIDUniqueTrendOverTimeChart         ClientsSectionID = 15
	ClientsSectionIDUniqueClientsTrendOverTimeTable  ClientsSectionID = 16
	ClientsSectionIDTopClientsByTrafficPercentile    ClientsSectionID = 17
	ClientsSectionIDTop5OSByClientsCount             ClientsSectionID = 18
	ClientsSectionIDTop10ManufacturersByClientsCount ClientsSectionID = 19
)

func AsClientsSectionID(in interface{}) (ClientsSectionID, error) {
	if i, err := toInt(in); err != nil {
		return ClientSectionIDUnknown, err
	} else if scid := ClientsSectionID(i); scid.Valid() {
		return scid, nil
	} else {
		return ClientSectionIDUnknown, fmt.Errorf("\"%d\" is not a valid ClientSectionID", i)
	}
}

func (id ClientsSectionID) String() string {
	if ClientsSectionIDOverview <= id && id <= ClientsSectionIDTop10ManufacturersByClientsCount {
		return strconv.Itoa(int(id))
	}
	panic(fmt.Sprintf("Unknown Clients Section ID seen: %d", id))
}

func (id ClientsSectionID) Valid() bool {
	return ClientsSectionIDOverview <= id && id <= ClientsSectionIDTop10ManufacturersByClientsCount
}

// TODO: expand this, maybe?
func toInt(in interface{}) (int, error) {
	var i int
	var err error
	if in == nil {
		return 0, errors.New("nil cannot be interpreted as a ReportID")
	} else {
		switch in.(type) {
		case string:
			if i, err = strconv.Atoi(in.(string)); err != nil {
				return unknownID, fmt.Errorf("unable to translate \"%s\" into ReportID: %s", in.(string), err)
			}
		case int:
			i = in.(int)
		case int8:
			i = int(in.(int8))
		case int16:
			i = int(in.(int16))
		case int32:
			i = int(in.(int32))
		case int64:
			i = int(in.(int64))
		case uint:
			i = int(in.(uint))
		case uint8:
			i = int(in.(uint8))
		case uint16:
			i = int(in.(uint16))
		case uint32:
			i = int(in.(uint32))
		case uint64:
			i = int(in.(uint64))

		default:
			return unknownID, fmt.Errorf("unable to translate \"%s\" into int", reflect.TypeOf(in))
		}
	}
	return i, nil
}
