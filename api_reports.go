package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Reports struct {
	c *Client
}
type (
	FindReportsGet200ResponseSlice []*FindReportsGet200Response

	FindReportsGet200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource string      `json:"filterDataSource,omitempty"`
		ID               int         `json:"id,omitempty"`
		Title            string      `json:"title,omitempty"`
		URLSegmentName   string      `json:"urlSegmentName,omitempty"`
	}
)

func (r *Reports) FindReportsGet(ctx context.Context, filter *Filter) (*http.Response, *FindReportsGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports", true)
	if err = request.SetFilterQueryParameter(filter); err != nil {
		return nil, nil, err
	}
	out := new(FindReportsGet200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FindReportByIDGet200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource string      `json:"filterDataSource,omitempty"`
		ID               int         `json:"id,omitempty"`
		Title            string      `json:"title,omitempty"`
		URLSegmentName   string      `json:"urlSegmentName,omitempty"`
	}
)

func (r *Reports) FindReportByIDGet(ctx context.Context, reportID int, filter *Filter) (*http.Response, *FindReportByIDGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports/{id}", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetFilterQueryParameter(filter); err != nil {
		return nil, nil, err
	}
	out := new(FindReportByIDGet200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportExistsByID200Response struct {
		Exists bool `json:"exists"`
	}
)

