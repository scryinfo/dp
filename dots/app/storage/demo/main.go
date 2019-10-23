package main

import (
    "fmt"
    "github.com/scryinfo/dp/dots/app/storage"
    "github.com/scryinfo/dp/dots/app/storage/definition"
    "strconv"
)

func main() {
    var (
        db *storage.SQLite
        err error
    )

    // init
    db.Init()

    // create
    for i := 1; i < 10; i++ {
        t := definition.DataList{
            PublishId: strconv.Itoa(i),
            Title: "db operate demo",
            Price: strconv.Itoa(i * 100),
            Keys: "test case " + strconv.Itoa(i),
            Description: "{description}",
            Seller: "0x40digits",
            SupportVerify: true,
            MetaDataExtension: ".txt",
            ProofDataExtensions: []byte{91,34,46,106,112,103,34,44,34,46,106,112,101,103,34,93},
            // CreateAt init by gorm
        }
        if err = db.Create(&t); err != nil {
            panic("error in create" + err.Error())
        }
    }
    fmt.Println("-------")

    // read
    // read all
    var dls []definition.DataList
    if err = db.ReadSome("publish_id desc", &dls); err != nil {
        panic("error in read some" + err.Error())
    }
    fmt.Println(len(dls))
    for i := range dls {
        fmt.Printf("%+v\n", dls[i])
    }
    fmt.Println("-------")

    // read one item
    var dl definition.DataList
    if err = db.Read(&dl, "Price = ?", "500"); err != nil {
        panic("error in read one" + err.Error())
    }
    fmt.Printf("%+v\n", dl)
    fmt.Println("-------")

    // update
    updateMap := make(map[string]interface{})
    updateMap["Description"] = "update test"
    if err = db.Update(updateMap, &dl, "Price = ?", "500"); err != nil {
        panic("error in update" + err.Error())
    }
    fmt.Printf("%+v\n", dl)
    fmt.Println("-------")

    // delete all, clean the record
    if err = db.Delete(definition.DataList{}, "publish_id < ?", 9999); err != nil {
        panic("error in delete" + err.Error())
    }

    // read all to check delete result
    if err = db.ReadSome("publish_id desc", &dls); err != nil {
        panic("error in read some" + err.Error())
    }
    if len(dls) == 0 {
        fmt.Println("test passed. ")
    } else {
        fmt.Println(len(dls))
        for i := range dls {
            fmt.Printf("%+v\n", dls[i])
        }
        panic("delete not run as expect, result is above not nil. ")
    }
}
