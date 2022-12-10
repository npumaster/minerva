package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/filecoin-project/chain/types"
	"github.com/filecoin-project/chain/wallet"
	"log"
	"strings"
)

func main() {
	ctx := context.Background()

	w, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {
		log.Fatal(err)
	}

	a, err := w.WalletNew(ctx, types.KTSecp256k1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(a)

	keyInfo, err := w.WalletExport(ctx, a)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(keyInfo)
	if err != nil {
		log.Fatal(err)
	}
	bs := hex.EncodeToString(b)
	log.Println(bs)

	data, err := hex.DecodeString(strings.TrimSpace(bs))
	if err != nil {
		log.Fatal(err)
	}

	var ki types.KeyInfo
	if err := json.Unmarshal(data, &ki); err != nil {
		log.Fatal(err)
	}

	addr, err := w.WalletImport(ctx, &ki)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addr)
}
