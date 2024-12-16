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
)

// Point2D represents a two dimensional point in a particular coordinate reference system.
type Point2D struct {
	X            float64
	Y            float64
	SpatialRefId uint32 // Id of coordinate reference system.
}

// Point3D represents a three dimensional point in a particular coordinate reference system.
type Point3D struct {
	X            float64
	Y            float64
	Z            float64
	SpatialRefId uint32 // Id of coordinate reference system.
}

// SerializeNeo4j serializes this point into the Packer.
func (p Point2D) SerializeNeo4j(packer packstream.Packer) {
	packer.StructHeader('X', 3)
	packer.Uint32(p.SpatialRefId)
	packer.Float64(p.X)
	packer.Float64(p.Y)
}

// SerializeNeo4j serializes this point into the Packer.
func (p Point3D) SerializeNeo4j(packer packstream.Packer) {
	packer.StructHeader('Y', 4)
	packer.Uint32(p.SpatialRefId)
	packer.Float64(p.X)
	packer.Float64(p.Y)
	packer.Float64(p.Z)
}

// String returns string representation of this point.
func (p Point2D) String() string {
	return fmt.Sprintf("Point{SpatialRefId=%d, X=%f, Y=%f}", p.SpatialRefId, p.X, p.Y)
}

// String returns string representation of this point.
func (p Point3D) String() string {
	return fmt.Sprintf("Point{SpatialRefId=%d, X=%f, Y=%f, Z=%f}", p.SpatialRefId, p.X, p.Y, p.Z)
}
