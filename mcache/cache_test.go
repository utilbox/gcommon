package mcache

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	Init(2*SizeKB, false, 0)
}

func TestIsUsable(t *testing.T) {
	Convey("Test whether mcache is initialized correctly", t, func() {
		usable := IsUsable()
		So(usable, ShouldBeTrue)
	})
}

func TestSetAndGet(t *testing.T) {
	Convey("Test mcache Set and Get", t, func() {
		key := "test_key"
		value := "test_value"
		Convey("Test mcache Set", func() {
			err := Set(key, []byte(value), 10)
			So(err, ShouldBeNil)
		})

		Convey("Test mcache Get", func() {
			raw, err := Get(key)
			So(err, ShouldBeNil)
			So(string(raw), ShouldEqual, value)
		})
	})
}

func TestSetAndGetInt(t *testing.T) {
	Convey("Test mcache SetInt and GetInt", t, func() {
		key := "key_set_int"
		value := 100
		Convey("Test mcache SetInt", func() {
			err := SetInt(key, value, 10)
			So(err, ShouldBeNil)
		})

		Convey("Test mcache GetInt", func() {
			v, err := GetInt(key)
			So(err, ShouldBeNil)
			So(v, ShouldEqual, value)
		})
	})
}

func TestSetAndGetString(t *testing.T) {
	Convey("Test SetString and GetString", t, func() {
		key := "key_set_string"
		value := "value_set_string"
		Convey("Test mcache SetString", func() {
			err := SetString(key, value, 10)
			So(err, ShouldBeNil)
		})

		Convey("Test mcache GetString", func() {
			v, err := GetString(key)
			So(err, ShouldBeNil)
			So(v, ShouldEqual, value)
		})
	})
}

func TestDel(t *testing.T) {
	Convey("Test mcache Del", t, func() {
		key := "key_test_del"
		value := "value_test_del"

		Convey("Test mcache SetString, should return nil", func() {
			err := SetString(key, value, 10)
			So(err, ShouldBeNil)
		})

		Convey("Test mcache GetString, should return no error when key not deleted", func() {
			v, err := GetString(key)
			So(err, ShouldBeNil)
			So(v, ShouldEqual, value)
		})

		Convey("Test mcache Del, should return true when key existed", func() {
			deleted := Del(key)
			So(deleted, ShouldBeTrue)
		})

		Convey("Test mcache GetString, should return error after key deleted", func() {
			v, err := GetString(key)
			So(err, ShouldNotBeNil)
			So(v, ShouldEqual, "")
		})
	})
}

func TestTTL(t *testing.T) {
	Convey("Test mcache TTL", t, func() {
		key := "key_test_TTL"
		value := "value_test_TTL"
		Convey("Test mcache SetString for TTL, should return no error when setting key", func() {
			err := SetString(key, value, 10)
			So(err, ShouldBeNil)
		})

		Convey("Test mcache TTL, should return none zero number", func() {
			ttl, err := TTL(key)
			So(err, ShouldBeNil)
			So(ttl, ShouldNotBeZeroValue)
		})

	})
}
