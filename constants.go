package api

// TODO: make truly distinct values?

import (
	"fmt"
	"strconv"
)

type ReportID uint8

const (
	ClientsReportID            ReportID = 1
	NetworkReportID            ReportID = 2
	ApplicationsReportID       ReportID = 3
	WLANsReportID              ReportID = 4
	AirtimeUtilizationReportID ReportID = 5
	SessionsSummaryReportID    ReportID = 6
	OverviewReportID           ReportID = 7
	APsRebootReportID          ReportID = 8
	AccessPointsReportID       ReportID = 9
)

// String returns a string version of the report ID
func (id ReportID) String() string {
	if ClientsReportID <= id && id <= AccessPointsReportID {
		return strconv.Itoa(int(id))
	}
	panic(fmt.Sprintf("Unknown Report ID seen: %d", id))
}

type ClientsReportSectionID uint8

const (
	ClientsOverviewSectionID                  ClientsReportSectionID = 12
	Top10UniqueClientsByTrafficSectionID      ClientsReportSectionID = 13
	ClientsDetailsSectionID                   ClientsReportSectionID = 14
	UniqueClientsTrendOverTimeChartSectionID  ClientsReportSectionID = 15
	UniqueClientsTrendOverTimeTableSectionID  ClientsReportSectionID = 16
	TopClientsByTrafficPercentileSectionID    ClientsReportSectionID = 17
	Top5OSByClientsCountSectionID             ClientsReportSectionID = 18
	Top10ManufacturersByClientsCountSectionID ClientsReportSectionID = 19
)

func (id ClientsReportSectionID) String() string {
	if ClientsOverviewSectionID <= id &&  id <= Top10ManufacturersByClientsCountSectionID {
		return strconv.Itoa(int(id))
	}
	panic(fmt.Sprintf("Unknown Clients Report Section ID seen: %d", id))
}


