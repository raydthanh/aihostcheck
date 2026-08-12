package model

const SchemaVersion = "1.0.0"

type Status string

const (
	Detected         Status = "detected"
	NotDetected      Status = "not_detected"
	Unknown          Status = "unknown"
	Unsupported      Status = "unsupported"
	Error            Status = "error"
	PermissionDenied Status = "permission_denied"
)

type Capability struct {
	Status   Status            `json:"status"`
	Value    string            `json:"value,omitempty"`
	Evidence []Evidence        `json:"evidence,omitempty"`
	Details  map[string]string `json:"details,omitempty"`
}

type Evidence struct {
	Source string `json:"source"`
	Detail string `json:"detail"`
}

type Report struct {
	SchemaVersion string                `json:"schema_version"`
	GeneratedAt   string                `json:"generated_at"`
	ToolVersion   string                `json:"tool_version"`
	Platform      string                `json:"platform"`
	Capabilities  map[string]Capability `json:"capabilities"`
}
