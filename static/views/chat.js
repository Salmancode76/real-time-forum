import { BasePage } from "./BasePage.js";
import {socket} from './socket.js'
import * as session from './Session.js'

let currentChatUser = null;
let currentUser = session.testCookie()

export class Chat extends BasePage {
  constructor() {
    super(Chat.getHtml());
   this.CheckAuth();
  }

  static getHtml() {
    return `
      <div class="container">
        <div class="column" id="DM">
          <h2>Column 1</h2>
          <p>Some text for the first column. This can be any HTML content.</p>
        </div>
        <div class="column">
          <h2>Column 2</h2>
          <div class="nested-row">
            <h3>Row 1</h3>
            <p>Content for the first row within the second column.</p>
            <input type="text" id="From" placeholder="From">
            <br>
            <input type="text" id="To" placeholder="To">
            <br>
            <input type="text" id="messageInput" placeholder="Enter a message">
            <button id="clicked">Send Message</button>
            <div id="output"></div>
          </div>
          <div class="nested-row" id="user-list">
            <h3>Row 2</h3>
            <p>Content for the second row within the second column.</p>
          </div>
        </div>
      </div>
    `;
  }


}
export function loadUsers(){
  

  console.log("connecting to users====>")
  socket.send(JSON.stringify({"type": "get_users"}));
}


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
         greenCircle.style.backgroundColor = "red";
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
           socket.send(JSON.stringify({"type": "get_chat_history","from":user.name,"to":currentUser}));
           console.log("Clicked on user: " + user.name);
           currentChatUser = user;
         });
 
         usersDiv.appendChild(userContainer);
        
     }
   
 }


 function sort(arr) {
  if (!Array.isArray(arr)) {
    return "Input is not an array.";
  }

  return arr.slice().sort((a, b) => {
    const nameA = a.name ? a.name.toLowerCase() : ""; // Handle potential missing or undefined names
    const nameB = b.name ? b.name.toLowerCase() : "";

    if (nameA < nameB) {
      return -1;
    }
    if (nameA > nameB) {
      return 1;
    }
    return 0; // Names are equal
  });
}

export function oldmessagesofserv(data){
  
  DM();
  showMessages(data,currentUser,currentChatUser.name);
}

function DM(){
  const dmDiv = document.getElementById("DM");
  dmDiv.innerHTML = "";
let template = `<div id="message-container">
    <div id="messages">
   
    </div>
    <form id="send-container">
      <input type="text" id="message-input" placeholder="Type your message...">
      <button id="submit" type="submit">send</button>
    </form>
  </div>`
  
  dmDiv.innerHTML = template;
 


}


//build messages 
function showMessages(data,from,to) {
console.log(data)
let messagesDiv = document.getElementById("messages");
messagesDiv.innerHTML = '';

for (let i = data.length - 1; i >= 0; --i) {
  messagesDiv.appendChild(buildMessageDiv(data[i]));
}
 
    const form = document.getElementById('send-container');
    const messageInput = document.getElementById('message-input');
    
    form.addEventListener('submit', function(event) {
      event.preventDefault();

      const message = messageInput.value;
      console.log('Message:', message);
      socket.send(JSON.stringify({
          "type": "message",
          "From": from,
          "To": to,
          "Text": message
      }));
      messageInput.value = '';

      const now = new Date();
      const formattedTime = formatTime(now);

      // Create a div element for the message
      const messageDiv = document.createElement('div');
      messageDiv.classList.add('message'); // add class for styling

      // Set the text content of the div
      messageDiv.textContent = from + " (" + formattedTime + "): " + message;

      // Append the div to the messagesDiv
      messagesDiv.appendChild(messageDiv);
  });
    
}


// Function to build message div
function buildMessageDiv(msgData) {
let div = document.createElement('div');
div.classList = 'message';
div.innerText = msgData.from + " (" + msgData.createdat + "): " + msgData.text;
return div;
}


function formatTime(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-indexed
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}





