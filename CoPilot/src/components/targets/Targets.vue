<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box">
          <div class="box-header">
            <!-- <router-link to="/createtarget"><i class="fa fa-user-plus" style="margin-right: 3px"></i> CREATE</router-link> -->
            <button data-toggle="modal" data-target="#myModal" @click="showModal">Add Target</button>
            <modal
              v-show="isModalVisible"
              @close="closeModal"
            />
          </div>
          <!-- /.box-header -->
          <div class="box-body">
            <div class="dataTables_wrapper form-inline dt-bootstrap" id="example1_wrapper">
              <div class="row">
                <div class="col-sm-6">
                  <div id="example1_length" class="dataTables_length">

                  </div>
                </div>
              </div>

              <div class="row">
                <div class="col-sm-12 table-responsive">
                  <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                    <thead>
                      <tr role="row">
                        <!-- <th style="width: 30px" aria-sort="ascending" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Id</th> -->
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Name</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Hosts</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Port List</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >Action</th>
                        <!-- <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0"></th> -->
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="target in targets" :key="target.id">
                        <!-- <td class="sorting_1">{{user.ID}}</td> -->
                        <td>
                          <router-link :to="{ name: 'Target Detail', params: {id: target.uuid }}">{{target.name}}</router-link>
                        </td>
                        <td>{{target.hosts}}</td>
                        <td>{{target.portlist}}</td>
                        <td class="action-edit">  
                          <updatemodal v-show="isModalVisible" @close="closeModal" :targetData="modalData" />                    
                          <a data-toggle="modal" data-target="#updateTarget" @click="showUpdateModal(target)" style="margin-right: 20px"><i class="fa fa-pencil" style="margin-right: 5px"></i>Edit</a>
                          <a @click="deleteTarget(target.id)"> <i class="fa fa-trash" style="margin-right: 5px"></i>Delete</a>
                        </td>
                      </tr>
                    </tbody>
                    
                    <tfoot>
                      
                    </tfoot>
                  </table>
                  <div class="pagination">
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.prev_page)" :disabled="pagination.prev_page == pagination.page">Previous</button>
                    <span>Page {{ pagination.page }} of {{ pagination.total_page }} </span>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.next_page)" :disabled="pagination.next_page == pagination.page">Next</button>
                  </div>
                </div>
              </div>
            </div>
            <!-- /.box-body -->
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
// import $ from 'jquery'
import axios from 'axios'
import modal from './Create.vue'
import updatemodal from './Update.vue'
// Require needed datatables modules
require('datatables.net')
require('datatables.net-bs')

export default {
  name: 'Tables',
  components: {
    modal,
    updatemodal
  },
  data() {
    return {
      targets: [],
      pagination: [],
      page: 1,
      isModalVisible: false,
      modalData: null
    }
  },
  mounted() {
    this.getTarget(this.page)
  },
  methods: {
    getTarget(page) {
      axios.get('http://localhost:8081/targets/page/' + page)
      .then(response => {
        let $this = this
        this.targets = response.data.records
        $this.makePagination(response.data)
      })
    },
    showModal() {
      this.isModalVisible = true
    },
    showUpdateModal(item) {
      this.isModalVisible = true
      this.modalData = item
    },
    closeModal() {
      this.selectedItem = undefined
      this.isModalVisible = false
    },
    createUser() {
      this.$router.push('/targets')
    },
    deleteTarget(id) {
      if (confirm('Do you really want to delete it?')) {
        axios({
          method: 'delete',
          url: 'http://localhost:8081/target/' + id
        })
        .then(response => {
          location.reload()
        })
      }
    },
    makePagination(data) {
      let pagination = {
        page: data.page,
        total_page: data.total_page,
        prev_page: data.prev_page,
        next_page: data.next_page
      }
      this.pagination = pagination
    },
    fetchPaginate(page) {
      this.getTarget(page)
    },
    openModal () {

    }
  }
}
</script>

<style>
/* Using the bootstrap style, but overriding the font to not draw in
   the Glyphicons Halflings font as an additional requirement for sorting icons.

   An alternative to the solution active below is to use the jquery style
   which uses images, but the color on the images does not match adminlte.

@import url('/static/js/plugins/datatables/jquery.dataTables.min.css');
*/

@import url('/static/js/plugins/datatables/dataTables.bootstrap.css');

table.dataTable thead .sorting:after,
table.dataTable thead .sorting_asc:after,
table.dataTable thead .sorting_desc:after {
  font-family: 'FontAwesome';
}

table.dataTable thead .sorting:after {
  content: '\f0dc';
}

table.dataTable thead .sorting_asc:after {
  content: '\f0dd';
}

table.dataTable thead .sorting_desc:after {
  content: '\f0de'
}
.pagination{
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-top: 30px
}
.pagination span{
  margin: 0px 10px
}
</style>
