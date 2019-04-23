<template>
  <div id="signup">
    <div class="col-sm-12">
      <!-- login form -->
      <form @submit.prevent="signUp">
        <div class="signup-field">
          <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-user"></i></span>
            <input class="form-control" name="name" placeholder="Username" type="text" v-model="name" v-validate="'required'">
          </div>
          <span v-if="errors.has('name')" class="error">
            {{ errors.first('name') }}
          </span>
        </div>
        <div class="signup-field">
            <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-envelope"></i></span>
            <input class="form-control" name="email" placeholder="Email" type="text" v-model="email" v-validate="'required|email'">
            </div>
            <span v-if="errors.has('email')" class="error">
            {{ errors.first('email') }}
            </span>
            <div>
                <span style="color:red"> {{ message }}</span>
            </div>
        </div>
        <div class="signup-field">
            <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-lock"></i></span>
            <input class="form-control" name="password" placeholder="Password" type="password" v-model="password" v-validate="'required|min:6'" ref="password">
            </div>
            <span v-if="errors.has('password')" class="error">
            {{ errors.first('password') }}
            </span>
        </div>
        <div class="signup-field">
            <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-lock"></i></span>
            <input class="form-control" name="confirmpassword" placeholder="Confirm Password" type="password" v-model="confirmpassword" v-validate="'required|confirmed:password'">
            </div>
            <span  v-if="errors.has('confirmpassword')" class="error">{{ errors.first('confirmpassword') }}</span>
        </div>
        <div class="signup-field">
          <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-user"></i></span>
            <input class="form-control" name="dob" placeholder="Date of birth" type="date" v-model="dob">
          </div>
          <!-- <span v-if="errors.has('dob')" class="error">
            {{ errors.first('dob') }}
          </span> -->
        </div>
        <div class="signup-field">
          <div class="input-group">
            <span class="input-group-addon"><i class="fa fa-user"></i></span>
            <input class="form-control" name="phone" placeholder="Phone" type="number" v-model="phone" v-validate="'required'">
          </div>
          <span v-if="errors.has('phone')" class="error">
            {{ errors.first('phone') }}
          </span>
        </div>
        <br>
        <div class="text-center">
            <button type="submit" v-bind:class="'text-center btn btn-primary btn-lg'">Sign up</button>
            <div class="signup">Already have an account ?
                <router-link to="/login">Login</router-link>
            </div>
        </div>
      </form>

      <!-- errors -->
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'AddUser',
  data() {
    return {
      name: '',
      email: '',
      password: '',
      confirmpassword: '',
      dob: '',
      phone: '',
      message: ''
    }
  },
  methods: {
    signUp() {
      this.$validator.validateAll().then(res => {
        if (res) {
          axios({
            method: 'post',
            url: 'http://localhost:8081/user',
            data: {
              name: this.name,
              email: this.email,
              password: this.password,
              dob: this.dob,
              phone: this.phone
            }
          })
          .then(response => {
            this.message = response.data
            if (this.message !== 'Email existed') {
              this.$router.push('/login')
            }
          })
        }
      })
    }
  }
}
</script>

<style>
#signup {
  padding: 10em 10em 10em 10em;
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
#signup button{
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
#signup .signup-field{
    margin-bottom: 1.3em;
}
#signup .signup-field span.error{
    color: red;
}
#signup .input-group {
  padding-bottom: 0px;
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
