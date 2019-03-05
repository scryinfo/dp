let DBoptions = {
    init: function (_this) {
        DBoptions.getDatalist(_this)
        DBoptions.getTransaction(_this)
    },
    getDatalist:function (_this) {
        astilectron.sendMessage({Name:"get.datalist",Payload:""},function (message) {
            for (let i=0;i<message.payload.length;i++) {
                let t = message.payload[i]
                let p = {}
                p.Title = t.Title;p.Price = t.Price;p.Keys = t.Keys;p.Description = t.Description;p.Owner = t.Owner
                dl_db.write(p,t.ID)
            }
            dl_db.init(_this)
        })
    },
    getTransaction:function (_this) {
        astilectron.sendMessage({Name:"get.transaction",Payload:_this.$store.state.account}, function (message) {
            for (let i=0;i<message.payload.length;i++) {
                let t = message.payload[i]
                let p = {}
                p.Title = t.Title;p.Seller = t.Seller;p.Buyer = t.Buyer;p.State = t.State
                mt_db.write(p,t.TransactionID)
            }
            mt_db.init(_this)
        })
    },
}

let dl_db = {
    init:function(_this){
        this.db_name = "Database"
        this.db_version = "1"
        this.db_store_name = "datalist"

        let request = indexedDB.open(this.db_name,this.db_version)
        request.onerror = function(event) {
            alert("open failed with error code: " + event.target.errorCode)
        }
        request.onupgradeneeded = function(event) {
            this.db = event.target.result
            this.db.createObjectStore(dl_db.db_store_name)
            this.db.createObjectStore("transaction")
            this.db.createObjectStore("accounts")
        }
        request.onsuccess = function(event) {
            _this.$store.state.datalist = []
            dl_db.db = event.target.result
            let s = dl_db.db.transaction(dl_db.db_store_name,"readonly").objectStore(dl_db.db_store_name)
            let c = s.openCursor()
            c.onsuccess = function(event) {
                let cursor = event.target.result
                if (cursor) {
                    _this.$store.state.datalist.push({
                        Title: cursor.value.Title,
                        Price: cursor.value.Price,
                        Keys: cursor.value.Keys,
                        Description: cursor.value.Description,
                        Owner: cursor.value.Owner,
                        ID: cursor.key
                    })
                    cursor.continue()
                }
            }
        }
    },
    write:function(params, key) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name)
        let request = store.put(params,key)
        request.onerror = function(event){
            console.log(event)
        }
    }
}
let mt_db = {
    init:function (_this) {
        this.db_name = "Database"
        this.db_version = "1"
        this.db_store_name = "transaction"

        let request = indexedDB.open(this.db_name,this.db_version)
        request.onerror = function(event) {
            alert("open failed with error code: " + event.target.errorCode)
        }
        request.onsuccess = function(event) {
            _this.$store.state.mytransaction = []
            mt_db.db = event.target.result
            let s = mt_db.db.transaction(mt_db.db_store_name,"readonly").objectStore(mt_db.db_store_name)
            let c = s.openCursor()
            c.onsuccess = function(event) {
                let cursor = event.target.result
                if (cursor) {
                    let state = ""
                    switch (cursor.value.State) {
                        case 0: state = "Created"; break
                        case 1: state = "Voted"; break
                        case 2: state = "Payed"; break
                        case 3: state = "ReadyForDownload"; break
                        case 4: state = "Closed"; break
                        default: state = "State parse error."; break
                    }
                    _this.$store.state.mytransaction.push({
                        Title: cursor.value.Title,
                        Seller: cursor.value.Seller,
                        Buyer: cursor.value.Buyer,
                        State: state,
                        TransactionID: cursor.key
                    })
                    cursor.continue()
                }
            }
        }
    },
    write:function(params, key) {
        let store = mt_db.db.transaction(mt_db.db_store_name, "readwrite").objectStore(mt_db.db_store_name)
        let request = store.put(params,key)
        request.onerror = function(event){
            console.log(event)
        }
    },
}
let acc_db = {
    init:function(_this){
        this.db_name = "Database"
        this.db_version = "1"
        this.db_store_name = "accounts"

        let request = indexedDB.open(this.db_name,this.db_version)
        request.onerror = function(event) {
            alert("open failed with error code: " + event.target.errorCode)
        }
        request.onsuccess = function(event) {
            _this.$store.state.accounts = [{ address: "" }]
            acc_db.db = event.target.result
            let s = acc_db.db.transaction(acc_db.db_store_name,"readonly").objectStore(acc_db.db_store_name)
            let c = s.openCursor()
            c.onsuccess = function(event) {
                let cursor = event.target.result
                if (cursor) {
                    _this.$store.state.accounts.push({ address: cursor.key })
                    cursor.continue()
                }
            }
        }
    },
    write:function(addr) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name)
        let request = store.put(addr,addr)
        request.onerror = function(event){
            console.log(event)
        }
    },
    // prepare a single remove function, for delete a account, maybe I can put a
    remove:function (key) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name)
        let request = store.delete(key)
        request.onerror = function (event) {
            console.log(event)
        }
    }
}

export {dl_db, mt_db, acc_db, DBoptions}
