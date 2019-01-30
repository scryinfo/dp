let dl_db = {};
dl_db.init = function(){
    this.db_name = "Datalist";
    this.db_version = "1";
    this.db_store_name = "datalist";

    let request = indexedDB.open(this.db_name,this.db_version);
    request.onerror = function(event) {
        alert("open failed with error code: " + event.target.errorCode);
    };
    request.onupgradeneeded = function(event) {
        this.db = event.target.result;
        let store = this.db.createObjectStore(dl_db.db_store_name);
        store.createIndex("ti","title",{unique:false});
    };

    request.onsuccess = function(event) {
        // 此处采用异步通知. 在使用curd的时候请通过事件触发
        dl_db.db = event.target.result;
        let t = dl_db.db.transaction(dl_db.db_store_name,"readonly");
        let s = t.objectStore(dl_db.db_store_name);
        let c = s.openCursor();
        c.onsuccess = function(event) {
            let cursor = event.target.result;
            if (cursor) {
                let p = {};
                p.title = cursor.value.title;p.price = cursor.value.price;p.keys = cursor.value.keys;
                p.description = cursor.value.description;p.owner = cursor.value.owner;
                main.insDL(p,cursor.key);
                cursor.continue();
            }
        };

    };
};

dl_db.write = function(params, key) {
    // here must explicit declaration transaction
    let transaction = dl_db.db.transaction(dl_db.db_store_name, "readwrite");
    let store = transaction.objectStore(dl_db.db_store_name);
    let request = store.put(params,key);
    request.onerror = function(event){
        console.log(event);
    };
};

dl_db.delete = function(key) {
    // dl_db.db.transaction.objectStore is not a function
    dl_db.db.transaction(dl_db.db_store_name, "readwrite").objectStore(dl_db.db_store_name).delete(key);
};

dl_db.select = function(key) {
    // 第二个参数可以省略
    let transaction = dl_db.db.transaction(dl_db.db_store_name,"readwrite");
    let store = transaction.objectStore(dl_db.db_store_name);
    let request;
    if (key) {
        request = store.get(key);
    }else {
        request = store.getAll();
    }

    request.onsuccess = function () {
        return request.result;
    };
};

dl_db.clear = function() {
    dl_db.db.transaction(dl_db.db_store_name,"readwrite").objectStore(dl_db.db_store_name).clear();
};
