package models

type ActionCollected struct {
	CountUse float64
	Name     string
}
type ToolsCollected struct {
	CountUse float64
	Name     string
}
type CollectedData struct {
	Platform struct {
		Os struct {
			Windows float64
			Linux   float64
			Mac     float64
			Other   float64
		}
		Version struct {
			Windows struct {
				V7    float64
				V8    float64
				V81   float64
				V10   float64
				Other float64
			}
			Linux struct {
				Ubuntu1404 float64
				Ubuntu1410 float64
				Ubuntu1504 float64
				Ubuntu1510 float64
				Ubuntu1604 float64
				Ubuntu1610 float64
				Ubuntu1704 float64
				Other      float64
			}
			Mac struct {
				V1012 float64
				Other float64
			}
		}
	}
	CPU struct {
		Architecture struct {
			X86_64 float64
			X86    float64 //на самом деле  x86_64 c маленькой буквы
			Other  float64
		}
		Cores struct {
			C1    float64
			C2    float64
			C3    float64
			C4    float64
			C6    float64
			C8    float64
			Other float64
		}
	}
	Compiler struct {
		Type struct {
			GCC   float64
			Clang float64
			MSVC  float64
			Other float64
		}
		// 	Version struct { //подумать что можно сделать
		// 		V5dot4 float64
		// 		Other  float64
		// 	}
	}
	Locale struct {
		Language struct {
			English float64
			Russian float64
			Other   float64
		}
	}

	Actions       []ActionCollected
	ToolsUse      []ToolsCollected
	ToolsActivate []ToolsCollected

	Tools struct {
		KisToolBrush float64
		KisToolLine  float64
	}
}
