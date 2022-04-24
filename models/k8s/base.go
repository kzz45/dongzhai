package k8s

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	corev1 "k8s.io/api/core/v1"
)

type MapString struct {
	Values map[string]string `json:"values" gorm:"type:TEXT"`
}

func (annotation *MapString) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), annotation)
	case []byte:
		return json.Unmarshal(val, annotation)
	default:
		return errors.New("not support")
	}
}

func (annotation MapString) Value() (driver.Value, error) {
	bytes, err := json.Marshal(annotation)
	return string(bytes), err
}

type Taints struct {
	Values []corev1.Taint `json:"values" gorm:"type:TEXT"`
}

func (taints *Taints) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), taints)
	case []byte:
		return json.Unmarshal(val, taints)
	default:
		return errors.New("not support")
	}
}

func (taints Taints) Value() (driver.Value, error) {
	bytes, err := json.Marshal(taints)
	return string(bytes), err
}

type KeyValue struct {
	Values []KeyValueData `json:"values" gorm:"type:TEXT"`
}
type KeyValueData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (taints *KeyValue) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), taints)
	case []byte:
		return json.Unmarshal(val, taints)
	default:
		return errors.New("not support")
	}
}

func (taints KeyValue) Value() (driver.Value, error) {
	bytes, err := json.Marshal(taints)
	return string(bytes), err
}
