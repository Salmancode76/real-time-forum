export class BasePage {
  constructor(html) {
    
    this.Html = html;
    const app = document.getElementById("app");
    app.innerHTML = this.Html;
  }

  

   async CheckAuth(divName) {
    try {
      const response = await fetch("/auth-check");

      if (response.status === 303) {
        alert(
          "You are not logged in to view this page. Redirecting to login page..."
        );
         window.location.href = "/login";

        return;
      }
    
         const nav = document.getElementById("nav");

         nav.innerHTML = `
            <nav>
                <ul>
                    <li class="header">  <a href="/" onclick="navigateTo('/');">  Community Forum </a> </li>
                    <div class="nav-links">
                        <li> <a href="/" onclick="navigateTo('/');"> Posts</a></li>
                        <li><a href="#">All users</a></li>
                        <li><a href="#">Private Messages</a></li>
                      <!--  <li><a href="#">Welcome, Stranger</a></li> -->
                        <li><a href="/logout" onclick="navigateTo('/login');" id="logout_link">Logout</a></li>
                    </div>
                    <img id="hamICON" src="/static/images/ham_menu.svg" alt="Menu">
                </ul>
            </nav>
    `;

        const hamIcon = document.getElementById("hamICON");
        if (hamIcon) {
          hamIcon.addEventListener("click",  ()=> {
            const links = document.querySelector(".nav-links");
            links.classList.toggle("active");
          });
        }

 const divElement = document.getElementById(divName);
 if (divElement) {
   divElement.style.display = "block";
 } else {
        console.error(`Element with ID ${divName} not found`);
 }    } catch (error) {
      console.error("Error checking authentication: ", error);
      document.getElementById("app").innerHTML =
        "Error checking authentication";
    }
  }
}