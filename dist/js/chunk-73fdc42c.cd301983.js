(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-73fdc42c"],{"129f":function(t,e){t.exports=Object.is||function(t,e){return t===e?0!==t||1/t===1/e:t!=t&&e!=e}},"13d5":function(t,e,r){"use strict";var n=r("23e7"),c=r("d58f").left,a=r("a640"),i=r("ae40"),o=a("reduce"),u=i("reduce",{1:0});n({target:"Array",proto:!0,forced:!o||!u},{reduce:function(t){return c(this,t,arguments.length,arguments.length>1?arguments[1]:void 0)}})},"14c3":function(t,e,r){var n=r("c6b6"),c=r("9263");t.exports=function(t,e){var r=t.exec;if("function"===typeof r){var a=r.call(t,e);if("object"!==typeof a)throw TypeError("RegExp exec method returned something other than an Object or null");return a}if("RegExp"!==n(t))throw TypeError("RegExp#exec called on incompatible receiver");return c.call(t,e)}},"1dde":function(t,e,r){var n=r("d039"),c=r("b622"),a=r("2d00"),i=c("species");t.exports=function(t){return a>=51||!n((function(){var e=[],r=e.constructor={};return r[i]=function(){return{foo:1}},1!==e[t](Boolean).foo}))}},"223a":function(t,e,r){"use strict";var n=r("b03b"),c=r.n(n);c.a},2532:function(t,e,r){"use strict";var n=r("23e7"),c=r("5a34"),a=r("1d80"),i=r("ab13");n({target:"String",proto:!0,forced:!i("includes")},{includes:function(t){return!!~String(a(this)).indexOf(c(t),arguments.length>1?arguments[1]:void 0)}})},"44e7":function(t,e,r){var n=r("861d"),c=r("c6b6"),a=r("b622"),i=a("match");t.exports=function(t){var e;return n(t)&&(void 0!==(e=t[i])?!!e:"RegExp"==c(t))}},"4de4":function(t,e,r){"use strict";var n=r("23e7"),c=r("b727").filter,a=r("1dde"),i=r("ae40"),o=a("filter"),u=i("filter");n({target:"Array",proto:!0,forced:!o||!u},{filter:function(t){return c(this,t,arguments.length>1?arguments[1]:void 0)}})},"5a34":function(t,e,r){var n=r("44e7");t.exports=function(t){if(n(t))throw TypeError("The method doesn't accept regular expressions");return t}},"65f0":function(t,e,r){var n=r("861d"),c=r("e8b5"),a=r("b622"),i=a("species");t.exports=function(t,e){var r;return c(t)&&(r=t.constructor,"function"!=typeof r||r!==Array&&!c(r.prototype)?n(r)&&(r=r[i],null===r&&(r=void 0)):r=void 0),new(void 0===r?Array:r)(0===e?0:e)}},"841c":function(t,e,r){"use strict";var n=r("d784"),c=r("825a"),a=r("1d80"),i=r("129f"),o=r("14c3");n("search",1,(function(t,e,r){return[function(e){var r=a(this),n=void 0==e?void 0:e[t];return void 0!==n?n.call(e,r):new RegExp(e)[t](String(r))},function(t){var n=r(e,t,this);if(n.done)return n.value;var a=c(t),u=String(this),s=a.lastIndex;i(s,0)||(a.lastIndex=0);var d=o(a,u);return i(a.lastIndex,s)||(a.lastIndex=s),null===d?-1:d.index}]}))},9263:function(t,e,r){"use strict";var n=r("ad6d"),c=r("9f7f"),a=RegExp.prototype.exec,i=String.prototype.replace,o=a,u=function(){var t=/a/,e=/b*/g;return a.call(t,"a"),a.call(e,"a"),0!==t.lastIndex||0!==e.lastIndex}(),s=c.UNSUPPORTED_Y||c.BROKEN_CARET,d=void 0!==/()??/.exec("")[1],l=u||d||s;l&&(o=function(t){var e,r,c,o,l=this,f=s&&l.sticky,p=n.call(l),v=l.source,h=0,m=t;return f&&(p=p.replace("y",""),-1===p.indexOf("g")&&(p+="g"),m=String(t).slice(l.lastIndex),l.lastIndex>0&&(!l.multiline||l.multiline&&"\n"!==t[l.lastIndex-1])&&(v="(?: "+v+")",m=" "+m,h++),r=new RegExp("^(?:"+v+")",p)),d&&(r=new RegExp("^"+v+"$(?!\\s)",p)),u&&(e=l.lastIndex),c=a.call(f?r:l,m),f?c?(c.input=c.input.slice(h),c[0]=c[0].slice(h),c.index=l.lastIndex,l.lastIndex+=c[0].length):l.lastIndex=0:u&&c&&(l.lastIndex=l.global?c.index+c[0].length:e),d&&c&&c.length>1&&i.call(c[0],r,(function(){for(o=1;o<arguments.length-2;o++)void 0===arguments[o]&&(c[o]=void 0)})),c}),t.exports=o},"9f7f":function(t,e,r){"use strict";var n=r("d039");function c(t,e){return RegExp(t,e)}e.UNSUPPORTED_Y=n((function(){var t=c("a","y");return t.lastIndex=2,null!=t.exec("abcd")})),e.BROKEN_CARET=n((function(){var t=c("^r","gy");return t.lastIndex=2,null!=t.exec("str")}))},a640:function(t,e,r){"use strict";var n=r("d039");t.exports=function(t,e){var r=[][t];return!!r&&n((function(){r.call(null,e||function(){throw 1},1)}))}},ab13:function(t,e,r){var n=r("b622"),c=n("match");t.exports=function(t){var e=/./;try{"/./"[t](e)}catch(r){try{return e[c]=!1,"/./"[t](e)}catch(n){}}return!1}},ac1f:function(t,e,r){"use strict";var n=r("23e7"),c=r("9263");n({target:"RegExp",proto:!0,forced:/./.exec!==c},{exec:c})},ad6d:function(t,e,r){"use strict";var n=r("825a");t.exports=function(){var t=n(this),e="";return t.global&&(e+="g"),t.ignoreCase&&(e+="i"),t.multiline&&(e+="m"),t.dotAll&&(e+="s"),t.unicode&&(e+="u"),t.sticky&&(e+="y"),e}},ae40:function(t,e,r){var n=r("83ab"),c=r("d039"),a=r("5135"),i=Object.defineProperty,o={},u=function(t){throw t};t.exports=function(t,e){if(a(o,t))return o[t];e||(e={});var r=[][t],s=!!a(e,"ACCESSORS")&&e.ACCESSORS,d=a(e,0)?e[0]:u,l=a(e,1)?e[1]:void 0;return o[t]=!!r&&!c((function(){if(s&&!n)return!0;var t={length:-1};s?i(t,1,{enumerable:!0,get:u}):t[1]=1,r.call(t,d,l)}))}},b03b:function(t,e,r){},b0c0:function(t,e,r){var n=r("83ab"),c=r("9bf2").f,a=Function.prototype,i=a.toString,o=/^\s*function ([^ (]*)/,u="name";n&&!(u in a)&&c(a,u,{configurable:!0,get:function(){try{return i.call(this).match(o)[1]}catch(t){return""}}})},b3b1:function(t,e,r){"use strict";r.r(e);var n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"list row"},[r("md-list",t._l(t.cart,(function(e,n){return r("md-list-item",{key:n,staticClass:"list-group-item"},[r("div",{staticClass:"md-list-item-text"},[r("span",[t._v(t._s(e.name))]),r("p",[t._v("Code "+t._s(e.code))])]),r("div",{staticClass:"md-list-item-text"},[r("span",[t._v("$ "+t._s(e.price))])]),r("div",{staticClass:"md-list-item-text"},[r("span",[t._v("$ "+t._s(e.cartQty))])]),r("div",[r("md-button",{staticClass:"md-icon-button md-list-action",on:{click:function(r){return t.addToCart(e)}}},[r("md-icon",{staticClass:"md"},[t._v("add")])],1),r("md-button",{staticClass:"md-icon-button md-list-action",on:{click:function(r){return t.rmFromCart(e)}}},[r("md-icon",{staticClass:"md"},[t._v("remove")])],1)],1),r("md-button",{staticClass:"md-icon-button md-list-action",on:{click:function(r){return t.dropFromCart(e.id)}}},[r("md-icon",{staticClass:"md"},[t._v("delete")])],1)],1)})),1),r("div",{staticClass:"fucntionBar"},[r("md-field",[r("md-input",{attrs:{placeholder:"byName"===t.serachType?"Search Name":"byPrice"===t.serachType?"Search Price":"Search Code",inputmode:"text"},model:{value:t.search,callback:function(e){t.search=e},expression:"search"}})],1)],1),r("div",{staticClass:"fucntionBar"},[r("b",[t._v("Cart Total : $ HKD "+t._s(t.cartTotal))]),r("md-button",{staticClass:"md-raised clear",on:{click:function(e){return t.clearCart()}}},[t._v("Clear")])],1),r("div",{staticClass:"fucntionBar"},[r("md-radio",{attrs:{value:"byPrice"},model:{value:t.serachType,callback:function(e){t.serachType=e},expression:"serachType"}},[t._v("by price")]),r("md-radio",{attrs:{value:"byName"},model:{value:t.serachType,callback:function(e){t.serachType=e},expression:"serachType"}},[t._v("by name")]),r("md-radio",{attrs:{value:"byCode"},model:{value:t.serachType,callback:function(e){t.serachType=e},expression:"serachType"}},[t._v("by code")])],1),r("md-list",t._l(t.filteredList,(function(e,n){return r("md-list-item",{key:n,staticClass:"list-group-item"},[r("div",{staticClass:"md-list-item-text"},[r("span",[t._v(t._s(e.name))]),r("span",[t._v("$ "+t._s(e.price))]),r("p",[t._v("Code "+t._s(e.code))])]),r("div",{staticClass:"md-list-item-text"},[r("md-button",{staticClass:"md-raised md-primary",on:{click:function(r){return t.addToCart(e)}}},[t._v("Add to Cart")])],1)])})),1)],1)},c=[],a=(r("4de4"),r("c740"),r("caad"),r("13d5"),r("b0c0"),r("ac1f"),r("2532"),r("841c"),r("b91f")),i={name:"Product-list",data:function(){return{search:"",products:[],cart:[],serachType:""}},methods:{retrieveProducts:function(){var t=this;a["a"].getAll().then((function(e){t.products=e.data,console.log(e.data)})).catch((function(e){t.$store.commit("errorMessage",e.message),console.log(e)}))},addToCart:function(t){this.adjustCart(t,"add")},rmFromCart:function(t){t.cartQty>1?this.adjustCart(t,"rm"):this.$store.commit("errorMessage","Cannot reduce product : only 1 left")},dropFromCart:function(t){this.cart=this.cart.filter((function(e){return e.id!==t}))},clearCart:function(){this.cart=[]},adjustCart:function(t,e){console.log({cart:this.cart});var r=this.cart.findIndex((function(e){return e.code===t.code}));return 0===this.cart.length||void 0===this.cart[r]?(t.cartQty=1,void this.cart.push(t)):this.cart.length>0?("add"===e?t.cartQty=this.cart[r].cartQty+1:"rm"===e&&this.cart[r].cartQty>0&&(t.cartQty=this.cart[r].cartQty-1),void this.$set(this.cart,r,t)):void 0}},mounted:function(){this.retrieveProducts()},computed:{filteredList:function(){var t=this;return this.products.filter((function(e){if(""===t.search)return!1;switch(t.serachType){case"byPrice":return String(e.price).includes(t.search.toLowerCase());case"byName":return e.name.toLowerCase().includes(t.search.toLowerCase());case"byCode":return String(e.code).includes(t.search.toLowerCase());default:return String(e.price).includes(t.search.toLowerCase())}}))},cartTotal:function(){if(0===this.cart.length)return 0;var t=this.cart.reduce((function(t,e){return t+e.price*e.cartQty}),0);return Math.round(10*t)/10}}},o=i,u=(r("223a"),r("2877")),s=Object(u["a"])(o,n,c,!1,null,null,null);e["default"]=s.exports},b727:function(t,e,r){var n=r("0366"),c=r("44ad"),a=r("7b0b"),i=r("50c4"),o=r("65f0"),u=[].push,s=function(t){var e=1==t,r=2==t,s=3==t,d=4==t,l=6==t,f=5==t||l;return function(p,v,h,m){for(var x,y,b=a(p),g=c(b),C=n(v,h,3),E=i(g.length),_=0,T=m||o,S=e?T(p,E):r?T(p,0):void 0;E>_;_++)if((f||_ in g)&&(x=g[_],y=C(x,_,b),t))if(e)S[_]=y;else if(y)switch(t){case 3:return!0;case 5:return x;case 6:return _;case 2:u.call(S,x)}else if(d)return!1;return l?-1:s||d?d:S}};t.exports={forEach:s(0),map:s(1),filter:s(2),some:s(3),every:s(4),find:s(5),findIndex:s(6)}},b91f:function(t,e,r){"use strict";var n=r("d4ec"),c=r("bee2"),a=r("c427"),i=function(){function t(){Object(n["a"])(this,t)}return Object(c["a"])(t,[{key:"getAll",value:function(){return a["a"].get("/product/get")}},{key:"get",value:function(t){return a["a"].get("/product/get/".concat(t))}},{key:"create",value:function(t){return a["a"].post("/product/insert",t)}},{key:"update",value:function(t,e){return a["a"].put("/product/update/".concat(t),e)}},{key:"delete",value:function(t){return a["a"].delete("/product/delete/".concat(t))}}]),t}();e["a"]=new i},c740:function(t,e,r){"use strict";var n=r("23e7"),c=r("b727").findIndex,a=r("44d2"),i=r("ae40"),o="findIndex",u=!0,s=i(o);o in[]&&Array(1)[o]((function(){u=!1})),n({target:"Array",proto:!0,forced:u||!s},{findIndex:function(t){return c(this,t,arguments.length>1?arguments[1]:void 0)}}),a(o)},caad:function(t,e,r){"use strict";var n=r("23e7"),c=r("4d64").includes,a=r("44d2"),i=r("ae40"),o=i("indexOf",{ACCESSORS:!0,1:0});n({target:"Array",proto:!0,forced:!o},{includes:function(t){return c(this,t,arguments.length>1?arguments[1]:void 0)}}),a("includes")},d58f:function(t,e,r){var n=r("1c0b"),c=r("7b0b"),a=r("44ad"),i=r("50c4"),o=function(t){return function(e,r,o,u){n(r);var s=c(e),d=a(s),l=i(s.length),f=t?l-1:0,p=t?-1:1;if(o<2)while(1){if(f in d){u=d[f],f+=p;break}if(f+=p,t?f<0:l<=f)throw TypeError("Reduce of empty array with no initial value")}for(;t?f>=0:l>f;f+=p)f in d&&(u=r(u,d[f],f,s));return u}};t.exports={left:o(!1),right:o(!0)}},d784:function(t,e,r){"use strict";r("ac1f");var n=r("6eeb"),c=r("d039"),a=r("b622"),i=r("9263"),o=r("9112"),u=a("species"),s=!c((function(){var t=/./;return t.exec=function(){var t=[];return t.groups={a:"7"},t},"7"!=="".replace(t,"$<a>")})),d=function(){return"$0"==="a".replace(/./,"$0")}(),l=a("replace"),f=function(){return!!/./[l]&&""===/./[l]("a","$0")}(),p=!c((function(){var t=/(?:)/,e=t.exec;t.exec=function(){return e.apply(this,arguments)};var r="ab".split(t);return 2!==r.length||"a"!==r[0]||"b"!==r[1]}));t.exports=function(t,e,r,l){var v=a(t),h=!c((function(){var e={};return e[v]=function(){return 7},7!=""[t](e)})),m=h&&!c((function(){var e=!1,r=/a/;return"split"===t&&(r={},r.constructor={},r.constructor[u]=function(){return r},r.flags="",r[v]=/./[v]),r.exec=function(){return e=!0,null},r[v](""),!e}));if(!h||!m||"replace"===t&&(!s||!d||f)||"split"===t&&!p){var x=/./[v],y=r(v,""[t],(function(t,e,r,n,c){return e.exec===i?h&&!c?{done:!0,value:x.call(e,r,n)}:{done:!0,value:t.call(r,e,n)}:{done:!1}}),{REPLACE_KEEPS_$0:d,REGEXP_REPLACE_SUBSTITUTES_UNDEFINED_CAPTURE:f}),b=y[0],g=y[1];n(String.prototype,t,b),n(RegExp.prototype,v,2==e?function(t,e){return g.call(t,this,e)}:function(t){return g.call(t,this)})}l&&o(RegExp.prototype[v],"sham",!0)}},e8b5:function(t,e,r){var n=r("c6b6");t.exports=Array.isArray||function(t){return"Array"==n(t)}}}]);
//# sourceMappingURL=chunk-73fdc42c.cd301983.js.map