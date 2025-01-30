import { Home } from './Handlers/Home.js';
import { signup } from './Handlers/signup.js';
import { s_test } from "./Handlers/s_testing.js";
import {login} from "./Handlers/login.js"

handleInitialLoad();
window.addEventListener('popstate', () => {
    const path = window.location.pathname;
    navigateTo(path);
});

function handleInitialLoad() {
    const path = window.location.pathname;
    console.log(path)
    navigateTo(path);
}
async function navigateTo(route) {
    const app = document.getElementById('app');

    try {
        switch (route) {
            case '/':
                new Home();
                break;
            case '/s':
                await new s_test();
                break;
            case '/sign':
                new signup();
                break;
            case '/login':
                new login();
                break
            default:
                app.innerHTML = "404 Not found :(";
        }
        

        history.pushState(null, '', route);
    } catch (error) {
        app.innerHTML ="500" + error
    } 
}
