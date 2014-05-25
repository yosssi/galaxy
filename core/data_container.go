package core

import "fmt"

// dataContainer represents a data container.
type dataContainer struct {
	data map[interface{}]interface{}
}

// SetData sets the data to the data container.
func (dc *dataContainer) SetData(key, value interface{}) error {
	if _, prs := dc.data[key]; prs {
		return fmt.Errorf(`the key has already been set [key: %+v][value: %+v]`, key, value)
	}

	dc.data[key] = value

	return nil
}

// SetForceData sets the data to the data container forcibly.
func (dc *dataContainer) SetForceData(key, value interface{}) {
	dc.data[key] = value
}

// GetData gets the data from the data container.
func (dc *dataContainer) GetData(key interface{}) (interface{}, bool) {
	value, ok := dc.data[key]

	return value, ok
}
