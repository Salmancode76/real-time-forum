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
      this.init();
         const nav = document.getElementById("nav");

         nav.innerHTML = `
            <nav>
                <ul>
                    <li class="header">  <a href="/" onclick="navigateTo('/');">  Community Forum </a> </li>
                    <div class="nav-links">
                        <li> <a href="/" onclick="navigateTo('/');"> Posts</a></li>
                        <li><a href="#">All users</a></li>
                        <li><a href="#">Private Messages</a></li>
                        <li><a href="#">Welcome, Stranger</a></li>
                        <li><a href="#">Logout</a></li>
                    </div>
                    <img id="hamICON" src="/static/images/ham_menu.svg" alt="Menu">
                </ul>
            </nav>
    `;

      document.getElementById(divName).style.display = "block";
    } catch (error) {
      console.error("Error checking authentication: ", error);
      document.getElementById("app").innerHTML =
        "Error checking authentication";
    }
  }
}