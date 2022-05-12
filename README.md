# zilkeystore

Very simple CLI utility to print the keys, public and private, associated to a zilliqa keystore account.


usage: ``zilkeystore <keystore.json file>``

## Example

```bash
$ zilkeystore zilliqa_wallet_unstoppabledomains_myself.json
Enter Passphrase (12 words separated by space):
Zilliqa Account "808A188886DF2A868153749D487E3A7B0863E3BB"
 public wallet address (hex)           : 808a188886df2a868153749d487e3a7b0863e3bb
 public wallet address (human readable): zil1sz9p3zyxmu4gdq2nwjw5sl360vyx8camrulx6j
 public wallet key (hex)               : 031d92ee2513e96880904634d1f0c995909cd85f69fb5a521bd86e159daaf47abe
WARNING! do not share the following private key. keep it in a safe place, like your passphrase. 
 private wallet key (hex)              : {---------your-private-key----------------------------------------}
```

Notice: the keying passphrase does not appear on the screen.

## Install

Requires `go`.

```bash
$ go install github.com/sunraylab/zilkeystore
```
