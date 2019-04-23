<template>
  <div>
    <div id="myModal" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              New User
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="createUser" class="create-user" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">Name</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                  <span v-if="message">{{message}}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">Comment</label>
                  <input class="form-control" v-model="comment" name="comment" type="text">
                </div>
                <div class="form-group">
                  <label class="control-label" for='password'>Password</label>
                  <input v-validate="'required|min:6'" v-model='password' name="password" type="password" class="form-control" ref="password">
                  <span v-if="errors.has('password')">{{ errors.first('password') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="confirm_password">Confirm Password</label>
                  <input v-model='confirm_password' name="confirm_password" v-validate="'required|confirmed:password'" type="password" class="form-control" data-vv-as="password">
                  <span  v-if="errors.has('confirm_password')">{{ errors.first('confirm_password') }}</span>
                </div>
                <div class="form-group">
                  <label for='role'>Roles</label>
                  <select class="form-control" v-model="role">
                    <option value="1">Admin</option>
                    <option value="5">User</option>
                    <option value="4">Monitor</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for="host_access">Host Access</label>
                  <input type="radio" name="host_access" v-model="host_access" value="0" checked>Allow all and deny<br>
                  <input type="radio" name="host_access" v-model="host_access" value="1" >Deny all and allow<br>
                  <input v-model='hosts' type="text" name="hosts" class="form-control">
                </div>
                <div class="form-group">
                  <label class="control-label" for="iface_allow">Interface Access</label>
                  <input type="radio" name="iface_allow" v-model="iface_allow" value="0" checked>Allow all and deny<br>
                  <input type="radio" name="iface_allow" v-model="iface_allow" value="1" >Deny all and allow<br>
                  <input v-model='ifaces' type="text" name="ifaces" class="form-control">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
              </form>
          </div>
          <div class="modal-footer">
              
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios'
  export default {
    name: 'Tables',
    data() {
      return {
        name: '',
        comment: '',
        password: '',
        comfirm_password: '',
        role: '',
        host_access: '',
        hosts: '',
        iface_allow: '',
        ifaces: '',
        message: ''
      }
    },
    methods: {
      createUser() {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'post',
              url: 'http://localhost:8081/user',
              data: {
                name: this.name,
                comment: this.comment,
                password: this.password,
                role: this.role,
                host_access: this.host_access,
                hosts: this.hosts,
                iface_allow: this.iface_allow,
                ifaces: this.ifaces
              }
            })
            .then(response => {
              this.message = response.data
              if (this.message !== 'User already exists') {
                this.$router.push('/users')
              }
            })
          }
        })
      }
    }
  }
</script>
<style>
#myModal .form-group{
    display: flex;
    align-items: center;
}
#myModal label{
    min-width: 100px;
}
</style>
