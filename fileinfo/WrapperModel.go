package fileinfo

type WrapperModel struct {
	RawPath     string
	IsDirectory bool
	IsFile      bool
	IsEmptyPath bool
	Separator   string
}
