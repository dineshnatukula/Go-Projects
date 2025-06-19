package strings

import "testing"

func TestMyAtoi(t *testing.T) {
	t.Run("TestMyAtoi1", func(t *testing.T) {
		if MyAtoi("-123") != -123 {
			t.Errorf("Expected -123, but got %d", MyAtoi("-123"))
		}
	})
	t.Run("TestMyAtoi2", func(t *testing.T) {
		if MyAtoi("") != 0 {
			t.Errorf("Expected 0, but got %d", MyAtoi(""))
		}
	})
	t.Run("TestMyAtoi3", func(t *testing.T) {
		if MyAtoi("___ASDF-123") != 0 {
			t.Errorf("Expected -123, but got %d", MyAtoi("___ASDF-123"))
		}
	})
	t.Run("TestMyAtoi4", func(t *testing.T) {
		if MyAtoi("___123") != 0 {
			t.Errorf("Expected 123, but got %d", MyAtoi("___123"))
		}
	})
	t.Run("TestMyAtoi5", func(t *testing.T) {
		if MyAtoi("-12323232323223") != -2147483648 {
			t.Errorf("Expected -123, but got %d", MyAtoi("-12323232323223"))
		}
	})
	t.Run("TestMyAtoi6", func(t *testing.T) {
		if MyAtoi("-1230083rgeert343453w452ser23rwer") != -1230083 {
			t.Errorf("Expected -123, but got %d", MyAtoi("-1230083rgeert343453w452ser23rwer"))
		}
	})
}
