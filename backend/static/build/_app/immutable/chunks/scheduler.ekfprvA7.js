function C(){}function S(t,n){for(const e in n)t[e]=n[e];return t}function P(t){return t()}function K(){return Object.create(null)}function B(t){t.forEach(P)}function Q(t){return typeof t=="function"}function R(t,n){return t!=t?n==n:t!==n||t&&typeof t=="object"||typeof t=="function"}function W(t){return Object.keys(t).length===0}function D(t,...n){if(t==null){for(const i of n)i(void 0);return C}const e=t.subscribe(...n);return e.unsubscribe?()=>e.unsubscribe():e}function V(t,n,e){t.$$.on_destroy.push(D(n,e))}function X(t,n,e,i){if(t){const r=w(t,n,e,i);return t[0](r)}}function w(t,n,e,i){return t[1]&&i?S(e.ctx.slice(),t[1](i(n))):e.ctx}function Y(t,n,e,i){if(t[2]&&i){const r=t[2](i(e));if(n.dirty===void 0)return r;if(typeof r=="object"){const s=[],c=Math.max(n.dirty.length,r.length);for(let u=0;u<c;u+=1)s[u]=n.dirty[u]|r[u];return s}return n.dirty|r}return n.dirty}function Z(t,n,e,i,r,s){if(r){const c=w(n,e,i,s);t.p(c,r)}}function $(t){if(t.ctx.length>32){const n=[],e=t.ctx.length/32;for(let i=0;i<e;i++)n[i]=-1;return n}return-1}let p=!1;function tt(){p=!0}function nt(){p=!1}function M(t,n,e,i){for(;t<n;){const r=t+(n-t>>1);e(r)<=i?t=r+1:n=r}return t}function O(t){if(t.hydrate_init)return;t.hydrate_init=!0;let n=t.childNodes;if(t.nodeName==="HEAD"){const l=[];for(let o=0;o<n.length;o++){const a=n[o];a.claim_order!==void 0&&l.push(a)}n=l}const e=new Int32Array(n.length+1),i=new Int32Array(n.length);e[0]=-1;let r=0;for(let l=0;l<n.length;l++){const o=n[l].claim_order,a=(r>0&&n[e[r]].claim_order<=o?r+1:M(1,r,j=>n[e[j]].claim_order,o))-1;i[l]=e[a]+1;const v=a+1;e[v]=l,r=Math.max(v,r)}const s=[],c=[];let u=n.length-1;for(let l=e[r]+1;l!=0;l=i[l-1]){for(s.push(n[l-1]);u>=l;u--)c.push(n[u]);u--}for(;u>=0;u--)c.push(n[u]);s.reverse(),c.sort((l,o)=>l.claim_order-o.claim_order);for(let l=0,o=0;l<c.length;l++){for(;o<s.length&&c[l].claim_order>=s[o].claim_order;)o++;const a=o<s.length?s[o]:null;t.insertBefore(c[l],a)}}function T(t,n){if(p){for(O(t),(t.actual_end_child===void 0||t.actual_end_child!==null&&t.actual_end_child.parentNode!==t)&&(t.actual_end_child=t.firstChild);t.actual_end_child!==null&&t.actual_end_child.claim_order===void 0;)t.actual_end_child=t.actual_end_child.nextSibling;n!==t.actual_end_child?(n.claim_order!==void 0||n.parentNode!==t)&&t.insertBefore(n,t.actual_end_child):t.actual_end_child=n.nextSibling}else(n.parentNode!==t||n.nextSibling!==null)&&t.appendChild(n)}function et(t,n,e){p&&!e?T(t,n):(n.parentNode!==t||n.nextSibling!=e)&&t.insertBefore(n,e||null)}function it(t){t.parentNode&&t.parentNode.removeChild(t)}function q(t){return document.createElement(t)}function g(t){return document.createTextNode(t)}function rt(){return g(" ")}function ct(){return g("")}function lt(t,n,e,i){return t.addEventListener(n,e,i),()=>t.removeEventListener(n,e,i)}function ut(t,n,e){e==null?t.removeAttribute(n):t.getAttribute(n)!==e&&t.setAttribute(n,e)}function st(t){return t.dataset.svelteH}function ot(t){return Array.from(t.childNodes)}function H(t){t.claim_info===void 0&&(t.claim_info={last_index:0,total_claimed:0})}function N(t,n,e,i,r=!1){H(t);const s=(()=>{for(let c=t.claim_info.last_index;c<t.length;c++){const u=t[c];if(n(u)){const l=e(u);return l===void 0?t.splice(c,1):t[c]=l,r||(t.claim_info.last_index=c),u}}for(let c=t.claim_info.last_index-1;c>=0;c--){const u=t[c];if(n(u)){const l=e(u);return l===void 0?t.splice(c,1):t[c]=l,r?l===void 0&&t.claim_info.last_index--:t.claim_info.last_index=c,u}}return i()})();return s.claim_order=t.claim_info.total_claimed,t.claim_info.total_claimed+=1,s}function I(t,n,e,i){return N(t,r=>r.nodeName===n,r=>{const s=[];for(let c=0;c<r.attributes.length;c++){const u=r.attributes[c];e[u.name]||s.push(u.name)}s.forEach(c=>r.removeAttribute(c))},()=>i(n))}function at(t,n,e){return I(t,n,e,q)}function L(t,n){return N(t,e=>e.nodeType===3,e=>{const i=""+n;if(e.data.startsWith(i)){if(e.data.length!==i.length)return e.splitText(i.length)}else e.data=i},()=>g(n),!0)}function ft(t){return L(t," ")}function _t(t,n){n=""+n,t.data!==n&&(t.data=n)}function dt(t,n){t.value=n??""}function ht(t,n,e,i){e==null?t.style.removeProperty(n):t.style.setProperty(n,e,i?"important":"")}function z(t,n,{bubbles:e=!1,cancelable:i=!1}={}){return new CustomEvent(t,{detail:n,bubbles:e,cancelable:i})}function mt(t,n){return new t(n)}let m;function b(t){m=t}function d(){if(!m)throw new Error("Function called outside component initialization");return m}function pt(t){d().$$.on_mount.push(t)}function bt(t){d().$$.after_update.push(t)}function yt(t){d().$$.on_destroy.push(t)}function xt(){const t=d();return(n,e,{cancelable:i=!1}={})=>{const r=t.$$.callbacks[n];if(r){const s=z(n,e,{cancelable:i});return r.slice().forEach(c=>{c.call(t,s)}),!s.defaultPrevented}return!0}}function gt(t,n){return d().$$.context.set(t,n),n}function vt(t){return d().$$.context.get(t)}const h=[],E=[];let _=[];const k=[],A=Promise.resolve();let x=!1;function F(){x||(x=!0,A.then(G))}function Et(){return F(),A}function U(t){_.push(t)}const y=new Set;let f=0;function G(){if(f!==0)return;const t=m;do{try{for(;f<h.length;){const n=h[f];f++,b(n),J(n.$$)}}catch(n){throw h.length=0,f=0,n}for(b(null),h.length=0,f=0;E.length;)E.pop()();for(let n=0;n<_.length;n+=1){const e=_[n];y.has(e)||(y.add(e),e())}_.length=0}while(h.length);for(;k.length;)k.pop()();x=!1,y.clear(),b(t)}function J(t){if(t.fragment!==null){t.update(),B(t.before_update);const n=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,n),t.after_update.forEach(U)}}function kt(t){const n=[],e=[];_.forEach(i=>t.indexOf(i)===-1?n.push(i):e.push(i)),e.forEach(i=>i()),_=n}export{W as A,U as B,kt as C,m as D,b as E,P as F,h as G,F as H,tt as I,nt as J,X as K,Z as L,$ as M,Y as N,gt as O,yt as P,vt as Q,xt as R,st as S,dt as T,lt as U,rt as a,ot as b,at as c,L as d,q as e,it as f,ft as g,T as h,et as i,_t as j,V as k,ct as l,bt as m,C as n,pt as o,ut as p,ht as q,E as r,R as s,g as t,mt as u,Et as v,B as w,K as x,G as y,Q as z};
