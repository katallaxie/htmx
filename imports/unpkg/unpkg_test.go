package unpkg

import (
	"testing"

	"github.com/katallaxie/htmx/imports"
	"github.com/stretchr/testify/require"
	"github.com/zeiss/pkg/jsonx"
)

func TestResolve(t *testing.T) {
	cdn := New()

	pkg := &imports.ExactPackage{
		Name:    "htmx.org",
		Version: "2.0.4",
	}

	err := cdn.Resolve(pkg)
	require.NoError(t, err)
	require.NotNil(t, pkg)

	require.Equal(t, "htmx.org", pkg.Name)
	require.Equal(t, "2.0.4", pkg.Version)
	require.Len(t, pkg.Files, 35)

	b, err := jsonx.Prettify(pkg)
	require.NoError(t, err)
	require.NotEmpty(t, b)
}
