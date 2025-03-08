import { socket } from "./socket.js";
import { DM,showMessages } from "./DM.js";

let currentChatUser = null;


export function showUsers(msg){
 // const data = JSON.parse(msg);
  //console.log(msg)
    let data =msg
    // update the div with the list of users
    const usersDiv = document.getElementById("user-list");
    usersDiv.innerHTML = "";
    const sorted = sort(data.users)
    for (const user of sorted) {
      const userContainer = document.createElement("div");
        userContainer.className = "user-container";
        userContainer.id = "user-" + user.name;

        const greenCircle = document.createElement("span");
        //to be changed to red
        greenCircle.style.backgroundColor = "green";
        greenCircle.style.width = "10px";
        greenCircle.style.height = "10px";
        greenCircle.style.borderRadius = "50%";
        greenCircle.style.display = "inline-block";
        greenCircle.style.marginRight = "10px";
        greenCircle.style.marginLeft = "10px";
        userContainer.appendChild(greenCircle);

        const usernameElem = document.createElement("span");
        usernameElem.className = "username";
        usernameElem.textContent = user.name;
        userContainer.appendChild(usernameElem);

        // If the user has a last message, display it in the user list (this is the last message they sent or received)
        if (currentChatUser === null || currentChatUser.id !== user.id) {
          const bubbleGif = document.createElement("img");
         // bubbleGif.src = "/style/bubble.gif";
          bubbleGif.className = "bubble";
          bubbleGif.style.marginLeft = "10px";
          bubbleGif.setAttribute("data-recipient", user.name);
          //console.log("Added recipient ID: " + user.username)
          userContainer.appendChild(bubbleGif);
        }

        userContainer.addEventListener("click", () => {
          //getHistoy()
          socket.send(JSON.stringify({"type": "get_chat_history","from":user.name,"to":"hamad.ar"}));
          console.log("Clicked on user: " + user.name);
          currentChatUser = user;
        });

        usersDiv.appendChild(userContainer);
       
    }
  
}
export function oldmessagesofserv(data){
  
  DM();
  showMessages(data,"hamad.ar",currentChatUser.name);
}


export function loadUsers(){
  

    console.log("connecting to users====>")
    socket.send(JSON.stringify({"type": "get_users"}));
}

function sort(arr) {
  if (!Array.isArray(arr)) {
    return "Input is not an array.";
  }

  return arr.slice().sort((a, b) => {
    const nameA = a.Name ? a.Name.toLowerCase() : ""; // Handle potential missing or undefined names
    const nameB = b.Name ? b.Name.toLowerCase() : "";

    if (nameA < nameB) {
      return -1;
    }
    if (nameA > nameB) {
      return 1;
    }
    return 0; // Names are equal
  });
}