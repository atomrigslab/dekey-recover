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

package main

import (
	"dekey-recover/utils"
	"flag"
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 5 {
		fmt.Println("Usage: dekey-recover -n <number-of-recovered-accounts> -m <normal | legacy>")
		return
	}

	nPt := flag.Int("n", 0, "number of accounts to be recovered")
	mPt := flag.String("m", "normal", "recovery mode, 'normal' or 'legacy'")
	flag.Parse()
	nOfAccounts := *nPt
	mode := *mPt

	if nOfAccounts <= 0 {
		fmt.Println("Usage: dekey-recover -n <number-of-recovered-accounts> -m <normal | legacy>")
		fmt.Println("<number-of-recovered-accounts> should be more than zero")
		return
	}

	if mode == "legacy" {
		err := utils.EmergentRecoveryLegacy(nOfAccounts)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if mode == "normal" {
		err := utils.EmergentRecovery(nOfAccounts)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Usage: dekey-recover -n <number-of-recovered-accounts> -m <normal | legacy>")
		fmt.Println("mode should be either 'normal' or 'legacy'.")
	}
}
