// Code generated by tool; DO NOT EDIT.
package config

import (
	"sync"
	"testing"
)

func TestAllParameters_LoadInitialValues(t *testing.T) {
	ap := &AllParameters{}
	if err := ap.LoadInitialValues(); err != nil {
		t.Errorf("LoadInitialValues failed. error:%v", err)
	}
}

func isconfigurationEqual(c1, c2 *configuration) bool {

	if c1.BoolSet1 != c2.BoolSet1 {
		return false
	}

	if c1.BoolSet2 != c2.BoolSet2 {
		return false
	}

	if c1.BoolSet3 != c2.BoolSet3 {
		return false
	}

	if c1.StringSet1 != c2.StringSet1 {
		return false
	}

	if c1.StringSet2 != c2.StringSet2 {
		return false
	}

	if c1.Int64set1 != c2.Int64set1 {
		return false
	}

	if c1.Int64set3 != c2.Int64set3 {
		return false
	}

	if c1.Int64Range1 != c2.Int64Range1 {
		return false
	}

	if c1.Float64set1 != c2.Float64set1 {
		return false
	}

	if c1.Float64set3 != c2.Float64set3 {
		return false
	}

	if c1.Float64Range1 != c2.Float64Range1 {
		return false
	}

	return true
}

func Test_configuration_LoadConfigurationFromString(t *testing.T) {
	t1 := `

boolSet1=true

boolSet2=false

boolSet3= false

stringSet1= "ss1"

stringSet2= ""

int64set1=1



int64set3= 0

int64Range1=1000

float64set1=1.0



float64set3= 0.0

float64Range1=1000.01
		
`
	t1_config := &configuration{
		rwlock: sync.RWMutex{},

		BoolSet1: true,

		BoolSet2: false,

		BoolSet3: false,

		StringSet1: "ss1",

		StringSet2: "",

		Int64set1: 1,

		Int64set3: 0,

		Int64Range1: 1000,

		Float64set1: 1.0,

		Float64set3: 0.0,

		Float64Range1: 1000.01,

		name2updatedFlags: nil,
	}

	type args struct {
		input  string
		config *configuration
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantErr2 bool
		wantErr3 bool
	}{
		{"t1", args{t1, t1_config}, false, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ap := &AllParameters{}
			if err := ap.LoadInitialValues(); err != nil {
				t.Errorf("LoadInitialValues failed.error %v", err)
			}
			config := &configuration{}
			if err := config.LoadConfigurationFromString(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigurationFromString() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				return
			}

			if err := ap.UpdateParametersWithConfiguration(config); (err != nil) != tt.wantErr2 {
				t.Errorf("UpdateParametersWithConfiguration failed. error:%v", err)
			}

			if (isconfigurationEqual(config, tt.args.config) != true) != tt.wantErr3 {
				t.Errorf("Configuration are not equal. %v vs %v ", config, tt.args.config)
				return
			}
		})
	}
}
