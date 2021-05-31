# Dekey self-recovery tool

## Introduction

When the Dekey service is not available, this tool helps our users recover their accounts (private keys) from backup materials.
There are two types of backup materials:

* Normal mode: Most users have this type of backups if they created accounts using the Dekey commercial edition. The backup includes a client share mnemonic (12 words) and a server share mnemonic (12 words)
* Legacy mode: Some long time users might have backup materials that are comprised of a client share mnemonic (12 words) and a hexstring-encoded ciphertext

This recovery tool supports these two types of backup materials.

## Get a program

See https://gitlab.com/atomrigsinc/dekey-recover/-/tree/master/binaries and download an appropriate executable file suitable for your operating system (Windows, MacOS, Linux)

### Build from sources

```
$ git clone https://gitlab.com/atomrigsinc/dekey-recover
$ cd dekey-recover/cmd
$ go build dekey-recover.go   (you need golang >= 1.15)
```

## How to recover

### Usage
$ dekey-recover -n [number-of-accounts] -m [normal | legacy]

Options (mandatory) 

-n:  the number of accounts to be recovered

-m:  recovery mode ('normal' or 'legacy')

### Test data

Client mnemonic: avocado law replace baby dawn winter equal guess miracle report civil spoil

Server backup (normal mode): million title unveil example awkward fault display throw noodle canvas fat solution

Server backup (legacy mode - old version of Dekey): 0x04db01c0fff5ed03e8df0f00c1b01b83f5966cc075d32e849d8009e48b9743921c88e11fe084f8044432e78ac156f9aaa32b7e74c510bac68e8504eaf1b7cd0558a6d30621d3a8da28843d93ca63eb49b78456716333122e1bcfe68a5e32f42cd5a2414f7d0a7733d7d892447c3c0c9674f0a045d410eb7b2003ee96ab9a1c589e02cd33f66d4ee4691d845532f50fd0ac79da088f05561b4d2bc2c79c238f91809ffc47f75d7cca3a71c2772370dfd2403885ae31eb25c7e2b5e3a234fccbb4bbe0d95fca5f386958678ff532ac6f55fb377fab62204bdeeee04552dc6dca8d6c88e34418f3f20925f116fccca7436888696aa0ac8d12a10543236fdb3ddf36bca935d51bf3d2493bc98e9e427d8f90edf59f2f71621e9c5fd60fd2d5be08e8ff98cdac40acfb8d30292ff2cb08b949ce7082bd936aa0c1f3a16d4b75f27ad003

### Windows example

```
c:\Temp>dekey-recover.exe -n 4 -m normal
Enter client mnemonic: avocado law replace baby dawn winter equal guess miracle report civil spoil
Enter server mnemonic: million title unveil example awkward fault display throw noodle canvas fat solution

[Recovered private keys and associated eth addresses]
[0] PrivateKey: 0x9738818fdd061ea7158e7c45a0db3178afbbe95a78f80345b6af540dd0c9e4e7   EthAddress: 0x78671D2DBe46C336E94dBde2DC6b68973Ba5D683
[1] PrivateKey: 0xb627a7b15e67fa8e3d9c5e033fda2557846cd266a4b1acbc424f43675181fd16   EthAddress: 0x58AF88873fEcF76FCf46A69dc3DF2912555b005B
[2] PrivateKey: 0x2863f2cdd58bd630d29988f6ff964a9c2c62c493d258d6bbe41440262188e87b   EthAddress: 0x45E024EF5b1d6EB82524D28935C7eF08B61D1A03
[3] PrivateKey: 0x17f1c0dd183feca2ddb86d113f80d7d378027d38d6ee1dd37ce2447ef5f4ba00   EthAddress: 0x3b5E8F142700fF3e53e97b231923B4794c41c6e3

c:\Temp>dekey-recover.exe -n 3 -m legacy
Enter client mnemonic: avocado law replace baby dawn winter equal guess miracle report civil spoil
Enter emergency seed cipher: 0x04db01c0fff5ed03e8df0f00c1b01b83f5966cc075d32e849d8009e48b9743921c88e11fe084f8044432e78ac156f9aaa32b7e74c510bac68e8504eaf1b7cd0558a6d30621d3a8da28843d93ca63eb49b78456716333122e1bcfe68a5e32f42cd5a2414f7d0a7733d7d892447c3c0c9674f0a045d410eb7b2003ee96ab9a1c589e02cd33f66d4ee4691d845532f50fd0ac79da088f05561b4d2bc2c79c238f91809ffc47f75d7cca3a71c2772370dfd2403885ae31eb25c7e2b5e3a234fccbb4bbe0d95fca5f386958678ff532ac6f55fb377fab62204bdeeee04552dc6dca8d6c88e34418f3f20925f116fccca7436888696aa0ac8d12a10543236fdb3ddf36bca935d51bf3d2493bc98e9e427d8f90edf59f2f71621e9c5fd60fd2d5be08e8ff98cdac40acfb8d30292ff2cb08b949ce7082bd936aa0c1f3a16d4b75f27ad003

[Recovered private keys and associated eth addresses]
[0] PrivateKey: 0x9738818fdd061ea7158e7c45a0db3178afbbe95a78f80345b6af540dd0c9e4e7   EthAddress: 0x78671D2DBe46C336E94dBde2DC6b68973Ba5D683
[1] PrivateKey: 0xb627a7b15e67fa8e3d9c5e033fda2557846cd266a4b1acbc424f43675181fd16   EthAddress: 0x58AF88873fEcF76FCf46A69dc3DF2912555b005B
[2] PrivateKey: 0x2863f2cdd58bd630d29988f6ff964a9c2c62c493d258d6bbe41440262188e87b   EthAddress: 0x45E024EF5b1d6EB82524D28935C7eF08B61D1A03
matched
```

## caveat

If the Dekey service is available, use 'client recovery' option in the extension.

Nevertheless if you recover your accounts by yourself, I recommend to use this tool in offline mode to protect the recovered private keys from any hacking attacks.

Unfortunatly, this library has not been verified by a third party.  There might be some bugs and
vulnerabilities. If you find one, please register an issue.