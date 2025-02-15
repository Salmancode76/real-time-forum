import { BasePage } from "./BasePage.js";

export class ViewPost extends BasePage{
    constructor(id){
        super(ViewPost.getHtml())
        this.CheckAuth("view");
        this.GetPost(id);
    }

  async  GetPost(id){
        const response = await fetch(`/post?id=${id}`, {
          headers: {
            Accept: "application/json",
          },
          cache: "no-store",
        });
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
        const data = await response.json();
        console.table(data);
      const post = data.Posts[0][0];
        document.getElementById("postview").innerHTML = `
        
        
        
        <div>
        <div class="post_title"> ${post.title} </div>
        
        <div class="post_info"> <span> Posted by: ${
          post.userId
        } </span> <span> Created At:  ${new Date(
          post.date
        ).toLocaleDateString()} </span>   </div>

        
        <div class="Viewcategory"> ${post.categories.map(
          (category) => `<div id="catgeory_viewpost">${category}</div>`
        ).join(" ")}
        </div>
       <div class="view_content"> ${post.content}</div>
       
     
       `;

       
    }
    static getHtml(){

        return `
        <div id="viewPost">
          <div id="postview"> 
          
          
          </div>

        
        </div>


          <div class="post_comment">
              <div class="submit_comment">

              <span class="header"> Comments </span>
              <br>
              <textarea name="commentsection" class="commentsection form-input" required placeholder="Write your post comment here"></textarea>
            <button class ="commentBtn">  Submit </button>

              </div>
          </div>
        `;
    }

}