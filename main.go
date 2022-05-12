package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"

	zil "github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
)

func isJSON(str string) bool {
    var js json.RawMessage
    return json.Unmarshal([]byte(str), &js) == nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("missing parameter\nusage: zilkeystore <keystore.json file>.")
		os.Exit(1)
	}

	// read the keystore.json file
	fname := os.Args[1]
	keystore, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	if !isJSON(string(keystore)) {
		fmt.Println("keystore file content must be a json format. It's not")
		os.Exit(1)
	}

	// ask for the passphrase
	fmt.Print("Enter Passphrase (12 words separated by space): ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
		log.Fatal(err)
    }
    passphrase := string(bytePassword)

	// build the wallet from the keystore and the passphrase
	w := zil.NewWallet()
	w.AddByKeyStore(string(keystore), passphrase)
	if len(w.Accounts) == 0 {
		fmt.Println("No accounts in this KeyStore file")
		os.Exit(1)
	}

	for accountid, zilaccount := range(w.Accounts) {
		fmt.Printf("\n\nZilliqa Account %q\n", accountid)
		if len(zilaccount.PrivateKey) == 0 {
			fmt.Println(" FAIL to decrypt. Check your passphrase")
		} else {
			fmt.Println(" public wallet address (hex)           :", zilaccount.Address)
			add, err := bech32.ToBech32Address(string(zilaccount.Address))
			if err != nil {
				fmt.Println(err)		
			} else {
				fmt.Println(" public wallet address (human readable):", add)
			}
			fmt.Println(" public wallet key (hex)               :", hex.EncodeToString(zilaccount.PublicKey))
			fmt.Println("WARNING! do not share the following private key. keep it in a safe place, like your passphrase.")
			fmt.Println(" private wallet key (hex)              :", hex.EncodeToString(zilaccount.PrivateKey))
		}
	}
}
