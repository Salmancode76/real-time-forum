import { initWebSocket, sendMessage } from "./WebSocket.js";
import { BasePage } from './BasePage.js';

export class Home extends BasePage{
    constructor(){

        super(Home.getHtml());
        this.init();
        
    }
    static getHtml() {
        return `
            <div id="home" class="page active">
                <input type="text" id="messageInput" placeholder="Enter a message">
                <button class="homebtn" onclick="sendMessage()">Send Message</button>
                <div id="output"></div>
                <a href="/s" onclick="navigateTo('/s'); return false;">Go to S Page</a>
                <a href="/sign" onclick="navigateTo('/sign'); return false;">Go to Signup Page</a>
               <a href="/login" onclick="navigateTo('/login'); return false;">Go to Login Page</a>
            </div>
        `;
    }
    init() {
        if (!window.socket || window.socket.readyState === WebSocket.CLOSED) {
            initWebSocket();
        }
        window.sendMessage = sendMessage;
        
    }
}