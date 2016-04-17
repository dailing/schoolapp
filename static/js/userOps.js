/**
 * Created by d on 4/12/16.
 */

function doAjaxCall(jsonObj, url) {
    var ajaxCall = new XMLHttpRequest();
    ajaxCall.open("Post", url, false);
    ajaxCall.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    ajaxCall.send(JSON.stringify(jsonObj));
    var resp = ajaxCall.responseText;
    console.info("Do Ajax Call");
    console.info(resp);
    return JSON.parse(resp)
}

function usrUpdate(jsonObj) {
    doAjaxCall(jsonObj, "/api/usr_update")
}

function user_add(name, psw) {
    console.info("adding names");
    var jsonObj= {
        "username": name,
        "password":psw
    };
    // var ajaxCall = new XMLHttpRequest();
    // ajaxCall.open("Post", "/api/usr_add", false);
    // ajaxCall.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    // ajaxCall.send(JSON.stringify(jsonObj));
    // resp = ajaxCall.responseText;
    // // document.getElementById("resp").innerText = resp
    // console.info(resp)
    var resp = doAjaxCall(jsonObj, "/api/usr_add");
    console.info(resp)
}

function user_login(name, psw) {
    console.info("logging in" + name + psw);
    var jsonObj = {
        "userinfo": {
            "username": name,
            "password": psw
        }
    };
    // var ajaxCall = new XMLHttpRequest();
    // ajaxCall.open("Post","/api/login", false)
    // ajaxCall.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    // ajaxCall.send(JSON.stringify(jsonObj));
    // var resp = ajaxCall.responseText;
    // console.info(resp);
    // var respJaon = JSON.parse(ajaxCall.responseText);
    var respJson = doAjaxCall(jsonObj, "/api/login");
    return respJson.token;
}

function user_get(token) {
    console.info("getting user info");
    var jsonObj = {
        "token": token
    };
    return doAjaxCall(jsonObj, "/api/usr_get");
}


function item_Add(token, iteminfo) {
    console.trace("enter the item_add");
    var jsonObj = {
        "token": token,
        "itemInfo": iteminfo
    };
    var response = doAjaxCall(jsonObj, "/api/item_add");
    console.trace(response);
    return response;
}

function item_get_list(token) {
    var jsonObj = {
        "token": token
    };
    return doAjaxCall(jsonObj, "/api/item_get_list");
}

function item_get(token, id) {

}