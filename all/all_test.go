package all

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	// 由于All函数会调用多个API，这里只测试返回值的结构
	result := All()

	// 检查返回结果的基本结构
	if code, ok := result["code"]; !ok || code != 200 {
		t.Errorf("Expected result to have code 200, got %v", code)
	}

	if obj, ok := result["obj"]; !ok {
		t.Error("Expected result to have obj field")
	} else {
		// obj应该是一个map
		objMap, ok := obj.(map[string]interface{})
		if !ok {
			t.Errorf("Expected obj to be a map, got %T", obj)
		} else {
			// obj可以是空的，因为API可能失败，但我们至少检查它的类型
			t.Logf("obj contains %d entries", len(objMap))
		}
	}
}

func TestGetAllSourceNames(t *testing.T) {
	sources := GetAllSourceNames()

	// 检查是否返回了非空切片
	if len(sources) == 0 {
		t.Error("Expected GetAllSourceNames to return non-empty slice")
	}

	// 检查一些预期的源是否在列表中
	expectedSources := []string{"baidu", "bilibili", "zhihu", "weibo", "toutiao", "douban", "github", "v2ex"}
	for _, expected := range expectedSources {
		found := false
		for _, source := range sources {
			if source == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find source '%s' in GetAllSourceNames result", expected)
		}
	}

	// 检查切片中是否有重复项
	seen := make(map[string]bool)
	for _, source := range sources {
		if seen[source] {
			t.Errorf("Duplicate source found: %s", source)
		}
		seen[source] = true
	}
}

func TestSourceFuncMapNotEmpty(t *testing.T) {
	// 检查sourceFuncMap是否非空
	if len(sourceFuncMap) == 0 {
		t.Error("Expected sourceFuncMap to be non-empty")
	}

	// 检查一些预期的键是否存在
	expectedKeys := []string{"baidu", "bilibili", "zhihu", "weibo"}
	for _, key := range expectedKeys {
		if _, exists := sourceFuncMap[key]; !exists {
			t.Errorf("Expected key '%s' to exist in sourceFuncMap", key)
		}
	}
}

func TestSourceFuncMapValuesAreFunctions(t *testing.T) {
	for key, fn := range sourceFuncMap {
		if fn == nil {
			t.Errorf("Function for key '%s' is nil", key)
		}
		// 尝试调用函数，检查类型是否正确
		// 注意：这里不实际调用API，只是验证函数类型
		if reflect.TypeOf(fn).Kind() != reflect.Func {
			t.Errorf("Value for key '%s' is not a function", key)
		}
	}
}
