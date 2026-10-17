package article

import "github.com/sirupsen/logrus"

// MathService checks numeric properties.
//
//go:generate mockery --config ../.mockery.yml
type MathService interface {
	IsEven(n int64) bool
}

// logInfo is the logging seam. Production calls logrus.Info; tests can replace it.
var logInfo = logrus.Info

type mathService struct{}

// NewMathService returns a MathService with no dependencies.
func NewMathService() MathService {
	return mathService{}
}

func (mathService) IsEven(n int64) bool {
	if n%2 == 0 {
		logInfo("even")
		return true
	}
	logInfo("odd")
	return false
}
