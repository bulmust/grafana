package recordingrule

import (
	"context"

	"github.com/grafana/grafana-app-sdk/app"
	"github.com/grafana/grafana-app-sdk/simple"
	"github.com/grafana/grafana/apps/alerting/rules/pkg/app/config"
)

func NewMutator(cfg config.RuntimeConfig) *simple.Mutator {
	return &simple.Mutator{
		MutateFunc: func(ctx context.Context, req *app.AdmissionRequest) (*app.MutatingResponse, error) {
			// TODO: cast to specific type and implement mutation
			// - mutate folder label to match folder from annotation
			// Add custom mutation logic here if needed
			return nil, nil
		},
	}
}
