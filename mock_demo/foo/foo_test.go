package foo

import (
	"github.com/golang/mock/gomock"
	mock_foo "golang_practice/mock_demo/foo/mock"
	"testing"
)

func TestBaz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_foo.NewMockFoo(ctrl)
	m.EXPECT().Bar(gomock.Eq(123)).Return(123)

	expected := 124
	if got := Baz(m, 123); got != expected {
		t.Errorf(`Baz(m, "abc") = %d, want %d`, got, expected)
	}
}
