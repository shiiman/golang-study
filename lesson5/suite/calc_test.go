package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// まず、テストスイートを作成します
type SumSuite struct {
	suite.Suite
}

func (suite *SumSuite) SetupSuite() {
	suite.T().Log("SetupSuite")
}

func (suite *SumSuite) TearDownSuite() {
	suite.T().Log("TearDownSuite")
}

func (suite *SumSuite) SetupTest() {
	suite.T().Log("SetupTest")
}

func (suite *SumSuite) TearDownTest() {
	suite.T().Log("TeardownTest")
}

func (suite *SumSuite) BeforeTest(suiteName, testName string) {
	suite.T().Log(testName + ": BeforeTest")
}

func (suite *SumSuite) AfterTest(suiteName, testName string) {
	suite.T().Log(testName + ": AfterTest")
}

// 実際のテスト(成功ケース).
// func (s *SumSuite) TestSuccess() {
// 	actual := sum(3, 5)
// 	expected := 8

// 	assert.Equal(s.T(), expected, actual)
// }

// // 実際のテスト(失敗ケース).
// func (s *SumSuite) TestFail() {
// 	actual := sum(3, 5)
// 	expected := 7

// 	assert.NotEqual(s.T(), expected, actual)
// }

// サブテストで実行する場合.
func (s *SumSuite) TestSub() {
	s.Run("TestSuccess", func() {
		actual := sum(3, 5)
		expected := 8

		assert.Equal(s.T(), expected, actual)
	})

	s.Run("TestFail", func() {
		actual := sum(3, 5)
		expected := 7

		assert.NotEqual(s.T(), expected, actual)
	})
}

// テストを実行する.
func TestSum(t *testing.T) {
	suite.Run(t, new(SumSuite))
}
