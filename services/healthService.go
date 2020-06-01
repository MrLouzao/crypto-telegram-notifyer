package services

// All healthchecks to be performed
type HealthChecks struct {
	Status   string `json:"status"`
	Database bool   `json:"database"`
	Message  string `json:"message"`
}

func GetSystemHealth() *HealthChecks {
	// Always true - TODO perform check over DB
	status := HealthChecks{
		Status:   "READY",
		Database: true,
		Message:  "Ready to be consumpted",
	}

	return &status
}
