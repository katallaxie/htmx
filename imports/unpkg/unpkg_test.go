package unpkg

import (
	"testing"

	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/pkg/jsonx"
	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg := &imports.ExactPackage{
		Name:    "htmx.org",
		Version: "4.0.1",
	}

	err := cdn.Resolve(pkg)
	require.NoError(t, err)
	require.NotNil(t, pkg)

	require.Equal(t, "htmx.org", pkg.Name)
	require.Equal(t, "4.0.1", pkg.Version)
	require.Len(t, pkg.Files, 73)

	b, err := jsonx.Prettify(pkg)
	require.NoError(t, err)
	require.NotEmpty(t, b)
}
