package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Joiner struct {
	items []string
}

func NewJoiner(capacity int) *Joiner {
	list := make(
		[]string,
		constants.Zero,
		capacity)

	return &Joiner{items: list}
}

func NewJoiner5() *Joiner {
	list := make(
		[]string,
		constants.Zero,
		constants.ArbitraryCapacity5)

	return &Joiner{items: list}
}

func EmptyJoiner() *Joiner {
	return &Joiner{items: []string{}}
}

func (receiver *Joiner) Length() int {
	return len(receiver.items)
}

func (receiver *Joiner) IsEmpty() bool {
	return len(receiver.items) == 0
}

func (receiver *Joiner) HasItems() bool {
	return len(receiver.items) > 0
}

func (receiver *Joiner) Add(addingPath string) *Joiner {
	receiver.items = append(
		receiver.items,
		addingPath)

	return receiver
}

func (receiver *Joiner) Adds(addingPaths ...string) *Joiner {
	if addingPaths == nil {
		return receiver
	}

	for _, curPath := range addingPaths {
		receiver.items = append(
			receiver.items,
			curPath)
	}

	return receiver
}

// ToString isNormalizePlusLongPathFix if true then adds UNC path fix for Windows
func (receiver *Joiner) ToString(sep string, isNormalizePlusLongPathFix bool) string {
	finalPath := strings.Join(receiver.items, sep)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		sep,
		finalPath)
}

// normalize and long path fix true
//
// Usages osconsts.PathSeparator as separator
func (receiver *Joiner) String() string {
	return receiver.ToString(
		osconsts.PathSeparator,
		true)
}
