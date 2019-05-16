<template>
  <section class="content">
    <div class="row">
      <div class="col-md-12">
        <div class="box">
          <div class="box-body">
            <div class="dataTables_wrapper form-inline dt-bootstrap" id="example1_wrapper">
              <div class="row">
                <div class="col-sm-6">
                  <div id="example1_length" class="dataTables_length">

                  </div>
                </div>
              </div>
              <div class="search">
                <form v-on:submit.prevent="searchNvts(1)" class="search-form">
                  <input style="width:250px" class="search-content form-control" v-model="search" type="text" name="search" v-validate="'required'" placeholder="Tìm kiếm theo tên">
                  <button type="submit" class="btn btn-primary"><i class="fa fa-search"></i></button>
                </form>
              </div>
              <div v-if="message" class="not-found">
                {{message}}
                <div>
                  <br>
                  <button @click="returnList" class="btn btn-primary">Xem danh sách NVT</button>
                </div>
              </div>
              <div class="row" v-if="!message">
                <div class="col-sm-12 table-responsive">
                  <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                    <thead>
                      <tr role="row">
                        <th style="width: 25%" aria-sort="ascending" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('nvts.nameMsg') }}</th>
                        <th style="width: 13%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('nvts.familyMsg') }}</th>
                        <th style="width: 13%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('createMsg') }}</th>
                        <th style="width: 13%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('modifyMsg') }}</th>
                        <th style="width: 23%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">CVE</th>
                        <th style="width: 6%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('severityMsg') }}</th>
                        <th style="width: 7%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('nvts.qodMsg') }}</th>
                        <!-- <th colspan="1" rowspan="1" aria-controls="example1" tabindex="0"></th> -->
                      </tr>
                    </thead>
                    <tbody>
                      <tr class="odd" role="row" v-for="nvt in nvts" :key="nvt.uuid">
                        <td class="sorting_1">
                          <router-link :to="{ name:'Chi tiết NVT', params: {id: nvt.uuid }}">{{nvt.name}}</router-link></td>
                        <td>{{nvt.family}}</td>
                        <td>{{nvt.created}}</td>
                        <td>{{nvt.modified}}</td>
                        <td><router-link class="cve" v-for="cve in nvt.cve" :key="cve" :to="{ name: 'Chi tiết CVE', params: {name: cve }}">{{cve}}&nbsp;&nbsp;&nbsp;</router-link></td>
                        <td>{{nvt.severity}}</td>
                        <td>{{nvt.qod}}</td>
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
      nvts: [],
      pagination: [],
      page: 1,
      search: '',
      searching: false,
      message: ''
    }
  },
  mounted() {
    this.getNvt(this.page)
  },
  methods: {
    getNvt(page) {
      axios.get('http://localhost:8081/nvts/page/' + page)
      .then(response => {
        let $this = this
        this.nvts = response.data.records
        $this.makePagination(response.data)
      })
    },
    searchNvts(page) {
      this.$validator.validateAll().then(res => {
        if (res) {
          axios({
            method: 'get',
            url: 'http://localhost:8081/nvts/page/' + page + '/search/' + this.search,
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
              this.nvts = response.data.records
              this.searching = true
              $this.makePagination(response.data)
            }
          })
        } else {
          this.getNvt(this.page)
        }
      })
    },
    returnList() {
      this.message = ''
      this.getNvt(this.page)
      this.search = ''
      this.searching = false
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
      this.getNvt(page)
    },
    fetchPaginate2(page) {
      this.searchNvts(page)
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
.cve{
  display: inline-flex;
}

.not-found{
  margin-top: 20px;
}
</style>
