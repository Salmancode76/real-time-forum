import { Home } from './views/Home.js';
import { signup } from './views/signup.js';
import { s_test } from "./views/s_testing.js";
import { login } from "./views/login.js";
import { CreatePost } from "./views/createPost.js";
import { ViewPost} from "./views/ViewPost.js"

handleInitialLoad();

document.getElementById('hamICON').addEventListener('click', ()=> {
    const links = document.querySelector('.nav-links');
    links.classList.toggle('active');
});

function handleInitialLoad() {
  //take the path with the parameters
    const path = window.location.pathname + window.location.search;
    console.log(path);
    navigateTo(path);
}

//handle navigation
export async function navigateTo(route) {
  
     console.log("Navigating to:", route);

     const url = new URL(route, window.location.origin);
     const path = url.pathname;

     console.log("Path:", path);


    const nav = document.getElementById("nav");
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
      
        
        switch (path) {
          case "/":
            new Home();
            break;
          case "/s":
            await new s_test();
            break;
          case "/sign":
            new signup();
            break;
          case "/login":
            new login();
            break;
          case "/logout":
            new login();
            break;
          case "/createPost":
            new CreatePost();
            break;
         case "/post":
            const id = url.searchParams.get("id");
            
          if (id) {
            await new ViewPost(id);
          } else {
            app.innerHTML = "Post ID is missing!";
          }
          break;

       
            
          default:
          
          app.innerHTML = `
            <div id="error_container">
            <div id="Error_title">
                404 
            </div>

            <div id="error_info">
Sorry, the page you're looking for doesn't exist
            </div>

          </div>
            `;
        }

        history.pushState(null, "", url.pathname + url.search);
    } catch (error) {
      
          app.innerHTML = `
            <div id="error_container">
            <div id="Error_title">
                500 
            </div>

            <div id="error_info">

            Internal server error ${error}
            </div>

          </div>
            `;
  console.error("500 Internal Server Error: ", error);

      
    }
}



//Making the router function golablly accessable
window.navigateTo = navigateTo;
