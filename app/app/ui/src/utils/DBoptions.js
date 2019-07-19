// Scry Info.  All rights reserved.
// license that can be found in the license file.

let dl_db = {
    db_name: "Utils",
    db_store_name: "datalist",
    db: IDBDatabase,
    init: function (_this){
        _this.$store.state.datalist = [];
        let c = dl_db.db.transaction(dl_db.db_store_name,"readwrite").objectStore(dl_db.db_store_name).openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                let sv = "";
                switch (cursor.value.SupportVerify) {
                    case true: sv = "支持验证"; break;
                    case false: sv = "不支持验证"; break;
                }
                _this.$store.state.datalist.push({
                    Title: cursor.value.Title,
                    Price: cursor.value.Price,
                    Keys: cursor.value.Keys,
                    Description: cursor.value.Description,
                    Seller: cursor.value.Seller,
                    SupportVerify: cursor.value.SupportVerify,
                    SVDisplay: sv,
                    MetaDataExtension: cursor.value.MetaDataExtension,
                    ProofDataExtensions: cursor.value.ProofDataExtensions,
                    PublishID: cursor.value.PublishID
                });
                cursor.continue();
            }
        };
    },
    write: function (params) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
    },
    read: function (key, cb) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name);
        let request = store.get(key);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    reset: function () {
        let store = dl_db.db.transaction(dl_db.db_store_name,"readwrite").objectStore(dl_db.db_store_name);
        let request = store.clear();
        request.onerror = function (err) {
            console.log(err);
        };
    }
};

let acc_db = {
    db_name: "Utils",
    db_store_name: "accounts",
    db_index_name: "nickname",
    db: IDBDatabase,
    init: function (_this){
        _this.$store.state.accounts = [{ address: ""}];
        let c = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name).openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                _this.$store.state.accounts.push({
                    address: cursor.value.address
                });
                cursor.continue();
            }
        };
    },
    write: function (params) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
    },
    read: function (addr, cb) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.get(addr);
        request.onerror = function(err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    readIndex: function (indexName, indexValue, cb) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.index(indexName).get(indexValue);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    readAll: function(cb) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.getAll();
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    remove: function (key) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.delete(key);
        request.onerror = function (err) {
            console.log(err);
        };
    },
    reset: function () {
        let store = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name);
        let c = store.openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                store.put({
                    address: cursor.value.address,
                    nickname: cursor.value.address,
                    fromBlock: 1,
                    isVerifier: false
                });
                cursor.continue();
            }
        };
    }
};

