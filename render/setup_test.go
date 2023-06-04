package render_test

import (
	"os"
	"testing"

	"github.com/CloudyKit/jet/v6"
	"github.com/roca/celeritas/render"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var testRender = render.Render{
	Renderer: "",
	RootPath: "",
	JetViews: views,
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
