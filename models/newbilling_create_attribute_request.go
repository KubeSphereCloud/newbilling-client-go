// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewbillingCreateAttributeRequest //属性 - 带value字段，用于存放订阅传入具体的属性值。
//
//	message AttributeWithValue {
//	   //属性ID
//	   google.protobuf.StringValue attr_id = 1;
//	   //属性编码
//	   google.protobuf.StringValue attr_code = 2;
//	   //产品ID
//	   google.protobuf.StringValue prod_id = 3;
//	   //属性名称
//	   google.protobuf.StringValue name = 4;
//	   //属性描述
//	   google.protobuf.StringValue description = 5;
//	   //属性值类型
//	   google.protobuf.StringValue value_type = 6;
//	   //属性值
//	   google.protobuf.StringValue value = 7;
//	   //属性值单位
//	   google.protobuf.StringValue unit = 8;
//	   //是否需要计量
//	   google.protobuf.UInt32Value is_need_meter = 9;
//	   //聚合方法
//	   google.protobuf.StringValue aggregation = 10;
//	   //该产品属性对应的计量服务中的指标
//	   google.protobuf.StringValue metric_in_acess_sys = 11;
//	   //计量数据的时间精度
//	   google.protobuf.UInt32Value metering_period = 12;
//	   //创建时间
//	   google.protobuf.Timestamp create_time = 13;
//	   //更新时间
//	   google.protobuf.Timestamp update_time = 14;
//	}
//
// 创建属性请求信息
//
// swagger:model newbillingCreateAttributeRequest
type NewbillingCreateAttributeRequest struct {

	// 接入系统ID仅super user有效，接入系统管理员/成员会默认用当前登陆的接入系统
	AccessSysID string `json:"access_sys_id,omitempty"`

	// 属性编码
	AttrCode string `json:"attr_code,omitempty"`

	// 属性描述
	Description string `json:"description,omitempty"`

	// 是否需要计量 gotags:valid:"OneOf(0,1)"
	IsNeedMeter int64 `json:"is_need_meter,omitempty"`

	// 属性名称  @gotags:valid:"Required;MaxSize(30)"
	Name string `json:"name,omitempty"`

	// 产品ID  @gotags:valid:"Required"
	ProdID string `json:"prod_id,omitempty"`

	// 属性值单位
	Unit string `json:"unit,omitempty"`

	// 属性值类型  @gotags:valid:"Required;OneOf(bool,string,number)"
	ValueType string `json:"value_type,omitempty"`
}

// Validate validates this newbilling create attribute request
func (m *NewbillingCreateAttributeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this newbilling create attribute request based on context it is used
func (m *NewbillingCreateAttributeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewbillingCreateAttributeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewbillingCreateAttributeRequest) UnmarshalBinary(b []byte) error {
	var res NewbillingCreateAttributeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
