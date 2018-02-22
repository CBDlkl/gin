// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

type Protobuf struct {
	Data interface{}
}

var pbContentType = []string{"application/x-protobuf"}

func (r Protobuf) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, pbContentType)
}

func (r Protobuf) Render(w http.ResponseWriter) error {
	writeContentType(w, pbContentType)
	encoded, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}
	if _, err = w.Write(encoded); err != nil {
		return err
	}
	return nil
}
