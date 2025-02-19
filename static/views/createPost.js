import { BasePage } from "./BasePage.js";

export class CreatePost extends BasePage {
  constructor() {
    super(CreatePost.getHtml());
      this.CheckAuth("home_main");
      this.setupFormSubmission();

  }


  static getHtml() {
    return `
            <div id="home_main" class="page active" style="display: none;">
                <div class="create_post"> 
                    <h1 class="header_create">Create New Post</h1>
                    <form action="/createPost" method="POST" id="PostForm">
                        <div class="Categories">
                            <span>Categories</span><br>
                            <label><input type="checkbox" name="category" value="Technology & Science"> Technology & Science</label><br>
                            <label><input type="checkbox" name="category" value="Health"> Health</label><br>
                            <label><input type="checkbox" name="category" value="Stories"> Stories</label><br>
                            <label><input type="checkbox" name="category" value="Others"> Others</label><br>
                        </div>
                        <div>
                            <label for="title">Title:</label><br>
                            <input type="text" name="title" id="title" required class="form-input" placeholder="Enter post title">
                        </div>
                        <div>
                            <label for="content">Content:</label>
                            <textarea name="content" id="content" required class="form-input" placeholder="Write your post content here"></textarea>
                        </div>
                        <button type="submit" class="submit_button">Create Post</button>
                    </form>
                </div>
            </div>
        `;
  }

  setupFormSubmission(){

    const form = document.getElementById("PostForm");
      if (!form) {
        console.error("Form not found!");
        return;
      }


      form.addEventListener("submit", async (e) => {
        e.preventDefault();
        const selectedCategories = Array.from( document.getElementsByName("category")).filter((box) => box.checked).map( (box)=> box.value )

        if(selectedCategories.length === 0){
          selectedCategories.push("Others");
        }

        const formData = {
          title : document.getElementById("title").value.trim(),
          content : document.getElementById("content").value.trim(),
          categories : selectedCategories
          

        }
       
       this.CheckAuth();


       await fetch("/createPost", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
          },
          credentials: "include",
          body : JSON.stringify(formData)
        });

        window.location.reload();


        
      });
  }

}
