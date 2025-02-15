import { initWebSocket, sendMessage } from "./WebSocket.js";
import { BasePage } from "./BasePage.js";

export class Home extends BasePage {
  constructor() {
    super(Home.getHtml());
    this.CheckAuth("home");
    this.init();
    this.FetchPosts();
  }

  static getHtml() {
    return `
            <div id="home" class="page active" style="display: none;">
            <div class="dashboard-header">
                  <h1>Dashboard</h1>
                  <a href="/createPost" class="btn">Create Post</a>
                  </div>
              <div id ="posts">  
              </div>
            </div>
        `;
  }

  async FetchPosts() {
    try {
      const response = await fetch("/api/posts");
      const data = await response.json();

      const posts = data.Posts[0]

      console.log(posts)

      document.getElementById("posts").innerHTML =
        posts.length > 0
          ? posts
              .map(
                (post) =>
                  `
            <div class="post_card">
              <span id="post_title">${post.title || "Untitled Post"}</span>
              <div class="user_date">
                  <div>Posted by: ${post.userId}</div>
                  <div>Created At: ${
                    post.date
                      ? new Date(post.date).toLocaleDateString()
                      : "Unknown date"
                  }</div>
              </div>
          <div class="category">
        ${post.categories
          .map((category) => `<div class="category-tag">${category}</div>`)
          .join(" ")}
      </div>

              <div class="content">${
                post.content?.replace(/\n/g, "<br>") || "No content available"
              }</div>
  
  
 
    <a href="/post?id=${post.id}" 
                       onclick="event.preventDefault(); navigateTo('/post?id=${
                         post.id
                       }');" 
                       class="read_more"> Read More</a>
</div>
          `
              )
              .join("")
          : `<div class="no-posts">No posts to display</div>`;
            
            
    } catch (error) {
      console.error("Fetch error:", error);
      document.getElementById("posts").innerHTML = `
          <div class="error">Error loading posts. Please refresh the page.</div>
      `;
    }
  }

  init() {
    if (!window.socket || window.socket.readyState === WebSocket.CLOSED) {
      initWebSocket();
    }
    window.sendMessage = sendMessage;
  }
}
