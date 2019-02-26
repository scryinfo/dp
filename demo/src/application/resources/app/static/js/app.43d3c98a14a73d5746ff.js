webpackJsonp([1],{100:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=a(48);e.default={name:"app",data:function(){return{}},methods:{c:function(){o.asticode.modaler.close()}}}},101:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"notFound.vue",data:function(){return{count:5}},created:function(){var t=this,e=window.setInterval(function(){--t.count<0&&(window.clearInterval(e),t.count=5,t.$router.push({path:"/"}))},1e3)}}},102:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=a(49),n=a(25);e.default={name:"Home.vue",data:function(){return{acc:""}},methods:{logout:function(){var t=this;this.$confirm("Are you sure to logout?","Tips:",{confirmButtonText:"Yes",cancelButtonText:"No",type:"warning"}).then(function(){t.$router.push("/")})}},created:function(){this.acc=this.$route.params.acc;var t=this;n.dl_db.init(this),n.mt_db.init(this),document.addEventListener("astilectron-ready",function(){o.lfg.listen(),n.options.init(t)})}}},103:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=a(49),n=a(25);e.default={name:"login",data:function(){return{account:"",password:"",showControl1:!1,showControl2:!1,buttonControl:!0,describe:""}},methods:{right:function(t){switch(this.showControl1=!0,this.showControl2=!1,t){case"Login":this.buttonControl=!0;break;case"New":this.buttonControl=!1}this.describe=t+":"},hide:function(){this.showControl1=!1,this.showControl2=!1,this.password=""},submit_login:function(){var t=this.password;this.password="";var e=this;astilectron.sendMessage({Name:"login.verify",Payload:{account:this.account,password:t}},function(t){t.payload?e.$router.push({name:"home",params:{acc:e.account}}):alert("account or password is wrong.")})},submit_new:function(){var t=this;astilectron.sendMessage({Name:"create.new.account",Payload:""},function(e){t.account=e.payload,t.showControl1=!1,t.showControl2=!0})},submit_keystore:function(){var t=this.password;this.password="";var e=this;astilectron.sendMessage({Name:"save.keystore",Payload:{account:this.$store.state.account,password:t}},function(t){t.payload?e.$router.push({name:"home",params:{acc:e.account}}):alert("save account information failed.")})}},created:function(){this.password="",this.describe="",this.account="";var t=this;n.dl_db.init(this),n.mt_db.init(this),document.addEventListener("astilectron-ready",function(){o.lfg.listen(),astilectron.sendMessage({Name:"get.accounts",Payload:""},function(e){for(var a=0;a<e.payload.length;a++)t.$store.state.accounts.push({address:e.payload[a]})}),n.options.init(t)})}}},104:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"DataList.vue",data:function(){return{selectsDL:[]}},methods:{buy:function(){astilectron.sendMessage({Name:"buy",Payload:{buyer:this.account,ids:this.selectsDL}},function(t){t.payload?console.log("Buy data succeed."):alert("Buy data failed.")})},selectedChange:function(t){this.selectsDL=t}}}},105:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={name:"MyTransaction.vue",data:function(){return{}}}},106:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=a(25);e.default={name:"PublishNewData",data:function(){return{pubData:{Title:"",Price:0,Keys:"",Description:"",Data:"",Proofs:[],Owner:this.$store.state.account}}},methods:{selFile:function(){var t=this.$refs.seldata,e=t.$refs.input.files[0],a=new FileReader,o=this;a.readAsDataURL(e),a.onload=function(t){o.pubData.Data=t.target.result}},selFiles:function(){for(var t=this.$refs.selproof,e=t.$refs.input.files,a=this,o=0;o<e.length;o++){var n=new FileReader;n.readAsDataURL(e[o]),n.onload=function(t){a.pubData.Proofs.push(t.target.result)}}},pub:function(){var t=this,e={Title:this.pubData.Title,Price:this.pubData.Price,Keys:this.pubData.Keys,Description:this.pubData.Description,Owner:this.pubData.Owner};astilectron.sendMessage({Name:"publish",Payload:this.pubData},function(a){null==a.err?(o.dl_db.write(e,a.payload),o.dl_db.init(t),console.log("Log: Publish new data success.")):alert("Publish data failed: ",a.payload)})}}}},107:function(t,e,a){"use strict";function o(t){return t&&t.__esModule?t:{default:t}}var n=a(2),s=o(n),i=a(75),r=o(i),l=a(73),c=o(l);a(74);var u=a(76),d=o(u),p=a(72),f=o(p),v=a(46),h=o(v),b=a(71),m=o(b);s.default.use(c.default),s.default.use(d.default),s.default.use(h.default);var _=new d.default({routes:m.default});new s.default({router:_,store:f.default,render:function(t){return t(r.default)}}).$mount("#app")},159:function(t,e){},160:function(t,e){},161:function(t,e){},162:function(t,e){},163:function(t,e){},164:function(t,e){},165:function(t,e){},175:function(t,e,a){a(160);var o=a(7)(a(101),a(182),null,null);t.exports=o.exports},176:function(t,e,a){a(159);var o=a(7)(a(102),a(181),null,null);t.exports=o.exports},177:function(t,e,a){a(165);var o=a(7)(a(103),a(187),null,null);t.exports=o.exports},178:function(t,e,a){a(161);var o=a(7)(a(104),a(183),"data-v-11faa61c",null);t.exports=o.exports},179:function(t,e,a){a(163);var o=a(7)(a(105),a(185),"data-v-630efc8e",null);t.exports=o.exports},180:function(t,e,a){a(164);var o=a(7)(a(106),a(186),"data-v-64bef98f",null);t.exports=o.exports},181:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",[a("el-col",{staticClass:"top",attrs:{span:24}},[a("el-col",{attrs:{span:20}},[t._v("My Astilectron demo")]),t._v(" "),a("el-col",{attrs:{span:4}},[a("el-dropdown",{staticClass:"top-dropdown"},[a("span",[t._v(t._s(t.acc))]),t._v(" "),a("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[a("el-dropdown-item",[t._v("Settings")]),t._v(" "),a("el-dropdown-item",{attrs:{divided:""},nativeOn:{click:function(e){return t.logout(e)}}},[t._v("Logout")])],1)],1)],1)],1),t._v(" "),a("el-col",{attrs:{span:24}},[a("el-col",{attrs:{span:4}},[a("aside",{staticClass:"aside"},[a("el-menu",{attrs:{"default-active":t.$route.path,"unique-opened":"",router:""}},[t._l(t.$router.options.routes,function(e){return t._l(e.children,function(e){return e.hidden?t._e():a("el-menu-item",{key:e.path,staticClass:"el-menu-item",attrs:{index:e.path}},[t._v(t._s(e.name))])})})],2)],1)]),t._v(" "),a("el-col",{attrs:{span:20}},[a("section",{staticClass:"section"},[a("div",[a("el-col",{attrs:{span:24}},[a("router-view")],1)],1)])])],1)],1)],1)},staticRenderFns:[]}},182:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("p",{staticClass:"page-container"},[t._v("\n    404: WebPage Not Found"),a("br"),t._v("\n    Redirect to login page after "+t._s(t.count)+" seconds.\n")])},staticRenderFns:[]}},183:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",[a("el-col",{staticStyle:{"padding-bottom":"0","background-color":"lightgrey"},attrs:{span:24}},[a("el-button",{on:{click:t.buy}},[t._v("Buy")])],1),t._v(" "),a("el-table",{attrs:{data:this.$store.state.datalist,"highlight-current-row":"",border:"",height:"400"},on:{"selection-change":t.selectedChange}},[a("el-table-column",{attrs:{type:"selection",width:"50"}}),t._v(" "),a("el-table-column",{attrs:{prop:"Title",label:"Title","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Price",label:"Price","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Keys",label:"Keys","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Description",label:"Description","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Owner",label:"Owner","show-overflow-tooltip":""}})],1)],1)},staticRenderFns:[]}},184:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("router-view"),t._v(" "),t._m(0),t._v(" "),a("div",{staticClass:"astimodaler",attrs:{id:"astimodaler"}},[a("div",{staticClass:"astimodaler-background"}),t._v(" "),a("div",{staticClass:"astimodaler-table"},[a("div",{staticClass:"astimodaler-wrapper"},[a("div",{staticClass:"astimodaler-body"},[a("i",{staticClass:"fa fa-close astimodaler-close",on:{click:t.c}}),t._v(" "),a("div",{attrs:{id:"astimodaler-content"}})])])])]),t._v(" "),a("div",{staticClass:"astinotifier",attrs:{id:"astinotifier"}})],1)},staticRenderFns:[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"astiloader",attrs:{id:"astiloader"}},[a("div",{staticClass:"astiloader-background"}),t._v(" "),a("div",{staticClass:"astiloader-table"},[a("div",{staticClass:"astiloader-content"},[a("i",{staticClass:"fa fa-spinner fa-spin fa-3x fa-fw"})])])])}]}},185:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("section",[a("el-col",{staticStyle:{"padding-bottom":"0","background-color":"lightgrey"},attrs:{span:24}},[t._v("\n        tool bar.\n    ")]),t._v(" "),a("el-table",{attrs:{data:this.$store.state.mytransaction,"highlight-current-row":"",border:"",height:"400"}},[a("el-table-column",{attrs:{type:"selection",width:"50"}}),t._v(" "),a("el-table-column",{attrs:{prop:"Title",label:"Title","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"TransactionID",label:"TransactionID","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Seller",label:"Seller","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"Buyer",label:"Buyer","show-overflow-tooltip":""}}),t._v(" "),a("el-table-column",{attrs:{prop:"State",label:"State","show-overflow-tooltip":""}})],1)],1)},staticRenderFns:[]}},186:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-form",{staticClass:"pubForm",attrs:{model:t.pubData,"label-width":"15%"}},[a("el-form-item",{attrs:{label:"Title"}},[a("el-input",{model:{value:t.pubData.Title,callback:function(e){t.$set(t.pubData,"Title",e)},expression:"pubData.Title"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"Price"}},[a("el-input",{attrs:{placeholder:"Unit is DDD"},model:{value:t.pubData.Price,callback:function(e){t.$set(t.pubData,"Price",e)},expression:"pubData.Price"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"Keys"}},[a("el-input",{attrs:{placeholder:"Separate each tag with a comma or semicolon",type:"textarea",rows:2},model:{value:t.pubData.Keys,callback:function(e){t.$set(t.pubData,"Keys",e)},expression:"pubData.Keys"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"Description"}},[a("el-input",{attrs:{type:"textarea",rows:3},model:{value:t.pubData.Description,callback:function(e){t.$set(t.pubData,"Description",e)},expression:"pubData.Description"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"Data"}},[a("el-input",{ref:"seldata",attrs:{type:"file"},on:{change:t.selFile}})],1),t._v(" "),a("el-form-item",{attrs:{label:"Proofs"}},[a("el-input",{ref:"selproof",attrs:{type:"file",multiple:""},on:{change:t.selFiles}})],1),t._v(" "),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.pub}},[t._v("Publish")])],1)],1)},staticRenderFns:[]}},187:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",{staticClass:"row"},[a("el-col",{attrs:{span:24}},[a("div",{staticClass:"top"},[t._v("My Astilectron demo")])]),t._v(" "),a("el-col",{attrs:{span:8}},[a("div",{staticClass:"left"},[a("div",{staticClass:"left-explain"},[t._v("select account:")]),t._v(" "),a("el-select",{staticClass:"left-account",attrs:{placeholder:"select account"},model:{value:t.account,callback:function(e){t.account=e},expression:"account"}},t._l(this.$store.state.accounts,function(t){return a("el-option",{key:t.address,attrs:{value:t.address,label:t.address}})}),1),t._v(" "),a("div",[a("button",{staticClass:"left-button",on:{click:function(e){return t.right("Login")}}},[t._v("Login")])]),t._v(" "),a("div",[a("button",{staticClass:"left-button",on:{click:function(e){return t.right("New")}}},[t._v("Create New Account")])])],1)]),t._v(" "),a("el-col",{attrs:{span:16}},[t.showControl1?a("div",{staticClass:"right",attrs:{id:"show"}},[a("div",{staticClass:"right-show"},[t._v(t._s(t.describe)+"\n                    "),a("el-input",{staticClass:"right-pwd",attrs:{placeholder:"password",type:"password",clearable:!0},model:{value:t.password,callback:function(e){t.password=e},expression:"password"}})],1),t._v(" "),a("div",[a("button",{staticClass:"right-button",on:{click:t.hide}},[t._v("Back")]),t._v(" "),t.buttonControl?a("button",{staticClass:"right-button",on:{click:t.submit_login}},[t._v("Submit")]):t._e(),t._v(" "),t.buttonControl?t._e():a("button",{staticClass:"right-button",on:{click:t.submit_new}},[t._v("Submit")])])]):t._e(),t._v(" "),t.showControl2?a("div",{staticClass:"right",attrs:{id:"show_new"}},[a("div",[t._v("Your account is created :  "+t._s(t.account)),a("br"),t._v("\n                    account information will saves at local :  keystore"),a("br"),t._v("\n                    please keep it properly."),a("br"),a("hr"),t._v("Do you want to save and login with this account?")]),t._v(" "),a("div",{staticClass:"right-pwd"},[a("button",{staticClass:"right-button",on:{click:t.hide}},[t._v("No")]),t._v(" "),a("button",{staticClass:"right-button",on:{click:t.submit_keystore}},[t._v("Yes")])])]):t._e()])],1)],1)},staticRenderFns:[]}},25:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o={init:function(t){o.getDatalist(t),o.getTransaction(t)},getDatalist:function(t){astilectron.sendMessage({Name:"get.datalist",Payload:""},function(e){for(var a=0;a<e.payload.length;a++){var o=e.payload[a],s={};s.Title=o.Title,s.Price=o.Price,s.Keys=o.Keys,s.Description=o.Description,s.Owner=o.Owner,n.write(s,o.ID)}n.init(t)})},getTransaction:function(t){astilectron.sendMessage({Name:"get.transaction",Payload:t.$store.state.account},function(e){for(var a=0;a<e.payload.length;a++){var o=e.payload[a],n={};n.Title=o.Title,n.Seller=o.Seller,n.Buyer=o.Buyer,n.State=o.State,s.write(n,o.TransactionID)}s.init(t)})}},n={init:function(t){this.db_name="Database",this.db_version="1",this.db_store_name="datalist";var e=indexedDB.open(this.db_name,this.db_version);e.onerror=function(t){alert("open failed with error code: "+t.target.errorCode)},e.onupgradeneeded=function(t){this.db=t.target.result,this.db.createObjectStore(n.db_store_name),this.db.createObjectStore("transaction")},e.onsuccess=function(e){t.$store.state.datalist=[],n.db=e.target.result,n.db.transaction(n.db_store_name,"readonly").objectStore(n.db_store_name).openCursor().onsuccess=function(e){var a=e.target.result;if(a){var o=a.value;o.ID=a.key,t.$store.dispatch("addDL",o),a.continue()}}}},write:function(t,e){n.db.transaction(n.db_store_name,"readwrite").objectStore(n.db_store_name).put(t,e).onerror=function(t){console.log(t)}}},s={init:function(t){this.db_name="Database",this.db_version="1",this.db_store_name="transaction";var e=indexedDB.open(this.db_name,this.db_version);e.onerror=function(t){alert("open failed with error code: "+t.target.errorCode)},e.onsuccess=function(e){t.$store.state.mytransaction=[],s.db=e.target.result,s.db.transaction(s.db_store_name,"readonly").objectStore(s.db_store_name).openCursor().onsuccess=function(e){var a=e.target.result;if(a){var o=a.value;switch(o.TransactionID=a.key,parseInt(a.State)){case 0:o.State="Created";break;case 1:o.State="Voted";break;case 2:o.State="Payed";break;case 3:o.State="ReadyForDownload";break;case 4:o.State="Closed"}t.$store.dispatch("addMT",o),a.continue()}}}},write:function(t,e){s.db.transaction(s.db_store_name,"readwrite").objectStore(s.db_store_name).put(t,e).onerror=function(t){console.log(t)}}};e.dl_db=n,e.mt_db=s,e.options=o},48:function(t,e,a){"use strict";if(Object.defineProperty(e,"__esModule",{value:!0}),void 0===o)var o={};o.loader={hide:function(){document.getElementById("astiloader").style.display="none"},show:function(){document.getElementById("astiloader").style.display="block"}},o.modaler={close:function(){void 0!==o.modaler.onclose&&null!==o.modaler.onclose&&o.modaler.onclose(),o.modaler.hide()},hide:function(){document.getElementById("astimodaler").style.display="none"},setContent:function(t){document.getElementById("astimodaler-content").innerHTML="",document.getElementById("astimodaler-content").appendChild(t)},show:function(){document.getElementById("astimodaler").style.display="block"}},o.notifier={error:function(t){this.notify("error",t)},info:function(t){this.notify("info",t)},notify:function(t,e){var a=document.createElement("div");a.className="astinotifier-wrapper";var o=document.createElement("div");o.className="astinotifier-item "+t;var n=document.createElement("div");n.className="astinotifier-label",n.innerHTML=e;var s=document.createElement("div");s.className="astinotifier-close",s.innerHTML='<i class="fa fa-close"></i>',s.onclick=function(){a.remove()},o.appendChild(n),o.appendChild(s),a.appendChild(o),document.getElementById("astinotifier").prepend(a),setTimeout(function(){s.click()},5e3)},success:function(t){this.notify("success",t)},warning:function(t){this.notify("warning",t)}},e.asticode=o},49:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.lfg=void 0;var o=a(48),n={listen:function(){astilectron.onMessage(function(t){switch(t.name){case"about":case"about2":return n.about(t.payload),{payload:"payload"};case"welcome":o.asticode.notifier.info(t.payload)}})},about:function(t){var e=document.createElement("div");e.innerHTML=t,o.asticode.modaler.setContent(e),o.asticode.modaler.show()}};e.lfg=n},71:function(t,e,a){"use strict";function o(t){return t&&t.__esModule?t:{default:t}}Object.defineProperty(e,"__esModule",{value:!0});var n=a(175),s=o(n),i=a(177),r=o(i),l=a(176),c=o(l),u=a(178),d=o(u),p=a(179),f=o(p),v=a(180),h=o(v),b=[{path:"/",component:r.default,name:"login",hidden:!0},{path:"/404",component:s.default,name:"not found",hidden:!0},{path:"/home",component:c.default,name:"home",children:[{path:"/dl",component:d.default,name:"Data list"},{path:"/mt",component:f.default,name:"My transaction"},{path:"/pd",component:h.default,name:"Publish new data"}]},{path:"*",redirect:{path:"/404"},hidden:!0}];e.default=b},72:function(t,e,a){"use strict";function o(t){return t&&t.__esModule?t:{default:t}}Object.defineProperty(e,"__esModule",{value:!0});var n=a(2),s=o(n),i=a(46),r=o(i);s.default.use(r.default);var l={datalist:[],mytransaction:[],accounts:[{address:""}]},c=new r.default.Store({state:l});e.default=c},74:function(t,e){},75:function(t,e,a){a(162);var o=a(7)(a(100),a(184),null,null);t.exports=o.exports}},[107]);
//# sourceMappingURL=app.43d3c98a14a73d5746ff.js.map