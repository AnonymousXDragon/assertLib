package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	a := &Assert{}
	a.t = t
	return a
}

func (a *Assert) Equal(actual, expected any, msg ...any) bool {
	if reflect.DeepEqual(actual, expected) {
		return true
	}

	a.fail(fmt.Sprintf("Not Equal expected %v but got %v", expected, actual), msg...)
	// fmt.Println("s")
	return false
}

func (a *Assert) NotEqual(actual, expected any, msg ...any) bool {
	if reflect.DeepEqual(actual, expected) {
		a.fail(fmt.Sprintf("Should not be equal %v", expected), msg...)
		return false
	}
	return true
}

func (a *Assert) True(value bool, msg ...any) bool {
	if value {
		return true
	}
	a.fail("should be true", msg...)
	return false
}

func (a *Assert) False(value bool, msg ...any) bool {
	if !value {
		return true
	}
	a.fail("should be false", msg...)
	return false
}

func (a *Assert) Nil(value any, msg ...any) bool {
	if value == nil {
		return true
	}
	a.fail(fmt.Sprintf("%v shoudl be nil", value), msg...)
	return false
}

func (a *Assert) NotNil(value any, msg ...any) bool {
	if value != nil {
		return true
	}
	a.fail(fmt.Sprintf("%v should not be nil", value), msg...)
	return false
}

func (a *Assert) Contains(stack any, required any, msg ...any) bool {
	ok, found := includeElement(stack, required)
	if ok && found {
		return true
	}
	a.fail(fmt.Sprintf("the value %v is not in %v", required, stack))
	return false
}

func (a *Assert) NoError(err error, msg ...any) bool {
	if err == nil {
		return true
	}
	a.fail(fmt.Sprintf("%v is not nil", err), msg...)
	return false
}

func (a *Assert) fail(fmsg string, msgs ...any) {
	msg := convertMsgArgs(msgs...)
	if len(msg) > 0 {
		fmt.Println("fishy", fmsg, msg)
		a.t.Errorf("%s\n\t\r%s", msg, fmsg)
	} else {
		a.t.Errorf("\r%s", fmsg)
	}

}
func convertMsgArgs(msgs ...any) string {

	var builder strings.Builder

	if len(msgs) == 0 || msgs == nil {
		return ""
	}

	if len(msgs) == 1 {
		return msgs[0].(string)
	}

	if len(msgs) > 1 {
		fmt.Println("called...!")
		for _, msg := range msgs {
			builder.WriteString(msg.(string))
		}
		fmt.Println(builder.String())
		return builder.String()
	}

	return ""
}

func includeElement(list any, element any) (ok, found bool) {
	listValue := reflect.ValueOf(list)
	elemtValue := reflect.ValueOf(element)

	defer func() {
		if err := recover(); err != nil {
			ok = false
			found = false
		}
	}()

	if listValue.Kind() == reflect.String {
		return true, strings.Contains(listValue.String(), elemtValue.String())
	}

	if listValue.Kind() != reflect.Slice && listValue.Kind() != reflect.Array {
		return false, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if reflect.DeepEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return false, false

}
