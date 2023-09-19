package osfileinfos

import "os"

type Collection struct {
	Items []os.FileInfo
}

func New(infos []os.FileInfo) *Collection {
	return &Collection{
		Items: infos,
	}
}

func NewUsingCap(cap int) *Collection {
	infos := make([]os.FileInfo, 0, cap)

	return &Collection{
		Items: infos,
	}
}

func (it *Collection) IsEmpty() bool {
	return it.Items == nil ||
		len(it.Items) == 0
}

func (it *Collection) HasItems() bool {
	return it.Items != nil &&
		len(it.Items) > 0
}

func (it *Collection) Length() int {
	if it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *Collection) AddInfo(info os.FileInfo) *Collection {
	if info == nil {
		return it
	}

	it.Items = append(
		it.Items,
		info)

	return it
}

func (it *Collection) Add(info os.FileInfo, err error) error {
	if err != nil || info == nil {
		return err
	}

	it.Items = append(
		it.Items,
		info)

	return err
}
