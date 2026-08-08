<template>

    <html lang="en">
    <head>
        <link rel="shortcut icon" href="">

    </head>
    <body>
        <div class="login-container">
            <div class="login-card">
                <div class="login-header">
                    <div class="neu-icon">
                        <div class="icon-inner">
                            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                                <circle cx="12" cy="7" r="4"/>
                            </svg>
                        </div>
                    </div>
                    <h2>Register </h2>
                </div>
                
                <form class="login-form" id="loginForm" @submit.prevent="Register" novalidate>
                    <div class="form-group">
                        <div class="input-group neu-input">
                            <input type="email" id="user_email" v-model="user_email" required placeholder=" ">
                            <label for="email">Email address</label>
                            <div class="input-icon">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
                                    <polyline points="22,6 12,13 2,6"/>
                                </svg>
                            </div>
                        </div>
                        <span class="error-message" id="emailError"></span>
                    </div>

                    <div class="form-group">
                        <div class="input-group neu-input password-group">
                            <input type="password" id="user_password" name="user_password" v-model="user_password" required  placeholder=" ">
                            <label for="user_password">Account password</label>
                            <div class="input-icon">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                                    <path d="M7 11V7a5 5 0 0110 0v4"/>
                                </svg>
                            </div>
                        </div>
                        <span class="error-message" id="passwordError"></span>
                    </div>

                    <div class="form-group">
                        <div class="input-group neu-input password-group">
                            <input type="password" id="daily_password" name="daily_password"v-model ="daily_password" required  placeholder=" ">
                            <label for="daily_password">Daily password</label>
                            <div class="input-icon">
                                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                                    <path d="M7 11V7a5 5 0 0110 0v4"/>
                                </svg>
                            </div>
                        </div>
                        <span class="error-message" id="passwordError"></span>
                    </div>

                    <button type="submit" class="neu-button login-btn" >
                        <span class="btn-text">Register</span>
                        <div class="btn-loader">
                            <div class="neu-spinner"></div>
                        </div>
                    </button>
                </form>


                <div class="signup-link">
                    <p>Already have an account? <a href="/login">Login</a></p>
                </div>

                <div class="success-message" id="successMessage">
                    <div class="success-icon neu-icon">
                        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3">
                            <polyline points="20 6 9 17 4 12"/>
                        </svg>
                    </div>
                    <h3>Success!</h3>
                    <p>Redirecting to your dashboard...</p>
                </div>
            </div>
        </div>

    </body>
    </html>
</template>

<script setup>
import { useRouter } from "vue-router";
import { ref} from "vue";
import "../styles/Auth.css";

const router = useRouter();

const user_email = ref("");
const user_password = ref("");
const daily_password = ref("");


async function Register() {
    const user = {
        UserEmail: user_email.value,
        UserPassword: user_password.value,
        DailyPassword: daily_password.value
    };
    
    try {
        const response = await fetch(`/auth/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(user)
        });
        if (response.ok) {
        alert("Registration successfull!");
        router.push('/login');
        } else {
        alert("Failed to register.");
        }
    } catch (error) {
        console.error(error);
        alert("Cannot connect to the server.");
    }

}
</script>