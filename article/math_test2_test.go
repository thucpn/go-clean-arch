package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIsEven_logsWithoutHook replaces logInfo with a recording func and does not use a logrus hook.
func TestIsEven_logsWithoutHook(t *testing.T) {
	var got []string
	original := logInfo
	logInfo = func(args ...interface{}) {
		got = append(got, args[0].(string))
	}
	t.Cleanup(func() { logInfo = original })

	svc := NewMathService()

	assert.True(t, svc.IsEven(4))
	require.Len(t, got, 1)
	assert.Equal(t, "even", got[0])

	assert.False(t, svc.IsEven(7))
	require.Len(t, got, 2)
	assert.Equal(t, "odd", got[1])
}
