let mt_db = {};
mt_db.init = function(){
    this.db_name = "Transaction";
    this.db_version = "1";
    this.db_store_name = "transaction";

    let request = indexedDB.open(this.db_name,this.db_version);
    request.onerror = function(event) {
        alert("open failed with error code: " + event.target.errorCode);
    };
    request.onupgradeneeded = function(event) {
        this.db = event.target.result;
        this.db.createObjectStore(mt_db.db_store_name);
    };

    request.onsuccess = function(event) {
        // 此处采用异步通知. 在使用curd的时候请通过事件触发
        mt_db.db = event.target.result;
        let t = mt_db.db.transaction(mt_db.db_store_name,"readonly");
        let s = t.objectStore(mt_db.db_store_name);
        let c = s.openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                let p = {};
                p.title = cursor.value.title;p.seller = cursor.value.seller;p.buyer = cursor.value.buyer;p.state = cursor.value.state;
                main.insMT(p,cursor.key);
                cursor.continue();
            }
        };
    };
};

mt_db.write = function(params, key) {
    // here must explicit declaration transaction
    let transaction = mt_db.db.transaction(mt_db.db_store_name, "readwrite");
    let store = transaction.objectStore(mt_db.db_store_name);
    let request = store.put(params,key);
    request.onerror = function(event){
        console.log(event);
    };
};

mt_db.delete = function(primaryKey) {
    // mt_db.db.transaction.objectStore is not a function
    mt_db.db.transaction(mt_db.db_store_name, "readwrite").objectStore(mt_db.db_store_name).delete(primaryKey);
};

mt_db.select = function(key) {
    // 第二个参数可以省略
    let transaction = mt_db.db.transaction(mt_db.db_store_name,"readwrite");
    let store = transaction.objectStore(mt_db.db_store_name);
    let request;
    if (key)
        request = store.get(key);
    else
        request = store.getAll();

    request.onsuccess = function () {
        console.log(request.result);
    };
};

mt_db.clear = function() {
    mt_db.db.transaction(mt_db.db_store_name,"readwrite").objectStore(mt_db.db_store_name).clear();
};