// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import (
	"encoding/json"
	"time"
)

type data struct {
	token            string
	digest           string
	updatedAtTime    time.Time        `db:"updated_at_time"`
	flightDateLayout string           `db:"flight_data_layout"`
	value            []VariFlightData `db:"-"`                 // for cache only
	valueJSONString  string           `db:"value_json_string"` // for store only
}

func newData(token, digest string, updatedAtTime time.Time, flightDateLayout string, value []VariFlightData, valueJSONString string) *data {
	return &data{
		token:            token,
		digest:           digest,
		updatedAtTime:    updatedAtTime,
		flightDateLayout: flightDateLayout,
		value:            value,
		valueJSONString:  valueJSONString,
	}
}

func (d *data) isUpdatedWithin(dur time.Duration) bool {
	return d.updatedAtTime.Add(dur).After(time.Now())
}

func (d *data) isDepPlanDateBeyond(dur time.Duration) bool {
	tData, _ := parseLocalToUTC(d.flightDateLayout, d.value[0].FlightDeptimePlanDate, d.value[0].OrgTimezone)
	tNow := nowToUTC()
	return tNow.Add(dur).Before(tData) || tData.Add(dur).Before(tNow)
}

func (d *data) isSameDigest(digestToCompare string) bool {
	return d.digest == digestToCompare
}

func (d *data) encodeValue() error {
	if d.value != nil {
		bytes, err := json.Marshal(d.value)
		if err != nil {
			return err
		}
		d.valueJSONString = string(bytes)
	}
	return nil
}

func (d *data) decodeJSONString() error {
	if d.valueJSONString != "" {
		var value []VariFlightData
		if err := json.Unmarshal([]byte(d.valueJSONString), &value); err != nil {
			return err
		}
		d.value = value
		return nil
	}
	return nil
}

type VariFlightData struct {
	Fcategory             string `json:"fcategory,omitempty"`
	FlightNo              string `json:"FlightNo,omitempty"`
	FlightCompany         string `json:"FlightCompany,omitempty"`
	FlightDepcode         string `json:"FlightDepcode,omitempty"`
	FlightArrcode         string `json:"FlightArrcode,omitempty"`
	FlightDeptimePlanDate string `json:"FlightDeptimePlanDate,omitempty"`
	FlightArrtimePlanDate string `json:"FlightArrtimePlanDate,omitempty"`
	FlightDeptimeDate     string `json:"FlightDeptimeDate,omitempty"`
	FlightArrtimeDate     string `json:"FlightArrtimeDate,omitempty"`
	FlightState           string `json:"FlightState,omitempty"`
	FlightHTerminal       string `json:"FlightHTerminal,omitempty"`
	FlightTerminal        string `json:"FlightTerminal,omitempty"`
	OrgTimezone           string `json:"org_timezone,omitempty"`
	DstTimezone           string `json:"dst_timezone,omitempty"`
	ShareFlightNo         string `json:"shareFlightNo,omitempty"`
	StopFlag              string `json:"StopFlag,omitempty"`
	ShareFlag             string `json:"ShareFlag,omitempty"`
	VirtualFlag           string `json:"VirtualFlag,omitempty"`
	LegFlag               string `json:"LegFlag,omitempty"`
	FlightDep             string `json:"FlightDep,omitempty"`
	FlightArr             string `json:"FlightArr,omitempty"`
	FlightDepAirport      string `json:"FlightDepAirport,omitempty"`
	FlightArrAirport      string `json:"FlightArrAirport,omitempty"`
}

type VariFlightErrorInfo struct {
	ErrorCode int    `json:"error_code", omitempty`
	Error     string `json:"error, omitempty"`
}
