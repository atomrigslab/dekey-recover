// Copyright 2021 Atomrigs Lab Inc.
// This library is part of Dekey, blockchain wallet software based on secure multi-party computation
//
// The library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The library is also distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this library. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"crypto/ecdsa"
	"math/big"
)

// Point -
type Point struct {
	X *big.Int
	Y *big.Int
}

// ToECDSA provides ecdsa.PublicKey using x, y coordinates
func (a *Point) ToECDSA() *ecdsa.PublicKey {

	if a.X == nil || a.Y == nil {
		return nil
	}
	return &ecdsa.PublicKey{
		Curve: BaseCurve,
		X:     a.X,
		Y:     a.Y,
	}
}

// NewPoint -
func NewPoint(x, y *big.Int) *Point {
	return &Point{X: x, Y: y}
}
