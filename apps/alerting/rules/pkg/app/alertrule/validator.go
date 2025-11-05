package alertrule

import (
	"context"

	"github.com/grafana/grafana-app-sdk/app"
	"github.com/grafana/grafana-app-sdk/simple"
	"github.com/grafana/grafana/apps/alerting/rules/pkg/app/config"
)

func NewValidator(cfg config.RuntimeConfig) *simple.Validator {
	return &simple.Validator{
		ValidateFunc: func(ctx context.Context, req *app.AdmissionRequest) error {
			// TODO: cast to specific type and implement validation
			// validations:
			// - check provenance status is valid
			// if !slices.Contains(model.AcceptedProvenanceStatuses, sourceProv) {
			// 	return nil, ngmodels.ProvenanceNone, fmt.Errorf("invalid provenance status: %s", sourceProv)
			// }
			// - validate group index label if group label is set
			// - validate that group and group index labels are not set on create
			// if p.Labels[model.GroupLabelKey] != "" || p.Labels[model.GroupIndexLabelKey] != "" {
			// 	return nil, k8serrors.NewBadRequest("cannot set group when creating alert rule")
			// }
			// - validate folder is set and valid
			// - validate notification settings if set
			// - enforce max name length (should be in the type definition?)
			// - ValidateRuleGroupInterval
			// - LabelsUserCannotSpecify
			// - For and Keep firing for > 0
			// Add custom validation logic here if needed
			return nil
		},
	}
}
