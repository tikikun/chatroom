var url = "wss://" + window.location.host + window.location.pathname + "/ws";
var ws = new WebSocket(url);
var names = ["Gau cho", "Tho chiahuahua", "Heo nai", "Khi dot", "Sieu nhan", "Nguoi nhen", "Than sam", "Nguoi doi", "Racoon"];
var name = names[Math.floor(Math.random() * names.length)] + "  ID: " + Math.floor(Math.random() * 1000);
var chat = document.getElementById("chat");
var text = document.getElementById("text");
var roomid = document.getElementById("roomid");
var button = document.getElementById("button");
var buttonlog = document.getElementById("showlog")
var roombutton = document.getElementById("roombutton");
var room = document.getElementsByClassName("room");
var roomName = window.location.pathname.split("/")[2];

//This function to calculate the moment the message is sent
var now = function () {
  var iso = new Date().toISOString();
  return iso.split("T")[1].split(".")[0];
};

function getLog() {
  $.get("https://" + window.location.host + window.location.pathname+ "/log", function (data) {
    chat.innerText = data + "\n\n--------------$--------------\n\n";
    text.scrollIntoView(false);
  });
}



chat.innerText = "The chat is loading please wait until you are connected\n\n";

ws.onopen = function () {
  text.value = "";
  for (i in room) {
    room[i].innerText = decodeURIComponent(roomName.toUpperCase());
  };
  ws.send(name + " has joint the chat on " + now());
};

// Get and display message
ws.onmessage = function (socket) {
  chat.innerText += socket.data + "\n ";
  text.scrollIntoView(false);
};

// Message
function sendMsg(username, textval) {
  val = now() + " " + "<" + username + "> :";
  val = val.padStart(40, '-');
  ws.send(val + textval);
  textval = "";
};

//Send message to the websocket
text.onkeydown = function (e) {
  if (e.keyCode === 13 && text.value !== "") {
    sendMsg(name, text.value)
    text.value = "";
  }
}

button.onclick = function () {
  if (text.value !== "") {
    sendMsg(name, text.value)
    text.value = "";
  }
};

roombutton.onclick = function () {
  window.location.replace("https://" + window.location.host + "/room/" + roomid.value);
};

buttonlog.onclick = function (){
  getLog()
  buttonlog.style.display = "none";
}