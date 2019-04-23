<template>
  <div>
    <div id="updateTarget" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Update Target
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="updateTarget" class="update-target" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">Name</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" placeholder="Enter ..." type="text">
                  <!-- <div>{{data.name}}</div> -->
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">Comment</label>
                  <input v-model='comment' class="form-control" name='comment'>
                </div>
                <div class="form-group">
                  <label class="control-label" for='password'>Host</label>
                  <input v-model='hosts' class="form-control" name='host'>
                </div>
                <div class="form-group">
                  <label class="control-label" for="port-list">Port List</label>
                  <select class="form-control" v-model="portlist">
                    <option v-for="p in port" :key="p.id" v-bind:value="p.id">{{p.name}}</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='alivetest'>Alive Test</label>
                  <select v-model="alivetest">
                    <option value = 0  selected>Scan Config Default</option>
                    <option value= 1>TCP-ACK Service Ping</option>
                    <option value= 2>TCP-SYN Service Ping</option>
                    <option value="3">ARP Ping</option>
                    <option value="4">ICMP & TCP-ACK Service Ping</option>
                    <option value="5"> ICMP & ARP Ping</option>
                    <option value="6">TCP-ACK Service & ARP Ping</option>
                    <option value="7">ICMP, TCP-ACK Service & ARP Ping</option>
                    <option value="8">Consider Alive</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='rlonly'>Reverse Lookup Only</label>
                  <input type="radio" name="reverse-only" v-model="rlonly" value=1>Yes<br>
                  <input type="radio" name="reverse-only" v-model="rlonly" value=0 checked>No<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='rlunify'>Reverse Lookup Unify</label>
                  <input type="radio" name="reverse-unify" v-model="rlunify" value=1>Yes<br>
                  <input type="radio" name="reverse-unify" v-model="rlunify" value=0 checked>No<br>
                </div>
                <button type="submit" class="btn btn-primary">Save</button>
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
    // props: {
    // //   passedObject
    //   targetData: {
    //     type: [Array, Object]
    //   }
    // },
    // },
    props: ['targetData'],
    data() {
      return {
        port: [],
        name: '',
        comment: '',
        hosts: '',
        portlist: '',
        alivetest: '',
        rlonly: '',
        rlunify: ''
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
      updateTarget(id) {
        axios({
          method: 'put',
          url: 'http://localhost:8081/target/' + this.targetData.id,
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
          this.$router.push('/targets')
        })
      }
    }
  }
</script>
<style>
#updateModal label{
    min-width: 100px;
}
#updateModal .form-group{
    display: flex;
    align-items: center;
}
.modal-open .modal {
    overflow-x: hidden;
    overflow-y: hidden;
}
</style>
