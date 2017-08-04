package models

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
	Tools struct {
	}
	Actions struct {
		Add_new_paint_layer             float64
		Clear                           float64
		Copy_layer_clipboard            float64
		Cut_layer_clipboard             float64
		Edit_cut                        float64
		Edit_redo                       float64
		Edit_undo                       float64
		File_new                        float64
		Fill_selection_background_color float64
		Fill_selection_foreground_color float64
		Fill_selection_pattern          float64
		Paste_at                        float64
		Paste_layer_from_clipboard      float64
		Paste_new                       float64
		Stroke_selection                float64
		View_show_canvas_only           float64
	}
}
