import{_ as t,a as n,b as r,c,i as e,d as o,S as a,s,t as u,l as i,n as f,z as l,g as p,D as d,w as m,p as v,Y as h,A as y}from"./client.667de7f7.js";import{t as b,_ as g}from"./stores.08bd5bd5.js";function w(t){var c=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(t){return!1}}();return function(){var e,o=n(t);if(c){var a=n(this).constructor;e=Reflect.construct(o,arguments,a)}else e=o.apply(this,arguments);return r(this,e)}}function R(t){var n;return{c:function(){n=u(t[0])},l:function(r){n=i(r,t[0])},m:function(t,r){f(t,n,r)},p:function(t,r){1&r&&l(n,t[0])},d:function(t){t&&p(n)}}}function x(t){var n,r=t[0]&&R(t);return{c:function(){r&&r.c(),n=d()},l:function(t){r&&r.l(t),n=d()},m:function(t,c){r&&r.m(t,c),f(t,n,c)},p:function(t,c){var e=m(c,1)[0];t[0]?r?r.p(t,e):((r=R(t)).c(),r.m(n.parentNode,n)):r&&(r.d(1),r=null)},i:v,o:v,d:function(t){r&&r.d(t),t&&p(n)}}}function D(t,n,r){var c;b.subscribe((function(t){r(0,c=t)}));return console.log(b),h(g(y.mark((function t(){return y.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:b.set(window.localStorage.getItem("token"));case 1:case"end":return t.stop()}}),t)})))),[c]}var S=function(n){t(u,a);var r=w(u);function u(t){var n;return c(this,u),n=r.call(this),e(o(n),t,D,x,s,{}),n}return u}();export default S;
