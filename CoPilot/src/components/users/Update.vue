<template>
  <div v-if="renderComponent">
    <div id="updateModal" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Tạo mới người dùng
            </span>
            <button type="button" @click="closeForm()" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="updateUser(userData.id)" class="create-user" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">{{ $t('users.nameMsg') }}</label>
                  <input class="form-control" v-model="userData.name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                  <input class="form-control" v-model="userData.comment" name="comment" type="text">
                </div>
                <!-- <div class="form-group">
                  <label class="control-label" for='password'>{{ $t('users.passwordMsg') }}</label>
                  <input v-validate="'required|min:6'" v-model='userData.password' name="password" type="password" class="form-control" ref="password">
                  <span v-if="errors.has('password')">{{ errors.first('password') }}</span>
                </div> -->
                <!-- <div class="form-group">
                  <label class="control-label" for="confirm_password">{{ $t('users.confirmPassword') }}</label>
                  <input v-model='confirm_password' name="confirm_password" v-validate="'required|confirmed:password'" type="password" class="form-control" data-vv-as="password">
                  <span  v-if="errors.has('confirm_password')">{{ errors.first('confirm_password') }}</span>
                </div> -->
                <div class="form-group">
                  <label for='role'>{{ $t('users.roleMsg') }}</label>
                  <select class="form-control" v-model="userData.role_id">
                    <option v-for="r in roles" :key="r.id" v-bind:value="r.id">{{r.name}}</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for="host_allow">{{ $t('users.hostAccessMsg') }}</label>
                  <div class="contentt">
                  <div class="optionn">
                  <input type="radio" name="host_allow" v-model="userData.host_allow_number" value="0" checked>{{ $t('users.allowanddeny') }}<br>
                  <input type="radio" name="host_allow" v-model="userData.host_allow_number" value="1" style="margin-left: 20px">{{ $t('users.denyandallow') }}<br>
                  </div>
                  <div class="inputcontent">
                  <input v-model='userData.hosts' type="text" name="hosts" class="form-control">
                  </div>
                  </div>
                </div>
                <div class="form-group">
                  <label class="control-label" for="iface_allow">{{ $t('users.interface') }}</label>
                  <div class="contentt">
                  <div class="optionn">
                  <input type="radio" name="iface_allow" v-model="userData.iface_allow_number" value="0" checked>{{ $t('users.allowanddeny') }}<br>
                  <input type="radio" name="iface_allow" v-model="userData.iface_allow_number" value="1" style="margin-left: 20px" >{{ $t('users.denyandallow') }}<br>
                  </div>
                  <div class="inputcontent">
                  <input v-model='userData.ifaces' type="text" name="ifaces" class="form-control">
                  </div>
                  </div>
                </div>
              </form>
          </div>
          <div class="modal-footer">
            <!-- <button class="btn btn-danger" v-if="message">{{message}}</button> -->
            <button class="btn btn-success" v-if="messageUpdate">{{messageUpdate}}</button>
            <button type="submit" @click="updateUser(userData.id)" class="btn btn-primary">Lưu</button>
            <!-- <span v-if="message">{{message}}</span>
            <span v-if="messageUpdate">{{messageUpdate}}</span> -->
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
    props: ['userData'],
    data() {
      return {
        // name: this.userData,
        // comment: '',
        // password: '',
        // // comfirm_password: '',
        // role: '',
        // host_allow: '',
        // hosts: [],
        // iface_allow: '',
        // ifaces: [],
        roles: [],
        // message: '',
        messageUpdate: '',
        renderComponent: true
      }
    },
    created() {
      axios({
        method: 'get',
        url: 'http://localhost:8081/roles'
      })
      .then(response => {
        this.roles = response.data
        // if (this.message !== 'User already exists') {
        //   this.$router.push('/users')
        // }
      })
    },
    methods: {
      updateUser(id) {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'put',
              url: 'http://localhost:8081/user/' + id,
              data: {
                name: this.userData.name,
                comment: this.userData.comment,
                // password: this.userData.password,
                role_id: this.userData.role_id,
                host_allow_number: this.userData.host_allow_number,
                hosts: this.userData.hosts,
                iface_allow_number: this.userData.iface_allow_number,
                ifaces: this.userData.ifaces
              }
            })
            .then(response => {
              if (response.data === 'Tên người dùng đã tồn tại!') {
                this.message = response.data
              }
              if (response.data === 'Cập nhật thông tin thành công!') {
                this.messageUpdate = response.data
                this.message = ''
                location.reload()
              }
              // if (this.message !== 'User already exists') {
              //   this.$router.push('/users')
              // }
            })
          }
        })
      },
      closeForm() {
        this.message = ''
        this.messageUpdate = ''
      }
    }
  }
</script>
<style>
#updateModal .form-group{
    display: flex;
    align-items: flex-start;
    flex-wrap: wrap;
    margin-bottom: 20px;
}

#updateModal label{
    min-width: 125px;
}

#updateModal .form-control {
  width: 100%;
}
.optionn {
  display: flex;
}

</style>
