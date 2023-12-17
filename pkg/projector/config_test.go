package projector_test

import (
	"reflect"
	"testing"

	"github.com/pcoelho00/projector_cli/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
	return opts
}

func TestConfigPrint(t *testing.T) {
	opts := getOpts([]string{})
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}

	if !reflect.DeepEqual([]string{}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}

}

func TestConfigPrintKey(t *testing.T) {
	opts := getOpts([]string{"foo"})
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}

	if !reflect.DeepEqual([]string{"foo"}, config.Args) {
		t.Errorf("expected args to be an empty string array but got %+v", config.Args)
	}

}
