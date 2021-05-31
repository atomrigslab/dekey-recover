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
	"bufio"
	"fmt"
	"math/big"
	"os"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

//////// test data
// clientMnemonic := "clerk wine book attract hover essence young vendor appear issue scissors burden"
// cipher, err := Hex0xToBytes("0x046f818a234b0129af2ff09ef526cba031c48db4373912658d3cd1c616257364ed73dd28cc862d242a85fc286b90e46fd901d03ccaca788c3b92899ee632a342f5402f1d8bc8e30c425975e274843e7e4873bdaad1c49e962eb57242cd984022309eec7d89106d6cd94ca801ee8cc679239395454bea32ffe635b8024c2454814648d9a9ac552109d2e0b8859c34aadf7f0bf9bae9cccd5cb88e3d4b7712e0b2f0105a76882f2b2150b29cfb088a7c61b10fd84220331ec8cddc6f5c69302a9d111e8ba13a4d6a0e2bfb7a4f1139b5fe6fa65fa99493a71641aafb7e5f3b79984f9de8b84a8e325ab7fd7cc5e043812180f941819373cf917e7b282afbe6691f32a222480f3b33c33b0ae061aa2afd16db8b90358340b8dcd0dc86adaa8e20e817557fa30be7a0aef85bd006261df3ebbd639c55c457603a")
//
// clientMnemonic := "avocado law replace baby dawn winter equal guess miracle report civil spoil"
// cipher, err := Hex0xToBytes("0x04db01c0fff5ed03e8df0f00c1b01b83f5966cc075d32e849d8009e48b9743921c88e11fe084f8044432e78ac156f9aaa32b7e74c510bac68e8504eaf1b7cd0558a6d30621d3a8da28843d93ca63eb49b78456716333122e1bcfe68a5e32f42cd5a2414f7d0a7733d7d892447c3c0c9674f0a045d410eb7b2003ee96ab9a1c589e02cd33f66d4ee4691d845532f50fd0ac79da088f05561b4d2bc2c79c238f91809ffc47f75d7cca3a71c2772370dfd2403885ae31eb25c7e2b5e3a234fccbb4bbe0d95fca5f386958678ff532ac6f55fb377fab62204bdeeee04552dc6dca8d6c88e34418f3f20925f116fccca7436888696aa0ac8d12a10543236fdb3ddf36bca935d51bf3d2493bc98e9e427d8f90edf59f2f71621e9c5fd60fd2d5be08e8ff98cdac40acfb8d30292ff2cb08b949ce7082bd936aa0c1f3a16d4b75f27ad003")

// EmergentRecoveryLegacy -
func EmergentRecoveryLegacy(nOfAccounts int) error {

	fmt.Printf("Enter client mnemonic: ")
	s1 := bufio.NewScanner(os.Stdin)
	s1.Scan()
	clientMnemonic := s1.Text()

	ucsk, err := RecoverPrivKeyFromMnemonic(clientMnemonic, Bip32SeedGenPassword)
	if err != nil {
		return err
	}

	fmt.Printf("Enter emergency seed cipher: ")
	s2 := bufio.NewScanner(os.Stdin)
	s2.Scan()
	cipher, err := Hex0xToBytes(s2.Text())
	if err != nil {
		return err
	}
	// decrypt a legacy cipher that was delivered by e-mail
	es, err := DecryptEciesES(ucsk, cipher,
		[]byte(ECIESSharedInfo1), []byte(ECIESSharedInfo2))
	if err != nil {
		return err
	}
	serverMnemonic := es.Mnemonic

	baseEthAddress, err := recoverKeys(clientMnemonic, serverMnemonic, nOfAccounts)
	if err != nil {
		fmt.Println("EmergentRecovery fails", err.Error())
		return err
	}
	if baseEthAddress != es.EthAddress {
		return fmt.Errorf("unmatched")
	}
	fmt.Println("matched")

	return err
}

// EmergentRecovery prints nOfAccounts of recovered private keys and returns base ethaddress and error
func EmergentRecovery(nOfAccounts int) error {

	fmt.Printf("Enter client mnemonic: ")
	s1 := bufio.NewScanner(os.Stdin)
	s1.Scan()
	clientMnemonic := s1.Text()

	fmt.Printf("Enter server mnemonic: ")
	s2 := bufio.NewScanner(os.Stdin)
	s2.Scan()
	serverMnemonic := s2.Text()

	_, err := recoverKeys(clientMnemonic, serverMnemonic, nOfAccounts)
	if err != nil {
		fmt.Println("EmergentRecovery fails", err.Error())
		return err
	}

	return nil
}

// recoverKeys prints nOfAccounts of recovered private keys and returns base ethaddress and error
func recoverKeys(clientMnemonic, serverMnemonic string, nOfAccounts int) (string, error) {

	x1, err := RecoverXiFromMnemonic(clientMnemonic, Bip32SeedGenPassword)
	if err != nil {
		return "", err
	}
	x2, err := RecoverXiFromMnemonic(serverMnemonic, Bip32SeedGenPassword)
	if err != nil {
		return "", err
	}

	// base account
	sk := new(big.Int).Mul(x1, x2)
	sk.Mod(sk, GroupOrder)
	pk := NewPoint(BaseCurve.ScalarBaseMult(sk.Bytes()))
	derivedEthAddress := ethcrypto.PubkeyToAddress(*pk.ToECDSA()).Hex()

	fmt.Printf("\n[Recovered private keys and associated eth addresses]\n")
	fmt.Printf("[%d] PrivateKey: %s   EthAddress: %s \n",
		0, IntToHex0x(sk), derivedEthAddress)

	if nOfAccounts > 1 {
		extPrivKey, err := RecoverExtendedKeyFromMnemonic(clientMnemonic, Bip32SeedGenPassword)
		if err != nil {
			return "", err
		}
		for aaID := 1; aaID < nOfAccounts; aaID++ {
			childx1, err := DerivePrivKey(extPrivKey, GetDerivationPath(aaID))
			if err != nil { // in case of 'err == hdkeychain.ErrInvalidChild'
				continue
			}
			childsk := new(big.Int).Mul(childx1, x2)
			childsk.Mod(childsk, GroupOrder)
			childpk := NewPoint(BaseCurve.ScalarBaseMult(childsk.Bytes()))
			fmt.Printf("[%d] PrivateKey: %s   EthAddress: %s \n",
				aaID, IntToHex0x(childsk), ethcrypto.PubkeyToAddress(*childpk.ToECDSA()).Hex())
		}
	}

	return derivedEthAddress, nil
}
