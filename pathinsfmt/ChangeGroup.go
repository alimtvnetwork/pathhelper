package pathinsfmt

type ChangeGroup struct {
	BaseIsRecursive
	GroupName string `json:"GroupName,omitempty"`
}
