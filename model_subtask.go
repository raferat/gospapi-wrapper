package kspapi_wrapper

import (
	"time"
)

// Podúloha a stav jejího odevzdání. U obecných endpointů uvádíme zatím nejlepší odevzdání, u endpointu `/tasks/submit` právě provedené odevzdání.
type Subtask struct {
	// ID podúlohy
	Id string `json:"id"`
	// Získané množství bodů za tuto konkrétní podúlohu
	Points float32 `json:"points"`
	// Maximální množství bodů, které lze za úlohu podúlohu získat.
	MaxPoints float32 `json:"max_points"`
	// Značí, zda je na serveru vstup vygenerován. Po vypršení platnosti vstupu se vrací do stavu `false`.
	InputGenerated bool `json:"input_generated"`
	// Pokud `input_generated` je `true`, tak `input_valid_until` se nutně v objektu nachází také. Ve formátu [dle RFC3339](https://tools.ietf.org/html/rfc3339#section-5.6). Může obsahovat data daleko v budoucnosti pro prakticky časově neomezené vstupy.
	InputValidUntil time.Time `json:"input_valid_until,omitempty"`
	// Pokud `input_generated` je `true`, z tohoto URL je možné stáhnout vstup pro úlohu. Na rozdíl od operace `tasks/input` není potřeba žádná autorizace.
	DownloadUrl string `json:"download_url,omitempty"`
	// Datumočas, kdy bylo odevzdáno aktuální řešení. Ve formátu [dle RFC3339](https://tools.ietf.org/html/rfc3339#section-5.6).
	SubmittedOn time.Time `json:"submitted_on,omitempty"`
	// Textová informace o vyhodnocení aktuálního řešení.
	Verdict string `json:"verdict,omitempty"`
}
