import { BasePage } from './BasePage.js';
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
                        <input id="usernametxt" type="text" placeholder="Enter your username">

                        <strong><label>Age</label></strong>
                        <input id="agetxt" type="number" placeholder="Enter your age">

                        <label for="gender"><strong>Gender</strong></label>
                        <select id="gender" name="gender" required>
                            <option value="" disabled selected>Select your gender</option>
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                        </select>

                        <label for="first-name"><strong>First Name</strong></label>
                        <input id="first-name" type="text" name="first_name" placeholder="Enter your first name" required>

                        <label for="last-name"><strong>Last Name</strong></label>
                        <input id="last-name" type="text" name="last_name" placeholder="Enter your last name" required>

                        <label for="email"><strong>Email</strong></label>
                        <input id="email" type="email" name="email" placeholder="Enter your email" required>

                        <label for="password"><strong>Password</strong></label>
                        <input id="password" type="password" name="password" placeholder="Enter your password" required>

                        <button type="submit">Sign Up</button>
                        <span class="sign-info">Already have an account? <a href="/login" onclick="navigateTo('/login');return false;">Login here</a></span>
                    </div>
                </div>
            </form>
        </div>
        `;
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
                username: document.getElementById('usernametxt').value,
                age: document.getElementById('agetxt').value,
                gender: document.getElementById('gender').value,
                firstName: document.getElementById('first-name').value,
                lastName: document.getElementById('last-name').value,
                email: document.getElementById('email').value,
                password: document.getElementById('password').value,
            };

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
                alert('Response from server: ' + JSON.stringify(data));
            } catch (error) {
                console.error("Form submission failed:", error);
                alert("Form submission failed. Please try again.");
            }
        });
    }
}