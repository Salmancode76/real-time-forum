const socket = new WebSocket("ws://localhost:8082/");


    
socket.addEventListener("message", (event) => {
    const data = JSON.parse(event.data);
  
    if (data.type === "users") {
      // update the div with the list of users
      const usersDiv = document.getElementById("user-list");
      usersDiv.innerHTML = "";
  
      for (const user of data.users) {
        const userElem = document.createElement("div");
        userElem.textContent = user.name;
        usersDiv.appendChild(userElem);
      }
    }

  });


export function sendMessage() {
    const messageInput = document.getElementById("messageInput");
    const messageFrom =document.getElementById("From")
    const messageTo =document.getElementById("To")
    message.text = messageInput.value;
    message.from= messageFrom.value
    message.to=messageTo.value
    socket.send(JSON.stringify(message));
    messageInput.value = "";
}

function loadUsers(){
    console.log("connecting to user")
    socket.send(JSON.stringify({"type": "get_users"}));
}