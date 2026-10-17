package article_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bxcodec/go-clean-arch/article"
)

func TestIsEven(t *testing.T) {
	svc := article.NewMathService()

	assert.True(t, svc.IsEven(0))
	assert.True(t, svc.IsEven(2))
	assert.True(t, svc.IsEven(-4))
	assert.False(t, svc.IsEven(1))
	assert.False(t, svc.IsEven(-3))
}

// TestIsEven_logsParity checks that IsEven logs "even" or "odd".
func TestIsEven_logsParity(t *testing.T) {
	logger := logrus.StandardLogger()
	previous := logger.ReplaceHooks(make(logrus.LevelHooks))
	recorder := &logRecorder{}
	logger.AddHook(recorder)
	t.Cleanup(func() { logger.ReplaceHooks(previous) })

	svc := article.NewMathService()

	assert.True(t, svc.IsEven(4))
	require.Len(t, recorder.messages, 1)
	assert.Equal(t, "even", recorder.messages[0])

	assert.False(t, svc.IsEven(7))
	require.Len(t, recorder.messages, 2)
	assert.Equal(t, "odd", recorder.messages[1])
}

type logRecorder struct {
	messages []string
}

func (r *logRecorder) Levels() []logrus.Level { return logrus.AllLevels }

func (r *logRecorder) Fire(entry *logrus.Entry) error {
	r.messages = append(r.messages, entry.Message)
	return nil
}
