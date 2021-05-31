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
	"encoding/hex"
	"math/big"
	"strings"
)

// Hex0xToBytes -
func Hex0xToBytes(src string) ([]byte, error) {

	if src == "" {
		return []byte{}, nil
	}
	if strings.HasPrefix(src, "0x") {
		src = src[2:]
	}
	return HexStrToBytes(src)
}

// HexStrToBytes -
func HexStrToBytes(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

// IntToHex0x -
func IntToHex0x(src *big.Int) string {
	if src == nil {
		return ""
	}
	return BytesToHex0x(src.Bytes())
}

// BytesToHex0x -
func BytesToHex0x(src []byte) string {
	if src == nil {
		return ""
	}
	hStr := BytesToHexStr(src)
	if hStr == "" || strings.HasPrefix(hStr, "0x") {
		return hStr
	}
	return "0x" + hStr
}

// BytesToHexStr -
func BytesToHexStr(src []byte) string {
	return hex.EncodeToString(src)
}
