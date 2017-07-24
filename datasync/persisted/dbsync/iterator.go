// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbsync

import (
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/syncbase"
	"github.com/ligato/cn-infra/db/keyval"
)

// NewIterator is a constructor
func NewIterator(delegate keyval.BytesKeyValIterator) *Iterator {
	return &Iterator{delegate: delegate}
}

// Iterator adapts the db_proto.KeyValIterator to the datasync.KeyValIterator
type Iterator struct {
	delegate keyval.BytesKeyValIterator
}

// GetNext just delegate GetNext
func (it *Iterator) GetNext() (kv datasync.KeyVal, stop bool) {
	kvbytes, stop := it.delegate.GetNext()
	if stop {
		return nil, stop
	}
	return syncbase.NewKeyValBytes(kvbytes.GetKey(), kvbytes.GetValue(), kvbytes.GetRevision()), stop
}
