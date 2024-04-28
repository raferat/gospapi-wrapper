package kspapi_wrapper

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type AuthApiService service

/*
AuthApiService Interní: Vytvoří dočasný token na základě cookie.
Vytvoří uživateli přihlášenému na webu KSP dočasný API token. Je určeno k použití z JavaScriptu na webu KSP, použití z jiných domén není kvůli cross-origin omezením možné. **Nepoužívejte bez domluvy. Chování se může kdykoliv změnit.**
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return InlineResponse200
*/
func (serv *AuthApiService) AuthXGetTokenPost(ctx context.Context) (InlineResponse200, *http.Response, error) {
	var (
		requestBody interface{}
		result      InlineResponse200
	)

	// create path and map variables
	requestPath := serv.client.cfg.BasePath + "/auth/x-get-token"

	requestHeaders := make(map[string]string)

	// set Accept header
	requestHeaders["Accept"] = "application/json"

	r, err := serv.client.prepareRequest(ctx, requestPath, "POST", requestBody, requestHeaders, url.Values{})
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
		if err == nil {
			return result, rawResponse, err
		}
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
