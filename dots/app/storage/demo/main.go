package main

import (
    "github.com/scryinfo/dp/dots/app/storage"
    "github.com/scryinfo/dp/dots/app/storage/definition"
)

func main() {
    var db storage.SQLite
    db.Init()
    db.Create(definition.Account{Address: "0x40", NickName: "", FromBlock: 10, IsVerifier: false, Verify: nil})
}
