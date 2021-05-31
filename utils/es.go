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
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// EmergencySeed is a server share info encrypted by uc's pubkey (ECIES)
//   Legacy version
type EmergencySeed struct {
	Mnemonic   string // mnemonic that derives Xi
	EthAddress string // Ethereum Address
	Sid        string // UC's Address
}

// EncodeES -
func (es *EmergencySeed) EncodeES(buf *bytes.Buffer) error {

	enc := json.NewEncoder(buf)
	err := enc.Encode(es)
	if err != nil {
		log.Println("EncodeES error:", err)
	}
	return err
}

// DecodeES -
func (es *EmergencySeed) DecodeES(buf *bytes.Buffer) error {

	dec := json.NewDecoder(buf)

	if err := dec.Decode(es); err != nil {
		return fmt.Errorf("DecodeES error: %s", err.Error())
	}
	return nil
}

// DecryptEciesES -
func DecryptEciesES(privKey *ecdsa.PrivateKey, hash, s1, s2 []byte) (*EmergencySeed, error) {

	privKey.Curve = ethcrypto.S256()
	plainBytes, err := ecies.ImportECDSA(privKey).Decrypt(hash, s1, s2)
	if err != nil {
		return nil, err
	}

	es := new(EmergencySeed)
	if err := es.DecodeES(bytes.NewBuffer(plainBytes)); err != nil {
		return nil, err
	}

	ethaddr := ethcrypto.PubkeyToAddress(privKey.PublicKey).Hex()

	if ethaddr != es.Sid {
		return nil, errors.New("decrypted sid is diff")
	}

	return es, nil
}
