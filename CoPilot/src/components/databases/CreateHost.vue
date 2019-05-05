<template>
  <div>
    <div id="myModal" class="modal fade" role="dialog" v-show='isModalVisible'>
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Tạo mới Host
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="createHost" class="create-host" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">{{ $t('hosts.nameMsg') }}</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                  <input class="form-control" v-model="comment" name="comment" type="text">
                </div>
                <button type="submit" class="btn btn-primary" >Lưu</button>
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
        isModalVisible: true
      }
    },
    methods: {
      createHost() {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'post',
              url: 'http://localhost:8081/host',
              data: {
                name: this.name,
                comment: this.comment
              }
            })
            .then(response => {
              this.isModalVisible = false
            })
          }
        })
      }
      // closeModel() {
      //   this.isModalVisible = false
      // }
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
