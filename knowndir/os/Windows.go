package os

import (
	"gitlab.com/evatix-go/core/osarchs"
	"gitlab.com/evatix-go/core/osconsts"
)

type Windows struct {
	X32, X64  string
	generated *string
}

func (receiver *Windows) Arch32() string {
	return receiver.X32
}

func (receiver *Windows) Arch64() string {
	return receiver.X64
}

func (receiver *Windows) Arch32Ptr() *string {
	return &receiver.X32
}

func (receiver *Windows) Arch64Ptr() *string {
	return &receiver.X64
}

func (receiver *Windows) GetDir(architecture osarchs.Architecture) string {
	if architecture.IsX32() {
		return receiver.Arch32()
	}

	return receiver.Arch64()
}

func (receiver *Windows) Generated() *string {
	if receiver.generated != nil {
		return receiver.generated
	}

	if osconsts.IsX64Architecture {
		receiver.generated = receiver.Arch64Ptr()
	} else {
		receiver.generated = receiver.Arch32Ptr()
	}

	return receiver.generated
}

func (receiver *Windows) String() string {
	return *receiver.Generated()
}
