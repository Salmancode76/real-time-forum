import { BasePage } from './BasePage.js';
import { navigateTo } from '../index.js';
export class signup  extends BasePage{
    constructor() {
       
        super(signup.getHtml());
        this.setupFormSubmission();
    }

   static  getHtml() {
        return `
        <div class="signup_main_form">
            <form method="post" action="/signup" id="signupForm">
                <div class="signup_main">
                    <div class="signup_container">
                        <span class="signup_title">Sign Up for Community Forum</span>
                        <strong><label for="usernametxt">Username</label></strong>
                        <input id="usernametxt" type="text" placeholder="Enter your username" required maxlength="50">

                        <strong><label>Age</label></strong>
                        <input id="agetxt" type="number" placeholder="Enter your age" required>

                        <label for="gender"><strong>Gender</strong></label>
                        <select id="gender" name="gender" required>
                            <option value="" disabled selected>Select your gender</option>
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                        </select>

                        <label for="first-name"><strong>First Name</strong></label>
                        <input id="first-name" type="text" name="first_name" placeholder="Enter your first name" required maxlength="50">

                        <label for="last-name"><strong>Last Name</strong></label>
                        <input id="last-name" type="text" name="last_name" placeholder="Enter your last name" required maxlength="50">

                        <label for="email"><strong>Email</strong></label>
                        <input id="email" type="email" name="email" placeholder="Enter your email" required required maxlength="254">

                        <label for="password"><strong>Password</strong></label>
                        <input id="password" type="password" name="password" placeholder="Enter your password" required maxlength="200">
                        <div id="res"> </div>

                        <button type="submit">Sign Up</button>
                        <span class="sign-info">Already have an account? <a href="/login" onclick="navigateTo('/login');return false;">Login here</a></span>
                    </div>
                </div>
            </form>
        </div>
        `;
    }
    validationSignup(formData) {
        let error_message = "";
    
        if (formData.username.length === 0) {
            error_message += "Username cannot be empty or just spaces.\n";
        }
    
        if (formData.firstName.length === 0) {
            error_message += "First Name cannot be empty or just spaces.\n";
        }
    
        if (formData.lastName.length === 0) {
            error_message += "Last Name cannot be empty or just spaces.\n";
        }
    
        if (formData.age < 0 || formData.age > 150) {
            error_message += "Please enter a valid age. Age should be between 0 and 150.\n";
        }
    
        const nameRegex = /^[A-Za-z\s]+$/;
        if (!nameRegex.test(formData.firstName)) {
            error_message += "First name cannot contain numbers or special characters.\n";
        }
    
        if (!nameRegex.test(formData.lastName)) {
            error_message += "Last name cannot contain numbers or special characters.\n";
        }
    
        const password = formData.password;
        const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
    
        if (!passwordRegex.test(password)) {
            error_message += "Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character (e.g., @, $, !, %, *, ?, &).\n";
        }
    
        if (error_message.length > 0) {
            return error_message;
        }
    
        return true;
    }

    setupFormSubmission() {
        const form = document.getElementById('signupForm');
        if (!form) {
            console.error("Form not found!");
            return;
        }

        form.addEventListener('submit', async (e) => {
            e.preventDefault();

            const formData = {
                username: document.getElementById('usernametxt').value.trim(),
                age: document.getElementById('agetxt').value.trim(),
                gender: document.getElementById('gender').value.trim(),
                firstName: document.getElementById('first-name').value.trim(),
                lastName: document.getElementById('last-name').value.trim(),
                email: document.getElementById('email').value.trim(),
                password: document.getElementById('password').value.trim(),
            };
            /* COMMENTED FOR TESTING
            const validationResult = this.validationSignup(formData);

            if (validationResult !== true) {
               // alert(validationResult);
                document.getElementById("res").innerText = "Error:  "+ validationResult;
                return;
            }
                */

            try {
                const response = await fetch('/sign', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json',
                    },
                    body: JSON.stringify(formData),
                });

                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                
                const data = await response.json();
                if (!data.Success){
                    document.getElementById("res").innerText = "ERROR:  "+ data.message;

                }else{
                    document.getElementById("res").innerText = "SUCCESS:  "+ data.message;
                    document.getElementById("res").style.color = "#2ECC71";
                    setTimeout(() => {
                        navigateTo('/login');
                    }, 2000);

                }
               // alert('Response from server: ' + (data.message || 'Unknown response'));   
             } catch (error) {
                console.error("Form submission failed:", error);
                alert("Form submission failed. Please try again.");
            }
        });
    }


}