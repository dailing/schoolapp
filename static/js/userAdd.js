/**
 * Created by d on 4/12/16.
 */
function user_add(name, psw) {
    console.info("adding names")
    var jsonObj= {
        "username": name,
        "password":psw
    };
    ajaxCall = new XMLHttpRequest();
    ajaxCall.open("Post","/api/usr_add",false)
    ajaxCall.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    ajaxCall.send(JSON.stringify(jsonObj))
    resp = ajaxCall.responseText
    // document.getElementById("resp").innerText = resp
    console.info(resp)
}
