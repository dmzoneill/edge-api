// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/images.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/redhatinsights/edge-api/pkg/models"
)

// MockImageServiceInterface is a mock of ImageServiceInterface interface.
type MockImageServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceInterfaceMockRecorder
}

// MockImageServiceInterfaceMockRecorder is the mock recorder for MockImageServiceInterface.
type MockImageServiceInterfaceMockRecorder struct {
	mock *MockImageServiceInterface
}

// NewMockImageServiceInterface creates a new mock instance.
func NewMockImageServiceInterface(ctrl *gomock.Controller) *MockImageServiceInterface {
	mock := &MockImageServiceInterface{ctrl: ctrl}
	mock.recorder = &MockImageServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageServiceInterface) EXPECT() *MockImageServiceInterfaceMockRecorder {
	return m.recorder
}

// AddUserInfo mocks base method.
func (m *MockImageServiceInterface) AddUserInfo(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserInfo", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserInfo indicates an expected call of AddUserInfo.
func (mr *MockImageServiceInterfaceMockRecorder) AddUserInfo(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserInfo", reflect.TypeOf((*MockImageServiceInterface)(nil).AddUserInfo), image)
}

// CheckImageName mocks base method.
func (m *MockImageServiceInterface) CheckImageName(name, account string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckImageName", name, account)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckImageName indicates an expected call of CheckImageName.
func (mr *MockImageServiceInterfaceMockRecorder) CheckImageName(name, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckImageName", reflect.TypeOf((*MockImageServiceInterface)(nil).CheckImageName), name, account)
}

// CreateImage mocks base method.
func (m *MockImageServiceInterface) CreateImage(image *models.Image, account string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImage", image, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateImage indicates an expected call of CreateImage.
func (mr *MockImageServiceInterfaceMockRecorder) CreateImage(image, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).CreateImage), image, account)
}

// CreateRepoForImage mocks base method.
func (m *MockImageServiceInterface) CreateRepoForImage(i *models.Image) (*models.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepoForImage", i)
	ret0, _ := ret[0].(*models.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRepoForImage indicates an expected call of CreateRepoForImage.
func (mr *MockImageServiceInterfaceMockRecorder) CreateRepoForImage(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepoForImage", reflect.TypeOf((*MockImageServiceInterface)(nil).CreateRepoForImage), i)
}

// GetImageByID mocks base method.
func (m *MockImageServiceInterface) GetImageByID(id string) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", id)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID.
func (mr *MockImageServiceInterfaceMockRecorder) GetImageByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockImageServiceInterface)(nil).GetImageByID), id)
}

// GetImageByOSTreeCommitHash mocks base method.
func (m *MockImageServiceInterface) GetImageByOSTreeCommitHash(commitHash string) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByOSTreeCommitHash", commitHash)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByOSTreeCommitHash indicates an expected call of GetImageByOSTreeCommitHash.
func (mr *MockImageServiceInterfaceMockRecorder) GetImageByOSTreeCommitHash(commitHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByOSTreeCommitHash", reflect.TypeOf((*MockImageServiceInterface)(nil).GetImageByOSTreeCommitHash), commitHash)
}

// RetryCreateImage mocks base method.
func (m *MockImageServiceInterface) RetryCreateImage(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryCreateImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetryCreateImage indicates an expected call of RetryCreateImage.
func (mr *MockImageServiceInterfaceMockRecorder) RetryCreateImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryCreateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).RetryCreateImage), image)
}

// SetErrorStatusOnImage mocks base method.
func (m *MockImageServiceInterface) SetErrorStatusOnImage(err error, i *models.Image) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetErrorStatusOnImage", err, i)
}

// SetErrorStatusOnImage indicates an expected call of SetErrorStatusOnImage.
func (mr *MockImageServiceInterfaceMockRecorder) SetErrorStatusOnImage(err, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetErrorStatusOnImage", reflect.TypeOf((*MockImageServiceInterface)(nil).SetErrorStatusOnImage), err, i)
}

// UpdateImage mocks base method.
func (m *MockImageServiceInterface) UpdateImage(image *models.Image, account string, previousImage *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImage", image, account, previousImage)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImage indicates an expected call of UpdateImage.
func (mr *MockImageServiceInterfaceMockRecorder) UpdateImage(image, account, previousImage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).UpdateImage), image, account, previousImage)
}

// UpdateImageStatus mocks base method.
func (m *MockImageServiceInterface) UpdateImageStatus(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImageStatus", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateImageStatus indicates an expected call of UpdateImageStatus.
func (mr *MockImageServiceInterfaceMockRecorder) UpdateImageStatus(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImageStatus", reflect.TypeOf((*MockImageServiceInterface)(nil).UpdateImageStatus), image)
}
