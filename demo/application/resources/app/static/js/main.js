let main = {
    init:function () {
        // Init
        database.init();

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // init account and lists
            document.getElementById("account").innerHTML = location.search.split("=")[1];
            main.getDatalist();
            main.getTransaction();
        });
    },
    getDatalist:function () {
        astilectron.sendMessage({Name:"get.datalist",Payload:""},function (message) {
            for (let i=0;i<message.payload.length;i++) {
                let t = message.payload[i];
                let p = {};
                p.title = t.Title;p.price = t.Price;p.keys = t.Keys;p.description = t.Description;p.owner = t.Owner;
                database.write(p,t.ID,dl_db);
            }
        });
    },
    getTransaction:function () {
        astilectron.sendMessage({Name:"get.transaction",Payload:location.search.split("=")[1]},
            function (message) {
            for (let i=0;i<message.payload.length;i++) {
                let t = message.payload[i];
                let p = {};
                p.title = t.Title;p.seller = t.Seller;p.buyer = t.Buyer;
                switch (parseInt(t.State)) {
                    case 0:p.state = "Created";break;
                    case 1:p.state = "Voted";break;
                    case 2:p.state = "Payed";break;
                    case 3:p.state = "ReadyForDownload";break;
                    case 4:p.state = "Closed";break;
                }
                database.write(p,t.TransactionID,mt_db);
            }
        });
    },
    onclick:function (id) {
        switch (id) {
            case "d":main.show('d');break;
            case "t":main.show('t');break;
            case "p":main.show('p');break;
            case "logout":window.location.href = "index.html";break;
            case "dl_buy":main.buy();break;
            case "dl_search":alert("Please input your search criteria.");break;
            case "show_proofs_name":
                document.getElementById("pub_proofs_show").innerHTML = "";
                let file = document.getElementById("pub_proofs").files;
                for(let i = 0 ;i<file.length;i++){
                    document.getElementById("pub_proofs_show").innerHTML += file[i].name+"\r\n";
                }break;
            case "pub_submit":main.publish();break;
        }
    },
    show:function (c) {
        document.getElementById("show_datalist").style.display = "none";
        document.getElementById("show_transaction").style.display = "none";
        document.getElementById("show_publish").style.display = "none";
        switch (c) {
            case 'd':document.getElementById("show_datalist").style.display = "block";dl_table.init();break;
            case 't':document.getElementById("show_transaction").style.display = "block";mt_table.init();break;
            case 'p':document.getElementById("show_publish").style.display = "block";break;
            default: alert("Illegal command " + c);break;
        }
    },
    buy:function () {
        let cbs = document.getElementsByName("dl_checkboxes");
        let purchase = [];
        for(let i = 0; i < cbs.length; i++){
            if(cbs[i].checked)
                purchase.push(cbs[i].value);
        }
        astilectron.sendMessage({Name:"buy",Payload:{buyer:location.search.split("=")[1],ids:purchase}},
            function (message) {
                if (message.payload) {
                    main.getTransaction();
                    console.log("Buy data succeed.");
                }else {
                    alert("Buy data failed.");
                }
        });
    },
    publish:function() {
        let elements = document.getElementsByClassName("right-publish-input");
        let values = [];
        for (let i=0;i<elements.length-2;i++){
            values.push(elements[i].value);
        }
        values[4] = [];
        let reader = new FileReader();
        reader.readAsDataURL(elements[4].files[0]);
        reader.onload = function (evt) {
            values[4].push(evt.target.result);
        };
        values[5] = [];
        for (let i=0;i<elements[5].files.length;i++) {
            let reader = new FileReader();
            reader.readAsDataURL(elements[5].files[i]);
            reader.onload = function (evt) {
                values[5].push(evt.target.result);
            }
        }
        let publish = {};
        publish.id="Qm462";// ID需要通过接口调用获取，这里先给测试数据，后面再调试
        publish.title=values[0];publish.price=parseInt(values[1]);publish.keys=values[2];publish.description=values[3];
        publish.data=values[4];publish.proofs=values[5];publish.owner=location.search.split("=")[1];
        let p = {};
        p.title=publish.title;p.price=publish.price;p.keys=publish.keys;p.description=publish.description;p.owner=publish.owner;
        astilectron.sendMessage({Name:"publish",Payload:publish}, function (message) {
            if (message.payload) {
                database.write(p,publish.id,dl_db);
                main.insDL(p,publish.id);
                console.log("Publish data succeed.");
            }else {
                alert("Publish data failed.");
            }
        });
    },
    insDL:function (dl,id) {
        let row = document.getElementById("dl_table").insertRow(1);
        row.insertCell(0).innerHTML =
            "<label style='width: 5%'><input type='checkbox' name='dl_checkboxes' id='1'/></label>";
        document.getElementById("1").value = id;
        row.insertCell(1).innerHTML = dl.title;
        row.insertCell(2).innerHTML = dl.price;
        row.insertCell(3).innerHTML = dl.keys;
        row.insertCell(4).innerHTML = dl.description;
        row.insertCell(5).innerHTML = dl.owner;
    },
    insMT:function (mt,id) {
        let row = document.getElementById("trans_table").insertRow(1);
        row.insertCell(0).innerHTML = mt.title;
        row.insertCell(1).innerHTML = id;
        row.insertCell(2).innerHTML = mt.seller;
        row.insertCell(3).innerHTML = mt.buyer;
        row.insertCell(4).innerHTML = mt.state;
    },
};

