package kspapi_wrapper

import (
	"time"
)

type SeriesTasks struct {
	// ID úlohy
	Id string `json:"id"`
	// Název úlohy
	Name string `json:"name"`
	// Typ úlohy (theory/open-data/serial/other)
	Type_ string `json:"type"`
	// Je povoleno odevzdávání úlohy?
	Enabled bool `json:"enabled,omitempty"`
	// Počet bodů za úlohu
	Points int32 `json:"points"`
	// Termín odevzdání řešení. Ve formátu [dle RFC3339](https://tools.ietf.org/html/rfc3339#section-5.6).
	Deadline time.Time `json:"deadline,omitempty"`
	// Termín odevzdání řešení za redukovaný počet bodů (druhý termín v KSP-Z).
	Deadline2 time.Time `json:"deadline2,omitempty"`
}
