let main = {
    init:function () {
        // Init

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // init account and lists
            document.getElementById("account").innerHTML = location.search.split("=")[1];
            main.getDatalist();
            main.getTransaction();
        });
    },
    getDatalist:function () {
        // 把当前最新的一条数据的ID传过去，go只读取和发送js没有的部分
        astilectron.sendMessage({Name:"get.datalist",Payload:""},function (message) {
            for (let i=0;i<message.payload.length;i++) {
                let t = message.payload[i];
                let p = {};
                p.title = t.Title;p.price = t.Price;p.keys = t.Keys;p.description = t.Description;p.owner = t.Owner;
                dl_db.write(p,t.ID);
                main.insDL(p,t.ID);
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
                switch (t.State) {
                    case '0':p.state = "Created";break;
                    case '1':p.state = "Voted";break;
                    case '2':p.state = "Payed";break;
                    case '3':p.state = "ReadyForDownload";break;
                    case '4':p.state = "Closed";break;
                }
                mt_db.write(p,t.TransactionID);
                main.insMT(p,t.TransactionID);
            }
        });
    },
    onclick:function (id) {
        switch (id) {
            case "d":main.show('d');break;
            case "t":main.show('t');break;
            case "p":main.show('p');break;
            case "logout":window.location.href = "index.html";break;
            case "dl_buy":
                let cbs = document.getElementsByName("dl_checkboxes");
                let parse = [];
                for(let i = 0; i < cbs.length; i++){
                    if(cbs[i].checked)
                        parse.push(cbs[i].value);
                }
                astilectron.sendMessage({Name:"buy",Payload:{buyer:location.search.split("=")[1],ids:parse}},
                    function (message) {
                    if (message.payload) {
                        main.getTransaction();
                        console.log("Buy data succeed.");
                    }else {
                        alert("Buy data failed.");
                    }
                });break;
            case "dl_search":alert("Please input your search criteria.");break;
            case "insDL":main.insDL();break;
            case "insMT":main.insMT();break;
            case "show":
                let file = document.getElementById("pub_proofs").files;
                document.getElementById("pub_proofs_show").innerHTML = "";
                for(let i = 0 ;i<file.length;i++){
                    document.getElementById("pub_proofs_show").innerHTML += file[i].name+"\r\n";
                }break;
            case "pub_submit":
                let elements = document.getElementsByClassName("right-publish-input");
                let values = [];
                for (let i=0;i<elements.length-2;i++){
                    values.push(elements[i].value);
                }
                let filestring = [];
                let reader = new FileReader();
                reader.readAsDataURL(elements[4].files[0]);
                reader.onload = function (evt) {
                    filestring.push(evt.target.result);
                };
                values[4] = filestring;
                let fileString = [];
                for (let i=0;i<elements[5].files.length;i++) {
                    let reader = new FileReader();
                    reader.readAsDataURL(elements[5].files[i]);
                    reader.onload = function (evt) {
                        fileString.push(evt.target.result);
                    }
                }
                values[5] = fileString;
                let publish = {};
                publish.id="Qm462";// ID需要通过接口调用获取，这里先给测试数据，后面再调试
                publish.title=values[0];publish.price=parseInt(values[1]);publish.keys=values[2];publish.description=values[3];
                publish.data=values[4];publish.proofs=values[5];publish.owner=location.search.split("=")[1];
                astilectron.sendMessage({Name:"publish",Payload:publish}, function (message) {
                    if (message.payload) {
                        dl_db.write({title:publish.title,price:publish.price,keys:publish.keys,
                            description:publish.description,owner:publish.owner},publish.id);
                        console.log("Publish data succeed.");
                    }else {
                        alert("Publish data failed.");
                    }
                });break;
        }
    },
    show:function (c) {
        document.getElementById("show_datalist").style.display = "none";
        document.getElementById("show_transaction").style.display = "none";
        document.getElementById("show_publish").style.display = "none";
        switch (c) {
            case 'd':document.getElementById("show_datalist").style.display = "block";break;
            case 't':document.getElementById("show_transaction").style.display = "block";break;
            case 'p':document.getElementById("show_publish").style.display = "block";break;
            default: alert("Illegal command " + c)
        }
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