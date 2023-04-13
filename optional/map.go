package optional

type Map map[string]interface{}

func NewMap(value map[string]interface{}) Map {
	return Map(value)
}

func EmptyMap() Map {
	return Map(nil)
}

func (m Map) IsSet() bool {
	return m != nil
}

func (m Map) Value() map[string]interface{} {
	return m
}

func (m Map) Default(defaultValue map[string]interface{}) map[string]interface{} {
	if m != nil {
		return m
	}
	return defaultValue
}
