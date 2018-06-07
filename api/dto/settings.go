package dto

type SettingsPayload struct {
	IsSetupTrackerEnabled  bool `json:"isSetupTrackerEnabled"`
	IsWelcomeDialogEnabled bool `json:"isWelcomeDialogEnabled"`
}
