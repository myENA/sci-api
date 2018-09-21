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
	ReportsListDefinitions200ResponseSlice []*ReportsListDefinitions200Response

	ReportsListDefinitions200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource *string     `json:"filterDataSource,omitempty"`
		ID               *int        `json:"id,omitempty"`
		Title            *string     `json:"title,omitempty"`
		URLSegmentName   *string     `json:"urlSegmentName,omitempty"`
	}
)

func (r Reports) ListDefinitions(ctx context.Context, filter *Filter) (*http.Response, ReportsListDefinitions200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports", true)
	if err = request.SetFilterQueryParameter(filter); err != nil {
		return nil, nil, err
	}
	out := make(ReportsListDefinitions200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	ReportsGetDefinition200Response struct {
		FacetID          interface{} `json:"facetId,omitempty"`
		FilterDataSource *string     `json:"filterDataSource,omitempty"`
		ID               *int        `json:"id,omitempty"`
		Title            *string     `json:"title,omitempty"`
		URLSegmentName   *string     `json:"urlSegmentName,omitempty"`
	}
)

func (r Reports) GetDefinition(ctx context.Context, reportID int, filter *Filter) (*http.Response, *ReportsGetDefinition200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports/{id}", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetFilterQueryParameter(filter); err != nil {
		return nil, nil, err
	}
	out := new(ReportsGetDefinition200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportsExistsByID200Response struct {
		Exists bool `json:"exists"`
	}
)

func (r Reports) Exists(ctx context.Context, reportID int) (*http.Response, *ReportsExistsByID200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("GET", "/reports/{id}/exists", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	out := new(ReportsExistsByID200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportsGetAllFacetData200ResponseDataApMacsSlice []string

	ReportsGetAllFacetData200ResponseDataSsidsSlice []*ReportsGetAllFacetData200ResponseDataSsids

	ReportsGetAllFacetData200ResponseDataSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsSlice []*ReportsGetAllFacetData200ResponseDataSystems

	ReportsGetAllFacetData200ResponseDataSystemsApMacsSlice []string

	ReportsGetAllFacetData200ResponseDataSystemsChildrenSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildren

	ReportsGetAllFacetData200ResponseDataSystemsChildrenApMacsSlice []string

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildrenChildren

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenApMacsSlice []string

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildren

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice []string

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsids

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildren struct {
		ApMacs ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenApMacsSlice `json:"apMacs,omitempty"`
		ID     *string                                                                         `json:"id,omitempty"`
		Ssids  ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSsidsSlice  `json:"ssids,omitempty"`
		Text   *string                                                                         `json:"text,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSsidsSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSsids

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsChildrenChildren struct {
		ApMacs   ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                                   `json:"id,omitempty"`
		Ssids    ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                                   `json:"text,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsChildrenSsidsSlice []*ReportsGetAllFacetData200ResponseDataSystemsChildrenSsids

	ReportsGetAllFacetData200ResponseDataSystemsChildrenSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsChildren struct {
		ApMacs   ReportsGetAllFacetData200ResponseDataSystemsChildrenApMacsSlice   `json:"apMacs,omitempty"`
		Children ReportsGetAllFacetData200ResponseDataSystemsChildrenChildrenSlice `json:"children,omitempty"`
		ID       *string                                                           `json:"id,omitempty"`
		Ssids    ReportsGetAllFacetData200ResponseDataSystemsChildrenSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                           `json:"text,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystemsSsidsSlice []*ReportsGetAllFacetData200ResponseDataSystemsSsids

	ReportsGetAllFacetData200ResponseDataSystemsSsids struct {
		Radio *string `json:"radio,omitempty"`
		Ssid  *string `json:"ssid,omitempty"`
	}

	ReportsGetAllFacetData200ResponseDataSystems struct {
		ApMacs   ReportsGetAllFacetData200ResponseDataSystemsApMacsSlice   `json:"apMacs,omitempty"`
		Children ReportsGetAllFacetData200ResponseDataSystemsChildrenSlice `json:"children,omitempty"`
		ID       *string                                                   `json:"id,omitempty"`
		Ssids    ReportsGetAllFacetData200ResponseDataSystemsSsidsSlice    `json:"ssids,omitempty"`
		Text     *string                                                   `json:"text,omitempty"`
	}

	ReportsGetAllFacetData200ResponseData struct {
		ApMacs  ReportsGetAllFacetData200ResponseDataApMacsSlice  `json:"apMacs,omitempty"`
		Ssids   ReportsGetAllFacetData200ResponseDataSsidsSlice   `json:"ssids,omitempty"`
		Systems ReportsGetAllFacetData200ResponseDataSystemsSlice `json:"systems,omitempty"`
	}

	ReportsGetAllFacetData200Response struct {
		Data *ReportsGetAllFacetData200ResponseData `json:"data,omitempty"`
	}
)

func (r Reports) GetAllFacetData(ctx context.Context, reportID int, start, end time.Time, granularity string) (*http.Response, *ReportsGetAllFacetData200Response, error) {
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
	out := new(ReportsGetAllFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportsGetAPMacFacetData200ResponseDataApMacsSlice []string

	ReportsGetAPMacFacetData200ResponseData struct {
		ApMacs ReportsGetAPMacFacetData200ResponseDataApMacsSlice `json:"apMacs,omitempty"`
	}

	ReportsGetAPMacFacetData200Response struct {
		Data *ReportsGetAPMacFacetData200ResponseData `json:"data,omitempty"`
	}
)

func (r Reports) GetAPMACFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, *ReportsGetAPMacFacetData200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/apmac", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(ReportsGetAPMacFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportsGetSSIDFacetData200ResponseDataSlice []string

	ReportsGetSSIDFacetData200Response struct {
		Data ReportsGetSSIDFacetData200ResponseDataSlice `json:"data,omitempty"`
	}
)

func (r Reports) GetSSIDFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, *ReportsGetSSIDFacetData200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(ReportsGetSSIDFacetData200Response)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportsGetSystemsFacetData200ResponseSlice []*ReportsGetSystemsFacetData200Response

	ReportsGetSystemsFacetData200ResponseChildrenSlice []*ReportsGetSystemsFacetData200ResponseChildren

	ReportsGetSystemsFacetData200ResponseChildren struct {
		Children   *bool   `json:"children,omitempty"`
		Data       *string `json:"data,omitempty"`
		FilterText *string `json:"filterText,omitempty"`
		ID         *string `json:"id,omitempty"`
		Text       *string `json:"text,omitempty"`
	}

	ReportsGetSystemsFacetData200ResponseState struct {
		Opened *bool `json:"opened,omitempty"`
	}

	ReportsGetSystemsFacetData200Response struct {
		Children ReportsGetSystemsFacetData200ResponseChildrenSlice `json:"children,omitempty"`
		ID       *string                                            `json:"id,omitempty"`
		State    *ReportsGetSystemsFacetData200ResponseState        `json:"state,omitempty"`
		Text     *string                                            `json:"text,omitempty"`
	}
)

func (r Reports) GetSystemsFacetData(ctx context.Context, reportID int, query *Query) (*http.Response, ReportsGetSystemsFacetData200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := make(ReportsGetSystemsFacetData200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	ReportsListSectionDefinitions200ResponseSlice []*ReportsListSectionDefinitions200Response

	ReportsListSectionDefinitions200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
		Metric      *string `json:"metric,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutColorsSlice []string

	ReportsListSectionDefinitions200ResponseLayoutColumnsSlice []*ReportsListSectionDefinitions200ResponseLayoutColumns

	ReportsListSectionDefinitions200ResponseLayoutColumns struct {
		Color           *string `json:"color,omitempty"`
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutHeadersSlice []*ReportsListSectionDefinitions200ResponseLayoutHeaders

	ReportsListSectionDefinitions200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutHeadersContent struct {
		Formats *ReportsListSectionDefinitions200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                              `json:"text,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutHeaders struct {
		Component *string                                                       `json:"component,omitempty"`
		Content   *ReportsListSectionDefinitions200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                                       `json:"name,omitempty"`
		Options   map[string]string                                             `json:"options,omitempty"`
		Query     *string                                                       `json:"query,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutSubSectionsSlice []*ReportsListSectionDefinitions200ResponseLayoutSubSections

	ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeriesSlice []*ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeries

	ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeries struct {
		Area   *bool   `json:"area,omitempty"`
		Color  *string `json:"color,omitempty"`
		Key    *string `json:"key,omitempty"`
		Values *string `json:"values,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayout struct {
		Series    ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayoutSeriesSlice `json:"series,omitempty"`
		Width     *string                                                                    `json:"width,omitempty"`
		XAxisType *string                                                                    `json:"xAxisType,omitempty"`
		YAxisType *string                                                                    `json:"yAxisType,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayoutSubSections struct {
		Component *string                                                          `json:"component,omitempty"`
		Layout    *ReportsListSectionDefinitions200ResponseLayoutSubSectionsLayout `json:"layout,omitempty"`
		Title     *string                                                          `json:"title,omitempty"`
	}

	ReportsListSectionDefinitions200ResponseLayout struct {
		Colors      ReportsListSectionDefinitions200ResponseLayoutColorsSlice      `json:"colors,omitempty"`
		Columns     ReportsListSectionDefinitions200ResponseLayoutColumnsSlice     `json:"columns,omitempty"`
		Format      *string                                                        `json:"format,omitempty"`
		Headers     ReportsListSectionDefinitions200ResponseLayoutHeadersSlice     `json:"headers,omitempty"`
		SubSections ReportsListSectionDefinitions200ResponseLayoutSubSectionsSlice `json:"subSections,omitempty"`
		Width       *string                                                        `json:"width,omitempty"`
	}

	ReportsListSectionDefinitions200Response struct {
		Component         *string                                                    `json:"component,omitempty"`
		DefaultParameters *ReportsListSectionDefinitions200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                                       `json:"id,omitempty"`
		Layout            *ReportsListSectionDefinitions200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                                       `json:"order,omitempty"`
		QueryName         *string                                                    `json:"queryName,omitempty"`
		ReportID          *int                                                       `json:"reportId,omitempty"`
		Title             *string                                                    `json:"title,omitempty"`
		URL               interface{}                                                `json:"url,omitempty"`
	}
)

func (r Reports) ListSectionDefinitions(ctx context.Context, reportID int, filter *Filter) (*http.Response, ReportsListSectionDefinitions200ResponseSlice, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetFilterQueryParameter(filter)
	out := make(ReportsListSectionDefinitions200ResponseSlice, 0)
	httpResponse, _, err := r.c.Ensure(ctx, request, http.StatusOK, &out)
	return httpResponse, out, err
}

type (
	ReportsGetSectionDefinition200ResponseDefaultParameters struct {
		Granularity *string `json:"granularity,omitempty"`
		Limit       *int    `json:"limit,omitempty"`
	}

	ReportsGetSectionDefinition200ResponseLayoutColumnsSlice []*ReportsGetSectionDefinition200ResponseLayoutColumns

	ReportsGetSectionDefinition200ResponseLayoutColumns struct {
		ColumnName      *string `json:"columnName,omitempty"`
		CustomComponent *string `json:"customComponent,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		Format          *string `json:"format,omitempty"`
		Hidden          *bool   `json:"hidden,omitempty"`
	}

	ReportsGetSectionDefinition200ResponseLayoutHeadersSlice []*ReportsGetSectionDefinition200ResponseLayoutHeaders

	ReportsGetSectionDefinition200ResponseLayoutHeadersContentFormats struct {
		Percentage   *string `json:"percentage,omitempty"`
		TotalTraffic *string `json:"totalTraffic,omitempty"`
		Traffic      *string `json:"traffic,omitempty"`
	}

	ReportsGetSectionDefinition200ResponseLayoutHeadersContent struct {
		Formats *ReportsGetSectionDefinition200ResponseLayoutHeadersContentFormats `json:"formats,omitempty"`
		Text    *string                                                            `json:"text,omitempty"`
	}

	ReportsGetSectionDefinition200ResponseLayoutHeaders struct {
		Component *string                                                     `json:"component,omitempty"`
		Content   *ReportsGetSectionDefinition200ResponseLayoutHeadersContent `json:"content,omitempty"`
		Name      *string                                                     `json:"name,omitempty"`
		Options   map[string]string                                           `json:"options,omitempty"`
		Query     *string                                                     `json:"query,omitempty"`
	}

	ReportsGetSectionDefinition200ResponseLayout struct {
		Columns ReportsGetSectionDefinition200ResponseLayoutColumnsSlice `json:"columns,omitempty"`
		Headers ReportsGetSectionDefinition200ResponseLayoutHeadersSlice `json:"headers,omitempty"`
		Width   *string                                                  `json:"width,omitempty"`
	}

	ReportsGetSectionDefinition200Response struct {
		Component         *string                                                  `json:"component,omitempty"`
		DefaultParameters *ReportsGetSectionDefinition200ResponseDefaultParameters `json:"defaultParameters,omitempty"`
		ID                *int                                                     `json:"id,omitempty"`
		Layout            *ReportsGetSectionDefinition200ResponseLayout            `json:"layout,omitempty"`
		Order             *int                                                     `json:"order,omitempty"`
		QueryName         *string                                                  `json:"queryName,omitempty"`
		ReportID          *int                                                     `json:"reportId,omitempty"`
		Title             *string                                                  `json:"title,omitempty"`
		URL               interface{}                                              `json:"url,omitempty"`
	}
)

func (r Reports) GetSectionDefinition(ctx context.Context, reportID, sectionID int, filter *Filter) (*http.Response, *ReportsGetSectionDefinition200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("GET", "/reports/{id}/sections/{fk}", true)
	request.SetPathParameter("id", strconv.Itoa(reportID))
	request.SetPathParameter("fk", strconv.Itoa(sectionID))
	request.SetFilterQueryParameter(filter)
	out := new(ReportsGetSectionDefinition200Response)
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
func (r Reports) GetSectionData(ctx context.Context, reportID, sectionID int, query *Query) (*http.Response, []byte, error) {
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
