package pathinsfmt

type BaseUserNamePlusGroupName struct {
	BaseGroupName
	UserName *string `json:"UserName,omitempty"` // Not define or empty string or * means keeping the existing one
}
