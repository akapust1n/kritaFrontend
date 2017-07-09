package models

type Tool struct {
	Tools []ToolsInternal `json:"Tools"`
}
type ToolsInternal struct {
	CountUse float64 `json:"countUse"`
	Time     float64 `json:"timeUseMSeconds"`
	ToolName string  `json:"toolName"`
}
