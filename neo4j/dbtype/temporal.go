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

package dbtype

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/packstream"
	"time"
)

// Cypher DateTime corresponds to Go time.Time

type (
	Time          time.Time // Time since start of day with timezone information
	Date          time.Time // Date value, without a time zone and time related components.
	LocalTime     time.Time // Time since start of day in local timezone
	LocalDateTime time.Time // Date and time in local timezone
)

// SerializeNeo4j serializes this LocalDateTime into the Packer.
func (t LocalDateTime) SerializeNeo4j(p packstream.Packer) {
	tt := time.Time(t)
	_, offset := tt.Zone()
	secs := tt.Unix() + int64(offset)
	p.StructHeader('d', 2)
	p.Int64(secs)
	p.Int(tt.Nanosecond())
}

// Time casts LocalDateTime to time.Time
//
// Note that the resulting time.Time will have its location set to time.Local.
// From the DBMS's perspective, however, a LocalDateTime is considered to not have any timezone information.
func (t LocalDateTime) Time() time.Time {
	return time.Time(t)
}

// String returns the string representation of this LocalDateTime in ISO-8601 compliant form:
// `YYYY-MM-DDThh:mm:ss.nnnnnnnnn`.
func (t LocalDateTime) String() string {
	return t.Time().Format("2006-01-02T15:04:05.999999999")
}

// SerializeNeo4j serializes this LocalTime into the Packer.
func (t LocalTime) SerializeNeo4j(p packstream.Packer) {
	tt := time.Time(t)
	nanos := int64(time.Hour)*int64(tt.Hour()) +
		int64(time.Minute)*int64(tt.Minute()) +
		int64(time.Second)*int64(tt.Second()) +
		int64(tt.Nanosecond())
	p.StructHeader('t', 1)
	p.Int64(nanos)
}

// Time casts LocalTime to time.Time
//
// Note that the resulting time.Time will have its location set to time.Local.
// From the DBMS's perspective, however, a LocalTime is considered to not have any timezone information.
func (t LocalTime) Time() time.Time {
	return time.Time(t)
}

// String returns the string representation of this LocalTime in ISO-8601 compliant form:
// `hh:mm:ss.nnnnnnnnn`.
func (t LocalTime) String() string {
	return t.Time().Format("15:04:05.999999999")
}

// SerializeNeo4j serializes this Date into the Packer.
func (t Date) SerializeNeo4j(p packstream.Packer) {
	tt := time.Time(t)
	secs := tt.Unix()
	_, offset := tt.Zone()
	secs += int64(offset)
	days := secs / (60 * 60 * 24)
	p.StructHeader('D', 1)
	p.Int64(days)
}

// Time casts Date to time.Time
func (t Date) Time() time.Time {
	return time.Time(t)
}

// String returns the string representation of this Date in ISO-8601 compliant form:
// `YYYY-MM-DD`.
func (t Date) String() string {
	return t.Time().Format("2006-01-02")
}

// SerializeNeo4j serializes this Time into the Packer.
func (t Time) SerializeNeo4j(p packstream.Packer) {
	tt := time.Time(t)
	_, tzOffsetSecs := tt.Zone()
	d := tt.Sub(
		time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, tt.Location()))
	p.StructHeader('T', 2)
	p.Int64(d.Nanoseconds())
	p.Int(tzOffsetSecs)
}

// Time casts Time to time.Time
func (t Time) Time() time.Time {
	return time.Time(t)
}

// String returns the string representation of this Time in ISO-8601 compliant form:
// `hh:mm:ss.nnnnnnnnn±Z/hh:mm`.
func (t Time) String() string {
	return t.Time().Format("15:04:05.999999999Z07:00")
}

// Duration represents temporal amount containing months, days, seconds and nanoseconds.
// Supports longer durations than time.Duration
type Duration struct {
	Months  int64
	Days    int64
	Seconds int64
	Nanos   int
}

// SerializeNeo4j serializes this Duration into the Packer.
func (d Duration) SerializeNeo4j(p packstream.Packer) {
	p.StructHeader('E', 4)
	p.Int64(d.Months)
	p.Int64(d.Days)
	p.Int64(d.Seconds)
	p.Int(d.Nanos)
}

// String returns the string representation of this Duration in ISO-8601 compliant form.
func (d Duration) String() string {
	sign := ""
	if d.Seconds < 0 && d.Nanos > 0 {
		d.Seconds++
		d.Nanos = int(time.Second) - d.Nanos

		if d.Seconds == 0 {
			sign = "-"
		}
	}

	timePart := ""
	if d.Nanos == 0 {
		timePart = fmt.Sprintf("%s%d", sign, d.Seconds)
	} else {
		timePart = fmt.Sprintf("%s%d.%09d", sign, d.Seconds, d.Nanos)
	}

	return fmt.Sprintf("P%dM%dDT%sS", d.Months, d.Days, timePart)
}

func (d1 Duration) Equal(d2 Duration) bool {
	return d1.Months == d2.Months && d1.Days == d2.Days && d1.Seconds == d2.Seconds && d1.Nanos == d2.Nanos
}
