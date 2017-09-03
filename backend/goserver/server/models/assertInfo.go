package models

type Assert struct {
	Asserts []AssertsInternal `json:"asserts"`
}
type AssertsInternal struct {
	AssertFile string  `json:"assertFile"`
	AssertLine float64 `json:"assertLine"`
	AssertText string  `json:"assertText"`
	Count      float64 `json:"count"`
	IsFatal    bool    `json:"isFatal"`
}
