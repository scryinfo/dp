let index = {
    init:function () {
        // Init

        // wait for ready
        document.addEventListener('astilectron-ready', function() {
            // -
        });
    },
    onclick:function (id) {
        switch (id) {
            case "login":
                index.prepare("Login");
                document.getElementById("button").innerHTML =
                    `<button class="right-button" id="submit_login">Submit</button>`;
                document.getElementById("submit_login").onclick = function () {
                    // 验证用户信息：func send(account,password) (bool) {}
                    if (true) {
                        window.location.href = "main.html";
                    }else {
                        alert("account or password is wrong.");
                    }
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
                // 将新建的账户信息保存到keystore：func send(account information) (bool) {}
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