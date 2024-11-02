import{Z as q,j as e,$ as h,Y as O,a0 as x,a1 as A,a2 as B,a3 as F,a4 as U,T as k,r as g,a5 as I,a6 as V,a7 as T,a8 as L,X as K,a9 as X,B as l,d as z,L as E,R as Y,aa as Z,U as J,n as R,b as Q,W as P,o as ee,V as se,H as oe,C as re}from"./index-CyQALytQ.js";import{g as te,f as ae,z as b,u as ne,t as le,a as ie,D as ce,F as v,b as de,L as ue}from"./form-provider-ByR0Juqh.js";import{a as pe,u as me}from"./model-DcVIY5nw.js";import"./index-BAAAdKb5.js";const xe=q(e.jsx("path",{d:"M6 10c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm12 0c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm-6 0c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"}),"MoreHoriz"),ge=["slots","slotProps"],he=h(O)(({theme:s})=>x({display:"flex",marginLeft:`calc(${s.spacing(1)} * 0.5)`,marginRight:`calc(${s.spacing(1)} * 0.5)`},s.palette.mode==="light"?{backgroundColor:s.palette.grey[100],color:s.palette.grey[700]}:{backgroundColor:s.palette.grey[700],color:s.palette.grey[100]},{borderRadius:2,"&:hover, &:focus":x({},s.palette.mode==="light"?{backgroundColor:s.palette.grey[200]}:{backgroundColor:s.palette.grey[600]}),"&:active":x({boxShadow:s.shadows[0]},s.palette.mode==="light"?{backgroundColor:A(s.palette.grey[200],.12)}:{backgroundColor:A(s.palette.grey[600],.12)})})),be=h(xe)({width:24,height:16});function fe(s){const{slots:o={},slotProps:n={}}=s,r=B(s,ge),a=s;return e.jsx("li",{children:e.jsx(he,x({focusRipple:!0},r,{ownerState:a,children:e.jsx(be,x({as:o.CollapsedIcon,ownerState:a},n.collapsedIcon))}))})}function ye(s){return U("MuiBreadcrumbs",s)}const je=F("MuiBreadcrumbs",["root","ol","li","separator"]),Ce=["children","className","component","slots","slotProps","expandText","itemsAfterCollapse","itemsBeforeCollapse","maxItems","separator"],ve=s=>{const{classes:o}=s;return L({root:["root"],li:["li"],ol:["ol"],separator:["separator"]},ye,o)},Me=h(k,{name:"MuiBreadcrumbs",slot:"Root",overridesResolver:(s,o)=>[{[`& .${je.li}`]:o.li},o.root]})({}),Be=h("ol",{name:"MuiBreadcrumbs",slot:"Ol",overridesResolver:(s,o)=>o.ol})({display:"flex",flexWrap:"wrap",alignItems:"center",padding:0,margin:0,listStyle:"none"}),we=h("li",{name:"MuiBreadcrumbs",slot:"Separator",overridesResolver:(s,o)=>o.separator})({display:"flex",userSelect:"none",marginLeft:8,marginRight:8});function Se(s,o,n,r){return s.reduce((a,t,i)=>(i<s.length-1?a=a.concat(t,e.jsx(we,{"aria-hidden":!0,className:o,ownerState:r,children:n},`separator-${i}`)):a.push(t),a),[])}const _e=g.forwardRef(function(o,n){const r=I({props:o,name:"MuiBreadcrumbs"}),{children:a,className:t,component:i="nav",slots:d={},slotProps:c={},expandText:u="Show path",itemsAfterCollapse:f=1,itemsBeforeCollapse:y=1,maxItems:j=8,separator:p="/"}=r,w=B(r,Ce),[D,H]=g.useState(!1),C=x({},r,{component:i,expanded:D,expandText:u,itemsAfterCollapse:f,itemsBeforeCollapse:y,maxItems:j,separator:p}),M=ve(C),W=V({elementType:d.CollapsedIcon,externalSlotProps:c.collapsedIcon,ownerState:C}),N=g.useRef(null),G=m=>{const _=()=>{H(!0);const $=N.current.querySelector("a[href],button,[tabindex]");$&&$.focus()};return y+f>=m.length?m:[...m.slice(0,y),e.jsx(fe,{"aria-label":u,slots:{CollapsedIcon:d.CollapsedIcon},slotProps:{collapsedIcon:W},onClick:_},"ellipsis"),...m.slice(m.length-f,m.length)]},S=g.Children.toArray(a).filter(m=>g.isValidElement(m)).map((m,_)=>e.jsx("li",{className:M.li,children:m},`child-${_}`));return e.jsx(Me,x({ref:n,component:i,color:"text.secondary",className:T(M.root,t),ownerState:C},w,{children:e.jsx(Be,{className:M.ol,ref:N,ownerState:C,children:Se(D||j&&S.length<=j?S:G(S),M.separator,p,C)})}))});function Re(s){return U("MuiCard",s)}F("MuiCard",["root"]);const ke=["className","raised"],Ie=s=>{const{classes:o}=s;return L({root:["root"]},Re,o)},Te=h(K,{name:"MuiCard",slot:"Root",overridesResolver:(s,o)=>o.root})(()=>({overflow:"hidden"})),Le=g.forwardRef(function(o,n){const r=I({props:o,name:"MuiCard"}),{className:a,raised:t=!1}=r,i=B(r,ke),d=x({},r,{raised:t}),c=Ie(d);return e.jsx(Te,x({className:T(c.root,a),elevation:t?8:void 0,ref:n,ownerState:d},i))}),De=["className","id"],Ne=s=>{const{classes:o}=s;return L({root:["root"]},te,o)},$e=h(k,{name:"MuiDialogTitle",slot:"Root",overridesResolver:(s,o)=>o.root})({padding:"16px 24px",flex:"0 0 auto"}),Ae=g.forwardRef(function(o,n){const r=I({props:o,name:"MuiDialogTitle"}),{className:a,id:t}=r,i=B(r,De),d=r,c=Ne(d),{titleId:u=t}=g.useContext(X);return e.jsx($e,x({component:"h2",className:T(c.root,a),ownerState:d,ref:n,variant:"h6",id:t??u},i))});function Fe({title:s,percent:o,total:n,model:r,icon:a,sx:t,...i}){return l,z,o<0,{...o<0&&{color:"error.main"}},o>0,ae(o),e.jsxs(Le,{sx:{p:2,pl:3,display:"flex",alignItems:"center",...t},...i,children:[e.jsxs(l,{sx:{flexGrow:1},children:[e.jsx(l,{sx:{color:"text.secondary",typography:"subtitle2"},children:s}),e.jsx(l,{sx:{my:1.5,typography:"h4"},children:r})]}),e.jsx(l,{sx:{width:50,height:50,lineHeight:0,borderRadius:"50%",bgcolor:"background.neutral"},children:a})]})}function Ue({link:s,activeLast:o,disabled:n}){const r={typography:"body2",alignItems:"center",color:"text.primary",display:"inline-flex",...n&&!o&&{cursor:"default",pointerEvents:"none",color:"text.disabled"}},a=e.jsxs(e.Fragment,{children:[s.icon&&e.jsx(l,{component:"span",sx:{mr:1,display:"inherit","& svg, & img":{width:20,height:20}},children:s.icon}),s.name]});return s.href?e.jsx(E,{component:Y,href:s.href,sx:r,children:a}):e.jsxs(l,{sx:r,children:[" ",a," "]})}function ze({links:s,action:o,heading:n,moreLink:r,activeLast:a,slotProps:t,sx:i,...d}){const c=s[s.length-1].name,u=e.jsx(k,{variant:"h4",sx:{mb:2,...t==null?void 0:t.heading},children:n}),f=e.jsx(_e,{separator:e.jsx(Ee,{}),sx:t==null?void 0:t.breadcrumbs,...d,children:s.map((p,w)=>e.jsx(Ue,{link:p,activeLast:a,disabled:p.name===c},p.name??w))}),y=e.jsxs(l,{sx:{flexShrink:0,...t==null?void 0:t.action},children:[" ",o," "]}),j=e.jsx(l,{component:"ul",children:r==null?void 0:r.map(p=>e.jsx(l,{component:"li",sx:{display:"flex"},children:e.jsx(E,{href:p,variant:"body2",target:"_blank",rel:"noopener",sx:t==null?void 0:t.moreLink,children:p})},p))});return e.jsxs(l,{gap:2,display:"flex",flexDirection:"column",sx:i,...d,children:[e.jsxs(l,{display:"flex",alignItems:"center",children:[e.jsxs(l,{sx:{flexGrow:1},children:[n&&u,!!s.length&&f]}),o&&y]}),!!r&&j]})}function Ee(){return e.jsx(l,{component:"span",sx:{width:4,height:4,borderRadius:"50%",bgcolor:"text.disabled"}})}const He=b.object({model_type:b.string().min(1,{message:"Model_type is required!"}),name:b.string(),model:b.string().min(1,{message:"Model is required!"}),api_key:b.string().min(1,{message:"Api_key is required!"}),base_url:b.string().min(1,{message:"Base_url is required!"})});function We({open:s,onClose:o}){const{trigger:n}=pe(),r={name:"",model:"taichu_llm",api_key:"ryvsk3zz73419gkgubrnvufp",base_url:"https://ai-cds.wair.ac.cn/maas/v1/",model_type:"llm"},a=ne({mode:"all",resolver:le(He),defaultValues:r}),{handleSubmit:t,formState:{isSubmitting:i}}=a,d=t(async c=>{try{const u={name:c.name||c.model,model:c.model,model_type:c.model_type,api_key:c.api_key,base_url:c.base_url};console.log("🚀 ~ onSubmit ~ params:",u),await n(u)}catch(u){console.error(u)}});return e.jsx(Z,{fullWidth:!0,maxWidth:"sm",open:s,onClose:o,children:e.jsxs(ie,{methods:a,onSubmit:d,children:[e.jsx(Ae,{children:"添加模型"}),e.jsx(ce,{dividers:!0,children:e.jsxs(J,{spacing:3,children:[e.jsx(v.RadioGroup,{row:!0,label:"Model Type",name:"model_type",options:[{label:"LLM",value:"llm"},{label:"Others",value:"others"}]}),e.jsx(v.Text,{name:"name",label:"Model alias"}),e.jsx(v.Text,{name:"model",label:"Model"}),e.jsx(v.Text,{name:"api_key",label:"Api Key"}),e.jsx(v.Text,{name:"base_url",label:"Base Url"})]})}),e.jsxs(de,{children:[e.jsx(l,{sx:{flexGrow:1},children:e.jsx(R,{variant:"soft",color:"error",onClick:o,children:"Delete"})}),e.jsx(R,{color:"inherit",variant:"outlined",onClick:o,children:"Cancel"}),e.jsx(ue,{type:"submit",variant:"contained",loading:i,children:"Create"})]})]})})}function Ge({title:s="Blank"}){const{modelList:o}=me();console.log("🚀 ~ MultiLLMChat ~ modelList:",o);const n=Q();return e.jsxs(P,{maxWidth:"xl",children:[e.jsx(ze,{heading:s,links:[{name:"basic ability"},{name:"models"}],action:e.jsx(R,{variant:"contained",startIcon:e.jsx(z,{icon:"mingcute:add-line"}),onClick:n.onTrue,children:"添加模型"}),sx:{mb:{xs:3,md:5}}}),e.jsx(l,{sx:{gap:3,display:"grid",gridTemplateColumns:{xs:"repeat(1, 1fr)",md:"repeat(4, 1fr)"}},children:ee.isArray(o)?o.map(r=>e.jsx(Fe,{title:"LLM",percent:2.6,model:r.model,sx:{p:a=>a.spacing(1.5,1,1.5,2)},icon:e.jsx(se,{src:r.brand,sx:{width:50,height:50}})},r.model)):null}),e.jsx(We,{open:n.value,onClose:n.onFalse})]})}const qe={title:`Page Models | Dashboard - ${re.appName}`};function Ye(){return e.jsxs(e.Fragment,{children:[e.jsx(oe,{children:e.jsxs("title",{children:[" ",qe.title]})}),e.jsx(Ge,{title:"模型列表"})]})}export{Ye as default};