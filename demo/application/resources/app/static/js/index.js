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
                document.getElementById("button").innerHTML =
                    `<button class="right-button" id="submit_login">Submit</button>`;
                document.getElementById("submit_login").onclick = function () {
                    astilectron.sendMessage({Name:"login.verify",Payload:
                            {account:document.getElementById("accounts").value ,
                             password:document.getElementById("password").value}},function (message) {
                        if (message.payload) {
                            window.location.href = "main.html";
                        }else {
                            alert("account or password is wrong.");
                        }
                    });
                };break;
            case "new_account":
                index.prepare("New");
                document.getElementById("button").innerHTML =
                    `<button class="right-button" id="submit_new">Submit</button>`;
                document.getElementById("submit_new").onclick = function () {
                    document.getElementById("show").style.display = "none";
                    document.getElementById("show_new").style.display = "block";
                };break;
            case "submit_keystore":
                astilectron.sendMessage({Name:"save.keystroe",Payload:
                        {account:document.getElementById("new_account").innerHTML ,
                         password:document.getElementById("password").value}},function (message) {
                    if (message.payload) {
                        window.location.href = "main.html";
                    }else {
                        alert("account or password is wrong.");
                    }
                });
                window.location.href = "main.html";break;
            case "back":document.getElementById("show").style.display = "none";break;
            case "back_new":document.getElementById("show_new").style.display = "none";break;
        }
    },
    prepare:function (describe) {
        document.getElementById("show_new").style.display = "none";
        document.getElementById("show").style.display = "block";
        document.getElementById("describe").innerHTML = describe;
    },
};