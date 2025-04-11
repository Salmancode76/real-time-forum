import { BasePage } from "./BasePage.js";
import {socket} from './socket.js'
import * as session from './Session.js'

let currentChatUser = null;
let currentUser = session.testCookie();

let currentname=null;
let userChatStates = {};

function getUserChateState(username) {
  if(!userChatStates[username]) {
    userChatStates[username] ={
      set:0,
      isPrependMessages:false
    };
  }
  return userChatStates[username];
}
  //let isPrependMessages = false;

   //let set = 0;


export class Chat extends BasePage {
  constructor() {
    super(Chat.getHtml("chat_container"));
   this.CheckAuth();
  }

  static getHtml() {
    return `
      <div class="container" id="chat_container">
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
    if (!currentChatUser) return;
  const userState = getUserChateState(currentChatUser.name);
  //set += 10;
  //isPrependMessages = true;

  userState.set += 10;
  userState.isPrependMessages = true;
  socket.send(
    JSON.stringify({
      type: "get_chat_history",
      from: currentChatUser.name,
      to: currentUser,
      set: userState.set,
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
        if(currentChatUser){
          const userState = getUserChateState(currentChatUser.name);
          userState.isPrependMessages = true;
        }
     // isPrependMessages = true;
          

   }
  });

  
    console.log("Scrolled to: ", messageContainer.scrollTop);
}



export function loadUsers(){
  

  console.log("connecting to users====>")
  socket.send(JSON.stringify({
    "type": "get_users",
    to: currentUser,
  }));
}

export function PM(msg){
  console.log("PM=====>"+msg.from)
  console.log("PM=====>"+msg.message)
 let messagesDiv = document.getElementById("messages");
  const now = new Date();
      const formattedTime = formatTime(now);

      // Create a div element for the message
      //const messageDiv = document.createElement('div');
      //messageDiv.classList.add('message'); // add class for styling

      // Set the text content of the div
      //messageDiv.textContent = msg.from + " (" + formattedTime + "): " + msg.message;

      // Append the div to the messagesDiv
      msg.createdat = formattedTime;
      msg.text = msg.message;
      
     messagesDiv.appendChild(buildMessageDiv(msg, msg.from));

      readMessage(currentChatUser.name, currentUser);
            const messageContainer =
              document.getElementById("message-container");

        messageContainer.scrollTop = messageContainer.scrollHeight; 

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
                document
                  .getElementById("user-" + user.name)
                  .querySelector(".username").style.color = "white";
           /*
           make text white after click
     
            */

           const userState = getUserChateState(user.name);
           userState.set = 0;
           userState.isPrependMessages = false;
           // isPrependMessages = false;
           console.log(
             `Clicked on user: ${user.name}, reset state to: set=${userState.set}`
           );

           const dmDiv = document.getElementById("DM");
           if (dmDiv) {
             dmDiv.innerHTML = "";
           }

           // Create fresh DM interface
           DM();
           currentChatUser = user;

           // set = 0;
           socket.send(
             JSON.stringify({
               type: "get_chat_history",
               from: user.name,
               to: currentUser,
               set: userState.set,
             })
           );

           console.log("Clicked on user: " + user.name);

           // let messagesDiv = document.getElementById("message-container");
           //messagesDiv.innerHTML = "";

           //adding
           //setTimeout(() => listenSroll(), 500);
           setTimeout(() => {
             const messageContainer =
               document.getElementById("message-container");
               messageContainer.scrollTop = messageContainer.scrollHeight;
             
           }, 10); 

           readMessage(user.name, currentUser);
         });
                  

 
         usersDiv.appendChild(userContainer);
        
     }
   
 }


 export function showFrinds(msg){
  // const data = JSON.parse(msg);
   //console.log(msg)
     let data =msg
     // update the div with the list of users
     const usersDiv = document.getElementById("user-list");
     usersDiv.innerHTML = "";
     //const sorted = sort(data.users)
     for (const user of data.users) {
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
       
           const userState = getUserChateState(user.name);
           userState.set = 0;
           userState.isPrependMessages = false;
           // isPrependMessages = false;
           console.log(
             `Clicked on user: ${user.name}, reset state to: set=${userState.set}`
           );

           const dmDiv = document.getElementById("DM");
           if (dmDiv) {
             dmDiv.innerHTML = "";
           }

           // Create fresh DM interface
           DM();
           currentChatUser = user;

           // set = 0;
           socket.send(
             JSON.stringify({
               type: "get_chat_history",
               from: user.name,
               to: currentUser,
               set: userState.set,
             })
           );

           console.log("Clicked on user: " + user.name);

           // let messagesDiv = document.getElementById("message-container");
           //messagesDiv.innerHTML = "";

           //adding
           //setTimeout(() => listenSroll(), 500);
           setTimeout(() => {
             const messageContainer =
               document.getElementById("message-container");
               messageContainer.scrollTop = messageContainer.scrollHeight;
             
           }, 10); 

           readMessage(user.name, currentUser);
         });
                  

 
         usersDiv.appendChild(userContainer);
        
     }
   
 }



 function readMessage(sender,receiver){



  socket.send(
    JSON.stringify({
      type: "read_message",
      from: sender,
      to: receiver,
    })
  );
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
  showMessages(data, currentUser, currentChatUser.name, getUserChateState(currentChatUser.name).set, getUserChateState(currentChatUser.name).isPrependMessages);
  readMessage(currentChatUser.name,currentUser);

}

function DM(){
  const dmDiv = document.getElementById("DM");
  dmDiv.innerHTML = "";
  let template = `<div id="message-container" tabindex="0">

    <div id="messages">
   
    </div>
    <div class="sender_div">
    <form id="send-container">
      <input type="text" id="message-input" placeholder="Type your message...">
      <button id="send_message" type="submit">send</button>
    </form>
    </div>
  </div>`;

  dmDiv.innerHTML = template;

  const messageContainer = document.getElementById("message-container");
  messageContainer.addEventListener("focus",()=>{
    if(currentChatUser) 
      {
         readMessage(currentChatUser.name,currentUser);
      }
      })

  listenSroll();






  }
  /*
 export function  updateReadStatus(sender){
      const messages = document.querySelectorAll(".message");
      messages.forEach((message) => {
        const messageSender = message.querySelector(".sender").textContent;
        if (messageSender === sender) {
          const tick = message.querySelector(".timestamp");
          tick.textContent = "✓✓";
          tick.style.color = "#427eff";
        }
      });
  }
      */



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

  form.addEventListener("submit",  (event)=> {
    event.preventDefault();
    let nameUser = getname()
    const message = (messageInput.value).trim();
    if(message === ""){
      alert("MESSAGE CAN'T BE EMPTY");
      location.reload();
      return;
    }
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

       readMessage(to, from);

    messageInput.value = "";

    const now = new Date();
    const formattedTime = formatTime(now);

    // Create a div element for the message
      let messagesDiv = document.getElementById("messages");

    //const messageDiv = document.createElement("div");
   // messageDiv.classList.add("message","self"); // add class for styling

    // Set the text content of the div
   // messageDiv.textContent =
    //  nameUser + " (" + formattedTime + "): " + message;
   const msg = {
     from: getname(),
     createdat: formattedTime,
     text: message.trim(),
     to: to,
   };

    //console.table(msg);
    messagesDiv.appendChild(buildMessageDiv(msg, to));

    // Append the div to the messagesDiv
    
  });
}

function getname(){
  const allCookies = document.cookie;
 // console.log(allCookies)
  // Split the cookies into an array
  const cookiesArray = allCookies.split(';');
  
  // Loop through the array to find the specific cookie
  let sessionCookieValue = null;
  cookiesArray.forEach(cookie => {
      const [name, value] = cookie.trim().split('=');
      if (name === 'userName') {
          sessionCookieValue = value;
      }
  });
  
  //.log('Session Cookie Value:', sessionCookieValue);
  return sessionCookieValue
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

  /*
  let tick = document.createElement("div");
  tick.classList.add("timestamp");
  if (msgData.isread && !isSelf) {
    tick.textContent = "✓✓";
    tick.style.color = "#427eff"; // Hex code for blue
  } 
     div.appendChild(tick);
     */

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

export function offline(msg){
  console.log(msg)
  const div = document.querySelectorAll('.user-container');
  for (const element of div) {
    if (element.textContent.trim() === msg.to) {
      const colorSpan = element.querySelector('span[style*="background-color: green"]');
  
      if (colorSpan) { // Check if the span was found.
        colorSpan.style.backgroundColor = "red";
        console.log("green printed should change");    }
  }
}

}

export function online(msg){
  const divs = document.querySelectorAll('.user-container');

  for (const key in msg.online) {
     const value = msg.online[key];
  
     for (let i = 0; i < divs.length; i++) {
      if (divs[i].textContent.trim() === value) {
        // Find the span with the inline background color style.
        const colorSpan = divs[i].querySelector('span[style*="background-color: red"]');
  
        if (colorSpan) { // Check if the span was found.
          colorSpan.style.backgroundColor = "green";
          console.log("green printed should change");
        }
  }
}
  }
}

export function Unread(msg) {
  console.log("Unread notification received:", msg);
  const divs = document.querySelectorAll(".user-container");

  // Check if msg.users exists and is an array
  if (!msg.users || !Array.isArray(msg.users)) {
    console.error("Invalid users data in notification:", msg);
    return;
  }

  for (const user of msg.users) {
    for (let i = 0; i < divs.length; i++) {
      // Compare with user.name instead of user object
      if (divs[i].textContent.includes(user.name)) {
        const usernameSpan = divs[i].querySelector(".username");
        if (usernameSpan) {
          usernameSpan.style.color = "red";
          console.log("Changed text color to red for user:", user.name);
        }
      }
    }
  }
}