package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockUserInterface struct {
	mock.Mock
}

func (_m *MockUserInterface) Get(id int) (*UserModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*UserModel), ret.Error(1)
}

func TestUser_UserName(t *testing.T) {
	testUser := &UserModel{ID: 1, Name: "Tom"}
	mockUser := new(MockUserInterface)
	mockUser.On("Get", testUser.ID).Return(testUser, nil)
	u := &User{
		Dt: mockUser,
	}
	_, err := u.UserName(testUser.ID)
	if err != nil {
		t.Error(err)
	}
}
