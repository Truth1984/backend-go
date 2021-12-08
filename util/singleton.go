package util

func SingletonSet(key string, value interface{}) interface{} {
	if singleton[key] == nil {
		singleton[key] = value
	}
	return singleton[key]
}

func SingletonGet(key string) interface{} {
	return singleton[key]
}

func SingletonGetAll() map[string]interface{} {
	return singleton
}

func SingletonHas(key string) bool {
	return singleton[key] != nil
}