func (r *Reports) ReportExistsByIDGet(ctx context.Context, reportID int) (*http.Response, *ReportExistsByID200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports/{id}/exists", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	out := new(ReportExistsByID200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FilterFacetsByReportIDGet200ResponseDataApMacsSlice []string

	FilterFacetsByReportIDGet200ResponseDataSsidsSlice []*FilterFacetsByReportIDGet200ResponseDataSsids

	FilterFacetsByReportIDGet200ResponseDataSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsSlice []*FilterFacetsByReportIDGet200ResponseDataSystems

	FilterFacetsByReportIDGet200ResponseDataSystemsApMacsSlice []string

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildren

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenApMacsSlice []string

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildren

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenApMacsSlice []string

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildren

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice []string

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSsids

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildren struct {
		ApMacs FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice `json:"apMacs,omitempty"`
		ID     *string                                                                            `json:"id,omitempty"`
		Ssids  FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice  `json:"ssids,omitempty"`
		Text   *string                                                                            `json:"text,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSsidsSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSsids

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildren struct {
		ApMacs   FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                                      `json:"id,omitempty"`
		Ssids    FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                                      `json:"text,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSsidsSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSsids

	FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsChildren struct {
		ApMacs   FilterFacetsByReportIDGet200ResponseDataSystemsChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children FilterFacetsByReportIDGet200ResponseDataSystemsChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                              `json:"id,omitempty"`
		Ssids    FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                              `json:"text,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystemsSsidsSlice []*FilterFacetsByReportIDGet200ResponseDataSystemsSsids

	FilterFacetsByReportIDGet200ResponseDataSystemsSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseDataSystems struct {
		ApMacs   FilterFacetsByReportIDGet200ResponseDataSystemsApMacsSlice   `json:"apMacs,omitempty"`
		Children FilterFacetsByReportIDGet200ResponseDataSystemsChildrenSlice `json:"children,omitempty"`
		ID       *string                                                      `json:"id,omitempty"`
		Ssids    FilterFacetsByReportIDGet200ResponseDataSystemsSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                      `json:"text,omitempty"`
	}

	FilterFacetsByReportIDGet200ResponseData struct {
		ApMacs  FilterFacetsByReportIDGet200ResponseDataApMacsSlice  `json:"apMacs,omitempty"`
		Ssids   FilterFacetsByReportIDGet200ResponseDataSsidsSlice   `json:"ssids,omitempty"`
		Systems FilterFacetsByReportIDGet200ResponseDataSystemsSlice `json:"systems,omitempty"`
	}

	FilterFacetsByReportIDGet200Response struct {
		Data *FilterFacetsByReportIDGet200ResponseData `json:"data,omitempty"`
	}
)

func (r *Reports) FilterFacetsByReportIDGet(ctx context.Context, reportID int, start, end time.Time, granularity string) (*http.Response, *FilterFacetsByReportIDGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports/{id}/facet/data", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetQueryParameter("start", start.Format(TimeFormat))
	request.SetQueryParameter("end", end.Format(TimeFormat))
	if granularity != "" {
		request.SetQueryParameter("granularity", granularity)
	}
	out := new(FilterFacetsByReportIDGet200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FilterFacetsByReportIDAPMACListPost200ResponseDataApMacsSlice []string

	FilterFacetsByReportIDAPMACListPost200ResponseData struct {
		ApMacs FilterFacetsByReportIDAPMACListPost200ResponseDataApMacsSlice `json:"apMacs,omitempty"`
	}

	FilterFacetsByReportIDAPMACListPost200Response struct {
		Data *FilterFacetsByReportIDAPMACListPost200ResponseData `json:"data,omitempty"`
	}
)

func (r *Reports) FilterFacetsByReportIDAPMACListPost(ctx context.Context, reportID int, query *Query) (*http.Response, *FilterFacetsByReportIDAPMACListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/apmac", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(FilterFacetsByReportIDAPMACListPost200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FilterFacetsByReportIDSSIDListPost200ResponseDataSlice []string

	FilterFacetsByReportIDSSIDListPost200Response struct {
		Data FilterFacetsByReportIDSSIDListPost200ResponseDataSlice `json:"data,omitempty"`
	}
)

func (r *Reports) FilterFacetsByReportIDSSIDListPost(ctx context.Context, reportID int, query *Query) (*http.Response, *FilterFacetsByReportIDSSIDListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(FilterFacetsByReportIDSSIDListPost200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FilterFacetsByReportIDSystemsListPost200ResponseSlice []*FilterFacetsByReportIDSystemsListPost200Response

	FilterFacetsByReportIDSystemsListPost200ResponseChildrenSlice []*FilterFacetsByReportIDSystemsListPost200ResponseChildren

	FilterFacetsByReportIDSystemsListPost200ResponseChildren struct {
		Children   *bool   `json:"children,omitempty"`
		Data       *string `json:"data,omitempty"`
		FilterText *string `json:"filterText,omitempty"`
		ID         *string `json:"id,omitempty"`
		Text       *string `json:"text,omitempty"`
	}

	FilterFacetsByReportIDSystemsListPost200ResponseState struct {
		Opened *bool `json:"opened,omitempty"`
	}

	FilterFacetsByReportIDSystemsListPost200Response struct {
		Children FilterFacetsByReportIDSystemsListPost200ResponseChildrenSlice `json:"children,omitempty"`
		ID       *string                                                       `json:"id,omitempty"`
		State    *FilterFacetsByReportIDSystemsListPost200ResponseState        `json:"state,omitempty"`
		Text     *string                                                       `json:"text,omitempty"`
	}
)

func (r *Reports) FilterFacetsByReportIDSystemsListPost(ctx context.Context, reportID int, query *Query) (*http.Response, *FilterFacetsByReportIDSystemsListPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(FilterFacetsByReportIDSystemsListPost200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	FindSectionsByReportIDGet200ResponseSlice []*FindSectionsByReportIDGet200Response

	FindSectionsByReportIDGet200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
		Metric      *string `json:"metric,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutColorsSlice []string

	FindSectionsByReportIDGet200ResponseLayoutColumnsSlice []*FindSectionsByReportIDGet200ResponseLayoutColumns

	FindSectionsByReportIDGet200ResponseLayoutColumns struct {
		Color           *string `json:"color,omitempty"`
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutHeadersSlice []*FindSectionsByReportIDGet200ResponseLayoutHeaders

	FindSectionsByReportIDGet200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutHeadersContent struct {
		Formats *FindSectionsByReportIDGet200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                          `json:"text,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutHeaders struct {
		Component *string                                                   `json:"component,omitempty"`
		Content   *FindSectionsByReportIDGet200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                                   `json:"name,omitempty"`
		Options   map[string]string                                         `json:"options,omitempty"`
		Query     *string                                                   `json:"query,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutSubSectionsSlice []*FindSectionsByReportIDGet200ResponseLayoutSubSections

	FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayoutSeriesSlice []*FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayoutSeries

	FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayoutSeries struct {
		Area   *bool   `json:"area,omitempty"`
		Color  *string `json:"color,omitempty"`
		Key    *string `json:"key,omitempty"`
		Values *string `json:"values,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayout struct {
		Series    FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayoutSeriesSlice `json:"series,omitempty"`
		Width     *string                                                                `json:"width,omitempty"`
		XAxisType *string                                                                `json:"xAxisType,omitempty"`
		YAxisType *string                                                                `json:"yAxisType,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayoutSubSections struct {
		Component *string                                                      `json:"component,omitempty"`
		Layout    *FindSectionsByReportIDGet200ResponseLayoutSubSectionsLayout `json:"layout,omitempty"`
		Title     *string                                                      `json:"title,omitempty"`
	}

	FindSectionsByReportIDGet200ResponseLayout struct {
		Colors      FindSectionsByReportIDGet200ResponseLayoutColorsSlice      `json:"colors,omitempty"`
		Columns     FindSectionsByReportIDGet200ResponseLayoutColumnsSlice     `json:"columns,omitempty"`
		Format      *string                                                    `json:"format,omitempty"`
		Headers     FindSectionsByReportIDGet200ResponseLayoutHeadersSlice     `json:"headers,omitempty"`
		SubSections FindSectionsByReportIDGet200ResponseLayoutSubSectionsSlice `json:"subSections,omitempty"`
		Width       *string                                                    `json:"width,omitempty"`
	}

	FindSectionsByReportIDGet200Response struct {
		Component         *string                                                `json:"component,omitempty"`
		DefaultParameters *FindSectionsByReportIDGet200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                                   `json:"id,omitempty"`
		Layout            *FindSectionsByReportIDGet200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                                   `json:"order,omitempty"`
		QueryName         *string                                                `json:"queryName,omitempty"`
		ReportID          *int                                                   `json:"reportId,omitempty"`
		Title             *string                                                `json:"title,omitempty"`
		URL               interface{}                                            `json:"url,omitempty"`
	}
)

func (r *Reports) FindSectionsByReportIDGet(ctx context.Context, reportID int, filter *Filter) (*http.Response, *FindSectionsByReportIDGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetFilterQueryParameter(filter)
	out := new(FindSectionsByReportIDGet200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	GetSectionByReportIDGet200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
	}

	GetSectionByReportIDGet200ResponseLayoutColumnsSlice []*GetSectionByReportIDGet200ResponseLayoutColumns

	GetSectionByReportIDGet200ResponseLayoutColumns struct {
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	GetSectionByReportIDGet200ResponseLayoutHeadersSlice []*GetSectionByReportIDGet200ResponseLayoutHeaders

	GetSectionByReportIDGet200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	GetSectionByReportIDGet200ResponseLayoutHeadersContent struct {
		Formats *GetSectionByReportIDGet200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                        `json:"text,omitempty"`
	}

	GetSectionByReportIDGet200ResponseLayoutHeaders struct {
		Component *string                                                 `json:"component,omitempty"`
		Content   *GetSectionByReportIDGet200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                                 `json:"name,omitempty"`
		Options   map[string]string                                       `json:"options,omitempty"`
		Query     *string                                                 `json:"query,omitempty"`
	}

	GetSectionByReportIDGet200ResponseLayout struct {
		Columns GetSectionByReportIDGet200ResponseLayoutColumnsSlice `json:"columns,omitempty"`
		Headers GetSectionByReportIDGet200ResponseLayoutHeadersSlice `json:"headers,omitempty"`
		Width   *string                                              `json:"width,omitempty"`
	}

	GetSectionByReportIDGet200Response struct {
		Component         *string                                              `json:"component,omitempty"`
		DefaultParameters *GetSectionByReportIDGet200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                                 `json:"id,omitempty"`
		Layout            *GetSectionByReportIDGet200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                                 `json:"order,omitempty"`
		QueryName         *string                                              `json:"queryName,omitempty"`
		ReportID          *int                                                 `json:"reportId,omitempty"`
		Title             *string                                              `json:"title,omitempty"`
		URL               interface{}                                          `json:"url,omitempty"`
	}
)

func (r *Reports) FindSectionByReportIDGet(ctx context.Context, reportID, sectionID int, filter *Filter) (*http.Response, *GetSectionByReportIDGet200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections/{fk}", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetPathParameter("fk", strconv.Itoa(sectionID))
	request.SetFilterQueryParameter(filter)
	out := new(GetSectionByReportIDGet200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

// GetDataPost will, on success, always return an object like:
//
//		{
//			"data": [],
//			"metadata": {}
//		}
//
// The exact structure of "data" and "metadata" vary greatly depending upon the report id and section queried, and
// therefore is not modeled here.  It is expected that the caller either doesn't care or knows which specific report
// they're interested in.
func (r *Reports) GetDataPost(ctx context.Context, reportID, sectionID int, query *Query) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/sections/{sectionId}/data", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetPathParameter("sectionId", strconv.Itoa(sectionID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	return r.c.Ensure(ctx, request, http.StatusOK, nil)
}
