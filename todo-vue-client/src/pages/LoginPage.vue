<template>
    <div>
        <div class="login-page">
            <div class="form">
                <!-- @submit.prevent is onClick event listener. Then .prevent prevents page reload. -->
                <form  @submit.prevent="submitForm">
                    <!-- v-model works as a two way information path. userEmail data goes in and we can return a empty string back into it. Reference data & methods -->
                    <input type="text" placeholder="email address" v-model="userEmail"/>
                    <input type="password" placeholder="password" v-model="userPassword"/>
                    <button>login</button>
                    <!-- TODO LINK THIS TO SIGNUP VIA VUE WAY -->
                    <p class="message">Not registered? <a href="#">Create an account</a></p>
                </form>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    data() {
        return {
            userEmail: '',
            userPassword: ''
        }
    },
    methods: {
        submitForm() {
            console.log("Email: " + this.userEmail)
            this.userEmail = ""
            console.log("Password: " + this.userPassword)
            this.userPassword = ""
        }
    }
}
</script>

<style scoped>

.login-page {
  width: 360px;
  padding: 8% 0 0;
  margin: auto;
}
.form {
  position: relative;
  z-index: 1;
  background: #FFFFFF;
  max-width: 360px;
  margin: 0 auto 100px;
  padding: 45px;
  text-align: center;
  box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
}
.form input {
  font-family: "monospace", sans-serif;
  outline: 0;
  background: #f2f2f2;
  width: 100%;
  border: 0;
  margin: 0 0 15px;
  padding: 15px;
  box-sizing: border-box;
  font-size: 14px;
}
.form button {
  font-family: "monospace", sans-serif;
  text-transform: uppercase;
  outline: 0;
  background: #4CAF50;
  width: 100%;
  border: 0;
  padding: 15px;
  color: #FFFFFF;
  font-size: 14px;
  cursor: pointer;
}
.form button:hover,.form button:active,.form button:focus {
  background: #43A047;
}
.form .message {
  margin: 15px 0 0;
  color: #b3b3b3;
  font-size: 12px;
}
.form .message a {
  color: blue;
  text-decoration: none;
}







</style>
