import { BasePage } from './BasePage.js';

export class login extends BasePage{
    constructor(){
        super(login.getHTml());
        this.setupFormSubmission();

    }
    static getHTml(){
        return `  
    <div class="login_main">
        <div class="login_container">
            <span class="login_title">Login to Community Forum</span>
            <form method="POST" post="/login" id="loginForm">
                <label for="UEnametxt">Username/Email</label>
                <input type="text" id="UEnametxt" name="UEnametxt" class="UEnametxt" placeholder="Enter your username or email">
                <label for="passwordtxt">Password</label>
                <input type="password" id="passwordtxt" name="passwordtxt" class="passwordtxt" placeholder="Enter your password">
                 <div id="res"> </div>
                <button type="submit">Login Up</button>
                <div class="login_info">Don't have an account? <a href="/sign" onclick="navigateTo('/sign');return false;">Sign up here</a></div>
            </form>
        </div>
    </div>
    `;
    }
    setupFormSubmission(){
        const form = document.getElementById('loginForm');
        if (!form) {
            console.error("Form not found!");
            return;
        }
        form.addEventListener('submit' , async (e)=>{
            e.preventDefault();

        const FormData ={
            uename :document.getElementById('UEnametxt').value,
            password : document.getElementById('passwordtxt').value
        }
        try{
            const response = await fetch("/login", {
              method: "POST",
              credentials: "include",
              headers: {
                "Content-Type": "application/json",
                Accept: "application/json",
              },
              body: JSON.stringify(FormData),
            });
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            if (!data.Success){
                document.getElementById("res").innerText = "ERROR:  "+ data.message;
            }else{
                   setTimeout(() => {
                                        navigateTo('/');
                                    }, 2000);
            }
            alert('Response from server: ' + JSON.stringify(data));

        }catch(error){
            console.error("Form submission failed:", error);
            alert("Form submission failed. Please try again.");
        }
        })
    }
}