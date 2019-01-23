let main = {
    listen:function () {
        document.getElementById("logout").onclick = function () {
            window.location.href = "index.html";
        };
        document.getElementById("pub_submit").onclick = function () {
            // 发布新数据 func publish(all details) (bool) {}
            if (true) {
                alert("Publish new data succeed.");
            }else {
                alert("Publish new data failed.");
            }
        };
        document.getElementById("dl_buy").onclick = function () {
            // 购买数据 func buy(all ids) (bool) {}
            if (true) {
                alert("Buy data succeed.");
            }else {
                alert("Buy data failed.");
            }
        };
        document.getElementById("dl_search").onclick = function () {
            alert("Please input your search criteria.")
        }
    },
    show:function (c) {
        main.closeAll();
        switch (c) {
            case 'd':document.getElementById("show_datalist").style.display = "block";break;
            case 't':document.getElementById("show_transaction").style.display = "block";break;
            case 'p':document.getElementById("show_publish").style.display = "block";break;
            default: alert("Illegal command " + c)
        }
    },
    closeAll:function () {
        document.getElementById("show_datalist").style.display = "none";
        document.getElementById("show_transaction").style.display = "none";
        document.getElementById("show_publish").style.display = "none";
    },
    reshow:function () {
        let file;
        //取得FileList取得的file集合
        for(let i = 0 ;i<document.getElementById("pub_proofs").files.length;i++){
            //file对象为用户选择的某一个文件
            file=document.getElementById("pub_proofs").files[i];
            //此时取出这个文件进行处理，这里只是显示文件名
            document.getElementById("pub_proofs_show").innerHTML += file.name + " ";
        }
    },
    insDL:function () {
        let row = document.getElementById("dl_table").insertRow(1);
        row.insertCell(0).innerHTML = "<label style='width: 5%'><input type='checkbox' /></label>";
        row.insertCell(1).innerHTML = "test2";
        row.insertCell(2).innerHTML = "test3";
        row.insertCell(3).innerHTML = "test4";
        row.insertCell(4).innerHTML = "test5";
        row.insertCell(5).innerHTML = "test6";
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