let txSeller_db = {
    db_name: "",
    db_store_name: "tx_seller",
    db: IDBDatabase,
    init: function (_this) {
        _this.$store.state.transactionsell = [];
        let c = txSeller_db.db.transaction(txSeller_db.db_store_name, "readwrite").objectStore(txSeller_db.db_store_name).openCursor();
        c.onsuccess = function (event) {
            let cursor = event.target.result;
            if (cursor) {
                _this.$store.state.transactionsell.push({
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
                    MetaDataIDEncWithArbitrator: cursor.value.MetaDataIDEncWithArbitrator,
                    ArbitrateResult: cursor.value.ArbitrateResult,
                    PublishID: cursor.value.PublishID,
                    TransactionID: cursor.value.TransactionID
                });
                cursor.continue();
            }
        };
    },
    write:function(params, cb) {
        let store = txSeller_db.db.transaction(txSeller_db.db_store_name, "readwrite").objectStore(txSeller_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
        request.onsuccess = function() {
            cb();
        };
    },
    read:function (key, cb) {
        let store = txSeller_db.db.transaction(txSeller_db.db_store_name, "readwrite").objectStore(txSeller_db.db_store_name);
        let request = store.get(key);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    closeDB: function () {
        txSeller_db.db.close();
    },
    reset:function() {
        let store = txSeller_db.db.transaction(txSeller_db.db_store_name,"readwrite").objectStore(txSeller_db.db_store_name);
        let request = store.clear();
        request.onerror = function (err) {
            console.log(err);
        };
    }
};

let txBuyer_db = {
    db_name: "",
    db_store_name: "tx_buy",
    db: IDBDatabase,
    init: function (_this) {
        _this.$store.state.transactionbuy = [];
        let c = txBuyer_db.db.transaction(txBuyer_db.db_store_name, "readwrite").objectStore(txBuyer_db.db_store_name).openCursor();
        c.onsuccess = function (event) {
            let cursor = event.target.result;
            if (cursor) {
                let sv = "";
                switch (cursor.value.SupportVerify) {
                    case true: sv = "支持验证"; break;
                    case false: sv = "不支持验证"; break;
                }
                _this.$store.state.transactionbuy.push({
                    Title: cursor.value.Title,
                    Price: cursor.value.Price,
                    Keys: cursor.value.Keys,
                    Description: cursor.value.Description,
                    Buyer: cursor.value.Buyer,
                    Seller: cursor.value.Seller,
                    State: cursor.value.State,
                    SupportVerify: cursor.value.SupportVerify,
                    SVDisplay: sv,
                    StartVerify: cursor.value.StartVerify,
                    MetaDataExtension: cursor.value.MetaDataExtension,
                    ProofDataExtensions: cursor.value.ProofDataExtensions,
                    MetaDataIDEncWithSeller: cursor.value.MetaDataIDEncWithSeller,
                    MetaDataIDEncWithBuyer: cursor.value.MetaDataIDEncWithBuyer,
                    MetaDataIDEncWithArbitrator: cursor.value.MetaDataIDEncWithArbitrator,
                    Verifier1Response: cursor.value.Verifier1Response,
                    Verifier2Response: cursor.value.Verifier2Response,
                    ArbitrateResult: cursor.value.ArbitrateResult,
                    PublishID: cursor.value.PublishID,
                    TransactionID: cursor.value.TransactionID
                });
                cursor.continue();
            }
        };
    },
    write: function (params, cb) {
        let store = txBuyer_db.db.transaction(txBuyer_db.db_store_name, "readwrite").objectStore(txBuyer_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
        request.onsuccess = function() {
            cb();
        };
    },
    read: function (key, cb) {
        let store = txBuyer_db.db.transaction(txBuyer_db.db_store_name, "readwrite").objectStore(txBuyer_db.db_store_name);
        let request = store.get(key);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    closeDB: function () {
        txBuyer_db.db.close();
    },
    reset: function () {
        let store = txBuyer_db.db.transaction(txBuyer_db.db_store_name,"readwrite").objectStore(txBuyer_db.db_store_name);
        let request = store.clear();
        request.onerror = function (err) {
            console.log(err);
        };
    }
};

let txVerifier_db = {
    db_name: "",
    db_store_name: "tx_verifier",
    db: IDBDatabase,
    init: function (_this) {
        _this.$store.state.transactionverifier = [];
        let c = txVerifier_db.db.transaction(txVerifier_db.db_store_name, "readwrite").objectStore(txVerifier_db.db_store_name).openCursor();
        c.onsuccess = function (event) {
            let cursor = event.target.result;
            if (cursor) {
                if (cursor.value.State === "Created" || cursor.value.State === "Voted") {
                    _this.$store.state.transactionverifier.push({
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
                        MetaDataIDEncWithArbitrator: cursor.value.MetaDataIDEncWithArbitrator,
                        Verifier1Response: cursor.value.Verifier1Response,
                        Verifier2Response: cursor.value.Verifier2Response,
                        ArbitrateResult: cursor.value.ArbitrateResult,
                        PublishID: cursor.value.PublishID,
                        TransactionID: cursor.value.TransactionID
                    });
                }
                cursor.continue();
            }
        };
    },
    write: function (params, cb) {
        let store = txVerifier_db.db.transaction(txVerifier_db.db_store_name, "readwrite").objectStore(txVerifier_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
        request.onsuccess = function() {
            cb();
        };
    },
    read: function (key, cb) {
        let store = txVerifier_db.db.transaction(txVerifier_db.db_store_name, "readwrite").objectStore(txVerifier_db.db_store_name);
        let request = store.get(key);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    closeDB: function () {
        txVerifier_db.db.close();
    },
    reset: function () {
        let store = txVerifier_db.db.transaction(txVerifier_db.db_store_name,"readwrite").objectStore(txVerifier_db.db_store_name);
        let request = store.clear();
        request.onerror = function (err) {
            console.log(err);
        };
    }
};

let txArbitrator_db = {
    db_name: "",
    db_store_name: "tx_arbitrator",
    db: IDBDatabase,
    init: function (_this) {
        _this.$store.state.transactionarbitrator = [];
        let c = txArbitrator_db.db.transaction(txArbitrator_db.db_store_name, "readwrite").objectStore(txArbitrator_db.db_store_name).openCursor();
        c.onsuccess = function (event) {
            let cursor = event.target.result;
            if (cursor) {
                if (cursor.value.State === "ReadyForDownload") {
                    _this.$store.state.transactionarbitrator.push({
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
                        MetaDataIDEncWithArbitrator: cursor.value.MetaDataIDEncWithArbitrator,
                        Verifier1Response: cursor.value.Verifier1Response,
                        Verifier2Response: cursor.value.Verifier2Response,
                        ArbitrateResult: cursor.value.ArbitrateResult,
                        PublishID: cursor.value.PublishID,
                        TransactionID: cursor.value.TransactionID
                    });
                }
                cursor.continue();
            }
        };
    },
    write: function (params, cb) {
        let store = txArbitrator_db.db.transaction(txArbitrator_db.db_store_name, "readwrite").objectStore(txArbitrator_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
        request.onsuccess = function() {
            cb();
        };
    },
    read: function (key, cb) {
        let store = txArbitrator_db.db.transaction(txArbitrator_db.db_store_name, "readwrite").objectStore(txArbitrator_db.db_store_name);
        let request = store.get(key);
        request.onerror = function (err) {
            console.log(err);
        };
        request.onsuccess = function (event) {
            cb(event.target.result);
        };
    },
    closeDB: function () {
        txArbitrator_db.db.close();
    },
    reset: function () {
        let store = txArbitrator_db.db.transaction(txArbitrator_db.db_store_name,"readwrite").objectStore(txArbitrator_db.db_store_name);
        let request = store.clear();
        request.onerror = function (err) {
            console.log(err);
        };
    }
};

let db_options = {
    utilsDBInit: function (_this) {
        let request = indexedDB.open("Utils", 1);
        request.onupgradeneeded = function (event) {
            dl_db.db = event.target.result;
            acc_db.db = event.target.result;
            event.target.result.createObjectStore(dl_db.db_store_name, {keyPath: "PublishID"});
            let store = event.target.result.createObjectStore(acc_db.db_store_name, {keyPath: "address"});
            store.createIndex(acc_db.db_index_name, acc_db.db_index_name, {unique: false});
        };
        request.onsuccess = function (event) {
            dl_db.db = event.target.result;
            acc_db.db = event.target.result;
            dl_db.init(_this);
            acc_db.init(_this);
        };
    },
    userDBInit: function (address) {
        txBuyer_db.db_name = address;
        txSeller_db.db_name = address;
        txVerifier_db.db_name = address;
        txArbitrator_db.db_name = address;
        let request = indexedDB.open(address, 1);
        request.onupgradeneeded = function (event) {
            txBuyer_db.db = event.target.result;
            txSeller_db.db = event.target.result;
            txVerifier_db.db = event.target.result;
            txArbitrator_db.db = event.target.result;
            event.target.result.createObjectStore(txBuyer_db.db_store_name, {keyPath: "TransactionID"});
            event.target.result.createObjectStore(txSeller_db.db_store_name, {keyPath: "TransactionID"});
            event.target.result.createObjectStore(txVerifier_db.db_store_name, {keyPath: "TransactionID"});
            event.target.result.createObjectStore(txArbitrator_db.db_store_name, {keyPath: "TransactionID"});
        };
        request.onsuccess = function (event) {
            txBuyer_db.db = event.target.result;
            txSeller_db.db = event.target.result;
            txVerifier_db.db = event.target.result;
            txArbitrator_db.db = event.target.result;
        };
        request.onerror = function (err) {
            console.log("数据库创建/打开失败，错误代码： ", err);
        };
    },
    txDBsDataUpdate: function (_this) {
        txBuyer_db.init(_this);
        txSeller_db.init(_this);
        txVerifier_db.init(_this);
        txArbitrator_db.init(_this);
    },
    userDBClose: function () {
        txBuyer_db.closeDB();
        txSeller_db.closeDB();
        txVerifier_db.closeDB();
        txArbitrator_db.closeDB();
    },
    deleteDB: function (db_name) {
        window.indexedDB.deleteDatabase(db_name);
    }
};

export {db_options, dl_db, acc_db, txBuyer_db, txSeller_db, txVerifier_db, txArbitrator_db};
