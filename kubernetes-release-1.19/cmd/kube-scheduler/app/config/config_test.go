package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	apiserver "k8s.io/apiserver/pkg/server"
)

func TestConfigComplete(t *testing.T) {
	scenarios := []struct {
		name   string
		want   *Config
		config *Config
	}{
		{
			name: "SetInsecureServingName",
			want: &Config{
				InsecureServing: &apiserver.DeprecatedInsecureServingInfo{
					Name: "healthz",
				},
			},
			config: &Config{
				InsecureServing: &apiserver.DeprecatedInsecureServingInfo{},
			},
		},
		{
			name: "SetMetricsInsecureServingName",
			want: &Config{
				InsecureMetricsServing: &apiserver.DeprecatedInsecureServingInfo{
					Name: "metrics",
				},
			},
			config: &Config{
				InsecureMetricsServing: &apiserver.DeprecatedInsecureServingInfo{},
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			cc := scenario.config.Complete()

			returnValue := cc.completedConfig.Config

			if diff := cmp.Diff(scenario.want, returnValue); diff != "" {
				t.Errorf("Complete(): Unexpected return value (-want, +got): %s", diff)
			}
		})
	}
}
