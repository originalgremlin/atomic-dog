package conf

// conf is a frozen map of dot-separated hierarchical string keys to any-type values
type conf map[string]interface{}

// Get TODO
func (c conf) Get(key string) interface{} {
	return c[key]
}

// Exists TODO
func (c conf) Exists(key string) bool {
	_, ok := c[key]
	return ok
}

// Conf is the single, application-wide configuration instance.
var Conf conf
