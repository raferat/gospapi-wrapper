/*
 * KSP API
 *
 * API pro interakci s webem KSP.
 *
 * API version: 1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kspapi_wrapper

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
/*var (
	_ context.Context
)*/

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
		requestBody    interface{}
		parsedResponse Subtask
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/generate"

	requestHeaders := make(map[string]string)
	requestQuery := url.Values{}

	requestQuery.Add("task", parameterToString(task, ""))
	requestQuery.Add("subtask", parameterToString(subtask, ""))

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, requestQuery, url.Values{}, "", []byte{})
	if err != nil {
		return parsedResponse, nil, err
	}

	rawResponse, err := serv.client.callAPI(r)
	if err != nil || rawResponse == nil {
		return parsedResponse, rawResponse, err
	}

	responseBody, err := io.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()
	if err != nil {
		return parsedResponse, rawResponse, err
	}

	if rawResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = serv.client.decode(&parsedResponse, responseBody, rawResponse.Header.Get("Content-Type"))
		return parsedResponse, rawResponse, err
	}

	newErr := GenericSwaggerError{
		body:  responseBody,
		error: rawResponse.Status,
	}
	return parsedResponse, rawResponse, newErr
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

	requestQeryParams.Add("task", parameterToString(task, ""))
	requestQeryParams.Add("subtask", parameterToString(subtask, ""))
	if requestOptions != nil && requestOptions.Generate.IsSet() {
		requestQeryParams.Add("generate", parameterToString(requestOptions.Generate.Value(), ""))
	}

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, requestQeryParams, url.Values{}, "", []byte{})
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
		requestBody    interface{}
		parsedResponse []string
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/tasks/list"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	if requestOptions != nil && requestOptions.Set.IsSet() {
		requestQueryParams.Add("set", parameterToString(requestOptions.Set.Value(), ""))
	}

	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "GET", requestBody, requestHeaders, requestQueryParams, url.Values{}, "", []byte{})
	if err != nil {
		return parsedResponse, nil, err
	}

	rawResponse, err := serv.client.callAPI(r)
	if err != nil || rawResponse == nil {
		return parsedResponse, rawResponse, err
	}

	responseBody, err := io.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()
	if err != nil {
		return parsedResponse, rawResponse, err
	}

	if rawResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = serv.client.decode(&parsedResponse, responseBody, rawResponse.Header.Get("Content-Type"))
		return parsedResponse, rawResponse, err
	}

	newErr := GenericSwaggerError{
		body:  responseBody,
		error: rawResponse.Status,
	}
	return parsedResponse, rawResponse, newErr
}

/*
TasksApiService Získání stavu úlohy
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param task textový identifikátor úlohy

@return Task
*/
func (a *TasksApiService) TasksStatusGet(ctx context.Context, task string) (Task, *http.Response, error) {
	var (
		requestBody interface{}
		result      Task
	)

	// create path and map variables
	requestPath := a.client.cfg.BasePath + "/tasks/status"

	requestHeaders := make(map[string]string)
	requestQueryParams := url.Values{}

	requestQueryParams.Add("task", parameterToString(task, ""))

	// set Accept header
	requestHeaders["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, requestPath, "GET", requestBody, requestHeaders, requestQueryParams, url.Values{}, "", []byte{})
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
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Subtask
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/submit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("task", parameterToString(task, ""))
	localVarQueryParams.Add("subtask", parameterToString(subtask, ""))

	localVarHeaderParams["Content-Type"] = "text/plain"

	// set Accept header
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Subtask
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode/100 == 4 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode/100 == 5 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
TasksApiService Interní: Souhrn úloh ve Cvičišti
Vypíše přehled všech úloh dostupných ve Cvičišti spolu se souhrnem bodů za všechny testy odevzdané jak v rámci Cvičiště, tak v sériích. **Nepoužívejte bez domluvy. Chování se může kdykoliv změnit.**
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Task
*/
func (a *TasksApiService) TasksXSummaryGet(ctx context.Context) ([]Task, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []Task
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/tasks/x-summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// set Accept header
	localVarHeaderParams["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []Task
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode/100 == 4 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode/100 == 5 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
