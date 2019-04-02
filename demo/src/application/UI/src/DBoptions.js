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
            this.db.createObjectStore("accounts", {keyPath: "address"})
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
                        Seller: cursor.value.Seller,
                        SupportVerify: cursor.value.SupportVerify,
                        MetaDataExtension: cursor.value.MetaDataExtension,
                        ProofDataExtensions: cursor.value.ProofDataExtensions,
                        PublishID: cursor.value.PublishID
                    })
                    cursor.continue()
                }
            }
        }
    },
    write:function(params, cb) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name)
        let request = store.put(params)
        request.onerror = function(err){
            console.log(err)
        }
        request.onsuccess = function () {
            cb()
        }
    },
    read:function (key, cb) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name)
        let request = store.get(key)
        request.onerror = function (err) {
            console.log(err)
        }
        request.onsuccess = function (event) {
            cb(event.target.result)
        }
    },
    reset:function() {
        let store = dl_db.db.transaction(dl_db.db_store_name,"readwrite").objectStore(dl_db.db_store_name)
        let request = store.clear()
        request.onerror = function (err) {
            console.log(err)
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
            _this.$store.state.transactionbuy = []
            _this.$store.state.transactionsell = []
            _this.$store.state.transactionverifier = []
            tx_db.db = event.target.result
            let curUser = _this.$store.state.account.toLowerCase()
            let s = tx_db.db.transaction(tx_db.db_store_name,"readonly").objectStore(tx_db.db_store_name)
            let c = s.openCursor()
            c.onsuccess = function(event) {
                let cursor = event.target.result
                if (cursor) {
                    let t = {
                        Title: cursor.value.Title,
                        Price: cursor.value.Price,
                        Keys: cursor.value.Keys,
                        Description: cursor.value.Description,
                        Buyer: cursor.value.Buyer,
                        Seller: cursor.value.Seller,
                        State: cursor.value.State,
                        SupportVerify: cursor.value.SupportVerify,
                        StartVerify: cursor.value.StartVerify,
                        MetaDataExtension: cursor.value.MetaDataExtension,
                        ProofDataExtensions: cursor.value.ProofDataExtensions,
                        MetaDataIDEncWithSeller: cursor.value.MetaDataIDEncWithSeller,
                        MetaDataIDEncWithBuyer: cursor.value.MetaDataIDEncWithBuyer,
                        Verifier1: cursor.value.Verifier1,
                        Verifier2: cursor.value.Verifier2,
                        Verifier1Response: cursor.value.Verifier1Response,
                        Verifier2Response: cursor.value.Verifier2Response,
                        ArbitrateResult: cursor.value.ArbitrateResult,
                        PublishID: cursor.value.PublishID,
                        TransactionID: cursor.value.TransactionID
                    }
                    if (t.Buyer.toLowerCase() === curUser) {
                        _this.$store.state.transactionbuy.push(t)
                    }
                    if (t.Seller.toLowerCase() === curUser) {
                        _this.$store.state.transactionsell.push(t)
                    }
                    if (t.Verifier1.toLowerCase() === curUser || t.Verifier2.toLowerCase() === curUser) {
                        _this.$store.state.transactionverifier.push(t)
                    }
                    cursor.continue()
                }
            }
        }
    },
    write:function(params, cb) {
        let store = tx_db.db.transaction(tx_db.db_store_name, "readwrite").objectStore(tx_db.db_store_name)
        let request = store.put(params)
        request.onerror = function(err){
            console.log(err)
        }
        request.onsuccess = function() {
            cb()
        }
    },
    read:function (key, cb) {
        let store = tx_db.db.transaction(tx_db.db_store_name, "readwrite").objectStore(tx_db.db_store_name)
        let request = store.get(key)
        request.onerror = function (err) {
            console.log(err)
        }
        request.onsuccess = function (event) {
            cb(event.target.result)
        }
    },
    reset:function() {
        let store = tx_db.db.transaction(tx_db.db_store_name,"readwrite").objectStore(tx_db.db_store_name)
        let request = store.clear()
        request.onerror = function (err) {
            console.log(err)
        }
    }
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
                    _this.$store.state.accounts.push({
                        address: cursor.value.address,
                        fromBlock: cursor.value.fromBlock,
                        isVerifier: cursor.value.isVerifier
                    })
                    cursor.continue()
                }
            }
        }
    },
    write:function(params) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name)
        let request = store.put(params)
        request.onerror = function(err){
            console.log(err)
        }
    },
    read: function(addr, cb) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name)
        let request = store.get(addr)
        request.onerror = function(err) {
            console.log(err)
        }
        request.onsuccess = function (event) {
            cb(event.target.result)
        }
    },
    // prepare a single remove function, for delete a account, maybe I can give out a button for user?
    remove:function (key) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name)
        let request = store.delete(key)
        request.onerror = function (err) {
            console.log(err)
        }
    },
    // Administrator function: reset acc_db when chain restart.
    reset:function () {
        let store = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name)
        let c = store.openCursor()
        c.onsuccess = function(event) {
            let cursor = event.target.result
            if (cursor) {
                store.put({
                    address: cursor.value.address,
                    fromBlock: 1,
                    isVerifier: false
                })
                cursor.continue()
            }
        }
    }
}

export {dl_db, tx_db, acc_db}
