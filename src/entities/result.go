package entities

type Slaoq struct {
	Comment string
	Nota    float32
	Result  map[string]float32
}

type Result map[string]Slaoq
