$(document).ready( function() {
    var date = new Date();

    var day = date.getDate();
    var month = date.getMonth() + 1;
    var year = date.getFullYear();

    if (month < 10) month = "0" + month;
    if (day < 10) day = "0" + day;

    var today = year + "-" + month + "-" + day;


    document.getElementById('ctime').value = today;
});

function valid(){
    v = document.getElementsByName("title")[0].value;
    if("" == v){
        alert("请输入标题");
        return false;
    }

    v = document.getElementsByName("content")[0].value;
    if("" == v){
        alert("请输入内容");
        return false;
    }

    v = document.getElementsByName("kind")[0].value;
    if("" == v){
        alert("请输入文章类型");
        return false;
    }


    v = document.getElementsByName("author")[0].value;
    if("" == v){
        alert("请输入作者");
        return false;
    }
    return true;
}
