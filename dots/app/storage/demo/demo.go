package main

import (
    "encoding/json"
    "fmt"
    "github.com/scryinfo/dp/dots/app/storage"
    "github.com/scryinfo/dp/dots/app/storage/definition"
    "strconv"
)

var (
    db *storage.SQLite
    dls []definition.DataList
)

func main() {
    // init
    db.Init()

    // delete all, clean the record
    deleteAll()

    // create
    create()

    // read
    // read with sql
    readWithSQL()

    // read all desc
    readAllDESC()

    // update
    update()

    // update without hooks
    updateWithoutHooks()

    // delete all, clean the record
    deleteAll()

    // read all to check delete result
    readAllToCheck()
}

func deleteAll() {
    fmt.Println("-------Delete all: ")

    deleteNum, err := db.Delete(definition.DataList{}, "publish_id < ?", 9999)
    if  err != nil {
        panic("error in delete. " + err.Error())
    }
    fmt.Println(deleteNum)

    return
}

func create() {
    fmt.Println("-------Create: ")

    r, err := serialize([]string{".jpg", ".avi"})
    if err != nil {
        panic("error in json serialize. " + err.Error())
    }

    data := make([]definition.DataList, 0)
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
            ProofDataExtensions: r,
            // CreateAt init by gorm
        }
        if i > 5 {
            t.Price = strconv.Itoa(500)
        }

        data = append(data, t)
    }

    createNum, err := db.Create(data)
    if err != nil {
        panic("error in create. " + err.Error())
    }

    fmt.Printf("Create %d items. \n", createNum)

    return
}

func readWithSQL() {
    fmt.Printf("-------Read with sql ")

    readNum, err := db.Read(&dls, "", "Price = ?", "500")
    if err != nil {
        panic("error in read with sql. " + err.Error())
    }
    fmt.Println(readNum, "items. ")
    for i := range dls {
        fmt.Printf("%+v\n", dls[i])
    }

    return
}

func readAllDESC() {
    fmt.Printf("-------Read all desc ")

    readNum, err := db.Read(&dls, "publish_id desc", "")
    if err != nil {
        panic("error in read all desc. " + err.Error())
    }
    fmt.Println(readNum, "items. ")
    for i := range dls {
        fmt.Printf("%+v\n", dls[i])
    }

    return
}

func update() {
    fmt.Println("-------Update: ")

    updateMap := make(map[string]interface{})
    updateMap["Title"] = "Update one will emit hooks"
    updateMap["Keys"] = "update, one, hooks"

    updateNum, err := db.Update(&dls, updateMap, "publish_id <= ?", "3")
    if err != nil {
        panic("error in update. " + err.Error())
    }
    fmt.Printf("Update %d items. \n", updateNum)

    db.Read(&dls, "", "")
    for i := range dls {
        fmt.Printf("%+v\n", dls[i])
    }

    return
}

func updateWithoutHooks() {
    fmt.Printf("-------Update without hooks ")

    updateMap := make(map[string]interface{})
    updateMap["Title"] = "Update"
    updateMap["Keys"] = "update"
    updateMap["Description"] = "update"

    updateNum, err := db.UpdateWithoutHooks(definition.DataList{}, updateMap, "Price = ?", "500")
    if  err != nil {
        panic("error in update without hooks. " + err.Error())
    }
    fmt.Println(updateNum, "items. ")

    db.Read(&dls, "", "")
    for i := range dls {
        fmt.Printf("%+v\n", dls[i])
    }

    return
}

func readAllToCheck() {
    fmt.Printf("-------Read all ")

    readNum, _ := db.Read(&dls, "", "")
    fmt.Println(readNum, "items. ")

    fmt.Println("-------Result: ")
    if readNum == 0 {
        fmt.Println("> demo passed. ")
    } else {
        fmt.Println(len(dls))
        for i := range dls {
            fmt.Printf("%+v\n", dls[i])
        }
        panic("> delete not run as expect, result is above not nil. ")
    }

    return
}

func serialize(i interface{}) ([]byte, error) {
    return json.Marshal(i)
}
