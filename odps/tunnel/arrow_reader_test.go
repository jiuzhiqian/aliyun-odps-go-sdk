// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tunnel

import (
	"bytes"
	"fmt"
	"github.com/jiuzhiqian/aliyun-odps-go-sdk/arrow/ipc"
	"github.com/jiuzhiqian/aliyun-odps-go-sdk/odps/datatype"
	"github.com/jiuzhiqian/aliyun-odps-go-sdk/odps/tableschema"
	"io"
	"testing"
)

var structTypeArrowData = []byte{
	0x78, 0x9c, 0x62, 0x60, 0x64, 0x60, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x13, 0x8c, 0x0c, 0x0c, 0x22,
	0x0c, 0x10, 0xc0, 0xc3, 0x20, 0xc6, 0x20, 0xc2, 0x20, 0xcc, 0xc0, 0xc3, 0xc0, 0xc2, 0xc0, 0xc3,
	0xc0, 0xc0, 0x10, 0x00, 0x15, 0x87, 0xca, 0x33, 0xb3, 0x30, 0x70, 0x31, 0x48, 0x30, 0xf0, 0x30,
	0x70, 0x30, 0x80, 0x58, 0x10, 0x71, 0x09, 0x46, 0x06, 0x90, 0x51, 0x70, 0x20, 0xc0, 0x80, 0x1b,
	0x70, 0xa0, 0xd1, 0x4c, 0x38, 0xf4, 0x08, 0xa0, 0xa9, 0x93, 0x40, 0x93, 0x97, 0x40, 0x93, 0x57,
	0x40, 0x93, 0xc7, 0xc5, 0x87, 0xa9, 0xd7, 0x40, 0x93, 0xd7, 0x40, 0x93, 0x37, 0x80, 0xd2, 0xdc,
	0x50, 0xda, 0x01, 0x4d, 0xbd, 0x03, 0x9a, 0x7a, 0x0f, 0x34, 0x3e, 0x08, 0xb0, 0x33, 0xa0, 0x86,
	0x0b, 0x03, 0x1d, 0xf8, 0x0c, 0xd0, 0x30, 0xcd, 0xb3, 0x84, 0xb0, 0x9d, 0x60, 0x82, 0xfc, 0x5f,
	0xfd, 0x6b, 0x19, 0x51, 0xf5, 0x80, 0xfc, 0x96, 0x9c, 0x93, 0x5f, 0x9a, 0x12, 0x5f, 0x9c, 0x99,
	0x93, 0x53, 0x09, 0xd7, 0x0f, 0xf2, 0x43, 0x5a, 0x7e, 0x7e, 0x49, 0x52, 0x62, 0x4e, 0x4e, 0xab,
	0x59, 0x97, 0x35, 0x20, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x58, 0x14, 0xc6,
}

type ReadClose struct {
	reader *bytes.Reader
}

func (r *ReadClose) Read(b []byte) (n int, err error) {
	return r.reader.Read(b)
}

func (r *ReadClose) Close() error {
	return nil
}

func TestRecordArrowReader_Read(t *testing.T) {
	br := &ReadClose{bytes.NewReader(structTypeArrowData)}
	reader := defaultDeflate().NewReader(br)

	c1 := tableschema.Column{
		Name: "name",
		Type: datatype.StringType,
	}

	c2 := tableschema.Column{
		Name: "score",
		Type: datatype.IntType,
	}

	structTypeStr := "struct<Address:array<string>, Hobby:string>"
	structType, _ := datatype.ParseDataType(structTypeStr)

	c3 := tableschema.Column{
		Name: "birthday",
		Type: datatype.DateTimeType,
	}

	c4 := tableschema.Column{
		Name: "extra",
		Type: structType,
	}

	sb := tableschema.NewSchemaBuilder()
	sb.Name("user_test").
		Columns(c1, c2, c3, c4).
		Lifecycle(2)

	schema := sb.Build()
	arrowSchema := schema.ToArrowSchema()

	sr := NewArrowStreamReader(reader)

	rar := RecordArrowReader{
		httpRes:           nil,
		recordBatchReader: ipc.NewRecordBatchReader(sr, arrowSchema),
		arrowReader:       sr,
	}

	record, err := rar.Read()

	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	for i, col := range record.Columns() {
		fmt.Printf("%s: %v\n", schema.Columns[i].Name, col)
	}

	record.Release()
}
