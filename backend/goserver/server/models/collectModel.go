package models

type CollectedData struct {
	Platform struct {
		Os struct {
			Windows float64
			Linux   float64
			Other   float64
		}
	}
}
