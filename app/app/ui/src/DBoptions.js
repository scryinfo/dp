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
                });
                cursor.continue();
            }
        };
    },
    write: function (params, cb) {
        let store = dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name);
        let request = store.put(params);
        request.onerror = function(err){
            console.log(err);
        };
        request.onsuccess = function () {
            cb();
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
    db: IDBDatabase,
    init: function (_this){
        _this.$store.state.accounts = [{ address: "", fromBlock: 1, isVerifier:false}];
        let c = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name).openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                _this.$store.state.accounts.push({
                    address: cursor.value.address,
                    fromBlock: cursor.value.fromBlock,
                    isVerifier: cursor.value.isVerifier
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
    // prepare a single remove function, for delete a account, maybe I can give out a button for user?
    remove: function (key) {
        let store = acc_db.db.transaction(acc_db.db_store_name, "readwrite").objectStore(acc_db.db_store_name);
        let request = store.delete(key);
        request.onerror = function (err) {
            console.log(err);
        };
    },
    // Administrator function: reset acc_db when chain restart anyhow.
    reset: function () {
        let store = acc_db.db.transaction(acc_db.db_store_name,"readwrite").objectStore(acc_db.db_store_name);
        let c = store.openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                store.put({
                    address: cursor.value.address,
                    fromBlock: 1,
                    isVerifier: false
                });
                cursor.continue();
            }
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
                _this.$store.state.transactionbuy.push({
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
    shut: function () {
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
                    Verifier1: cursor.value.Verifier1,
                    Verifier2: cursor.value.Verifier2,
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
    shut: function () {
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
                    Verifier1: cursor.value.Verifier1,
                    Verifier2: cursor.value.Verifier2,
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
    shut: function () {
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

let db_options = {
    utilsDBInit: function (_this) {
        let request = indexedDB.open("Utils", 1);
        request.onupgradeneeded = function (event) {
            dl_db.db = event.target.result;
            acc_db.db = event.target.result;
            event.target.result.createObjectStore(dl_db.db_store_name, {keyPath: "PublishID"});
            event.target.result.createObjectStore(acc_db.db_store_name, {keyPath: "address"});
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
        let request = indexedDB.open(address, 1);
        request.onupgradeneeded = function (event) {
            txBuyer_db.db = event.target.result;
            txSeller_db.db = event.target.result;
            txVerifier_db.db = event.target.result;
            event.target.result.createObjectStore(txBuyer_db.db_store_name, {keyPath: "TransactionID"});
            event.target.result.createObjectStore(txSeller_db.db_store_name, {keyPath: "TransactionID"});
            event.target.result.createObjectStore(txVerifier_db.db_store_name, {keyPath: "TransactionID"});
        };
        request.onsuccess = function (event) {
            txBuyer_db.db = event.target.result;
            txSeller_db.db = event.target.result;
            txVerifier_db.db = event.target.result;
        };
        request.onerror = function (err) {
            console.log("数据库打开/创建失败，错误代码： ", err);
        };
    },
    txDBsDataUpdate: function (_this) {
        txBuyer_db.init(_this);
        txSeller_db.init(_this);
        txVerifier_db.init(_this);
    },
    userDBClose: function () {
        txBuyer_db.shut();
        txSeller_db.shut();
        txVerifier_db.shut();
    }
};

export {db_options, dl_db, acc_db, txBuyer_db, txSeller_db, txVerifier_db};
