(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[18],{B1RH:function(e,t,a){e.exports={table:"table___rkUGq"}},zDn2:function(e,t,a){"use strict";a.r(t);a("14J3");var n,r,i,l,o,c,s,m,d,u,p,h,y,f,g=a("BMrR"),E=(a("jCWc"),a("kPKH")),k=(a("P2fV"),a("NJEC")),b=(a("+L6B"),a("2/Rp")),v=a("p0pE"),T=a.n(v),C=a("2Taf"),w=a.n(C),F=a("vZ4D"),_=a.n(F),I=a("l4Ni"),O=a.n(I),D=a("ujKo"),V=a.n(D),R=a("MhPg"),M=a.n(R),j=a("q1tI"),x=a.n(j),A=(a("17x9"),a("7Qib")),S=a("MuoO"),P=a("ZD0w"),q=a("Kvkj"),N=a("Qyje"),Y=(a("g9YV"),a("wCAj")),z=a("jehZ"),K=a.n(z),U=a("Y/ft"),B=a.n(U),Z=(a("2qtc"),a("kLXV")),J=a("mOP9"),L=a("B1RH"),Q=a.n(L),H=Z["a"].confirm,G=(n=Object(P["withI18n"])(),n((i=function(e){function t(){var e,a;w()(this,t);for(var n=arguments.length,r=new Array(n),i=0;i<n;i++)r[i]=arguments[i];return a=O()(this,(e=V()(t)).call.apply(e,[this].concat(r))),a.handleMenuClick=function(e,t){var n=a.props,r=n.onDeleteItem,i=n.onEditItem,l=n.i18n;"1"===t.key?i(e):"2"===t.key&&H({title:l._("Are you sure delete this record?"),onOk:function(){r(e.id)}})},a}return M()(t,e),_()(t,[{key:"render",value:function(){var e=this,t=this.props,a=(t.onDeleteItem,t.onEditItem,t.i18n),n=B()(t,["onDeleteItem","onEditItem","i18n"]),r=[{title:x.a.createElement(P["Trans"],{id:"ID"}),dataIndex:"id",key:"id",render:function(e,t){return x.a.createElement(J["a"],{to:"user/".concat(t.id)},e)}},{title:x.a.createElement(P["Trans"],{id:"Name"}),dataIndex:"name",key:"name",render:function(e,t){return x.a.createElement(J["a"],{to:"user/".concat(t.id)},e)}},{title:x.a.createElement(P["Trans"],{id:"CreateTime"}),dataIndex:"createTime",key:"createTime"},{title:x.a.createElement(P["Trans"],{id:"Operation"}),key:"operation",fixed:"right",render:function(t,n){return x.a.createElement(q["a"],{onMenuClick:function(t){return e.handleMenuClick(n,t)},menuOptions:[{key:"1",name:a._("Update")},{key:"2",name:a._("Delete")}]})}}];return x.a.createElement(Y["a"],K()({},n,{pagination:T()({},n.pagination,{showTotal:function(e){return a._("Total {total} Items",{total:e})}}),className:Q.a.table,bordered:!0,scroll:{x:1200},columns:r,simple:!0,rowKey:function(e){return e.id}}))}}]),t}(j["PureComponent"]),r=i))||r),W=G,X=(a("y8nQ"),a("Vl3Y")),$=(a("iQDF"),a("+eQT")),ee=(a("5NDa"),a("5rEg")),te=a("wd/R"),ae=a.n(te),ne=a("Lo71"),re=ee["a"].Search,ie=$["a"].RangePicker,le={xs:24,sm:12,style:{marginBottom:16}},oe=T()({},le,{xl:96}),ce=(l=Object(P["withI18n"])(),o=X["a"].create(),l(c=o((s=function(e){function t(){var e,a;w()(this,t);for(var n=arguments.length,r=new Array(n),i=0;i<n;i++)r[i]=arguments[i];return a=O()(this,(e=V()(t)).call.apply(e,[this].concat(r))),a.handleFields=function(e){var t=e.createTime;return t.length&&(e.createTime=[ae()(t[0]).format("YYYY-MM-DD"),ae()(t[1]).format("YYYY-MM-DD")]),e},a.handleSubmit=function(){var e=a.props,t=e.onFilterChange,n=e.form,r=n.getFieldsValue,i=r();i=a.handleFields(i),t(i)},a.handleReset=function(){var e=a.props.form,t=e.getFieldsValue,n=e.setFieldsValue,r=t();for(var i in r)({}).hasOwnProperty.call(r,i)&&(r[i]instanceof Array?r[i]=[]:r[i]=void 0);n(r),a.handleSubmit()},a.handleChange=function(e,t){var n=a.props,r=n.form,i=n.onFilterChange,l=r.getFieldsValue,o=l();o[e]=t,o=a.handleFields(o),i(o)},a}return M()(t,e),_()(t,[{key:"componentDidUpdate",value:function(e,t){0===Object.keys(e.filter).length&&this.handleReset()}},{key:"render",value:function(){var e=this.props,t=e.onAdd,a=e.filter,n=e.form,r=e.i18n,i=n.getFieldDecorator,l=a.name,o=(a.address,[]);return a.createTime&&a.createTime[0]&&(o[0]=ae()(a.createTime[0])),a.createTime&&a.createTime[1]&&(o[1]=ae()(a.createTime[1])),x.a.createElement(g["a"],{gutter:24},x.a.createElement(E["a"],K()({},le,{xl:{span:4},md:{span:8}}),i("name",{initialValue:l})(x.a.createElement(re,{placeholder:r._("Search Name"),onSearch:this.handleSubmit}))),x.a.createElement(E["a"],K()({},le,{xl:{span:6},md:{span:8},sm:{span:12},id:"createTimeRangePicker"}),x.a.createElement(q["c"],{label:r._("CreateTime")},i("createTime",{initialValue:o})(x.a.createElement(ie,{style:{width:"100%"},onChange:this.handleChange.bind(this,"createTime"),getCalendarContainer:function(){return document.getElementById("createTimeRangePicker")}})))),x.a.createElement(E["a"],K()({},oe,{xl:{span:10},md:{span:24},sm:{span:24}}),x.a.createElement(g["a"],{type:"flex",align:"middle",justify:"space-between"},x.a.createElement("div",null,x.a.createElement(b["a"],{type:"primary",className:"margin-right",onClick:this.handleSubmit},x.a.createElement(P["Trans"],{id:"Search"})),x.a.createElement(b["a"],{onClick:this.handleReset},x.a.createElement(P["Trans"],{id:"Reset"}))),x.a.createElement(b["a"],{type:"ghost",onClick:t},x.a.createElement(P["Trans"],{id:"Create"})))))}}]),t}(j["Component"]),c=s))||c)||c),se=ce,me=(a("6UJt"),a("DFOY")),de=(a("giR+"),a("fyUT")),ue=(a("7Kak"),a("9yH6")),pe=X["a"].Item,he={labelCol:{span:6},wrapperCol:{span:14}},ye=(m=Object(P["withI18n"])(),d=X["a"].create(),m(u=d((p=function(e){function t(){var e,a;w()(this,t);for(var n=arguments.length,r=new Array(n),i=0;i<n;i++)r[i]=arguments[i];return a=O()(this,(e=V()(t)).call.apply(e,[this].concat(r))),a.handleOk=function(){var e=a.props,t=e.item,n=void 0===t?{}:t,r=e.onOk,i=e.form,l=i.validateFields,o=i.getFieldsValue;l(function(e){if(!e){var t=T()({},o(),{key:n.key});t.address=t.address.join(" "),r(t)}})},a}return M()(t,e),_()(t,[{key:"render",value:function(){var e=this.props,t=e.item,a=void 0===t?{}:t,n=(e.onOk,e.form),r=e.i18n,i=B()(e,["item","onOk","form","i18n"]),l=n.getFieldDecorator;return x.a.createElement(Z["a"],K()({},i,{onOk:this.handleOk}),x.a.createElement(X["a"],{layout:"horizontal"},x.a.createElement(pe,K()({label:r._("Name"),hasFeedback:!0},he),l("name",{initialValue:a.name,rules:[{required:!0}]})(x.a.createElement(ee["a"],null))),x.a.createElement(pe,K()({label:r._("NickName"),hasFeedback:!0},he),l("nickName",{initialValue:a.nickName,rules:[{required:!0}]})(x.a.createElement(ee["a"],null))),x.a.createElement(pe,K()({label:r._("Gender"),hasFeedback:!0},he),l("isMale",{initialValue:a.isMale,rules:[{required:!0,type:"boolean"}]})(x.a.createElement(ue["a"].Group,null,x.a.createElement(ue["a"],{value:!0},x.a.createElement(P["Trans"],{id:"Male"})),x.a.createElement(ue["a"],{value:!1},x.a.createElement(P["Trans"],{id:"Female"}))))),x.a.createElement(pe,K()({label:r._("Age"),hasFeedback:!0},he),l("age",{initialValue:a.age,rules:[{required:!0,type:"number"}]})(x.a.createElement(de["a"],{min:18,max:100}))),x.a.createElement(pe,K()({label:r._("Phone"),hasFeedback:!0},he),l("phone",{initialValue:a.phone,rules:[{required:!0,pattern:/^1[34578]\d{9}$/,message:r._("The input is not valid phone!")}]})(x.a.createElement(ee["a"],null))),x.a.createElement(pe,K()({label:r._("Email"),hasFeedback:!0},he),l("email",{initialValue:a.email,rules:[{required:!0,pattern:/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/,message:r._("The input is not valid E-mail!")}]})(x.a.createElement(ee["a"],null))),x.a.createElement(pe,K()({label:r._("Address"),hasFeedback:!0},he),l("address",{initialValue:a.address&&a.address.split(" "),rules:[{required:!0}]})(x.a.createElement(me["a"],{style:{width:"100%"},options:ne["a"],placeholder:r._("Pick an address")})))))}}]),t}(j["PureComponent"]),u=p))||u)||u),fe=ye,ge=(h=Object(P["withI18n"])(),y=Object(S["connect"])(function(e){var t=e.strategy,a=e.loading;return{strategy:t,loading:a}}),h(f=y(f=function(e){function t(){return w()(this,t),O()(this,V()(t).apply(this,arguments))}return M()(t,e),_()(t,[{key:"render",value:function(){var e=this.props,t=e.location,a=e.dispatch,n=e.strategy,r=e.loading,i=e.i18n,l=t.query,o=t.pathname,c=n.list,s=n.pagination,m=n.currentItem,d=n.modalVisible,u=n.modalType,p=n.selectedRowKeys,h=function(e){A["k"].push({pathname:o,search:Object(N["stringify"])(T()({},l,e),{arrayFormat:"repeat"})})},y={item:"create"===u?{}:m,visible:d,maskClosable:!1,confirmLoading:r.effects["strategy/".concat(u)],title:"".concat("create"===u?i._("Create User"):i._("Update User")),centered:!0,onOk:function(e){a({type:"strategy/".concat(u),payload:e}).then(function(){h()})},onCancel:function(){a({type:"strategy/hideModal"})}},f={dataSource:c,loading:r.effects["strategy/query"],pagination:s,onChange:function(e){h({page:e.current,pageSize:e.pageSize})},onDeleteItem:function(e){a({type:"strategy/delete",payload:e}).then(function(){h({page:1===c.length&&s.current>1?s.current-1:s.current})})},onEditItem:function(e){a({type:"strategy/showModal",payload:{modalType:"update",currentItem:e}})},rowSelection:{selectedRowKeys:p,onChange:function(e){a({type:"strategy/updateState",payload:{selectedRowKeys:e}})}}},v={filter:T()({},l),onFilterChange:function(e){h(T()({},e,{page:1}))},onAdd:function(){a({type:"strategy/showModal",payload:{modalType:"create"}})}},C=function(){a({type:"strategy/multiDelete",payload:{ids:p}}).then(function(){h({page:c.length===p.length&&s.current>1?s.current-1:s.current})})};return x.a.createElement(q["f"],{inner:!0},x.a.createElement(se,v),p.length>0&&x.a.createElement(g["a"],{style:{marginBottom:24,textAlign:"right",fontSize:13}},x.a.createElement(E["a"],null,"Selected ".concat(p.length," items "),x.a.createElement(k["a"],{title:"Are you sure delete these items?",placement:"left",onConfirm:C},x.a.createElement(b["a"],{type:"primary",style:{marginLeft:8}},"Remove")))),x.a.createElement(W,f),d&&x.a.createElement(fe,y))}}]),t}(j["PureComponent"]))||f)||f);t["default"]=ge}}]);