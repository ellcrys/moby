package volumes // import "github.com/ellcrys/docker/integration/plugin/volumes"

import (
	"fmt"
	"os"
	"testing"

	"github.com/ellcrys/docker/internal/test/environment"
)

var (
	testEnv *environment.Execution
)

func TestMain(m *testing.M) {
	var err error
	testEnv, err = environment.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if testEnv.OSType != "linux" {
		os.Exit(0)
	}
	err = environment.EnsureFrozenImagesLinux(testEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	testEnv.Print()
	os.Exit(m.Run())
}
