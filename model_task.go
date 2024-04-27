/*
 * KSP API
 *
 * API pro interakci s webem KSP.
 *
 * API version: 1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kspapi_wrapper

type Task struct {
	// ID úlohy
	Id string `json:"id"`
	// Jméno úlohy
	Name string `json:"name"`
	// Získané body za úlohu
	Points float32 `json:"points"`
	// Maximální množství bodů, které lze za úlohu získat.
	MaxPoints float32   `json:"max_points"`
	Subtasks  []Subtask `json:"subtasks"`
}
