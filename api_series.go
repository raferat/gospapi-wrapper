package kspapi_wrapper

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/antihax/optional"
)

type Series struct {
	// ID série
	Id string `json:"id"`
	// Čas zveřejnění zadání úloh. Ve formátu [dle RFC3339](https://tools.ietf.org/html/rfc3339#section-5.6).
	TasksPublished time.Time `json:"tasks_published,omitempty"`
	// Termín odevzdání řešení. Ve formátu [dle RFC3339](https://tools.ietf.org/html/rfc3339#section-5.6).
	Deadline time.Time `json:"deadline,omitempty"`
	// Termín odevzdání řešení za redukovaný počet bodů (druhý termín v KSP-Z).
	Deadline2 time.Time     `json:"deadline2,omitempty"`
	Tasks     []SeriesTasks `json:"tasks,omitempty"`
}

type SeriesApiService service

/*
SeriesApiService Seznam sérií a úloh v nich (nevyžaduje autentikaci)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SeriesApiTasksCatalogGetOpts - Optional Parameters:
     * @param "Year" (optional.String) -  Ročník, jehož série se mají vypsat. Default: aktuální ročník. Pozor, na rozhraní školních roků mohou být aktuální dva ročníky současně.
     * @param "Tasks" (optional.Bool) -  Přepíná, zda se k sériím mají uvést i všechny úlohy.
@return []Series
*/

type SeriesApiTasksCatalogGetOpts struct {
	Year  optional.String
	Tasks optional.Bool
}

func (serv *SeriesApiService) TasksCatalogGet(ctx context.Context, requestOptions *SeriesApiTasksCatalogGetOpts) ([]Series, *http.Response, error) {
	var (
		requestBody interface{}
		result      []Series
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/catalog"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	if requestOptions != nil && requestOptions.Year.IsSet() {
		requestQueryParams.Add("year", parameterToString(requestOptions.Year.Value(), ""))
	}
	if requestOptions != nil && requestOptions.Tasks.IsSet() {
		requestQueryParams.Add("tasks", parameterToString(requestOptions.Tasks.Value(), ""))
	}

	// set Accept header
	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(
		ctx,
		requestPath,
		"GET",
		requestBody,
		requestHeaders,
		requestQueryParams,
		url.Values{},
		"",
		[]byte{},
	)
	if err != nil {
		return result, nil, err
	}

	rawResponse, err := serv.client.callAPI(r)
	if err != nil || rawResponse == nil {
		return result, rawResponse, err
	}

	responseBody, err := io.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()
	if err != nil {
		return result, rawResponse, err
	}

	if rawResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = serv.client.decode(&result, responseBody, rawResponse.Header.Get("Content-Type"))
		return result, rawResponse, err
	}

	newErr := GenericSwaggerError{
		body:  responseBody,
		error: rawResponse.Status,
	}
	var v ModelError
	err = serv.client.decode(&v, responseBody, rawResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr.error = err.Error()
		return result, rawResponse, newErr
	}
	newErr.model = v
	return result, rawResponse, newErr
}
