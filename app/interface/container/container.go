package container

import (
	useCaseSample "project-name/app/usecase/sample"
)

type Container struct {
	SampleUsecase useCaseSample.UseCase
}

func SetupContainer() Container {
	// init utils
	// init repo

	// init use case
	sampleUseCase := useCaseSample.NewSampleUseCase()

	return Container{
		SampleUsecase: sampleUseCase,
	}
}
