package models

type Assert struct {
	Assert AssertssInternal `json:"asserts"`
}
type AssertssInternal struct {
	AssertFile string `json:"assertFile"`
	AssertLine string `json:"assertLine"`
	AssertText string `json:"assertText"`
}
