<template>
  <div>
    <div id="myModal" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Tạo mới target
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="createTarget" class="create-user" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">{{ $t('targets.nameMsg') }}</label>
                  <input class="form-control modal-input" v-model="name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')" style="display: block">{{ errors.first('name') }}</span>
                </div>
                
                <div class="form-group">
                  <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                  <input v-model='comment' class="form-control" name='comment'>
                </div>
                <div class="form-group">
                  <label class="control-label" for='password'>{{ $t('targets.hostMsg') }}</label>
                  <input v-model='hosts' class="form-control" name='host'>
                </div>
                <div class="form-group">
                  <label class="control-label" for="port-list">{{ $t('targets.portMsg') }}</label>
                  <select class="form-control" v-model="portlist">
                    <option v-for="p in port" :key="p.id" v-bind:value="p.id">{{p.name}}</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='alivetest'>{{ $t('targets.aliveTestMsg') }}</label>
                  <select class="form-control" v-model="alivetest">
                    <option value = "0">Scan Config Default</option>
                    <option value= "1">TCP-ACK Service Ping</option>
                    <option value= "2">TCP-SYN Service Ping</option>
                    <option value="3">ARP Ping</option>
                    <option value="4">ICMP & TCP-ACK Service Ping</option>
                    <option value="5"> ICMP & ARP Ping</option>
                    <option value="6">TCP-ACK Service & ARP Ping</option>
                    <option value="7">ICMP, TCP-ACK Service & ARP Ping</option>
                    <option value="8">Consider Alive</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='rlonly' style="margin-right: 30px;">{{ $t('targets.rlOnlyMsg') }}</label>
                  <input type="radio" name="reverse-only" v-model="rlonly" value=1 >Có<br>
                  <input type="radio" name="reverse-only" v-model="rlonly" value=0 checked style="margin-left: 30px;">Không<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='rlunify' style="margin-right: 30px;">{{ $t('targets.rlUnifyMsg') }}</label>
                  <input type="radio" name="reverse-unify" v-model="rlunify" value=1>Có<br>
                  <input type="radio" name="reverse-unify" v-model="rlunify" value=0 checked style="margin-left: 30px;">Không<br>
                </div>
                <button type="submit" class="btn btn-primary">Lưu</button>
              </form>
          </div>
          
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios'
  // import $ from 'jquery'
  export default {
    name: 'Tables',
    data() {
      return {
        port: [],
        name: '',
        comment: '',
        hosts: '',
        portlist: '1',
        alivetest: '0',
        rlonly: '0',
        rlunify: '0'
      }
    },
    mounted() {
      this.getPortList()
    },
    methods: {
      getPortList() {
        axios.get('http://localhost:8081/portlist')
        .then(response => {
          this.port = response.data
        })
      },
      createTarget() {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'post',
              url: 'http://localhost:8081/target',
              data: {
                name: this.name,
                comment: this.comment,
                hosts: this.hosts,
                portlist: this.portlist,
                alivetest: this.alivetest,
                rlonly: this.rlonly,
                rlunify: this.rlunify
              }
            })
            .then(response => {
              // // this.$router.push('/targets')
              // this.$router.go(0)
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
    min-width: 125px;
}
</style>
