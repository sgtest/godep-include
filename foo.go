package include // import "github.com/sgtest/godepinclude"

import "github.com/stretchr/testify/require"

type testingT struct{}

func (t *testingT) FailNow()                                  {}
func (t *testingT) Errorf(format string, args ...interface{}) {}

func foo() {
	t := &testingT{}
	require.Equal(t, 1, 1, "test")
}
