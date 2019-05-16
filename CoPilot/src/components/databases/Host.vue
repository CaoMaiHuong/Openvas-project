<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box">
          <div class="box-header">
            <button data-toggle="modal" class="btn btn-primary" data-target="#myModal" @click="showModal"><i class="fa fa-plus" style="margin-right: 3px"></i> {{ $t('action.createMsg') }}</button>
            <modal v-show="isModalVisible"/>
          </div>
          <div class="box-body">
            <div class="dataTables_wrapper form-inline dt-bootstrap" id="example1_wrapper">
              <div class="row">
                <div class="col-sm-12 table-responsive">
                  <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                    <thead>
                      <tr role="row">
                        <th style="width: 10%" aria-sort="ascending" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('hosts.nameMsg') }}</th> 
                        <th style="width: 13%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('hosts.hostnameMsg') }}</th>
                        <th style="width: 7%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('hosts.ipMsg') }}</th>
                        <th style="width: 6%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('severityMsg') }}</th>
                        <th style="width: 15%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('modifyMsg') }}</th>
                        <th style="width: 12%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('action.nameMsg') }}</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="(host, index) in hosts" :key="host.id">
                        <td class="sorting_1">
                          <router-link :to="{ name: 'Chi tiết Host', params: {id: host.uuid }}">{{host.name}}</router-link>
                        </td>
                        <td>{{host.hostname.String}}</td>
                        <td>{{host.ipaddress.String}}</td>
                        <td>{{host.severity.String}}</td>
                        <td>{{host.modified}}</td>
                        <td class="action-edit">
                          <updatemodal v-show="isModalVisible" :hostData="modalData" />                    
                          <a data-toggle="modal" data-target="#updateHost" @click="showUpdateModal(host)" style="margin-right: 20px"><i class="fa fa-pencil" style="margin-right: 5px"></i>{{ $t('action.editMsg') }}</a>
                          <a @click="deleteHost(index, host.id)" class="action-delete"> <i class="fa fa-trash" style="margin-right: 5px"></i>{{ $t('action.deleteMsg') }}</a>
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
import modal from './CreateHost.vue'
import updatemodal from './UpdateHost.vue'
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
      hosts: [],
      pagination: [],
      page: 1,
      isModalVisible: false,
      modalData: null
    }
  },
  mounted() {
    this.getHost(this.page)
  },
  methods: {
    getHost(page) {
      axios.get('http://localhost:8081/hosts/page/' + page)
      .then(response => {
        let $this = this
        this.hosts = response.data.records
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
    deleteHost(index, id) {
      if (confirm('Bạn có chắc chắn muốn xóa?')) {
        axios.delete('http://localhost:8081/host/' + id)
        .then(response => {
          this.hosts.splice(index, 1)
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
      this.getHost(page)
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
a:hover {
  cursor: pointer;
}
</style>
