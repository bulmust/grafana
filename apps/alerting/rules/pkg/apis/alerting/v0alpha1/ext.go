package v0alpha1

const (
	InternalPrefix                = "grafana.com/"
	GroupLabelKey                 = InternalPrefix + "group"
	GroupIndexLabelKey            = GroupLabelKey + "-index"
	ProvenanceStatusAnnotationKey = InternalPrefix + "provenance"
	// Copy of the max title length used in legacy validation path
	AlertRuleMaxTitleLength = 190
	// Annotation key used to store the folder UID on resources
	FolderAnnotationKey = "grafana.app/folder"
	FolderLabelKey      = FolderAnnotationKey
)

const (
	ProvenanceStatusNone = ""
	ProvenanceStatusAPI  = "api"
)

var (
	AcceptedProvenanceStatuses = []string{ProvenanceStatusNone, ProvenanceStatusAPI}
)
