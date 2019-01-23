let index = {
    listen:function () {
        document.getElementById("login").onclick = function () {
            document.getElementById("show_new").style.display = "none";
            document.getElementById("show").style.display = "block";
            document.getElementById("describe").innerHTML = `Login`;
            document.getElementById("button").innerHTML =
                `<button class="right-button" id="submit_login">Submit</button>`;
            document.getElementById("submit_login").onclick = function () {
                // 验证用户信息：func send(account,password) (bool) {}
                if (true) {
                    index.login();
                }else {
                    alert("account or password is wrong.");
                }
            };
        };
        document.getElementById("new_account").onclick = function () {
            document.getElementById("show_new").style.display = "none";
            document.getElementById("show").style.display = "block";
            document.getElementById("describe").innerHTML = `New`;
            document.getElementById("button").innerHTML =
                `<button class="right-button" id="submit_new">Submit</button>`;
            document.getElementById("submit_new").onclick = function () {
                document.getElementById("show").style.display = "none";
                document.getElementById("show_new").style.display = "block";
                document.getElementById("submit_keystore").onclick = function () {
                    // 将新建的账户信息保存到keystore：func send(account information) (bool) {}
                    index.login();
                };
            };
        };
        document.getElementById("back").onclick = function () {
            document.getElementById("show").style.display = "none";
        };
        document.getElementById("back_new").onclick = function () {
            document.getElementById("show_new").style.display = "none";
        };
    },
    login:function () {
        window.location.href = "main.html";
    },
};