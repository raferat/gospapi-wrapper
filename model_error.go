package kspapi_wrapper

type ModelError struct {
	// Kód chyby pro strojové zpracování:  - `bad-params`: chybné parametry požadavku - `internal-error`: interní chyba serveru - `no-set`: tato množina úloh neexistuje - `no-subtask`: tato podúloha neexistuje - `no-task`: tato úloha neexistuje - `no-year`: tento ročník neexistuje - `not-enrolled`: nejsi přihlášený do ročníku a kategorie - `not-submittable`: tato úloha nemá povolené odevzdávání - `opendata-not-generated`: open-datový vstup ještě nebyl vygenerován - `submit-failed`: obecná chyba při odevzdávání řešení - `unauthorized`: k provedení této operace nemáte právo - `unsupported-task-type`: tato operace nepodporuje tento typ úloh  V budoucnosti mohou přibývat nové kódy chyb, a to i pro stávající operace.
	ErrorCode string `json:"errorCode"`
	// Lidsky čitelná chybová zpráva. Není určena pro strojové zpracování, v budoucnu se může libovolně změnit.
	ErrorMsg string `json:"errorMsg"`
}
