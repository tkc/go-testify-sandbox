package gomock

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Does not make any assertions. Returns 101 when Bar is invoked with 99.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(101).
		AnyTimes()

	// Does not make any assertions. Returns 103 when Bar is invoked with 101.
	m.
		EXPECT().
		Bar(gomock.Eq(101)).
		Return(103).
		AnyTimes()

}
