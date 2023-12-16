package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
	gofrLog "gofr.dev/pkg/log"

	"sample/model"
	"sample/store"
)

func newMock(t *testing.T) (gofrLog.Logger, *store.MockCar) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockCar(ctrl)
	mockLogger := gofrLog.NewMockLogger(io.Discard)

	return mockLogger, mockStore
}

func createContext(method string, params map[string]string, car interface{}, logger gofrLog.Logger, t *testing.T) *gofr.Context {
	body, err := json.Marshal(car)
	if err != nil {
		t.Fatalf("Error while marshalling model: %v", err)
	}

	r := httptest.NewRequest(method, "/dummy", bytes.NewBuffer(body))
	query := r.URL.Query()

	for key, value := range params {
		query.Add(key, value)
	}

	r.URL.RawQuery = query.Encode()

	req := request.NewHTTPRequest(r)

	return gofr.NewContext(nil, req, nil)
}

func Test_Create(t *testing.T) {
	mockLogger, mockStore := newMock(t)
	h := New(mockStore)
	car := model.Car{
		ID:    1,
		Name:  "test car",
		Color: "test color",
	}

	testCases := []struct {
		desc      string
		input     interface{}
		mockCalls []*gomock.Call
		expRes    interface{}
		expErr    error
	}{
		{"success case", car, []*gomock.Call{
			mockStore.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), &car).Return(&car, nil).Times(1),
		}, &car, nil},
		{"failure case", car, []*gomock.Call{
			mockStore.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), &car).Return(nil, errors.Error("test error")).Times(1),
		}, nil, errors.Error("test error")},
		{"failure case-bind error", "test", []*gomock.Call{}, nil, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := createContext(http.MethodPost, nil, tc.input, mockLogger, t)
			res, err := h.Create(ctx)

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}

func Test_GetByID(t *testing.T) {
	mockLogger, mockStore := newMock(t)
	h := New(mockStore)
	car := model.Car{
		ID:    1,
		Name:  "test car",
		Color: "test color",
	}

	testCases := []struct {
		desc      string
		id        string
		mockCalls []*gomock.Call
		expRes    interface{}
		expErr    error
	}{
		{"success case", "1", []*gomock.Call{
			mockStore.EXPECT().GetByID(gomock.AssignableToTypeOf(&gofr.Context{}), 1).Return(&car, nil).Times(1),
		}, &car, nil},
		{"failure case", "1", []*gomock.Call{
			mockStore.EXPECT().GetByID(gomock.AssignableToTypeOf(&gofr.Context{}), 1).Return(nil, errors.EntityNotFound{Entity: "car", ID: "1"}).Times(1),
		}, nil, errors.EntityNotFound{Entity: "car", ID: "1"}},
		{"failure case-missing id", "", []*gomock.Call{}, nil, errors.MissingParam{Param: []string{"id"}}},
		{"failure case-invalid id", "test", []*gomock.Call{}, nil, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := createContext(http.MethodGet, nil, nil, mockLogger, t)

			ctx.SetPathParams(map[string]string{"id": tc.id})

			res, err := h.GetByID(ctx)

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}

func Test_Update(t *testing.T) {
	mockLogger, mockStore := newMock(t)
	h := New(mockStore)
	car := model.Car{
		ID:    1,
		Name:  "test car",
		Color: "test color",
	}

	testCases := []struct {
		desc      string
		id        string
		input     interface{}
		mockCalls []*gomock.Call
		expRes    interface{}
		expErr    error
	}{
		{"success case", "1", car, []*gomock.Call{
			mockStore.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), &car).Return(&car, nil).Times(1),
		}, &car, nil},
		{"failure case", "1", car, []*gomock.Call{
			mockStore.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), &car).Return(nil, errors.Error("test error")).Times(1),
		}, nil, errors.Error("test error")},
		{"failure case-missing id", "", "test", []*gomock.Call{}, nil, errors.MissingParam{Param: []string{"id"}}},
		{"failure case-invalid id", "test", "test", []*gomock.Call{}, nil, errors.InvalidParam{Param: []string{"id"}}},
		{"failure case-bind error", "1", "test", []*gomock.Call{}, nil, errors.InvalidParam{Param: []string{"body"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := createContext(http.MethodPut, nil, tc.input, mockLogger, t)

			ctx.SetPathParams(map[string]string{"id": tc.id})

			res, err := h.Update(ctx)

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}

func Test_Delete(t *testing.T) {
	mockLogger, mockStore := newMock(t)
	h := New(mockStore)

	testCases := []struct {
		desc      string
		id        string
		mockCalls []*gomock.Call
		expErr    error
	}{
		{"success case", "1", []*gomock.Call{
			mockStore.EXPECT().Delete(gomock.AssignableToTypeOf(&gofr.Context{}), 1).Return(nil).Times(1),
		}, nil},
		{"failure case", "1", []*gomock.Call{
			mockStore.EXPECT().Delete(gomock.AssignableToTypeOf(&gofr.Context{}), 1).Return(errors.Error("test error")).Times(1),
		}, errors.Error("test error")},
		{"failure case-missing id", "", []*gomock.Call{}, errors.MissingParam{Param: []string{"id"}}},
		{"failure case-invalid id", "test", []*gomock.Call{}, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := createContext(http.MethodDelete, nil, nil, mockLogger, t)

			ctx.SetPathParams(map[string]string{"id": tc.id})

			_, err := h.Delete(ctx)

			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}

func Test_validateID(t *testing.T) {
	testCases := []struct {
		desc   string
		id     string
		expID  int
		expErr error
	}{
		{"success case", "1", 1, nil},
		{"failure case-empty ID", "", 0, errors.MissingParam{Param: []string{"id"}}},
		{"failure case-invalid ID", "test", 0, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			id, err := validateID(tc.id)

			assert.Equal(t, tc.expID, id, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}
