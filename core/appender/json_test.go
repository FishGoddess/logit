// Copyright 2022 FishGoddess. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package appender

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

// go test -v -cover -run=^TestJsonAppender$
func TestJsonAppender(t *testing.T) {
	appender := new(jsonAppender)

	buffer := make([]byte, 0, 512)
	buffer = appender.Begin(buffer)
	buffer = appender.AppendAny(buffer, "any", map[string]interface{}{"AppendAnyKey1": "value", "AppendAnyKey2": 123, "AppendAnyKey3": 666.666})
	buffer = appender.AppendBool(buffer, "bool", true)
	buffer = appender.AppendByte(buffer, "byte", 'b')
	buffer = appender.AppendRune(buffer, "rune", 'r')
	buffer = appender.AppendInt(buffer, "int", -1)
	buffer = appender.AppendInt8(buffer, "int8", -8)
	buffer = appender.AppendInt16(buffer, "int16", -16)
	buffer = appender.AppendInt32(buffer, "int32", -32)
	buffer = appender.AppendInt64(buffer, "int64", -64)
	buffer = appender.AppendUint(buffer, "uint", 1)
	buffer = appender.AppendUint8(buffer, "uint8", 8)
	buffer = appender.AppendUint16(buffer, "uint16", 16)
	buffer = appender.AppendUint32(buffer, "uint32", 32)
	buffer = appender.AppendUint64(buffer, "uint64", 64)
	buffer = appender.AppendFloat32(buffer, "float32", 32.32)
	buffer = appender.AppendFloat64(buffer, "float64", 64.64)
	buffer = appender.AppendString(buffer, "string", "value")
	buffer = appender.AppendTime(buffer, "time", time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local), "2006-01-02 15:04:05")
	buffer = appender.AppendError(buffer, "error", errors.New("ops"))
	buffer = appender.AppendStringer(buffer, "stringer", net.IPv4bcast)
	buffer = appender.AppendBools(buffer, "bools", []bool{true, false})
	buffer = appender.AppendBytes(buffer, "bytes", []byte{'b', 'y', 't', 'e'})
	buffer = appender.AppendRunes(buffer, "runes", []rune{'r', 'u', 'n', 'e'})
	buffer = appender.AppendInts(buffer, "ints", []int{1, 2, 3})
	buffer = appender.AppendInt8s(buffer, "int8s", []int8{8, 88, 123})
	buffer = appender.AppendInt16s(buffer, "int16s", []int16{16, 1616, 1234})
	buffer = appender.AppendInt32s(buffer, "int32s", []int32{32, 323232, 12345})
	buffer = appender.AppendInt64s(buffer, "int64s", []int64{64, 64646464, 123456})
	buffer = appender.AppendUints(buffer, "uints", []uint{1, 2, 3})
	buffer = appender.AppendUint8s(buffer, "uint8s", []uint8{8, 88, 123})
	buffer = appender.AppendUint16s(buffer, "uint16s", []uint16{16, 1616, 1234})
	buffer = appender.AppendUint32s(buffer, "uint32s", []uint32{32, 323232, 12345})
	buffer = appender.AppendUint64s(buffer, "uint64s", []uint64{64, 64646464, 123456})
	buffer = appender.AppendFloat32s(buffer, "float32s", []float32{32.32, 3232.3232, 323232.323232})
	buffer = appender.AppendFloat64s(buffer, "float64s", []float64{64.64, 64.64, 64.64})
	buffer = appender.AppendStrings(buffer, "strings", []string{"value1", "value2", "value3"})
	buffer = appender.AppendTimes(buffer, "times", []time.Time{time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local), time.Date(2021, 10, 27, 0, 38, 0, 0, time.Local)}, "2006-01-02 15:04:05")
	buffer = appender.AppendErrors(buffer, "errors", []error{errors.New("ops1"), errors.New("ops2"), errors.New("ops3")})
	buffer = appender.AppendStringers(buffer, "stringers", []fmt.Stringer{net.IPv4zero, net.IPv4bcast})
	buffer = appender.End(buffer)

	result := string(buffer)
	want := `{"any":{"AppendAnyKey1":"value","AppendAnyKey2":123,"AppendAnyKey3":666.666},"bool":true,"byte":"b","rune":"r","int":-1,"int8":-8,"int16":-16,"int32":-32,"int64":-64,"uint":1,"uint8":8,"uint16":16,"uint32":32,"uint64":64,"float32":32.31999969482422,"float64":64.64,"string":"value","` +
		`time":"2006-01-02 15:04:05","error":"ops","stringer":"255.255.255.255","bools":[true,false],"bytes":["b","y","t","e"],"runes":["r","u","n","e"],"ints":[1,2,3],"int8s":[8,88,123],"int16s":[16,1616,1234],"int32s":[32,323232,12345],"int64s":[64,64646464,123456],"uints":[1,2,3],"uint8s":[8,88,123],"uint16s":[16,161` +
		`6,1234],"uint32s":[32,323232,12345],"uint64s":[64,64646464,123456],"float32s":[32.31999969482422,3232.3232421875,323232.3125],"float64s":[64.64,64.64,64.64],"strings":["value1","value2","value3"],"times":["2006-01-02 15:04:05","2021-10-27 00:38:00"],"errors":["ops1","ops2","ops3"],"stringers":["0.0.0.0","255.25` +
		`5.255.255"]}` + "\n"
	if result != want {
		t.Errorf("result %s != want %s", result, want)
	}
}
