<template>
    <div class="row">
      <div class="col-md-12">
        <div class="box"  v-for="dt in detail" :key="dt.id">
          <div class="box-header">
            <div class="cve-list">
                <h4><router-link :to="{name: 'Cpes'}"><i class="fa fa-list"><span  style="margin-left:5px">CPE List</span></i></router-link></h4>
            </div>
          </div>
          <div class="box-body">
            <div class="cpe-infomation">
              <h3>{{dt.name}}</h3>
                <span>ID: {{dt.name}}</span><br>
                <span v-if='dt.title.Valid'>Title: {{dt.title.String}} <br></span>
                <span v-if='dt.nvd_id.Valid'>NVD ID: {{dt.nvd_id.Int64}}<br></span>
                <span>Created: {{dt.created}}</span> <br>
                <span>Last updated: {{dt.modified}}</span><br>
                <span v-if='dt.status.Valid'>Status: {{dt.status.String}} <br></span>
                <span>Severity: {{dt.severity.String}}</span>
            </div>
            <div class="cpe">
              
              <div>
                <h3>Reported vulnerabilites</h3>
                <div v-if='dt.reportedVulnerabilites != null'>
                  <div class="col-sm-12 table-responsive">
                    <table aria-describedby="example1_info" role="grid" id="example1" class="table table-bordered table-striped dataTable">
                      <thead>
                        <tr role="row">
                          <th style="width: 13%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0">Name</th>
                          <th style="width: 7%" colspan="1" rowspan="1" aria-controls="example1" tabindex="0" >Severity</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr class="odd" role="row" v-for="rp in dt.reportedVulnerabilites" :key="rp.cve">
                          <td class="sorting_1">
                            <router-link :to="{ name: 'Cve Detail', params: {name: rp.cve }}">{{rp.cve}}</router-link></td>
                          <td>{{rp.severity}}</td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
                <div v-else>None</div>
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
    this.getCpe(this.id)
  },
  methods: {
    getCpe(id) {
      axios.get('http://localhost:8081/cpe/' + id)
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

</style>
