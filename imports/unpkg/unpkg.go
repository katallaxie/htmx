package unpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/katallaxie/htmx/imports"
)

const (
	DefaultURL    = "https://unpkg.com/%s@%s/?meta"
	DefaultCdnURL = "https://unpkg.com/%s@%s/%s"
)

var _ imports.Resolver = (*client)(nil)

type client struct{}

// Response is the response from the unpkg API.
type Response struct {
	Package string `json:"package"`
	Version string `json:"version"`
	Prefix  string `json:"prefix"`
	Files   []struct {
		Path      string `json:"path"`
		Size      int    `json:"size"`
		Integrity string `json:"integrity"`
		Type      string `json:"type"`
	} `json:"files"`
}

// New returns a new unpkg provider.
func New() imports.Resolver {
	return &client{}
}

// Resolve resolves the package to a URL.
func (c *client) Resolve(pkg *imports.ExactPackage) error {
	url := fmt.Sprintf(DefaultURL, pkg.Name, pkg.Version)

	//nolint:noctx
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unpkg API responded with status %d", resp.StatusCode)
	}

	var meta Response
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return err
	}

	for _, f := range meta.Files {
		switch filepath.Ext(f.Path) {
		case ".js":
			pkg.Files = append(pkg.Files, &imports.FileJS{
				Path:      fmt.Sprintf(DefaultCdnURL, meta.Package, meta.Version, strings.TrimPrefix(f.Path, "/")),
				Integrity: f.Integrity,
				LocalPath: strings.TrimPrefix(f.Path, "/"),
			})
		case ".css":
			pkg.Files = append(pkg.Files, &imports.FileCSS{
				Path:      fmt.Sprintf(DefaultCdnURL, meta.Package, meta.Version, strings.TrimPrefix(f.Path, "/")),
				Integrity: f.Integrity,
				LocalPath: strings.TrimPrefix(f.Path, "/"),
			})
		default:
			pkg.Files = append(pkg.Files, &imports.FileUnkown{
				Path:      fmt.Sprintf(DefaultCdnURL, meta.Package, meta.Version, strings.TrimPrefix(f.Path, "/")),
				Integrity: f.Integrity,
				LocalPath: strings.TrimPrefix(f.Path, "/"),
			})
		}
	}

	return nil
}
