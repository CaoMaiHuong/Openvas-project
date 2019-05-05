<template>
  <div>
    <div id="updateHost" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Sửa đổi Host
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="updateHost(hostData.id)" class="update-host" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">{{ $t('hosts.nameMsg') }}</label>
                  <input class="form-control" v-model="hostData.name" name="name" v-validate="'required'" type="text" disabled>
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                  <input class="form-control" v-model="hostData.comment" name="comment" type="text">
                </div>
                <button type="submit" class="btn btn-primary">Lưu</button>
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
    props: ['hostData'],
    data() {
      return {
        // name: '',
        // comment: ''
      }
    },
    methods: {
      updateHost(id) {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'put',
              url: 'http://localhost:8081/host/' + id,
              data: {
                name: this.hostData.name,
                comment: this.hostData.comment
              }
            })
            .then(response => {
              this.$router.push('/hosts')
            })
          }
        })
      }
    }
  }
</script>
<style>
#updateHost .form-group{
    display: flex;
    align-items: center;
    margin-bottom: 15px;
}
#updateHost .form-control{
  width: 100%;
}

#updateHost label{
    min-width: 125px;
}
#updateHost button.btn-primary{
  margin-top: 20px;
}
</style>
