package bili_test

import (
	"reflect"
	"testing"

	"github.com/rongyuio/bilibili"
	"github.com/rongyuio/bilibili/bili"
)

// TestNewWithSharesRoot 验证聚合入口与子包入口复用同一个根客户端。
func TestNewWithSharesRoot(t *testing.T) {
	root := bilibili.New()
	client := bili.NewWith(root)

	if client.Client != root {
		t.Fatal("聚合入口没有复用根客户端")
	}
	if client.Video.Client != root {
		t.Fatal("video 入口没有复用根客户端")
	}
	if client.User.Client != root {
		t.Fatal("user 入口没有复用根客户端")
	}
}

// TestFieldsInitialized 验证生成器覆盖到的每个域都已初始化。
func TestFieldsInitialized(t *testing.T) {
	client := bili.New()
	v := reflect.ValueOf(client).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Pointer && field.IsNil() {
			t.Errorf("字段 %s 未初始化", v.Type().Field(i).Name)
		}
	}
}

// TestNewWithNil 验证 nil 根客户端退回默认构造，不返回 nil 字段。
func TestNewWithNil(t *testing.T) {
	client := bili.NewWith(nil)
	if client == nil || client.Video == nil {
		t.Fatal("NewWith(nil) 没有退回默认构造")
	}
}
