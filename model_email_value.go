package decta_api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var _ MappedNullable = &EmailValue{}

type EmailValue struct {
	Value string `json:"value"`
}

type _EmailValue EmailValue

func NewEmailValue(value string) *EmailValue {
	this := EmailValue{}
	this.Value = value
	return &this
}

func NewEmailValueWithDefaults() *EmailValue {
	this := EmailValue{}
	return &this
}

// GetValue returns the Value field value
func (o *EmailValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EmailValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EmailValue) SetValue(v string) {
	o.Value = v
}

func (o EmailValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *EmailValue) UnmarshalJSON(data []byte) (err error) {
	requiredProperties := []string{
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEmailValue := _EmailValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailValue)

	if err != nil {
		return err
	}

	*o = EmailValue(varEmailValue)

	return err
}

type NullableEmailValue struct {
	value *EmailValue
	isSet bool
}

func (v NullableEmailValue) Get() *EmailValue {
	return v.value
}

func (v *NullableEmailValue) Set(val *EmailValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailValue(val *EmailValue) *NullableEmailValue {
	return &NullableEmailValue{value: val, isSet: true}
}

func (v NullableEmailValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
