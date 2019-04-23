<template>
  <div>
    <div id="updateHost" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Update Host
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="updateHost" class="update-host" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">Name</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">Comment</label>
                  <input class="form-control" v-model="comment" name="comment" type="text">
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
    props: ['hostData'],
    data() {
      return {
        name: '',
        comment: ''
      }
    },
    methods: {
      updateHost(id) {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'put',
              url: 'http://localhost:8081/host/' + this.hostData.id,
              data: {
                name: this.name,
                comment: this.comment
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
#myModal .form-group{
    display: flex;
    align-items: center;
}
#myModal label{
    min-width: 100px;
}
</style>
