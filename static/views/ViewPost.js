import { BasePage } from "./BasePage.js";

export class ViewPost extends BasePage {
  constructor(id) {
    super(ViewPost.getHtml());
    this.id = id;
    this.CheckAuth("app");
    this.initialize()
    
  }
  async initialize(){
    try{
       await  this.GetPost(); 
    await  this.setUpForm();
    }catch(error){
      if (error.status === 404){
           document.getElementById("app").innerHTML = `
                    <div id="error_container">
                        <div id="Error_title">404</div>
                        <div id="error_info">Sorry, the post you're looking for doesn't exist</div>
                    </div>
                `;
           return;
      }
    }
  }

  async GetPost() {
    // Use this.id here
    const response = await fetch(`/post?id=${this.id}`, {
      headers: { Accept: "application/json" },
      cache: "no-store",
    });
    console.log(response.status)

     if (!response.ok) {
       const error = new Error(`HTTP error! status: ${response.status}`);
       error.status = response.status;
       throw error;
     }

    const data = await response.json();
    console.table(data);

    const post = data.Posts[0][0];
    document.getElementById("postview").innerHTML = `
      <div>
        <div class="post_title">${post.title}</div>
        <div class="post_info">
          <span>Posted by: ${post.userId}</span>
          <span>Created At: ${new Date(post.date).toLocaleDateString()}</span>
        </div>
        <div class="Viewcategory">
          ${post.categories
            .map(
              (category) => `<div class="category_viewpost">${category}</div>`
            )
            .join(" ")}
        </div>
        <div class="view_content">${post.content}</div>
      </div>
    `;
  post.Comments = post.Comments || [];

if (post.Comments.length > 0) {
  document.getElementById("userComments").innerHTML = `
    <div id="comments">
        ${post.Comments.map(
          (comment) => `
            <div class="comment_item">
            <div class="comment_data">
              <span> Posted by: ${comment.Username}</span>
              <spam>Created at: ${new Date(
                comment.date
              ).toLocaleDateString()} ${new Date(
            comment.date
          ).toLocaleTimeString()}   </span> 
            </div>
              <span id="comment_content">${comment.content}</span>
            </div>`
        ).join("")}
      </div>
    
    `;
}
  }

  static getHtml() {
    return `
      <div id="viewPost">
        <div id="postview"></div>
      </div>
      <form action="/post" method="POST" id="CommentForm">
        <div class="post_comment">
          <div class="submit_comment">
            <span class="header">Comments</span>
            <br>
            <textarea name="commentsection" id="commentsection" required placeholder="Write your post comment here"></textarea>
           <div class="commentBtnSec">
            <button class="commentBtn" type="submit">Submit</button>
            </div>
            </div>
        </div>
      </form>
    </div>
      <div id="userComments">

      </div>
    `;
  }

  async setUpForm() {
    const form = document.getElementById("CommentForm");
    if (!form) {
      console.error("Form not found!");
      return;
    }

    form.addEventListener("submit", (e) => {
      e.preventDefault(); 
      const Comment = {
        PostID: this.id,
        content: document.getElementById("commentsection").value,
      };

      this.CheckAuth()

      fetch("/createComment", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(Comment),
      });

      window.location.reload()
    });
  }
}
