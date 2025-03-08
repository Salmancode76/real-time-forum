
import { socket } from "./socket.js";

export function DM(){
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

export function showMessages(data,from,to) {
  console.log(data)
  let messagesDiv = document.getElementById("messages");
  messagesDiv.innerHTML = '';

      for (let i = 0; i < data.length; ++i) {

          messagesDiv.appendChild(buildMessageDiv(data[i]));
      }
   
      const form = document.getElementById('send-container');
      const messageInput = document.getElementById('message-input');
      
      form.addEventListener('submit', function(event) {
        event.preventDefault(); // Prevent page refresh
        
        const message = messageInput.value;
        console.log('Message:', message);
          socket.send(JSON.stringify({
           "type": "message",
           "From":from,
           "To":to,
          "Text":message}));
        console.log("it shouldnet reload")
        // Optionally clear the input field
        messageInput.value = '';
 
      });

      
  }


// Function to build message div
function buildMessageDiv(msgData) {
  let div = document.createElement('div');
  div.classList = 'message';
  div.innerText = msgData.from + " (" + msgData.createdat + "): " + msgData.text;
  return div;
}


