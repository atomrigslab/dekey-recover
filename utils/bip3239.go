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
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/tyler-smith/go-bip39"
)

// RecoverPrivKeyFromMnemonic -
func RecoverPrivKeyFromMnemonic(mnemonic, password string) (*ecdsa.PrivateKey, error) {

	extKey, err := RecoverExtendedKeyFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	sk, err := extKey.ECPrivKey()
	if err != nil {
		return nil, err
	}
	return sk.ToECDSA(), nil
}

// RecoverExtendedKeyFromMnemonic -
func RecoverExtendedKeyFromMnemonic(mnemonic, password string) (*hdkeychain.ExtendedKey, error) {

	if mnemonic == "" {
		return nil, errors.New("mnemonic is required")
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("mnemonic is invalid")
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		return nil, err
	}

	return hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
}

// RecoverXiFromMnemonic -
func RecoverXiFromMnemonic(mnemonic, password string) (*big.Int, error) {

	sk, err := RecoverPrivKeyFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return sk.D, nil
}

// DerivePrivKey is used by uc  (extended privkey --> hardened/non-hardened child privkey )
func DerivePrivKey(ekey *hdkeychain.ExtendedKey, pathStr string) (*big.Int, error) {

	// // check pathStr is valid (and also all are non-hardened)
	// if strings.Contains(pathStr, "'") {
	// 	return nil, errors.New("we do not support hardened derivation yet")
	// }
	path, err := ParseDerivationPath(pathStr)
	if err != nil {
		return nil, err
	}
	// extkey is a privkey extkey
	if !ekey.IsPrivate() {
		return nil, errors.New("extkey is not a private key")
	}
	key := ekey
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return privKey.D, nil
}

// ParseDerivationPath -
func ParseDerivationPath(path string) (accounts.DerivationPath, error) {
	return accounts.ParseDerivationPath(path)
}

// GetDerivationPath -
func GetDerivationPath(id int) string {
	return fmt.Sprintf("m/44'/60'/%d'/0/0", id)
}
