package pathext

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/extensionsconst"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/pathhelper/internal/isstr"
	"gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"
)

type Wrapper struct {
	fullPath                 string
	fileInfo                 *os.FileInfo
	fileInfoError            *errorwrapper.Wrapper
	dotExtension, extension  *string
	filteringExt             *string
	dotIndex                 *int
	fileNameWithExtension    *string
	baseDir                  *string
	fileNameWithoutExtension *string
}

func New(path string) Wrapper {
	return Wrapper{
		fullPath: path,
	}
}

func NewPtr(path string) *Wrapper {
	return &Wrapper{
		fullPath: path,
	}
}

func (receiver *Wrapper) IsPathEquals(path string, ignoreCase bool) bool {
	if ignoreCase {
		return strings.EqualFold(receiver.fullPath, path)
	}

	return receiver.fullPath == path
}

func (receiver *Wrapper) IsPathContains(path string) bool {
	return strings.Contains(receiver.fullPath, path)
}

func (receiver *Wrapper) ExtDotIndex() int {
	if receiver.dotIndex != nil {
		return *receiver.dotIndex
	}

	invalid := constants.InvalidNotFoundCase
	p := receiver.fullPath
	if p == "" {
		receiver.dotIndex = &invalid

		return *receiver.dotIndex
	}

	// doesn't look good on line break
	for i := len(p) - 1; i >= 0 && !(p[i] == constants.BackwardChar || p[i] == constants.ForwardChar); i-- {
		if p[i] == constants.DotChar {
			receiver.dotIndex = &i

			return i
		}
	}

	receiver.dotIndex = &invalid

	return *receiver.dotIndex
}

func (receiver *Wrapper) FileNameWithExtension() string {
	if receiver.fileNameWithExtension != nil {
		return *receiver.fileNameWithExtension
	}

	receiver.initializeProperties()

	return *receiver.fileNameWithExtension
}

func (receiver *Wrapper) FileNameWithoutExtension() string {
	if receiver.fileNameWithoutExtension != nil {
		return *receiver.fileNameWithoutExtension
	}

	receiver.initializeProperties()

	return *receiver.fileNameWithoutExtension
}

func (receiver *Wrapper) initializeProperties() {
	baseDir, fileNameWithExtension := pathsplitinternal.GetWithoutSlash(
		receiver.fullPath)
	fileNameWithoutExt := ""

	if receiver.HasExtension() {
		fileNameWithoutExt = strings.Replace(
			fileNameWithExtension,
			*receiver.DotExtension(),
			"",
			1)
	}

	receiver.fileNameWithExtension = &fileNameWithExtension
	receiver.baseDir = &baseDir
	receiver.fileNameWithoutExtension = &fileNameWithoutExt
}

// FilteringExt is ext and one char more from left.
// Panics if char is not there
func (receiver *Wrapper) FilteringExt() string {
	if receiver.filteringExt != nil {
		return *receiver.filteringExt
	}

	filteringExt := receiver.GetMoreThanExt(1)
	receiver.filteringExt = &filteringExt

	return *receiver.filteringExt
}

func (receiver *Wrapper) FileInfoWrapper() *os.FileInfo {
	if receiver.fileInfo != nil ||
		receiver.fileInfoError != nil &&
			receiver.fileInfoError.HasError() {
		return receiver.fileInfo
	}

	fileInfo, err :=
		os.Stat(receiver.fullPath)

	if err != nil {
		receiver.fileInfoError =
			errorwrapper.NewErrorPtr(err)
	} else {
		receiver.fileInfo = &fileInfo
	}

	return receiver.fileInfo
}

func (receiver *Wrapper) IsFile() bool {
	fileInfo := receiver.FileInfoWrapper()
	err := receiver.fileInfoError

	if err != nil && err.HasError() {
		return false
	}

	return fileInfo != nil &&
		!(*fileInfo).IsDir()
}

func (receiver *Wrapper) IsDir() bool {
	fileInfo := receiver.FileInfoWrapper()
	err := receiver.fileInfoError

	if err != nil && err.HasError() {
		return false
	}

	return fileInfo != nil &&
		(*fileInfo).IsDir()
}

