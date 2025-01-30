export class BasePage{
    constructor(html) {
        this.Html =html
        const app = document.getElementById('app');
        app.innerHTML = this.Html;
    }
    
  
}