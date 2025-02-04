
export class s_test  {
    constructor() {
        this.init();
    }

    async init() {
        try {
            // Fetch and render the HTML content
            const html = await this.getHtml();
            this.renderHtml(html);
        } catch (error) {
            console.error('Failed to initialize s_test page:', error);
            this.renderHtml(`<div>Error: ${error.message}</div>`); // Render an error message
        }
    }

    async getHtml() {
        const response = await fetch('/s', {
            headers: { 
                'Accept': 'application/json',
            },
            cache: 'no-store'
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json(); // Parse the JSON response

        return `
            <div id="s-page" class="page">
                <h1>S Page</h1>
                <div id="s-content">
                    <p><strong>Message:</strong> ${data.message}</p>
                    <p><strong>Route Name:</strong> ${data.routeName}</p>
                    <p><strong>Show Request Made:</strong> ${data.showRequestMade}</p>
                </div>
                <a href="/" onclick="navigateTo('/'); return false;">Back Home</a>
            </div>
        `;
    }

    renderHtml(html) {
        const app = document.getElementById('app');
        if (app) {
            app.innerHTML = html;
        } else {
            console.error("Element with id 'app' not found.");
        }
    }
}