func (receiver *Wrapper) GetMoreThanExt(moreIndex int) string {
	dotIndex := receiver.ExtDotIndex()

	if dotIndex == constants.InvalidValue {
		return ""
	}

	newIndex := dotIndex - moreIndex

	return receiver.fullPath[newIndex:]
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (receiver *Wrapper) DotExtension() *string {
	if receiver.dotExtension == nil {
		dotExt := receiver.GetMoreThanExt(0)
		receiver.dotExtension = &dotExt
	}

	return receiver.dotExtension
}

func (receiver *Wrapper) HasExtension() bool {
	return receiver.ExtDotIndex() > -1
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (receiver *Wrapper) Extension() *string {
	if receiver.extension != nil {
		return receiver.extension
	}

	dotExt := *receiver.DotExtension()

	if len(dotExt) > 0 && dotExt[0] == constants.Dot[0] {
		ext := dotExt[1:]
		receiver.extension = &ext
	} else {
		receiver.extension = &dotExt
	}

	return receiver.extension
}

func (receiver *Wrapper) IsExtension(extension string) bool {
	return *receiver.Extension() == extension
}

func (receiver *Wrapper) IsExtOrDotExt(extOrDotExt string) bool {
	return *receiver.Extension() == extOrDotExt ||
		*receiver.DotExtension() == extOrDotExt
}

func (receiver *Wrapper) IsDotExtension(dotExtension string) bool {
	return *receiver.DotExtension() == dotExtension
}

func (receiver *Wrapper) IsAnyOfExtension(extensions ...string) bool {
	if extensions == nil {
		return false
	}

	return receiver.IsAnyOfExtensionPtr(&extensions)
}

func (receiver *Wrapper) IsAnyOfExtensionPtr(
	extensions *[]string,
) bool {
	if extensions == nil || len(*extensions) == 0 {
		return false
	}

	currentExt := *receiver.Extension()
	for _, ext := range *extensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}

func (receiver *Wrapper) IsAnyOfDotExtension(
	dotExtensions ...string,
) bool {
	if dotExtensions == nil {
		return false
	}

	return receiver.IsAnyOfDotExtensionPtr(&dotExtensions)
}

func (receiver *Wrapper) IsAnyOfDotExtensionPtr(
	dotExtensions *[]string,
) bool {
	if dotExtensions == nil || len(*dotExtensions) == 0 {
		return false
	}

	currentExt := *receiver.DotExtension()
	for _, ext := range *dotExtensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}

func (receiver *Wrapper) IsExtensionFilterMatch(
	extensionFilter string,
) bool {
	if extensionFilter == extensionsconst.AllFiles {
		return true
	}

	length := len(extensionFilter)
	if length == 0 {
		return false
	}

	currentDotExt := *receiver.DotExtension()

	if length == 2 {
		// *.
		firstChar := extensionFilter[coreindexes.First]
		secondChar := extensionFilter[coreindexes.Second]

		// a file name ends with dot.
		isFileEndsWithDot := firstChar == constants.AstrekChar &&
			secondChar == constants.DotChar &&
			currentDotExt == constants.Dot

		return isFileEndsWithDot ||
			firstChar == constants.DotChar &&
				currentDotExt == extensionFilter
	}

	if length < 3 {
		return false
	}

	firstChar := extensionFilter[coreindexes.First]
	secondChar := extensionFilter[coreindexes.Second]

	// *.ext
	if firstChar == constants.AstrekChar &&
		secondChar == constants.DotChar {
		dotExt := extensionFilter[coreindexes.Second:]

		if dotExt == currentDotExt {
			return true
		}
	}

	// .whatever
	if firstChar == constants.DotChar &&
		extensionFilter == currentDotExt {
		return true
	}

	// {don't care}what.ever, ends with extension
	return isstr.EndsWith(
		receiver.fullPath,
		extensionFilter,
		true)
}

func (receiver *Wrapper) IsExtensionFiltersMatch(
	extensionsFilter *[]string,
	extensionsLength int,
) bool {
	if extensionsLength == 0 {
		return false
	}

	for _, extFilter := range *extensionsFilter {
		if receiver.IsExtensionFilterMatch(extFilter) {
			return true
		}
	}

	return false
}

func (receiver *Wrapper) IsNameWithExtensionMatches(
	fullPath string,
) bool {
	_, fileName := pathsplitinternal.Get(fullPath)

	return isstr.EndsWith(
		receiver.fullPath,
		fileName,
		true)
}
