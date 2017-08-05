package main

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_parseWSPHeaders(t *testing.T) {
	bytes, _ := hex.DecodeString("01100d850c048083ff7f048183ff7f028314806170706c69636174696f6e2f782d7761702d70726f762e62726f777365722d73657474696e677300806170706c69636174696f6e2f782d6e6f6b69612e73657474696e6773008094809580878086809880ae806170706c69636174696f6e2f782d7761702d70726f762e62726f777365722d626f6f6b6d61726b730080746578742f782d636f2d6465736300809d809e80696d6167652f6a70670080696d6167652f626d700080a080a180696d6167652f766e642e6e6f6b2d77616c6c70617065720080696d6167652f766e642d6e6f6b2d63616d6572612d736e61700080696d6167652f766e642d6e6f6b2d63616d6572612d736e737000806170706c69636174696f6e2f766e642e7761702e6d6d732d6d6573736167650080746578742f766e642e73756e2e6a326d652e6170702d64657363726970746f7200806170706c69636174696f6e2f766e642e6e6f6b69612e72696e67696e672d746f6e650080617564696f2f6d6964690080617564696f2f6d69640080617564696f2f782d6d6964690080617564696f2f782d6d69640080617564696f2f73702d6d69646900806170706c69636174696f6e2f6a6176610080696d6167652f766e642e6e6f6b2d6f706c6f676f2d636f6c6f7200806170706c69636174696f6e2f6a6176612d6172636869766500806170706c69636174696f6e2f782d6a6176612d617263686976650083d281848102ea5181040203e83db5687474703a2f2f6e64732e6e6f6b69612e636f6d2f756170726f662f4e3335313069723130302e786d6c00a94e6f6b696133353130692f312e30202830352e3330292050726f66696c652f4d4944502d312e3020436f6e66696775726174696f6e2f434c44432d312e3000582d5741502e544f44000100")
	type args struct {
		in []byte
	}
	tests := []struct {
		name  string
		args  args
		want  wirelessSessionProtocolHeaders
		want1 int
	}{
		{
			name: "test input",
			args: args{
				in: bytes,
			},
			want: wirelessSessionProtocolHeaders{
				PDUType:            1,
				MajorVersion:       1,
				MinorVersion:       0,
				CapabilitiesLength: 13,
				HeadersLength:      652,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := parseWSPHeaders(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseWSPHeaders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSubPDULength(t *testing.T) {
	bytes, _ := hex.DecodeString("82a20a")
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "test parse",
			args: args{
				in: bytes,
			},
			want: 674,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSubPDULength(tt.args.in)
			if got != tt.want {
				t.Errorf("getSubPDULength() got = %v, want %v", got, tt.want)
			}
		})
	}
}
