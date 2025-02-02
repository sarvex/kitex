// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package thriftgo

const file = `
{{define "file"}}
// Code generated by Kitex {{Version}}. DO NOT EDIT.

package {{.PkgName}}

import (
	"fmt"
	"bytes"
	"strings"
	"reflect"

	"github.com/apache/thrift/lib/go/thrift"
	{{if GenerateDeepCopyAPIs -}}
	kutils "github.com/cloudwego/kitex/pkg/utils"
	{{- end}}
	{{if GenerateFastAPIs}}
	"{{ImportPathTo "pkg/protocol/bthrift"}}"
	{{- end}}
	{{- range $path, $alias := .Imports}}
	{{$alias }}"{{$path}}"
	{{- end}}
)

{{InsertionPoint "KitexUnusedProtection"}}

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	{{- if GenerateFastAPIs}}
	_ = bthrift.BinaryWriter(nil)
	{{- end}}
	{{- range .Imports | ToPackageNames}}
	_ = {{.}}.KitexUnusedProtection
	{{- end}}
)

{{template "body" .}}

{{end}}{{/* define "file" */}}
`

const body = `
{{define "body"}}

{{- range .Scope.StructLikes}}
{{template "StructLikeCodec" .}}
{{- end}}

{{- range .Scope.Services}}
{{template "Processor" .}}
{{- end}}

{{template "ArgsAndResult" .}}

{{template "ExtraTemplates" .}}
{{- end}}{{/* define "body" */}}
`

const patchArgsAndResult = `
{{define "ArgsAndResult"}}
{{range $svc := .Scope.Services}}
{{range .Functions}}

{{$argType := .ArgType}}
{{$resType := .ResType}}

{{- if GenerateArgsResultTypes}}
{{template "StructLike" $argType}}
{{- end}}{{/* if GenerateArgsResultTypes */}}
func (p *{{$argType.GoName}}) GetFirstArgument() interface{} {
	return {{if .Arguments}}p.{{(index $argType.Fields 0).GoName}}{{else}}nil{{end}}
}

{{if not .Oneway}}
{{- if GenerateArgsResultTypes}}
{{template "StructLike" $resType}}
{{- end}}{{/* if GenerateArgsResultTypes */}}
func (p *{{$resType.GoName}}) GetResult() interface{} {
	return {{if .Void}}nil{{else}}p.Success{{end}}
}
{{- end}}{{/* if not .Oneway */}}
{{- end}}{{/* range Functions */}}
{{- end}}{{/* range .Scope.Service */}}
{{- end}}{{/* define "FirstResult" */}}
`

var basicTemplates = []string{
	patchArgsAndResult,
	file,
	body,
}
