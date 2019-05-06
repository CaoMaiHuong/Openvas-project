<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box">
          <div class="box-header">
            <!-- <router-link to="/createtarget"><i class="fa fa-user-plus" style="margin-right: 3px"></i> CREATE</router-link> -->
            <button data-toggle="modal" class="btn btn-primary" data-target="#addTask" @click="showModal" id="btn-modal"><i class="fa fa-plus"></i>&nbsp;&nbsp;{{ $t('action.createMsg') }}</button>
            <addTask v-show="isModalVisible"
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
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('tasks.nameMsg') }}</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('tasks.status') }}</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('tasks.report') }}</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('tasks.lastReport') }}</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('tasks.severity') }}</th>
                        <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('action.nameMsg') }}</th>
                        <!-- <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0"></th> -->
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="(task, index) in tasks" :key="task.id">
                        <!-- <td class="sorting_1">{{user.ID}}</td> -->
                        <td>{{task.name}}</td>
                        <td>{{task.status}}</td>
                        <td><router-link :to="{ name: 'Danh sách báo cáo', params: {id: task.uuid}}">{{task.rpnumber.String}}</router-link></td>
                        <td>{{task.last_report}}</td>
                        <td>{{task.severity.String}}</td>
                        <td>
                          <updatetask v-show="isModalVisible" @close="closeModal" :taskData="modalData" />                    
                          <a data-toggle="modal" data-target="#updateTask" @click="showUpdateModal(task,index)" style="margin-right: 20px"><i class="fa fa-pencil" style="margin-right: 5px"></i>{{ $t('action.editMsg') }}</a>
                          <a @click="deleteTask(task.id, index)" class="action-delete"> <i class="fa fa-trash" style="margin-right: 5px"></i>{{ $t('action.deleteMsg') }}</a>
                        </td>
                      </tr>
                    </tbody>
                    
                    <tfoot>
                      
                    </tfoot>
                  </table>
                  <div class="pagination">
                    <button class="btn btn-primary" v-on:click="fetchPaginate(1)" :disabled="pagination.page == 1"><i class="fa fa-angle-double-left"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.prev_page)" :disabled="pagination.page == 1"><i class="fa fa-angle-left"></i></button>
                    <span>Trang {{ pagination.page }} / {{ pagination.total_page }} </span>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.next_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-right"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.total_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-double-right"></i></button>
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
import addTask from './Create.vue'
import updatetask from './Update.vue'
// Require needed datatables modules
require('datatables.net')
require('datatables.net-bs')

export default {
  name: 'Tables',
  components: {
    addTask,
    updatetask
  },
  data() {
    return {
      tasks: [],
      pagination: [],
      page: 1,
      isModalVisible: false,
      modalData: null
    }
  },
  mounted() {
    this.getTask(this.page)
  },
  methods: {
    getTask(page) {
      axios.get('http://localhost:8081/tasks/page/' + page)
      .then(response => {
        let $this = this
        this.tasks = response.data.records
        $this.makePagination(response.data)
      })
    },
    showModal() {
      this.isModalVisible = true
    },
    showUpdateModal(item, index) {
      this.isModalVisible = true
      this.modalData = item
    },
    closeModal() {
      this.isModalVisible = false
    },
    deleteTask(id, index) {
      if (confirm('Bạn có chắc chắn muốn xóa?')) {
        axios.delete('http://localhost:8081/task/' + id)
        .then(response => {
          this.tasks.splice(index, 1)
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
      this.getTask(page)
    },
    openModal () {

    }
  }
}
</script>

<style>

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
#myModal .form-group{
    display: flex;
    align-items: center;
}
#myModal label{
    min-width: 125px;
}
</style>
