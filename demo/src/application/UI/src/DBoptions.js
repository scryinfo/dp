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
            this.db.createObjectStore(dl_db.db_store_name, {keyPath: "PublishID"})
            this.db.createObjectStore("transaction", {keyPath: "TransactionID"})
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
                    let sv = ""
                    switch (cursor.value.SupportVerify) {
                        case true: sv = "Yes."; break
                        case false: sv = "No."; break
                    }
                    _this.$store.state.datalist.push({
                        Title: cursor.value.Title,
                        Price: cursor.value.Price,
                        Keys: cursor.value.Keys,
                        Description: cursor.value.Description,
                        SupportVerify: sv,
                        PublishID: cursor.key
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
let tx_db = {
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
            tx_db.db = event.target.result
            let s = tx_db.db.transaction(tx_db.db_store_name,"readonly").objectStore(tx_db.db_store_name)
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
                        default: state = "undefined state."; break
                    }
                    let t = {
                        Title: cursor.value.Title,
                        Price: cursor.value.Price,
                        Seller: cursor.value.Seller,
                        Buyer: cursor.value.Buyer,
                        State: state,
                        PublishID: cursor.value.PublishID,
                        TransactionID: cursor.key,
                        Verifier1Response: cursor.value.Verifier1Response,
                        Verifier2Response: cursor.value.Verifier2Response,
                        ArbitrateResult: cursor.value.ArbitrateResult
                    }
                    let acc = _this.$store.state.account
                    if (cursor.value.Buyer === acc) {
                        _this.$store.state.transactionbuy.push(t)
                    }
                    if (cursor.value.Seller === acc) {
                        _this.$store.state.transactionsell.push(t)
                    }
                    cursor.continue()
                }
            }
        }
    },
    write:function(params, key) {
        let store = tx_db.db.transaction(tx_db.db_store_name, "readwrite").objectStore(tx_db.db_store_name)
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

export {dl_db, tx_db, acc_db}
