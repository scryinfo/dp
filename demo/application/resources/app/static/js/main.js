let main = {
    init:function () {
        // Init

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // -
            main.getUser();
            main.getDatalist();
            main.getTransaction();
        });
    },
    getUser:function () {
        let acc = location.search.split("=")[1];
        document.getElementById("account").innerHTML = acc;
    },
    getDatalist:function () {
        astilectron.sendMessage({Name:"get.datalist",Payload:""},function (message) {
            for (let i=0;i<message.payload.length;i++) {
                main.insDL(message.payload[i]);
            }
        });
    },
    getTransaction:function () {
        astilectron.sendMessage({Name:"get.transaction",Payload:""},function (message) {
            for (let i=0;i<message.payload.length;i++) {
                main.insMT(message.payload[i]);
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
                // 购买数据 func buy(all ids) (bool) {}
                if (true) {
                    alert("Buy data succeed.");
                }else {
                    alert("Buy data failed.");
                }
                break;
            case "dl_search":alert("Please input your search criteria.");break;
            case "insDL":main.insDL();break;
            case "insMT":main.insMT();break;
            case "show":
                let file;
                for(let i = 0 ;i<document.getElementById("pub_proofs").files.length;i++){
                    file=document.getElementById("pub_proofs").files[i];
                    document.getElementById("pub_proofs_show").innerHTML += file.name + " ";
                }
                break;
            case "pub_submit":
                // 发布新数据 func publish(all details) (bool) {}
                if (true) {
                    alert("Publish new data succeed.");
                }else {
                    alert("Publish new data failed.");
                }
                break;
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
    insDL:function (dl) {
        let row = document.getElementById("dl_table").insertRow(1);
        row.insertCell(0).innerHTML =
            "<label style='width: 5%'><input type='checkbox' name='dl_checkboxes' id='1'/></label>";
        document.getElementById("1").value = dl.ID;
        row.insertCell(1).innerHTML = dl.Title;
        row.insertCell(2).innerHTML = dl.Price;
        row.insertCell(3).innerHTML = dl.Keys;
        row.insertCell(4).innerHTML = dl.Description;
        row.insertCell(5).innerHTML = dl.Owner;
    },
    insMT:function (mt) {
        switch (mt.State) {
            case '0':mt.State = "Created";break;
            case '1':mt.State = "Voted";break;
            case '2':mt.State = "Payed";break;
            case '3':mt.State = "ReadyForDownload";break;
            case '4':mt.State = "Closed";break;
        }
        let row = document.getElementById("trans_table").insertRow(1);
        row.insertCell(0).innerHTML = mt.Title;
        row.insertCell(1).innerHTML = mt.TransactionID;
        row.insertCell(2).innerHTML = mt.Seller;
        row.insertCell(3).innerHTML = mt.Buyer;
        row.insertCell(4).innerHTML = mt.State;
    },
};