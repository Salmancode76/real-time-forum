import { BasePage } from "./BasePage.js";


export class Chat extends BasePage {
  constructor() {
    super(Chat.getHtml());
   this.CheckAuth("home_main");
  }

  static getHtml() {
    return `
  
      <div class="container" >
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
