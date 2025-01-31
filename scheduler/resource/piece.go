/*
 *     Copyright 2022 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resource

import (
	"time"

	commonv2 "d7y.io/api/pkg/apis/common/v2"
)

// IsPieceBackToSource returns whether the piece is downloaded back-to-source.
func IsPieceBackToSource(parentID string) bool {
	return parentID == ""
}

// Piece represents information of piece.
type Piece struct {
	// Piece number.
	Number uint32
	// Parent peer id.
	ParentID string
	// Piece offset.
	Offset uint64
	// Piece size.
	Size uint64
	// Digest of the piece data, for example md5:xxx or sha256:yyy.
	Digest string
	// Traffic type.
	TrafficType commonv2.TrafficType
	// Downloading piece costs time.
	Cost time.Duration
	// Piece create time.
	CreatedAt time.Time
}
