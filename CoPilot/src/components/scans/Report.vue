<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box" v-if="report != null">
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
                      <tr>
                        <!-- <th style="width: 30px" aria-sort="ascending" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Id</th> -->
                        <th rowspan="2" aria-controls="example1" tabindex="0">{{ $t('tasks.vul') }}</th>
                        <th rowspan="2" aria-controls="example1" tabindex="0">{{ $t('tasks.severity') }}</th>
                        <th rowspan="2" aria-controls="example1" tabindex="0">QoD</th>
                        <th colspan="2" aria-controls="example1" tabindex="0" style="text-align: center">Host</th>
                        <th rowspan="2" aria-controls="example1" tabindex="0" >{{ $t('tasks.location') }}</th>
                        <th rowspan="2" aria-controls="example1" tabindex="0" >{{ $t('createMsg') }}</th>
                        <!-- <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0"></th> -->
                      </tr>
                      <tr>
                        <th>IP</th>
                        <th class="report-hostname" style="border-right-width: thin;">Name</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="r in report" :key="r.id">
                        <!-- <td class="sorting_1">{{user.ID}}</td> -->
                        <td><router-link :to="{ name: 'Báo cáo', params: {id: report.uuid}}">{{r.vulnerability}}</router-link></td>
                        <td>{{r.severity}}</td>
                        <td>{{r.qod}}</td>
                        <td>{{r.host.ip}}</td>
                        <td>{{r.host.name}}</td>
                        <td>{{r.location}}</td>
                        <td>{{r.created}}</td>
                      </tr>
                    </tbody>
                    
                    <tfoot>
                      
                    </tfoot>
                  </table>
                  <div class="pagination">
                    <button class="btn btn-primary" v-on:click="fetchPaginate(1)" :disabled="pagination.page == 1"><i class="fa fa-angle-double-left"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.prev_page)" :disabled="pagination.page == 1"><i class="fa fa-angle-left"></i></button>
                    <span>Page {{ pagination.page }} of {{ pagination.total_page }} </span>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.next_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-right"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.total_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-double-right"></i></button>
                  </div>
                </div>
              </div>
            </div>
            <!-- /.box-body -->
          </div>
        </div>
        <div v-if="report == null">
            <span> Not report </span>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
// import $ from 'jquery'
import axios from 'axios'
// Require needed datatables modules
require('datatables.net')
require('datatables.net-bs')

export default {
  name: 'Tables',
  props: ['id'],
  data() {
    return {
      report: [],
      pagination: [],
      page: 1,
      isModalVisible: false
    }
  },
  mounted() {
    this.getReport(this.page, this.id)
  },
  methods: {
    getReport(page, id) {
      axios.get('http://localhost:8081/report/' + id + '/page/' + page)
      .then(response => {
        let $this = this
        this.report = response.data.records
        $this.makePagination(response.data)
      })
    },
    showModal() {
      this.isModalVisible = true
    },
    closeModal() {
      this.isModalVisible = false
    },
    // deleteTarget: function(id) {
    //   if (confirm('Do you really want to delete it?')) {
    //     axios.delete('http://localhost:8081/user/' + id)
    //     .then(response => {
    //       location.reload()
    //     })
    //   }
    // },
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
      this.getReport(page, this.id)
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
.report-hostname {
  border-right-width: thin;
}
</style>
