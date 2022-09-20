package sample

import "fmt"

type sampleUseCase struct {
}

func NewSampleUseCase() UseCase {
	return &sampleUseCase{}
}

func (uc *sampleUseCase) MultiplicationTable(in int) (out map[string]interface{}, err error) {
	var res string

	i := 1
	for {
		if i > 99 {
			break
		}
		res += fmt.Sprintln(in, " X ", i, " = ", in*i)
		i++
	}

	out = map[string]interface{}{
		"response": res,
	}
	return
}
