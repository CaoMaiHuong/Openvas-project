<template>
  <div id="login">
    <div class="text-center col-sm-12">
      <!-- login form -->
      <form @submit.prevent="checkCreds">
        <div class="input-group">
          <span class="input-group-addon"><i class="fa fa-envelope"></i></span>
          <input class="form-control" name="name" placeholder="Tên đăng nhập" type="text" v-model="name">
        </div>

        <div class="input-group">
          <span class="input-group-addon"><i class="fa fa-lock"></i></span>
          <input class="form-control" name="password" placeholder="Mật khẩu" type="password" v-model="password">
        </div>
        <span style="color:red"> {{ message }}</span>
        <br>
        <button type="submit" v-bind:class="'btn btn-primary btn-lg'">Đăng nhập</button>
        <div class="signup">Bạn chưa có tài khoản?
          <router-link to="/signup">Đăng ký</router-link>
        </div>
      </form>

      <!-- errors -->
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Login',
  data(router) {
    return {
      name: '',
      password: '',
      message: ''
    }
  },
  methods: {
    checkCreds() {
      axios({
        method: 'post',
        url: 'http://localhost:8081/login',
        // withCredentials: true,
        data: {
          password: this.password,
          name: this.name
        }
      })
      .then(response => {
        if (response.data === 'Login successful') {
          let name = this.name
          window.localStorage.setItem('name', name)
          this.$router.push('/')
        } else {
          this.message = 'Tên đăng nhập hoặc mât khẩu không chính xác'
        }
      })
    }
  }
}
</script>

<style>
#login {
  padding: 13em 10em 10em 10em;
}

html,
body,
.container-table {
  height: 100%;
  background-color: #f1f1f5;
}
.container-table {
  display: table;
  color: white;
}
#login button{
  margin-top: 10px;
  padding-top: 7px;
  padding-bottom: 7px;
  margin-bottom: 10px;
}

.vertical-center-row {
  display: table-cell;
  vertical-align: middle;
}
.vertical-20p {
  padding-top: 20%;
}
.vertical-10p {
  padding-top: 10%;
}
.vertical-5p {
  padding-top: 5%;
}
.logo {
  width: 15em;
  padding: 3em;
}

.input-group {
  padding-bottom: 1.3em;
  height: 4em;
  width: 100%;
}

.input-group span.input-group-addon {
  width: 2em;
  height: 4em;
}

@media (max-width: 1241px) {
  .input-group input {
    height: 4em;
  }
}
@media (min-width: 1242px) {
  form {
    padding-left: 20em;
    padding-right: 20em;
  }

  .input-group input {
    height: 4em;
  }
}

.input-group-addon i {
  height: 15px;
  width: 15px;
}
</style>
