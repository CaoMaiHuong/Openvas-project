<template>
  <div class="row">
    <div class="col-md-12">
      <div class="box"  v-for="dt in detail" :key="dt.id">
        <div class="box-header">
          <div class="cve-list">
              <h4><router-link :to="{name: 'Hosts'}"><i class="fa fa-list"><span  style="margin-left:5px">{{ $t('hosts.hostList') }}</span></i></router-link></h4>
          </div>
        </div>
        <div class="box-body">
          <div class="host-name">
            <h3>Host: {{dt.name}}</h3>
            <span>ID: {{dt.uuid}}</span><br>
            <span>{{ $t('createMsg') }}: {{dt.created}}</span> <br>
            <span>{{ $t('lastUpdate') }}: {{dt.modified}}</span><br>
          </div>
          <div class="host-infomation">
            <span> {{ $t('hosts.hostnameMsg') }}:  {{dt.hostname.String}}<br></span>
            <span> {{ $t('hosts.ipMsg') }}:  {{dt.ipaddress.String}}<br></span>
            <span> {{ $t('commentMsg') }}: {{dt.comment}}<br></span>
            <span> {{ $t('severityMsg') }}:  {{dt.severity.String}}<br></span>
            <div class="latest-identifiers">
              <h3>{{ $t('hosts.lastestId') }}</h3>
              <div v-if='dt.identifier != null'>
                <div class="col-sm-12 table-responsive">
                  <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                    <thead>
                      <tr role="row">
                        <th style="width: 15%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">{{ $t('hosts.nameMsg') }}</th>
                        <th style="width: 25%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('hosts.value') }}</th>
                        <th style="width: 25%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('createMsg') }}</th>
                        <th style="width: 30%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >{{ $t('hosts.source') }}</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="i in dt.identifier" :key="i.uuid">
                        <td>{{i.name}}</td>
                        <td>{{i.value}}</td>
                        <td>{{i.created}}</td>
                        <td>
                          <span style="display: flex">{{ $t('hosts.reportMsg') }} &nbsp; &nbsp;<router-link :to="{ name: 'Báo cáo', params: {id: i.source.report}}">{{i.source.report}}</router-link></span>
                          <span v-if='i.source.nvt !=""'>(NVT  &nbsp; &nbsp;<router-link :to="{ name: 'Chi tiết NVT', params: {id: i.source.nvt }}">{{i.source.nvt}}</router-link>)</span>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div v-else>Không</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import $ from 'jquery'
import axios from 'axios'
// Require needed datatables modules

export default {
  name: 'Tables',
  props: ['id'],
  data() {
    return {
      detail: []
    }
  },
  mounted() {
    this.getHost(this.id)
  },
  methods: {
    getHost(id) {
      axios.get('http://localhost:8081/host/' + id)
      .then(response => {
        this.detail = response.data
      })
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
.latest-identifiers td{
    word-break: break-all;
}
</style>
