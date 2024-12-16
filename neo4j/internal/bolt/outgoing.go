/*
 * Copyright (c) "Neo4j"
 * Neo4j Sweden AB [https://neo4j.com]
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bolt

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	idb "github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/packstream"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
	pubpackstream "github.com/neo4j/neo4j-go-driver/v5/neo4j/packstream"
	"io"
	"reflect"
	"time"
)

type outgoing struct {
	chunker    chunker
	packer     packstream.Packer
	onPackErr  func(error)
	onIoErr    func(context.Context, error)
	boltLogger log.BoltLogger
	logId      string
}

func (o *outgoing) begin() {
	o.chunker.beginMessage()
	o.packer.Begin(o.chunker.buf)
}

func (o *outgoing) end() {
	buf, err := o.packer.End()
	o.chunker.buf = buf
	o.chunker.endMessage()
	if err != nil {
		o.onPackErr(err)
	}
}

func (o *outgoing) appendHello(hello map[string]any) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "HELLO %s", loggableDictionary(hello))
	}
	o.begin()
	o.packer.StructHeader(msgHello, 1)
	o.packMap(hello)
	o.end()
}

func (o *outgoing) appendLogoff() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "LOGOFF")
	}
	o.begin()
	o.packer.StructHeader(msgLogoff, 0)
	o.end()
}

func (o *outgoing) appendLogon(logon map[string]any) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "LOGON %s", loggableDictionary(logon))
	}
	o.begin()
	o.packer.StructHeader(msgLogon, 1)
	o.packMap(logon)
	o.end()
}

func (o *outgoing) appendBegin(meta map[string]any) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "BEGIN %s", loggableDictionary(meta))
	}
	o.begin()
	o.packer.StructHeader(msgBegin, 1)
	o.packMap(meta)
	o.end()
}

func (o *outgoing) appendCommit() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "COMMIT")
	}
	o.begin()
	o.packer.StructHeader(msgCommit, 0)
	o.end()
}

func (o *outgoing) appendRollback() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "ROLLBACK")
	}
	o.begin()
	o.packer.StructHeader(msgRollback, 0)
	o.end()
}

func (o *outgoing) appendRun(cypher string, params, meta map[string]any) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "RUN %q %s %s", cypher, loggableDictionary(params), loggableDictionary(meta))
	}
	o.begin()
	o.packer.StructHeader(msgRun, 3)
	o.packer.String(cypher)
	o.packMap(params)
	o.packMap(meta)
	o.end()
}

func (o *outgoing) appendPullN(n int) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "PULL %s", loggableDictionary{"n": n})
	}
	o.begin()
	o.packer.StructHeader(msgPullN, 1)
	o.packer.MapHeader(1)
	o.packer.String("n")
	o.packer.Int(n)
	o.end()
}

func (o *outgoing) appendPullNQid(n int, qid int64) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "PULL %s", loggableDictionary{"n": n, "qid": qid})
	}
	o.begin()
	o.packer.StructHeader(msgPullN, 1)
	o.packer.MapHeader(2)
	o.packer.String("n")
	o.packer.Int(n)
	o.packer.String("qid")
	o.packer.Int64(qid)
	o.end()
}

func (o *outgoing) appendDiscardN(n int) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "DISCARD %s", loggableDictionary{"n": n})
	}
	o.begin()
	o.packer.StructHeader(msgDiscardN, 1)
	o.packer.MapHeader(1)
	o.packer.String("n")
	o.packer.Int(n)
	o.end()
}

func (o *outgoing) appendDiscardNQid(n int, qid int64) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "DISCARD %s", loggableDictionary{"n": n, "qid": qid})
	}
	o.begin()
	o.packer.StructHeader(msgDiscardN, 1)
	o.packer.MapHeader(2)
	o.packer.String("n")
	o.packer.Int(n)
	o.packer.String("qid")
	o.packer.Int64(qid)
	o.end()
}

func (o *outgoing) appendPullAll() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "PULL ALL")
	}
	o.begin()
	o.packer.StructHeader(msgPullAll, 0)
	o.end()
}

// Only valid for V4.3
func (o *outgoing) appendRouteToV43(context map[string]string, bookmarks []string, database string) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "ROUTE %s %s %q", loggableStringDictionary(context), loggableStringList(bookmarks), database)
	}
	o.begin()
	o.packer.StructHeader(msgRoute, 3)
	o.packer.MapHeader(len(context))
	for k, v := range context {
		o.packer.String(k)
		o.packer.String(v)
	}
	o.packer.ArrayHeader(len(bookmarks))
	for _, bookmark := range bookmarks {
		o.packer.String(bookmark)
	}
	if database == idb.DefaultDatabase {
		o.packer.Nil()
	} else {
		o.packer.String(database)
	}
	o.end()
}

func (o *outgoing) appendRoute(context map[string]string, bookmarks []string, what map[string]any) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "ROUTE %s %s %s", loggableStringDictionary(context), loggableStringList(bookmarks), loggableDictionary(what))
	}
	o.begin()
	o.packer.StructHeader(msgRoute, 3)
	o.packer.MapHeader(len(context))
	for k, v := range context {
		o.packer.String(k)
		o.packer.String(v)
	}
	o.packer.ArrayHeader(len(bookmarks))
	for _, bookmark := range bookmarks {
		o.packer.String(bookmark)
	}
	o.packMap(what)
	o.end()
}

func (o *outgoing) appendReset() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "RESET")
	}
	o.begin()
	o.packer.StructHeader(msgReset, 0)
	o.end()
}

func (o *outgoing) appendGoodbye() {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "GOODBYE")
	}
	o.begin()
	o.packer.StructHeader(msgGoodbye, 0)
	o.end()
}

func (o *outgoing) appendTelemetry(api int) {
	if o.boltLogger != nil {
		o.boltLogger.LogClientMessage(o.logId, "TELEMETRY %d", api)
	}
	o.begin()
	o.packer.StructHeader(msgTelemetry, 1)
	o.packer.Int(api)
	o.end()
}

// For tests
func (o *outgoing) appendX(tag byte, fields ...any) {
	o.begin()
	o.packer.StructHeader(tag, len(fields))
	for _, f := range fields {
		o.packX(f)
	}
	o.end()
}

func (o *outgoing) send(ctx context.Context, wr io.Writer) {
	err := o.chunker.send(ctx, wr)
	if err != nil {
		o.onIoErr(ctx, err)
	}
}

func (o *outgoing) packMap(m map[string]any) {
	o.packer.MapHeader(len(m))
	for k, v := range m {
		o.packer.String(k)
		o.packX(v)
	}
}

func (o *outgoing) packStruct(x any) {
	switch v := x.(type) {
	case time.Time:
		o.packer.Time(v)
	default:
		o.onPackErr(&db.UnsupportedTypeError{Type: reflect.TypeOf(x)})
	}
}

func (o *outgoing) packX(x any) {
	if x == nil {
		o.packer.Nil()
		return
	}

	if v, ok := x.(interface{ SerializeNeo4j(pubpackstream.Packer) }); ok {
		v.SerializeNeo4j(&o.packer)
		return
	}

	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Bool:
		o.packer.Bool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		o.packer.Int64(v.Int())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		o.packer.Uint32(uint32(v.Uint()))
	case reflect.Uint64, reflect.Uint:
		o.packer.Uint64(v.Uint())
	case reflect.Float32, reflect.Float64:
		o.packer.Float64(v.Float())
	case reflect.String:
		o.packer.String(v.String())
	case reflect.Ptr:
		if v.IsNil() {
			o.packer.Nil()
			return
		}
		o.packX(v.Elem().Interface())
	case reflect.Struct:
		o.packStruct(x)
	case reflect.Slice:
		// Optimizations
		switch s := x.(type) {
		case []byte:
			o.packer.Bytes(s) // Not just optimization
		case []int:
			o.packer.Ints(s)
		case []int64:
			o.packer.Int64s(s)
		case []string:
			o.packer.Strings(s)
		case []float64:
			o.packer.Float64s(s)
		case []any:
			o.packer.ArrayHeader(len(s))
			for _, e := range s {
				o.packX(e)
			}
		default:
			num := v.Len()
			o.packer.ArrayHeader(num)
			for i := 0; i < num; i++ {
				o.packX(v.Index(i).Interface())
			}
		}
	case reflect.Map:
		// Optimizations
		switch m := x.(type) {
		case map[string]int:
			o.packer.IntMap(m)
		case map[string]string:
			o.packer.StringMap(m)
		case map[string]any:
			o.packer.MapHeader(len(m))
			for k, v := range m {
				o.packer.String(k)
				o.packX(v)
			}
		default:
			t := reflect.TypeOf(x)
			if t.Key().Kind() != reflect.String {
				o.onPackErr(&db.UnsupportedTypeError{Type: reflect.TypeOf(x)})
				return
			}
			o.packer.MapHeader(v.Len())
			r := v.MapRange()
			for r.Next() {
				o.packer.String(r.Key().String())
				o.packX(r.Value().Interface())
			}
		}
	default:
		o.onPackErr(&db.UnsupportedTypeError{Type: reflect.TypeOf(x)})
	}
}
