package dynakube

import (
	"os"
	"testing"

	"github.com/Dynatrace/dynatrace-operator/api/v1alpha1"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	dynaKube := v1alpha1.DynaKube{}

	assert.NotNil(t, dynaKube.Spec.RoutingSpec)
}

func TestMkDirAll(t *testing.T) {
	const path = "/tmp/foo/bar/buzz"
	_ = os.MkdirAll(path, os.ModePerm)
	file, _ := os.Open(path)
	_ = file.Chmod(os.ModePerm)

	fs := afero.NewBasePathFs(afero.NewOsFs(), "/tmp")
	_ = fs.Mkdir("/quack", os.ModePerm)
	_ = fs.Chmod("/quack", os.ModePerm)
}
