// Code generated by GoJay. DO NOT EDIT.

package event 

import "github.com/francoispqt/gojay"

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *NormalizedEvent) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "conn_id":
		return dec.Uint64(&v.ConnID)
	case "event_id":
		return dec.String(&v.EventID)
	case "timestamp":
		return dec.String(&v.Timestamp)
	case "sensor":
		return dec.String(&v.Sensor)
	case "plugin_id":
		return dec.Int(&v.PluginID)
	case "plugin_sid":
		return dec.Int(&v.PluginSID)
	case "product":
		return dec.String(&v.Product)
	case "category":
		return dec.String(&v.Category)
	case "subcategory":
		return dec.String(&v.SubCategory)
	case "src_ip":
		return dec.String(&v.SrcIP)
	case "src_port":
		return dec.Int(&v.SrcPort)
	case "dst_ip":
		return dec.String(&v.DstIP)
	case "dst_port":
		return dec.Int(&v.DstPort)
	case "protocol":
		return dec.String(&v.Protocol)
	case "custom_data1":
		return dec.String(&v.CustomData1)
	case "custom_label1":
		return dec.String(&v.CustomLabel1)
	case "custom_data2":
		return dec.String(&v.CustomData2)
	case "custom_label2":
		return dec.String(&v.CustomLabel2)
	case "custom_data3":
		return dec.String(&v.CustomData3)
	case "custom_label3":
		return dec.String(&v.CustomLabel3)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *NormalizedEvent) NKeys() int { return 20 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *NormalizedEvent) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Uint64KeyOmitEmpty("conn_id", v.ConnID)
	enc.StringKey("event_id", v.EventID)
	enc.StringKey("timestamp", v.Timestamp)
	enc.StringKey("sensor", v.Sensor)
	enc.IntKey("plugin_id", v.PluginID)
	enc.IntKey("plugin_sid", v.PluginSID)
	enc.StringKey("product", v.Product)
	enc.StringKey("category", v.Category)
	enc.StringKeyOmitEmpty("subcategory", v.SubCategory)
	enc.StringKey("src_ip", v.SrcIP)
	enc.IntKey("src_port", v.SrcPort)
	enc.StringKey("dst_ip", v.DstIP)
	enc.IntKey("dst_port", v.DstPort)
	enc.StringKey("protocol", v.Protocol)
	enc.StringKeyOmitEmpty("custom_data1", v.CustomData1)
	enc.StringKeyOmitEmpty("custom_label1", v.CustomLabel1)
	enc.StringKeyOmitEmpty("custom_data2", v.CustomData2)
	enc.StringKeyOmitEmpty("custom_label2", v.CustomLabel2)
	enc.StringKeyOmitEmpty("custom_data3", v.CustomData3)
	enc.StringKeyOmitEmpty("custom_label3", v.CustomLabel3)
}

// IsNil returns wether the structure is nil value or not
func (v *NormalizedEvent) IsNil() bool { return v == nil }
