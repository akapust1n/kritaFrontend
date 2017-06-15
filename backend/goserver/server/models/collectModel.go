package models

type CollectedData struct {
	Platform struct {
		Os struct {
			Windows float64
			Linux   float64
			Mac     float64
			Other   float64
		}
	}
	CPU struct {
		Architecture struct {
			X86_64 float64
			X86    float64 //на самом деле  x86_64 c маленькой буквы
			Other  float64
		}
		Cores struct {
			One   float64
			Two   float64
			Three float64
			Four  float64
			Six   float64
			Eight float64
			Other float64
		}
	}
	Compiler struct {
		Type struct {
			GCC   float64
			Other float64
		}
		Version struct {
			V5dot4 float64
			Other  float64
		}
	}
}
