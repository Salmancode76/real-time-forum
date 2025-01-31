import { Home } from './Handlers/Home.js';
import { signup } from './Handlers/signup.js';
import { s_test } from "./Handlers/s_testing.js";
import {login} from "./Handlers/login.js"

handleInitialLoad();

document.getElementById('hamICON').addEventListener('click', function() {
    const links = document.querySelector('.nav-links');
    links.classList.toggle('active');
});

function handleInitialLoad() {
    const path = window.location.pathname;
    console.log(path)
    navigateTo(path);
}
async function navigateTo(route) {
    const nav = document.getElementById('nav');

    nav.innerHTML = `
    <nav>
        <ul>
            <li class="header">  <a href="/" onclick="navigateTo('/'); return false;">  Community Forum </a> </li>
            <div class="nav-links">
                <li> <a href="/" onclick="navigateTo('/'); return false;"> Posts</a></li>

                <li><a href="#">All users</a></li>
                <li><a href="#">Private Messages</a></li>
                <li><a href="#">Welcome, Stranger</a></li>
                <li><a href="#">Logout</a></li>
            </div>
            <img id="hamICON" src="/static/icons8-hamburger.svg" alt="Menu">
        </ul>
    </nav>
    
    
    `
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
