package kspapi_wrapper

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/antihax/optional"
)

type TasksApiService service

/*
TasksApiService Vygeneruje vstup pro úlohu
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param task
  - @param subtask

@return Subtask
*/
func (serv *TasksApiService) TasksGeneratePost(ctx context.Context, task string, subtask string) (Subtask, *http.Response, error) {
	var (
		requestBody interface{}
		result      Subtask
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/generate"

	requestHeaders := make(map[string]string)
	requestQuery := url.Values{}

	requestQuery.Add("task", task)
	requestQuery.Add("subtask", subtask)

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, requestQuery)
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

/*
TasksApiService Stáhne vstup pro úlohu
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param task
 * @param subtask ID podúlohy získané například z &#x60;/tasks/status&#x60;
 * @param optional nil or *TasksApiTasksInputPostOpts - Optional Parameters:
     * @param "Generate" (optional.Bool) -  Při stahování vstupu rovnou vygenerovat nový s novou platností.

*/

type TasksApiTasksInputPostOpts struct {
	Generate optional.Bool
}

func (serv *TasksApiService) TasksInputPost(ctx context.Context, task string, subtask string, requestOptions *TasksApiTasksInputPostOpts) ([]byte, *http.Response, error) {
	var requestBody interface{}

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/input"

	requestHeaders := make(map[string]string)
	requestQeryParams := url.Values{}

	requestQeryParams.Add("task", task)
	requestQeryParams.Add("subtask", subtask)
	if requestOptions != nil && requestOptions.Generate.IsSet() {
		requestQeryParams.Add("generate", strconv.FormatBool(requestOptions.Generate.Value()))
	}

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, requestQeryParams)
	if err != nil {
		return nil, nil, err
	}

	rawResponse, err := serv.client.callAPI(r)
	if err != nil || rawResponse == nil {
		return nil, rawResponse, err
	}

	responseBody, err := io.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()
	if err != nil {
		return responseBody, rawResponse, err
	}

	if rawResponse.StatusCode < 300 {
		return responseBody, rawResponse, nil
	}

	newErr := GenericSwaggerError{
		body:  responseBody,
		error: rawResponse.Status,
	}
	var v ModelError
	err = serv.client.decode(&v, responseBody, rawResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr.error = err.Error()
		return responseBody, rawResponse, newErr
	}
	newErr.model = v
	return responseBody, rawResponse, newErr
}

/*
TasksApiService Seznam aktuálně řešitelných úloh (nevyžaduje autentikaci)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TasksApiTasksListGetOpts - Optional Parameters:
     * @param "Set" (optional.String) -  Název množiny úloh, která se má zobrazit. Default jsou otevřené úlohy z kategorií H a Z, explicitně si můžete vybrat &#x60;cviciste&#x60;. Pokud chcete o úlohách vědět více, může se hodit end-point &#x60;catalog&#x60;.
@return []string
*/

type TasksApiTasksListGetOpts struct {
	Set optional.String
}

func (serv *TasksApiService) TasksListGet(ctx context.Context, requestOptions *TasksApiTasksListGetOpts) ([]string, *http.Response, error) {
	var (
		requestBody interface{}
		result      []string
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/list"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	if requestOptions != nil && requestOptions.Set.IsSet() {
		requestQueryParams.Add("set", requestOptions.Set.Value())
	}

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "GET", requestBody, requestHeaders, requestQueryParams)
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

/*
TasksApiService Získání stavu úlohy
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param task textový identifikátor úlohy

@return Task
*/
func (serv *TasksApiService) TasksStatusGet(ctx context.Context, task string) (Task, *http.Response, error) {
	var (
		requestBody interface{}
		result      Task
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/status"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	requestQueryParams.Add("task", task)

	// set Accept header
	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "GET", requestBody, requestHeaders, requestQueryParams)
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

/*
TasksApiService Odevzdá řešení úlohy ke kontrole
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Řešení úlohy. Opendatovky vyžadují content-type &#x60;text/plain&#x60; nebo &#x60;application/binary&#x60;, teoretické úlohy přijímají i PDF.
  - @param task
  - @param subtask

@return Subtask
*/
func (a *TasksApiService) TasksSubmitPost(ctx context.Context, body string, task string, subtask string) (Subtask, *http.Response, error) {
	var (
		requestBody interface{}
		result      Subtask
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/tasks/submit"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	requestQueryParams.Add("task", task)
	requestQueryParams.Add("subtask", subtask)

	requestHeaders["Content-Type"] = "text/plain"

	// set Accept header
	requestHeaders["Accept"] = "application/json"

	// body params
	requestBody = &body
	r, err := a.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, requestQueryParams)
	if err != nil {
		return result, nil, err
	}

	rawResponse, err := a.client.callAPI(r)
	if err != nil || rawResponse == nil {
		return result, rawResponse, err
	}

	responseBody, err := io.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()
	if err != nil {
		return result, rawResponse, err
	}

	if rawResponse.StatusCode < 300 {
		err = a.client.decode(&result, responseBody, rawResponse.Header.Get("Content-Type"))
		return result, rawResponse, err
	}

	newErr := GenericSwaggerError{
		body:  responseBody,
		error: rawResponse.Status,
	}
	var v ModelError
	err = a.client.decode(&v, responseBody, rawResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr.error = err.Error()
		return result, rawResponse, newErr
	}
	newErr.model = v
	return result, rawResponse, newErr
}

/*
TasksApiService Interní: Souhrn úloh ve Cvičišti
Vypíše přehled všech úloh dostupných ve Cvičišti spolu se souhrnem bodů za všechny testy odevzdané jak v rámci Cvičiště, tak v sériích. **Nepoužívejte bez domluvy. Chování se může kdykoliv změnit.**
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Task
*/
func (serv *TasksApiService) TasksXSummaryGet(ctx context.Context) ([]Task, *http.Response, error) {
	var (
		requestBody interface{}
		result      []Task
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/x-summary"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "GET", requestBody, requestHeaders, requestQueryParams)
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
