package config

import "context"

// RuntimeConfig holds configuration values needed at runtime by the alerting/rules app from the running Grafana instance.
type RuntimeConfig struct {
	// function to check folder existence given its uid
	FolderValidator func(ctx context.Context, folderUID string) (bool, error)
	// base evaluation interval in seconds
	BaseEvaluationInterval uint
	// list of strings which are illegal for label keys on rules
	ReservedLabelKeys             []string
	NotificationSettingsValidator func(ctx context.Context, receiver string) (bool, error)
}
