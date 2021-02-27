package os

import (
	"gitlab.com/evatix-go/core/osarchs"
	"gitlab.com/evatix-go/core/osconsts"
)

type Mac struct {
	X32, X64  string
	generated *string
}

func (receiver *Mac) Arch32() string {
	return receiver.X32
}

func (receiver *Mac) Arch64() string {
	return receiver.X64
}

func (receiver *Mac) Arch32Ptr() *string {
	return &receiver.X32
}

func (receiver *Mac) Arch64Ptr() *string {
	return &receiver.X64
}

func (receiver *Mac) GetDir(architecture osarchs.Architecture) string {
	if architecture.IsX32() {
		return receiver.Arch32()
	}

	return receiver.Arch64()
}

func (receiver *Mac) Generated() *string {
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

func (receiver *Mac) String() string {
	return *receiver.Generated()
}
