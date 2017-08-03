package models

type Action struct {
	Actions []ActionsInternal `json:"actions"`
}
type ActionsInternal struct {
	CountUse   float64 `json:"countUse"`
	Sources    float64 `json:"timeUseMSeconds"`
	ActionName string  `json:"actionName"`
}
