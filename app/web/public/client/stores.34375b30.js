import{w as t}from"./client.d472693b.js";const o=t("");async function s(t){const s=await t.json();if(401===t.status)window.localStorage.removeItem("token"),o.set("");else{if(200!==t.status)throw s;if(void 0!==s.errors)throw s.errors[0]}return s}export{s as c,o as t};
