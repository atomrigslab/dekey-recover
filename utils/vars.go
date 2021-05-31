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

import btcec "github.com/btcsuite/btcd/btcec"

const (
	// Bip32SeedGenPassword -
	Bip32SeedGenPassword = "Atomrigs_Lab_BIP32_Password" // non 25th word setting"Atomrigs_Lab_BIP32_Password"
	// ECIESSharedInfo1 is fed into key derivation in ECIES
	ECIESSharedInfo1 = "ATOMRIGS_LAB_DKEYS001"
	// ECIESSharedInfo2 is fed into the MAC in ECIES
	ECIESSharedInfo2 = "ATOMRIGS_LAB_DKEYS002"
)

var (
	// Bitcoin/Ethereum curve (secp256k1) 2^256 - 2^32 - 2^9 - 2^8 - 2^7 - 2^6 - 2^4 - 1
	BaseCurve = btcec.S256()
	// GroupOrder is the  group order of an elliptic curve group
	GroupOrder = BaseCurve.Params().N
)
