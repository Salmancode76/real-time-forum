import { Home } from './views/Home.js';
import { signup } from './views/signup.js';
import { s_test } from "./views/s_testing.js";
import { login } from "./views/login.js";
import { CreatePost } from "./views/createPost.js";

handleInitialLoad();

document.getElementById('hamICON').addEventListener('click', ()=> {
    const links = document.querySelector('.nav-links');
    links.classList.toggle('active');
});

function handleInitialLoad() {
    const path = window.location.pathname;
    console.log(path);
    navigateTo(path);
}

export async function navigateTo(route) {
    const nav = document.getElementById('nav');

    nav.innerHTML = `
    <nav>
        <ul>
            <li class="header">  <a href="/" onclick="navigateTo('/');">  Community Forum </a> </li>
            <div class="nav-links">
                <li><a <a href="/sign" onclick="navigateTo('/sign');"> Sign-Up</a></li>
                <li><a <a href="/login" onclick="navigateTo('/login');"> Login</a></li>

            </div>
            <img id="hamICON" src="/static/images/ham_menu.svg" alt="Menu">
        </ul>
    </nav>
    `;

    const app = document.getElementById('app');


 
    try {
        switch (route) {
          case "/":
            new Home();
            break;
          case "/s":
            await new s_test();
            break;
          case "/sign":
            new signup();
            break;
          case "/login" || "/logout":
            new login();
            break;
          case "/createPost":
            new CreatePost();
            break;

          default:
            app.innerHTML = "404 Not found :(";
        }

        history.pushState(null, '', route);
    } catch (error) {
        app.innerHTML = "500" + error;
    }
}



window.navigateTo = navigateTo;
