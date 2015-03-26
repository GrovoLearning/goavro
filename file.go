// Copyright 2015 LinkedIn Corp. Licensed under the Apache License,
// Version 2.0 (the "License"); you may not use this file except in
// compliance with the License.  You may obtain a copy of the License
// at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.Copyright [201X] LinkedIn Corp. Licensed under the Apache
// License, Version 2.0 (the "License"); you may not use this file
// except in compliance with the License.  You may obtain a copy of
// the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.

package goavro

const (
	magicBytes     = "Obj\x01"
	syncLength     = 16
	metadataSchema = `{"type":"map","values":"bytes"}`
)

const (
	CompressionNull    = "null"
	CompressionDeflate = "deflate"
	CompressionSnappy  = "snappy"
)

var (
	metadataCodec Codec
)

func init() {
	metadataCodec, _ = NewCodec(metadataSchema)
}

// IsCompressionCodecSupported returns true if and only if the specified codec
// string is supported by this library.
func IsCompressionCodecSupported(someCodec string) bool {
	switch someCodec {
	case CompressionNull, CompressionDeflate, CompressionSnappy:
		return true
	default:
		return false
	}
}

// // FileSetter functions are those those which are used to instantiate
// // a new FileReader or FileWriter.
// type FileSetter func(*File) error