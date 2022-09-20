package transport

import (
	"project-name/app/interface/container"
)

type Tp struct {
	SampleTransport *sampleTransport
}

func SetupTransport(container *container.Container) *Tp {
	sampleTp := NewSampleTransport(container.SampleUsecase)

	return &Tp{
		SampleTransport: sampleTp,
	}
}
