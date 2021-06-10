let input_name = document.getElementById("name");
let input_text = document.getElementById("text");
let output = document.getElementById("msg");
let socket = new WebSocket("ws://127.0.0.1:80/comment-echo");
let l_socket = new WebSocket("ws://127.0.0.1:80/comment-load");

socket.onopen = function () {
};
socket.onerror = function(e){
  alert(e)
};
socket.onclose = function(){
  alert("connection lost")
};
socket.onmessage = function (e) {
    let x = JSON.parse(e.data)
    //output.innerHTML += x.name + " : " + x.text + "\n"
    create_new_msg(x)
};

function create_new_msg(x) {
    let box_div = document.createElement('div');
    let name_div = document.createElement('div');
    let text_div = document.createElement('div');
    let date_div = document.createElement('div');
    box_div.setAttribute('class' , 'box');
    name_div.textContent = x.name;
    name_div.setAttribute('class', 'name');
    text_div.textContent = x.text;
    text_div.setAttribute('class' , 'comment-content');
    date_div.textContent = x.date;
    date_div.setAttribute('class' , 'comment-date');
    box_div.appendChild(name_div);
    box_div.appendChild(text_div);
    box_div.appendChild(date_div);
    output.appendChild(box_div);
    scroll_down(output);

}

function send() {
    socket.send(JSON.stringify(
        {
            name : input_name.value,
            text : input_text.value,
            date : currdate(),
            url : window.location.href,
            confirmed : 0
        }
    ));
    clear()
}
function clear() {
    input_text.value = "";
    input_name.value = "";
}
function currdate() {
    let today = new Date();
    let date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
    let time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    let dateTime = date+' '+time;
    return dateTime;

}

l_socket.onopen = function () {
    l_socket.send(window.location.href)
};
l_socket.onmessage = function (e) {
    let sts = JSON.parse(e.data);
    for(let i = 0 ; i < sts.length ; i++){
        create_new_msg(sts[i])
    }
};




function scroll_down(x) {
    x.scrollTop = x.scrollHeight
}