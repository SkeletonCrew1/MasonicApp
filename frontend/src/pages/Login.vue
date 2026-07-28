<template>

    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Neumorphism Login Form</title>
        <meta name="description" content="Soft-UI login form with realistic dual shadows and ambient mouse-tracked lighting.">
        <meta name="author" content="Aigars Silkalns / Colorlib">
        <link rel="canonical" href="https://preview.colorlib.com/theme/login-forms/forms/neumorphism/">
        <link rel="icon" type="image/svg+xml" href="data:image/svg+xml,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'><circle cx='50' cy='50' r='48' fill='%236366f1'/><text x='50' y='68' font-size='60' text-anchor='middle' fill='white' font-family='system-ui,sans-serif'>L</text></svg>">
        <meta property="og:type" content="website">
        <meta property="og:title" content="Neumorphism Login Form">
        <meta property="og:description" content="Soft-UI login form with realistic dual shadows and ambient mouse-tracked lighting.">
        <meta property="og:url" content="https://preview.colorlib.com/theme/login-forms/forms/neumorphism/">
        <meta property="og:image" content="https://preview.colorlib.com/theme/login-forms/assets/screenshots/neumorphism.png">
        <meta name="twitter:card" content="summary_large_image">
        <meta name="twitter:title" content="Neumorphism Login Form">
        <meta name="twitter:description" content="Soft-UI login form with realistic dual shadows and ambient mouse-tracked lighting.">
        <meta name="twitter:image" content="https://preview.colorlib.com/theme/login-forms/assets/screenshots/neumorphism.png">
        <link rel="stylesheet" href="style.css">
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
                    <h2>Login </h2>
                </div>
                
                <form class="login-form" id="loginForm" novalidate>
                    <div class="form-group">
                        <div class="input-group neu-input">
                            <input type="email" id="email" name="email" required autocomplete="email" placeholder=" ">
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
                            <input type="password" id="daily_password" name="daily_password" required  placeholder=" ">
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

                    <div class="form-group">
                        <div class="input-group neu-input password-group">
                            <input type="password" id="daily_password" name="daily_password" required  placeholder=" ">
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

                    <button type="submit" class="neu-button login-btn">
                        <span class="btn-text">Sign In</span>
                        <div class="btn-loader">
                            <div class="neu-spinner"></div>
                        </div>
                    </button>
                </form>


                <div class="signup-link">
                    <p>Don't have an account? <a href="#">Register</a></p>
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
import "../styles/Auth.css";
</script>