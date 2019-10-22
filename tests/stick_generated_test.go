// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestStickNormal0(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("stick on src table pop if !localhost")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestStickNormal1(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("stick match src table pop if !localhost")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestStickNormal2(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("stick store-request src table pop if !localhost")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestStickFail0(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("stick")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestStickFail1(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestStickFail2(t *testing.T) {
	parser := &parsers.Stick{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
