package transport

import (
	"github.com/gin-gonic/gin"
	"project-name/app/usecase/sample"
	"strconv"
)

type sampleTransport struct {
	sampleUC sample.UseCase
}

func NewSampleTransport(useCase sample.UseCase) *sampleTransport {
	return &sampleTransport{
		sampleUC: useCase,
	}
}

func (tp *sampleTransport) MultiplicationTable(c *gin.Context) {
	param := c.Request.URL.Query().Get("value")
	in, err := strconv.Atoi(param)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := tp.sampleUC.MultiplicationTable(in)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}
