package kspapi_wrapper

type InlineResponse200 struct {
	// Dočasný API token
	Token string `json:"token"`
	// Platnost tokenu v sekundách
	ValiditySeconds int32 `json:"validity_seconds"`
}
