import { BasePage } from "./BasePage.js";
import {socket} from './socket.js'
import * as session from './Session.js'

let currentChatUser = null;
let currentUser = session.testCookie();

let isPrependMessages= false;
let currentname=null;

   let set = 0;


export class Chat extends BasePage {
  constructor() {
    super(Chat.getHtml());
   this.CheckAuth();
  }

  static getHtml() {
    return `
      <div class="container">
        <div class="chats" id="DM">
              <!--  

          <h2>Column 1</h2>
          <p>Some text for the first column. This can be any HTML content.</p>
          -->
        </div>

      <!--   <div class="column"> 
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
           
                     </div>

            -->


          <div class="user_lists" id="user-list">
            <h3>Row 2</h3>
            <p>Content for the second row within the second column.</p>
          </div>
        </div>
      </div>
    `;
  }


}

const throttledGetHistory = throttle(() => {
  set += 10;
  isPrependMessages = true;
  socket.send(
    JSON.stringify({
      type: "get_chat_history",
      from: currentChatUser.name,
      to: currentUser,
      set: set,
    })
  );
}, 100);

function throttle(func,delay){
  let LastTime=0;
  return function(...args){
    let now =  Date.now();
    if(now - LastTime >= delay){
      LastTime=now;
      func.apply(this, args);
    }
  }
}


function listenSroll(){
  const messageContainer = document.getElementById("message-container");

  messageContainer.addEventListener("scroll", () => {
    
   if (messageContainer.scrollTop <= 0 ) {

     //alert("Time for new messages " + set);
       //showMessages(data, currentUser, currentChatUser.name, set);
 
      throttledGetHistory();
      messageContainer.scrollTop = 10;
           
         // alert(currentChatUser.name);
          //alert(currentname);
      isPrependMessages = true;
          

   }
  });
    console.log("Scrolled to: ", messageContainer.scrollTop);
}



export function loadUsers(){
  

  console.log("connecting to users====>")
  socket.send(JSON.stringify({"type": "get_users"}));
}

export function PM(msg){
  console.log("PM=====>"+msg.from)
  console.log("PM=====>"+msg.message)
 let messagesDiv = document.getElementById("messages");
  const now = new Date();
      const formattedTime = formatTime(now);

      // Create a div element for the message
      const messageDiv = document.createElement('div');
      messageDiv.classList.add('message'); // add class for styling

      // Set the text content of the div
      messageDiv.textContent = msg.from + " (" + formattedTime + "): " + msg.message;

      // Append the div to the messagesDiv
      messagesDiv.appendChild(messageDiv);

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

           isPrependMessages = false;
           set = 0;
           socket.send(
             JSON.stringify({
               type: "get_chat_history",
               from: user.name,
               to: currentUser,
               set: set,
             })
           );
           console.log("Clicked on user: " + user.name);
           currentChatUser = user;

           let messagesDiv = document.getElementById("message-container");
           messagesDiv.innerHTML = "";

           //adding
          setTimeout(() => listenSroll(), 200);

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
  
  if (!document.getElementById("messages")) {
    DM();

    
  }  
  showMessages(data, currentUser, currentChatUser.name, set, isPrependMessages);

}

function DM(){
  const dmDiv = document.getElementById("DM");
  dmDiv.innerHTML = "";
  let template = `<div id="message-container">
    <div id="messages">
   
    </div>
    <form id="send-container">
      <input type="text" id="message-input" placeholder="Type your message...">
      <button id="send_message" type="submit">send</button>
    </form>
  </div>`;

  dmDiv.innerHTML = template;

  listenSroll();



  }


//build messages 
function showMessages(data, from, to, set, isPrependMessages) {
  console.log(data);
  let messagesDiv = document.getElementById("messages");
  //messagesDiv.innerHTML = '';


  
  if(isPrependMessages) {
  for (let i = 0; i <= data.length - 1; i++) {
    messagesDiv.prepend(buildMessageDiv(data[i], to));
  }
   
  }else{
    
  for (let i = data.length - 1; i >= 0; --i) {
    messagesDiv.appendChild(buildMessageDiv(data[i], to));
  }
      //messageContainer.scrollTop = messageContainer.scrollHeight;

  }

  // scrollToBottom();

  //Scrolling down when load


      //messageContainer.scrollTop = 10;

  


  const form = document.getElementById("send-container");
  const messageInput = document.getElementById("message-input");

  form.addEventListener("submit", function (event) {
    event.preventDefault();

    const message = messageInput.value;
    console.log("Message:", message);
    if(message.trim()=== "") {
      return;
    }
    socket.send(
      JSON.stringify({
        type: "message",
        From: from,
        To: to,
        Text: message,
        set: set,
      })
    );
    messageInput.value = "";

    const now = new Date();
    const formattedTime = formatTime(now);

    // Create a div element for the message
    const messageDiv = document.createElement("div");
    messageDiv.classList.add("message"); // add class for styling

    // Set the text content of the div
    messageDiv.textContent =
      currentname + " (" + formattedTime + "): " + message;

    // Append the div to the messagesDiv
    messagesDiv.appendChild(messageDiv);
  });
}


// Function to build message div
function buildMessageDiv(msgData, to) {
  console.log(to);
  const isSelf = to === msgData.from;
;

  let div = document.createElement("div");
  div.classList.add("message", !isSelf ? "self" : "other");

  let divName = document.createElement("div");
  divName.classList.add("sender");
  divName.textContent = msgData.from;

  let divTime = document.createElement("div");
  divTime.classList.add("timestamp");
  divTime.textContent = msgData.createdat;

  let divText = document.createElement("div");
  divText.classList.add("message-content");
  divText.textContent = msgData.text;

    
    div.appendChild(divName);
  div.appendChild(divText);
  div.appendChild(divTime);

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





