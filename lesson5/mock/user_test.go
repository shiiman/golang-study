package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// まず、テストスイートを作成します
type UserSuite struct {
	suite.Suite
}

// モックの定義.
type MockUser struct {
	mock.Mock
}

func (m *MockUser) getAge() int {
	args := m.Called()
	return args.Int(0)
}

// 実際のテスト(成功ケース).
func (s *UserSuite) TestSuccess() {
	m := new(MockUser)
	m.On("getAge").Return(20)

	svc := NewUserService(m)
	expected := 20

	assert.Equal(s.T(), expected, svc.User.getAge())
}

// 実際のテスト(失敗ケース).
func (s *UserSuite) TestFail() {
	m := new(MockUser)
	m.On("getAge").Return(8)

	svc := NewUserService(m)
	expected := 20

	assert.NotEqual(s.T(), expected, svc.User.getAge())
}

// テストを実行する.
func TestUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
