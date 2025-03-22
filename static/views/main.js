import {socket} from './socket.js'
import * as Chat from './chat.js'
import * as Session from './Session.js'



export function usersLoadfunc(){
socket.addEventListener("open",() => {
    Chat.loadUsers()
    console.log("showing users")
})
}
  
socket.addEventListener("message",(e) => {
    
//console.log(e.data)   
    
const msg = JSON.parse(e.data)
    
//console.log(msg)
    // sends messages to appropriate functions
    switch (msg.type) {
        case "users":
       Chat.showUsers(msg)
       Session.testCookie()
        break
        case "oldmessages":
        Chat.oldmessagesofserv(msg.chathistory)
        console.log(msg.chathistory)
        break
        case "chat":
        Chat.receiveChatMsg(msg)
        break
        case "post":
        Post.receivePostMsg(msg)
        break
        case "comment":
        Post.receiveCommentMsg(msg)
        break
        case "PM":
        Chat.PM(msg)
        break
        case "online":
        Chat.online(msg)
        break
        case "offline":
        console.log("offiline tringered")
        Chat.offline(msg)
        break
        case "message_read":
        //updateReadStatus(msg.from);
        //do nothing in the frontend 
         break
         
        default:
        console.log("Msg type not supported : " + msg.type)
        break
    }
}
)


