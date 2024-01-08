<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <form @submit.prevent="submitForm" class="form-signin">
          <h2 class="form-signin-heading mb-3">Welcome Back</h2>
          <input type="email" id="inputEmail" class="form-control mb-3" placeholder="Email address" required autofocus v-model="email">
          <input type="password" id="inputPassword" class="form-control mb-3" placeholder="Password" required v-model="password">
          <button class="btn btn-md btn-primary btn-block w-100 mb-3" type="submit">Sign in</button>
          <a class="sign-up-link form-control mb-3" href="/auth/sign_up">Sign up</a>
          <button type="button" class="btn btn-md btn-danger btn-block w-100" @click="signInWithGoogle">Sign in with Google</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LoginForm',
  data() {
    return {
      username: '',
      password: '',
      passwordFieldType: 'password',
      clientId: '337616075879-edmpl5ejbik5ns8sd10es6l966bs4nus.apps.googleusercontent.com'
    }
  },
  mounted() {
    window.handleCredentialResponse = this.handleCredentialResponse
  },
  methods: {
    submitForm() {
      axios.post('http://your-server-url.com/auth', {
        username: this.username,
        password: this.password
      })
          .then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
    },
    handleCredentialResponse(response) {
      console.log(response)
      axios.post('http://your-server-url.com/auth', {
        credential: response.credential
      })
          .then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
    },
    togglePasswordVisibility() {
      this.passwordFieldType = this.passwordFieldType === 'password' ? 'text' : 'password';
    }
  }
}
</script>

<style scoped>
.container {
  padding-top: 40px;
}

.form-signin {
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}

.form-signin .form-control {
  position: relative;
  height: auto;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
  padding: 10px;
  font-size: 16px;
}

.form-signin .form-control:focus {
  z-index: 2;
}

.form-signin input[type="email"],
.form-signin input[type="password"] {
  margin-bottom: 20px;
}

.sign-up-link {
  color: #007bff;
  text-decoration: none;
}

.sign-up-link:hover {
  color: #0056b3;
  text-decoration: underline;
}
</style>