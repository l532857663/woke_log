package dao

import (
	"backend-gateway/dao/mocks"

	"github.com/golang/mock/gomock"
	logging "github.com/op/go-logging"
)

var (
	mockLogger = logging.MustGetLogger("dao/mock_dao")
)

// 创建一个 DAO 并返回对象
func NewMock() Storage {
	// mock控制器通过NewController接口生成，是mock生态系统的顶层控制，定义了mock对象的作用域和生命周期，以及mock对象的期望
	// 多个协程同时调用控制器的方法是安全的
	// 当用例结束后，控制器会检查所有剩余期望的调用是否满足条件
	ctl := gomock.NewController(mockLogger)
	// 进行 mock 用例的期望值断言
	defer ctl.Finish()

	// mock 对象注入控制器
	mockDB := mocks.NewMockStorage(ctl)

	return mockDB
}

// Close mock
func CloseMockExpect(mockDB Storage) {
	// 直接返回
	mockDB.(*mocks.MockStorage).EXPECT().Close().Return()

	return
}

func getFloat64Pointer(value float64) *float64 {
	return &value
}

func getUint64Pointer(value uint64) *uint64 {
	return &value
}
