<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box">
          <!-- <div class="box-header">
            <router-link to="/createuser"><i class="fa fa-user-plus" style="margin-right: 3px"></i> CREATE</router-link>
          </div> -->
          <!-- /.box-header -->
          <div class="box-body">
            <div class="dataTables_wrapper form-inline dt-bootstrap" id="example1_wrapper">
              <div class="row">
                <div class="col-sm-6">
                  <div id="example1_length" class="dataTables_length">

                  </div>
                </div>
              </div>
              <div class="search">
                <form v-on:submit.prevent="searchCpes(1)" class="search-form">
                  <input style="width:250px" class="search-content form-control" v-model="search" name="search" v-validate="'required'" type="text" placeholder="Tìm kiếm theo tiêu đề">
                  <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i></button>
                </form>
              </div>
              <div v-if="message" class="not-found">
                {{message}}
              </div>
              <div class="row" v-if="!message">
                <div class="col-sm-12 table-responsive">
                  <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                    <thead>
                      <tr role="row">
                        <th style="width: 30%" aria-sort="ascending" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('cpes.nameMsg') }}</th>
                        <th style="width: 23%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('cpes.titleMsg') }}</th>
                        <th style="width: 27%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('modifyMsg') }}</th>
                        <th style="width: 10%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Cves</th>
                        <th style="width: 10%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('severityMsg') }}</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="cpe in cpes" :key="cpe.id">
                        <td class="sorting_1">
                          <router-link :to="{ name: 'Chi tiết CPE', params: {id: cpe.id }}">{{cpe.name}}</router-link></td>
                        <td>{{cpe.title.String}}</td>
                        <td>{{cpe.modified}}</td>
                        <td>{{cpe.cves}}</td>
                        <td>{{cpe.severity.String}}</td>
                      </tr>
                    </tbody>
                    <tfoot>
                      
                    </tfoot>
                  </table>
                  <div class="pagination" v-if="!this.searching">
                    <button class="btn btn-primary" v-on:click="fetchPaginate(1)" :disabled="pagination.page == 1"><i class="fa fa-angle-double-left"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.prev_page)" :disabled="pagination.page == 1"><i class="fa fa-angle-left"></i></button>
                    <span>Trang {{ pagination.page }} / {{ pagination.total_page }} </span>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.next_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-right"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate(pagination.total_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-double-right"></i></button>
                  </div>
                  <div class="pagination" v-if="this.searching">
                    <button class="btn btn-primary" v-on:click="fetchPaginate2(1)" :disabled="pagination.page == 1"><i class="fa fa-angle-double-left"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate2(pagination.prev_page)" :disabled="pagination.page == 1"><i class="fa fa-angle-left"></i></button>
                    <span>Trang {{ pagination.page }} / {{ pagination.total_page }} </span>
                    <button class="btn btn-primary" v-on:click="fetchPaginate2(pagination.next_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-right"></i></button>
                    <button class="btn btn-primary" v-on:click="fetchPaginate2(pagination.total_page)" :disabled="pagination.page == pagination.total_page"><i class="fa fa-angle-double-right"></i></button>
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
// Require needed datatables modules
require('datatables.net')
require('datatables.net-bs')

export default {
  name: 'Tables',
  data() {
    return {
      cpes: [],
      pagination: [],
      page: 1,
      search: '',
      searching: false,
      message: ''
    }
  },
  mounted() {
    this.getCpes(this.page)
  },
  methods: {
    getCpes(page) {
      axios({
        method: 'get',
        url: 'http://localhost:8081/cpes/page/' + page
      })
      .then(response => {
        let $this = this
        this.cpes = response.data.records
        $this.makePagination(response.data)
        this.searching = false
      })
    },
    searchCpes(page) {
      this.$validator.validateAll().then(res => {
        if (res) {
          axios({
            method: 'get',
            url: 'http://localhost:8081/cpes/page/' + page + '/search/' + this.search,
            data: {
              search: this.search
            }
          })
          .then(response => {
            if (response.data.records === null) {
              this.message = 'Không tìm thấy kết quả!'
            } else {
              this.message = ''
              let $this = this
              this.cpes = response.data.records
              this.searching = true
              $this.makePagination(response.data)
            }
          })
        } else {
          this.getCpes(this.page)
        }
      })
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
      this.getCpes(page)
    },
    fetchPaginate2(page) {
      this.searchCpes(page)
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
.search-form{
  padding: 0px;
}
</style>
