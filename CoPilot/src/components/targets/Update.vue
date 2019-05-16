<template>
  <div>
    <div id="updateTarget" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Sửa đổi Target
            </span>
            <button type="button" class="close" @click="closeForm()" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
            <form v-on:submit.prevent="updateTarget(targetData.id)" class="update-target" style="padding: 0px">
              <div class="form-group">
                <label class="control-label" for="name">{{ $t('targets.nameMsg') }}</label>
                <input class="form-control" v-model="targetData.name" name="name" v-validate="'required'" type="text">
                <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
              </div>
              <div class="form-group">
                <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                <input v-model='targetData.comment' class="form-control" name='comment'>
              </div>
              <div class="form-group">
                <label class="control-label" for='host'>{{ $t('targets.hostMsg') }}</label>
                <input v-model='targetData.hosts' class="form-control" name='host' v-validate="'required'" type="text">
                <span v-if="errors.has('host')">{{ errors.first('host') }}</span>
              </div>
              <div class="form-group">
                <label class="control-label" for="port-list">{{ $t('targets.portMsg') }}</label>
                <select class="form-control" v-model="targetData.portlist_id">
                  <option v-for="p in port" :key="p.id" v-bind:value="p.id">{{p.name}}</option>
                </select>
              </div>
              <div class="form-group">
                <label class="control-label" for='alivetest'>{{ $t('targets.aliveTestMsg') }}</label>
                <select class="form-control" v-model="targetData.alivetest">
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
                <input type="radio" name="reverse-only" v-model="targetData.rlonly" value=1>Có<br>
                <input type="radio" name="reverse-only" v-model="targetData.rlonly" value=0 style="margin-left: 30px;" checked>Không<br>
              </div>
              <div class="form-group">
                <label class="control-label" for='rlunify' style="margin-right: 30px;">{{ $t('targets.rlUnifyMsg') }}</label>
                <input type="radio" name="reverse-unify" v-model="targetData.rlunify" value=1>Có<br>
                <input type="radio" name="reverse-unify" v-model="targetData.rlunify" value=0  style="margin-left: 30px;" checked>Không<br>
              </div>
            </form>
          </div>
          <!-- <div>{{targetData.id}}</div> -->
          <div class="modal-footer">
            <button class="btn btn-success" v-if="messageUpdate">{{messageUpdate}}</button>
             <button type="submit" @click="updateTarget(targetData.id)" class="btn btn-primary button-save">Lưu</button>
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
        messageUpdate: ''
        // name: '',
        // comment: '',
        // hosts: '',
        // portlist: '',
        // alivetest: '',
        // rlonly: '',
        // rlunify: ''
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
          url: 'http://localhost:8081/target/' + id,
          data: {
            name: this.targetData.name,
            comment: this.targetData.comment,
            hosts: this.targetData.hosts,
            portlist_id: this.targetData.portlist_id,
            alivetest: this.targetData.alivetest,
            rlonly: this.targetData.rlonly,
            rlunify: this.targetData.rlunify
          }
        })
        .then(response => {
          if (response.data === 'Cập nhật thông tin thành công!') {
            this.messageUpdate = response.data
            location.reload()
          }
        })
      },
      closeForm() {
        this.messageUpdate = ''
      }
    }
  }
</script>
<style>
#updateTarget .form-control{
  display: unset;
  width: 100%;
}
#updateTarget .form-group{
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  margin-bottom: 15px;
}
#updateTarget label {
    min-width: 125px;
}
/* .modal-open .modal {
    overflow-x: hidden; 
    overflow-y: hidden;
} */
</style>
