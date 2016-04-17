/**
 * Created by d on 4/17/16.
 */

function upload(file, callbackfunc) {
    var request = new XMLHttpRequest();
    var fd = new FormData();
    request.open('POST', '/api/img_upload', true);
    fd.append("binaryFile", file);
    console.log(file.type);
    request.send(fd);
    request.onload = function () {
        if (request.readyState == 4) {
            console.log("finished load");
            console.info(request.responseText);
        }
        var resp = JSON.parse(request.responseText);
        callbackfunc(resp)
    };
}