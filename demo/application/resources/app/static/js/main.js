let main = {
    init:function () {
        // Init

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // -
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
    insDL:function () {
        let row = document.getElementById("dl_table").insertRow(1);
        row.insertCell(0).innerHTML = "<label style='width: 5%'><input type='checkbox' /></label>";
        row.insertCell(1).innerHTML = "test1";
        row.insertCell(2).innerHTML = "test2";
        row.insertCell(3).innerHTML = "test3";
        row.insertCell(4).innerHTML = "test4";
        row.insertCell(5).innerHTML = "test5";
    },
    insMT:function () {
        let row = document.getElementById("trans_table").insertRow(1);
        row.insertCell(0).innerHTML = "test1";
        row.insertCell(1).innerHTML = "test2";
        row.insertCell(2).innerHTML = "test3";
        row.insertCell(3).innerHTML = "test4";
        row.insertCell(4).innerHTML = "test5";
    },
};