let dl_db = {
    init:function(){
        console.log("node:init");
        this.db_name = "Database";
        this.db_version = "1";
        this.db_store_name = "datalist";

        let request = indexedDB.open(this.db_name,this.db_version);
        request.onerror = function(event) {
            alert("open failed with error code: " + event.target.errorCode);
        };
        request.onupgradeneeded = function(event) {
            this.db = event.target.result;
            this.db.createObjectStore(dl_db.db_store_name);
            this.db.createObjectStore("transaction");
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
    },
};
let mt_db = {
    init:function () {
        this.db_name = "Database";
        this.db_version = "1";
        this.db_store_name = "transaction";

        let request = indexedDB.open(this.db_name,this.db_version);
        request.onerror = function(event) {
            alert("open failed with error code: " + event.target.errorCode);
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
    },
};
let database = {
    init:function() {
        dl_db.init();
        mt_db.init();
    },
    write:function(params, key,DB) {
        let transaction = DB.db.transaction(DB.db_store_name, "readwrite");
        let store = transaction.objectStore(DB.db_store_name);
        let request = store.put(params,key);
        request.onerror = function(event){
            console.log(event);
        };
    },
    deleteItem:function(key,DB) {
        DB.db.transaction(DB.db_store_name, "readwrite").objectStore(DB.db_store_name).delete(key);
    },
    select:function(key,DB) {
        let transaction = DB.db.transaction(DB.db_store_name,"readwrite");
        let store = transaction.objectStore(DB.db_store_name);
        let request;
        if (key) {
            request = store.get(key);
        }else {
            request = store.getAll();
        }
        request.onsuccess = function () {
            return request.result;
        };
    },
    clear:function(DB) {
        DB.db.transaction(DB.db_store_name,"readwrite").objectStore(DB.db_store_name).clear();
    },
};

let dl_table = {
    init: function () {
        dl_table.tableData = document.getElementById("dl_table");
        dl_table.firstSpan = document.getElementById("dl_First");
        dl_table.preSpan = document.getElementById("dl_Pre");
        dl_table.nextSpan = document.getElementById("dl_Next");
        dl_table.lastSpan = document.getElementById("dl_Last");
        dl_table.pageNumSpan = document.getElementById("dl_TotalPage");
        dl_table.currPageSpan = document.getElementById("dl_PageNum");

        dl_table.pageCount = 5;
        dl_table.numCount = dl_table.tableData.rows.length - 1;
        dl_table.pageNum = parseInt(dl_table.numCount / dl_table.pageCount);
        if (dl_table.numCount % dl_table.pageCount !== 0) {
            dl_table.pageNum += 1;
        }
        dl_table.showPage(1);
    },
};
dl_table.showPage=function(page) {
    for(let i = 1; i < dl_table.numCount + 1; i ++){
        dl_table.tableData.rows[i].style.display = "none";
    }
    dl_table.currPageNum = page;
    dl_table.currPageSpan.innerHTML = dl_table.currPageNum;
    dl_table.pageNumSpan.innerHTML = dl_table.pageNum;
    let firstR = dl_table.pageCount*(dl_table.currPageNum - 1) + 1;
    let lastR = (firstR + dl_table.pageCount)<=(dl_table.numCount + 1) ? (firstR + dl_table.pageCount) : (dl_table.numCount + 1);
    for(let i = firstR; i < lastR; i++){
        dl_table.tableData.rows[i].style.display = "";
    }
    if(1 === dl_table.currPageNum){
        aheadText();
        hinderLink();
    }else if(dl_table.pageNum === dl_table.currPageNum){
        aheadLink();
        hinderText();
    }else{
        aheadLink();
        hinderLink();
    }
    function aheadText() {dl_table.firstSpan.innerHTML = "First";dl_table.preSpan.innerHTML = "Pre";}
    function aheadLink() {
        dl_table.firstSpan.innerHTML = "<a href='javascript:dl_table.showPage(1);'>First</a>";
        dl_table.preSpan.innerHTML = "<a href='javascript:dl_table.showPage(dl_table.currPageNum-1);'>Pre</a>";
    }
    function hinderText() {dl_table.nextSpan.innerHTML = "Next";dl_table.lastSpan.innerHTML = "Last";}
    function hinderLink() {
        dl_table.nextSpan.innerHTML = "<a href='javascript:dl_table.showPage(dl_table.currPageNum+1);'>Next</a>";
        dl_table.lastSpan.innerHTML = "<a href='javascript:dl_table.showPage(dl_table.pageNum)'>Last</a>";
    }
};
let mt_table = {
    init:function () {
        mt_table.tableData = document.getElementById("trans_table");
        mt_table.firstSpan = document.getElementById("mt_First");
        mt_table.preSpan = document.getElementById("mt_Pre");
        mt_table.nextSpan = document.getElementById("mt_Next");
        mt_table.lastSpan = document.getElementById("mt_Last");
        mt_table.pageNumSpan = document.getElementById("mt_TotalPage");
        mt_table.currPageSpan = document.getElementById("mt_PageNum");

        mt_table.pageCount = 5;
        mt_table.numCount = mt_table.tableData.rows.length - 1;
        mt_table.pageNum = parseInt(mt_table.numCount / mt_table.pageCount);
        if (mt_table.numCount % mt_table.pageCount !== 0) {
            mt_table.pageNum += 1;
        }
        mt_table.showPage(1);
    },
};
mt_table.showPage=function(page) {
    for(let i = 1; i < mt_table.numCount + 1; i ++){
        mt_table.tableData.rows[i].style.display = "none";
    }
    mt_table.currPageNum = page;
    mt_table.currPageSpan.innerHTML = mt_table.currPageNum;
    mt_table.pageNumSpan.innerHTML = mt_table.pageNum;
    let firstR = mt_table.pageCount*(mt_table.currPageNum - 1) + 1;
    let lastR = (firstR + mt_table.pageCount)<=(mt_table.numCount + 1) ? (firstR + mt_table.pageCount) : (mt_table.numCount + 1);
    for(let i = firstR; i < lastR; i++){
        mt_table.tableData.rows[i].style.display = "";
    }
    if(1 === mt_table.currPageNum){
        aheadText();
        hinderLink();
    }else if(mt_table.pageNum === mt_table.currPageNum){
        aheadLink();
        hinderText();
    }else{
        aheadLink();
        hinderLink();
    }
    function aheadText() {mt_table.firstSpan.innerHTML = "First";mt_table.preSpan.innerHTML = "Pre";}
    function aheadLink() {
        mt_table.firstSpan.innerHTML = "<a href='javascript:mt_table.showPage(1);'>First</a>";
        mt_table.preSpan.innerHTML = "<a href='javascript:mt_table.showPage(mt_table.currPageNum-1);'>Pre</a>";
    }
    function hinderText() {mt_table.nextSpan.innerHTML = "Next";mt_table.lastSpan.innerHTML = "Last";}
    function hinderLink() {
        mt_table.nextSpan.innerHTML = "<a href='javascript:mt_table.showPage(mt_table.currPageNum+1);'>Next</a>";
        mt_table.lastSpan.innerHTML = "<a href='javascript:mt_table.showPage(mt_table.pageNum)'>Last</a>";
    }
};
