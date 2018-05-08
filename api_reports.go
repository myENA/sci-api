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
	ListReports200ResponseSlice []*ListReports200Response

	ListReports200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource *string     `json:"filterDataSource,omitempty"`
		ID               *int        `json:"id,omitempty"`
		Title            *string     `json:"title,omitempty"`
		URLSegmentName   *string     `json:"urlSegmentName,omitempty"`
	}
)

func (r *Reports) ListDefinitions(ctx context.Context, filter *Filter) (*http.Response, ListReports200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports", true)
	if err = request.SetFilterQueryParameter(filter); err != nil {
		return nil, nil, err
	}
	out := make(ListReports200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	FindReportByIDGet200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource *string     `json:"filterDataSource,omitempty"`
		ID               *int        `json:"id,omitempty"`
		Title            *string     `json:"title,omitempty"`
		URLSegmentName   *string     `json:"urlSegmentName,omitempty"`
	}
)

func (r *Reports) GetDefinition(ctx context.Context, reportID int, filter *Filter) (*http.Response, *FindReportByIDGet200Response, error) {
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

func (r *Reports) Exists(ctx context.Context, reportID int) (*http.Response, *ReportExistsByID200Response, error) {
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
	GetAllFacetData200ResponseDataApMacsSlice []string

	GetAllFacetData200ResponseDataSsidsSlice []*GetAllFacetData200ResponseDataSsids

	GetAllFacetData200ResponseDataSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsSlice []*GetAllFacetData200ResponseDataSystems

	GetAllFacetData200ResponseDataSystemsApMacsSlice []string

	GetAllFacetData200ResponseDataSystemsChildrenSlice []*GetAllFacetData200ResponseDataSystemsChildren

	GetAllFacetData200ResponseDataSystemsChildrenApMacsSlice []string

	GetAllFacetData200ResponseDataSystemsChildrenChildrenSlice []*GetAllFacetData200ResponseDataSystemsChildrenChildren

	GetAllFacetData200ResponseDataSystemsChildrenChildrenApMacsSlice []string

	GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSlice []*GetAllFacetData200ResponseDataSystemsChildrenChildrenChildren

	GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice []string

	GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice []*GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsids

	GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsChildrenChildrenChildren struct {
		ApMacs GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice `json:"apMacs,omitempty"`
		ID     *string                                                                  `json:"id,omitempty"`
		Ssids  GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice  `json:"ssids,omitempty"`
		Text   *string                                                                  `json:"text,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsChildrenChildrenSsidsSlice []*GetAllFacetData200ResponseDataSystemsChildrenChildrenSsids

	GetAllFacetData200ResponseDataSystemsChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsChildrenChildren struct {
		ApMacs   GetAllFacetData200ResponseDataSystemsChildrenChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children GetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                            `json:"id,omitempty"`
		Ssids    GetAllFacetData200ResponseDataSystemsChildrenChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                            `json:"text,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsChildrenSsidsSlice []*GetAllFacetData200ResponseDataSystemsChildrenSsids

	GetAllFacetData200ResponseDataSystemsChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsChildren struct {
		ApMacs   GetAllFacetData200ResponseDataSystemsChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children GetAllFacetData200ResponseDataSystemsChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                    `json:"id,omitempty"`
		Ssids    GetAllFacetData200ResponseDataSystemsChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                    `json:"text,omitempty"`
	}

	GetAllFacetData200ResponseDataSystemsSsidsSlice []*GetAllFacetData200ResponseDataSystemsSsids

	GetAllFacetData200ResponseDataSystemsSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	GetAllFacetData200ResponseDataSystems struct {
		ApMacs   GetAllFacetData200ResponseDataSystemsApMacsSlice   `json:"apMacs,omitempty"`
		Children GetAllFacetData200ResponseDataSystemsChildrenSlice `json:"children,omitempty"`
		ID       *string                                            `json:"id,omitempty"`
		Ssids    GetAllFacetData200ResponseDataSystemsSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                            `json:"text,omitempty"`
	}

	GetAllFacetData200ResponseData struct {
		ApMacs  GetAllFacetData200ResponseDataApMacsSlice  `json:"apMacs,omitempty"`
		Ssids   GetAllFacetData200ResponseDataSsidsSlice   `json:"ssids,omitempty"`
		Systems GetAllFacetData200ResponseDataSystemsSlice `json:"systems,omitempty"`
	}

	GetAllFacetData200Response struct {
		Data *GetAllFacetData200ResponseData `json:"data,omitempty"`
	}
)

func (r *Reports) GetAllFacetData(ctx context.Context, reportID int, start, end time.Time, granularity string) (*http.Response, *GetAllFacetData200Response, error) {
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
	out := new(GetAllFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	GetAPMacFacetData200ResponseDataApMacsSlice []string

	GetAPMacFacetData200ResponseData struct {
		ApMacs GetAPMacFacetData200ResponseDataApMacsSlice `json:"apMacs,omitempty"`
	}

	GetAPMacFacetData200Response struct {
		Data *GetAPMacFacetData200ResponseData `json:"data,omitempty"`
	}
)

func (r *Reports) GetAPMACFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, *GetAPMacFacetData200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/apmac", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(GetAPMacFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	GetSSIDFacetData200ResponseDataSlice []string

	GetSSIDFacetData200Response struct {
		Data GetSSIDFacetData200ResponseDataSlice `json:"data,omitempty"`
	}
)

func (r *Reports) GetSSIDFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, *GetSSIDFacetData200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(GetSSIDFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	GetSystemsFacetData200ResponseSlice []*GetSystemsFacetData200Response

	GetSystemsFacetData200ResponseChildrenSlice []*GetSystemsFacetData200ResponseChildren

	GetSystemsFacetData200ResponseChildren struct {
		Children   *bool   `json:"children,omitempty"`
		Data       *string `json:"data,omitempty"`
		FilterText *string `json:"filterText,omitempty"`
		ID         *string `json:"id,omitempty"`
		Text       *string `json:"text,omitempty"`
	}

	GetSystemsFacetData200ResponseState struct {
		Opened *bool `json:"opened,omitempty"`
	}

	GetSystemsFacetData200Response struct {
		Children GetSystemsFacetData200ResponseChildrenSlice `json:"children,omitempty"`
		ID       *string                                     `json:"id,omitempty"`
		State    *GetSystemsFacetData200ResponseState        `json:"state,omitempty"`
		Text     *string                                     `json:"text,omitempty"`
	}
)

func (r *Reports) GetSystemsFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, GetSystemsFacetData200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := make(GetSystemsFacetData200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	ListSectionDefinitions200ResponseSlice []*ListSectionDefinitions200Response

	ListSectionDefinitions200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
		Metric      *string `json:"metric,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutColorsSlice []string

	ListSectionDefinitions200ResponseLayoutColumnsSlice []*ListSectionDefinitions200ResponseLayoutColumns

	ListSectionDefinitions200ResponseLayoutColumns struct {
		Color           *string `json:"color,omitempty"`
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutHeadersSlice []*ListSectionDefinitions200ResponseLayoutHeaders

	ListSectionDefinitions200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutHeadersContent struct {
		Formats *ListSectionDefinitions200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                       `json:"text,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutHeaders struct {
		Component *string                                                `json:"component,omitempty"`
		Content   *ListSectionDefinitions200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                                `json:"name,omitempty"`
		Options   map[string]string                                      `json:"options,omitempty"`
		Query     *string                                                `json:"query,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutSubSectionsSlice []*ListSectionDefinitions200ResponseLayoutSubSections

	ListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeriesSlice []*ListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeries

	ListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeries struct {
		Area   *bool   `json:"area,omitempty"`
		Color  *string `json:"color,omitempty"`
		Key    *string `json:"key,omitempty"`
		Values *string `json:"values,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutSubSectionsLayout struct {
		Series    ListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeriesSlice `json:"series,omitempty"`
		Width     *string                                                             `json:"width,omitempty"`
		XAxisType *string                                                             `json:"xAxisType,omitempty"`
		YAxisType *string                                                             `json:"yAxisType,omitempty"`
	}

	ListSectionDefinitions200ResponseLayoutSubSections struct {
		Component *string                                                   `json:"component,omitempty"`
		Layout    *ListSectionDefinitions200ResponseLayoutSubSectionsLayout `json:"layout,omitempty"`
		Title     *string                                                   `json:"title,omitempty"`
	}

	ListSectionDefinitions200ResponseLayout struct {
		Colors      ListSectionDefinitions200ResponseLayoutColorsSlice      `json:"colors,omitempty"`
		Columns     ListSectionDefinitions200ResponseLayoutColumnsSlice     `json:"columns,omitempty"`
		Format      *string                                                 `json:"format,omitempty"`
		Headers     ListSectionDefinitions200ResponseLayoutHeadersSlice     `json:"headers,omitempty"`
		SubSections ListSectionDefinitions200ResponseLayoutSubSectionsSlice `json:"subSections,omitempty"`
		Width       *string                                                 `json:"width,omitempty"`
	}

	ListSectionDefinitions200Response struct {
		Component         *string                                             `json:"component,omitempty"`
		DefaultParameters *ListSectionDefinitions200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                                `json:"id,omitempty"`
		Layout            *ListSectionDefinitions200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                                `json:"order,omitempty"`
		QueryName         *string                                             `json:"queryName,omitempty"`
		ReportID          *int                                                `json:"reportId,omitempty"`
		Title             *string                                             `json:"title,omitempty"`
		URL               interface{}                                         `json:"url,omitempty"`
	}
)

func (r *Reports) ListSectionDefinitions(ctx context.Context, reportID int, filter *Filter) (*http.Response, ListSectionDefinitions200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetFilterQueryParameter(filter)
	out := make(ListSectionDefinitions200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	GetSectionDefinition200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
	}

	GetSectionDefinition200ResponseLayoutColumnsSlice []*GetSectionDefinition200ResponseLayoutColumns

	GetSectionDefinition200ResponseLayoutColumns struct {
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	GetSectionDefinition200ResponseLayoutHeadersSlice []*GetSectionDefinition200ResponseLayoutHeaders

	GetSectionDefinition200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	GetSectionDefinition200ResponseLayoutHeadersContent struct {
		Formats *GetSectionDefinition200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                     `json:"text,omitempty"`
	}

	GetSectionDefinition200ResponseLayoutHeaders struct {
		Component *string                                              `json:"component,omitempty"`
		Content   *GetSectionDefinition200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                              `json:"name,omitempty"`
		Options   map[string]string                                    `json:"options,omitempty"`
		Query     *string                                              `json:"query,omitempty"`
	}

	GetSectionDefinition200ResponseLayout struct {
		Columns GetSectionDefinition200ResponseLayoutColumnsSlice `json:"columns,omitempty"`
		Headers GetSectionDefinition200ResponseLayoutHeadersSlice `json:"headers,omitempty"`
		Width   *string                                           `json:"width,omitempty"`
	}

	GetSectionDefinition200Response struct {
		Component         *string                                           `json:"component,omitempty"`
		DefaultParameters *GetSectionDefinition200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                              `json:"id,omitempty"`
		Layout            *GetSectionDefinition200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                              `json:"order,omitempty"`
		QueryName         *string                                           `json:"queryName,omitempty"`
		ReportID          *int                                              `json:"reportId,omitempty"`
		Title             *string                                           `json:"title,omitempty"`
		URL               interface{}                                       `json:"url,omitempty"`
	}
)

func (r *Reports) GetSectionDefinition(ctx context.Context, reportID, sectionID int, filter *Filter) (*http.Response, *GetSectionDefinition200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections/{fk}", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetPathParameter("fk", strconv.Itoa(sectionID))
	request.SetFilterQueryParameter(filter)
	out := new(GetSectionDefinition200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

// GetSectionData will, on success, always return an object like:
//
//		{
//			"data": [],
//			"metadata": {}
//		}
//
// The exact structure of "data" and "metadata" vary greatly depending upon the report id and section queried, and
// therefore is not modeled here.  It is expected that the caller either doesn't care or knows which specific report
// they're interested in.
func (r *Reports) GetSectionData(ctx context.Context, reportID, sectionID int, query *Query) (*http.Response, []byte, error) {
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
