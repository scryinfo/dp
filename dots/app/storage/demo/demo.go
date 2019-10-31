package main

import (
	"encoding/json"
	"fmt"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/dp/dots/app/storage"
	"github.com/scryinfo/dp/dots/app/storage/definition"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var (
	db  *storage.SQLite
	dls []definition.DataList
)

// use: go build demo.go, and run 'demo.exe';
//      or config debug items, mainly 'out dir' and 'exec file name'
func main() {
	// init
	var (
		l   dot.Line
		err error
	)
	{
		l, err = line.BuildAndStart(func(l dot.Line) error {
			l.PreAdd(storage.SQLiteTypeLive())
			return nil
		})

		if err != nil {
			dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
			return
		}

		var d dot.Dot
		d, err = l.ToInjecter().GetByLiveId(dot.LiveId(storage.SQLiteTypeId))
		if err != nil {
			dot.Logger().Errorln("load SQLite component failed.")
		}

		if sql, ok := d.(*storage.SQLite); !ok {
			dot.Logger().Errorln("load SQLite component failed.", zap.Any("d", d))
		} else {
			db = sql
		}
	}

	defer line.StopAndDestroy(l, true)

	go start()

	ssignal.WaitCtrlC(func(s os.Signal) bool {
		return false //quit
	})
}

func start() {
	// delete all, clean the record
	deleteAll()

	// create
	create()

	// read
	{
		// read with sql
		readWithSQL()

		// read all desc
		readAllDESC()
	}

	// update
	{
		// update
		update()

		// update without hooks
		updateWithoutHooks()
	}

	// delete all, clean the record
	deleteAll()

	// read all to check delete result
	readAllToCheck()
}

func deleteAll() {
	fmt.Printf("-------Delete all ")

	deleteNum, err := db.Delete(definition.DataList{}, "")
	if err != nil {
		panic("error in delete. " + err.Error())
	}
	fmt.Println(deleteNum, "items. ")
	dot.Logger().Infoln("-------Delete all " + strconv.FormatInt(deleteNum, 10) + " items.")

	return
}

func create() {
	fmt.Println("-------Create: ")
	dot.Logger().Infoln("-------Create: ")

	r, err := serialize([]string{".jpg", ".avi"})
	if err != nil {
		panic("error in json serialize. " + err.Error())
	}

	data := make([]definition.DataList, 0)
	for i := 1; i < 10; i++ {
		t := definition.DataList{
			PublishId:           strconv.Itoa(i),
			Title:               "db operate demo",
			Price:               strconv.Itoa(i * 100),
			Keys:                "test case " + strconv.Itoa(i),
			Description:         "{description}",
			Seller:              "0x40digits",
			SupportVerify:       true,
			MetaDataExtension:   ".txt",
			ProofDataExtensions: r,
			// CreateAt init by gorm
		}
		if i > 5 {
			t.Price = strconv.Itoa(500)
		}

		data = append(data, t)
	}

	createNum, err := db.Insert(data)
	if err != nil {
		panic("error in create. " + err.Error())
	}

	fmt.Printf("Create %d items. \n", createNum)
	dot.Logger().Infoln("Create " + strconv.FormatInt(createNum, 10) + " items.")

	return
}

func readWithSQL() {
	fmt.Printf("-------Read with sql ")

	readNum, err := db.Read(&dls, "", "Price = ?", "500")
	if err != nil {
		panic("error in read with sql. " + err.Error())
	}
	fmt.Println(readNum, "items. ")
	dot.Logger().Infoln("-------Read with sql " + strconv.FormatInt(readNum, 10) + " items.")
	for i := range dls {
		fmt.Printf("%+v\n", dls[i])
		dot.Logger().Infoln("", zap.Any("", dls[i]))
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
	dot.Logger().Infoln("-------Read all desc " + strconv.FormatInt(readNum, 10) + " items.")
	for i := range dls {
		fmt.Printf("%+v\n", dls[i])
		dot.Logger().Infoln("", zap.Any("", dls[i]))
	}

	return
}

func update() {
	fmt.Println("-------Update: ")
	dot.Logger().Infoln("-------Update: ")

	updateMap := make(map[string]interface{})
	updateMap["Title"] = "Update one will emit hooks"
	updateMap["Keys"] = "update, one, hooks"

	updateNum, err := db.Update(&dls, updateMap, "publish_id <= ?", "3")
	if err != nil {
		panic("error in update. " + err.Error())
	}
	fmt.Printf("Update %d items. \n", updateNum)
	dot.Logger().Infoln("Update " + strconv.FormatInt(updateNum, 10) + " items.")

	db.Read(&dls, "", "")
	for i := range dls {
		fmt.Printf("%+v\n", dls[i])
		dot.Logger().Infoln("", zap.Any("", dls[i]))
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
	if err != nil {
		panic("error in update without hooks. " + err.Error())
	}
	fmt.Println(updateNum, "items. ")
	dot.Logger().Infoln("-------Update without hooks " + strconv.FormatInt(updateNum, 10) + " items.")

	db.Read(&dls, "", "")
	for i := range dls {
		fmt.Printf("%+v\n", dls[i])
		dot.Logger().Infoln("", zap.Any("", dls[i]))
	}

	return
}

func readAllToCheck() {
	fmt.Printf("-------Read all ")

	readNum, _ := db.Read(&dls, "", "")
	fmt.Println(readNum, "items. ")
	dot.Logger().Infoln("-------Read all " + strconv.FormatInt(readNum, 10) + " items.")

	fmt.Println("-------Result: ")
	dot.Logger().Infoln("-------Result: ")
	if readNum == 0 {
		fmt.Println("> demo passed. ")
		dot.Logger().Infoln("> demo passed. ")
	} else {
		fmt.Println(len(dls))
		dot.Logger().Infoln("> demo failed! ", zap.Int("len(dls) not 0 as expect but is: ", len(dls)))
		for i := range dls {
			fmt.Printf("%+v\n", dls[i])
			dot.Logger().Infoln("", zap.Any("", dls[i]))
		}
		panic("> delete not run as expect, result is above not nil. ")
	}

	return
}

func serialize(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}
