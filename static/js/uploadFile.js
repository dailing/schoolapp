/**
 * Created by d on 4/17/16.
 */

function upload(file) {
    var request = new XMLHttpRequest();
    var fd = new FormData();
    request.open('POST', '/api/upload');
    request.responseType = 'arraybuffer';
    fd.append("binaryFile", file);
    console.log(file.type);
    request.send(fd);
    // request.onload = function () {
    //     if(request.readyState == 4){
    //         var arr = new Uint8Array(this.response);
    //         // var raw = String.fromCharCode.apply(null,arr);
    //         var raw = '';
    //         var i,j,subArray,chunk = 5000;
    //         for (i=0,j=arr.length; i<j; i+=chunk) {
    //             subArray = arr.subarray(i,i+chunk);
    //             raw += String.fromCharCode.apply(null, subArray);
    //         }
    //         var b64=btoa(raw);
    //         var dataURL="data:image/jpeg;base64,"+b64;
    //         document.getElementById("vesselseg_out_img").src = dataURL
    //         console.log("finished load")
    //         return
    //     }
    // }
}