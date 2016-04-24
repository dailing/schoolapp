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

function addComments(token, comment) {
    var reqObj = {
        "token": token,
        "comment": comment
    };
    doAjaxCall(reqObj, "/api/item_add_comment")
}

function getComments(token, id) {
    var reqObj = {
        "token": token,
        "comment": {
            "itemID": id
        }
    };
    return doAjaxCall(reqObj, "/api/item_get_comment")
}

function char_add(token, buyerID, itemID, text) {
    var reqObj = {
        "token": token,
        "chat": {
            "content": text,
            "buyerID": buyerID,
            "itemID": itemID
        }
    };
    doAjaxCall(reqObj, "/api/item_add_chart");
}

function item_get_all() {
    var reqObj = {
        "token":""
    };
    return doAjaxCall(reqObj, "/api/item_get_all");
}

function getImage(imgid, element) {
    console.info("imgid is ", imgid);
    var request = new XMLHttpRequest();
    var fd = new FormData();
    request.open('POST', "/api/img_get");
    request.responseType = 'arraybuffer';
    jsonObj = {
        "imageID": imgid
    };
    request.send(JSON.stringify(jsonObj));

    request.onload = function () {
        if (request.readyState == 4) {
            var arr = new Uint8Array(this.response);
            var raw = '';
            var i, j, subArray, chunk = 5000;
            for (i = 0, j = arr.length; i < j; i += chunk) {
                subArray = arr.subarray(i, i + chunk);
                raw += String.fromCharCode.apply(null, subArray);
            }
            var b64 = btoa(raw);
            element.src = "data:image/jpeg;base64," + b64;
            console.log("finished load");
        }
    }
}

function gen_image_id(index, imageID) {
    return "" + index + "_" + imageID;
}
