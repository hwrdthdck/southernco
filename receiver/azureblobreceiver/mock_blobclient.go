// Code generated by mockery v2.10.0. DO NOT EDIT.

package azureblobreceiver

// import (
// 	mock "github.com/stretchr/testify/mock"
// 	config "go.opentelemetry.io/collector/config"
// )

// // BlobClient is an autogenerated mock type for the BlobClient type
// type MockBlobClient struct {
// 	mock.Mock
// }

// // UploadData provides a mock function with given fields: data, dataType
// func (_m *MockBlobClient) UploadData(data []byte, dataType config.Type) error {
// 	ret := _m.Called(data, dataType)

// 	var r0 error
// 	if len(ret) > 0 {
// 		if rf, ok := ret.Get(0).(func([]byte, config.Type) error); ok {
// 			r0 = rf(data, dataType)
// 		} else {
// 			r0 = ret.Error(0)
// 		}
// 	}

// 	return r0
// }

// func NewMockBlobClient() *MockBlobClient {
// 	blobClient := &MockBlobClient{}
// 	blobClient.On("UploadData", mock.Anything, mock.Anything)
// 	return blobClient
// }
