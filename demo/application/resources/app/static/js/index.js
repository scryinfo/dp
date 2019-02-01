let index = {
    init:function () {
        // Init

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // -
            index.getAccounts();
        });
    },
    getAccounts:function () {
        astilectron.sendMessage({Name : "get.accounts",Payload: ""},function (message) {
            let str = `<option></option>`;
            for(let i=0;i<message.payload.length;i++){
                str += `<option>`+message.payload[i]+`</option>`;
            }
            document.getElementById("accounts").innerHTML = str;
        });
    },
    onclick:function (id) {
        switch (id) {
            case "login":
                index.prepare("Login");
                document.getElementById("submit_login").onclick = function () {
                    let acc = document.getElementById("accounts").value;
                    astilectron.sendMessage({Name:"login.verify",Payload:
                            {account:acc,
                             password:document.getElementById("password").value}},function (message) {
                        if (message.payload) {
                            window.location.href = "main.html?acc="+acc;
                        }else {
                            alert("account or password is wrong.");
                        }
                    });
                };break;
            case "new_account":
                index.prepare("New");
                document.getElementById("submit_new").onclick = function () {
                    astilectron.sendMessage({Name:"create.new.account",Payload:""},function(message) {
                        document.getElementById("show_new_account").innerHTML = message.payload;
                        document.getElementById("show").style.display = "none";
                        document.getElementById("show_new").style.display = "block";
                    });
                };break;
            case "submit_keystore":
                let acc = document.getElementById("show_new_account").innerHTML;
                astilectron.sendMessage({Name:"save.keystore",Payload:
                        {account: acc,
                         password:document.getElementById("password").value}},function (message) {
                    if (message.payload) {
                        window.location.href = "main.html?acc="+acc;
                    }else {
                        alert("save account information failed.");
                    }
                });break;
            case "back":document.getElementById("show").style.display = "none";break;
            case "back_new":document.getElementById("show_new").style.display = "none";break;
        }
    },
    prepare:function (describe) {
        let bh = "";
        switch (describe) {
            case "Login":bh = `<button class="right-button" id="submit_login">Submit</button>`;break;
            case "New":bh = `<button class="right-button" id="submit_new">Submit</button>`;break;
        }
        document.getElementById("show_new").style.display = "none";
        document.getElementById("show").style.display = "block";
        document.getElementById("describe").innerHTML = describe;
        document.getElementById("button").innerHTML = bh;
    },
};