package common

import "fmt"

type PropertyItem struct {
	Code      string `json:"code"`
	Extension string `json:"extension"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Unit      string `json:"unit,omitempty"`
	ValueType int    `json:"valueType,omitempty"`
}

func (o *PropertyItem) ToString() string {
	return fmt.Sprintf("PropertyItem [Key=%s, Value=%v, Unit=%s]", o.Key, o.Value, o.Unit)
}

type PropertyGroup struct {
	Group string          `json:"group"`
	Items []*PropertyItem `json:"items"`
}
