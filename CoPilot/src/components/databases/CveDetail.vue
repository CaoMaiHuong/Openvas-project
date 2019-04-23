<template>
<div>
  <div class="row">
    <div class="col-md-12">
      <div class="box"  v-for="dt in detail" :key="dt.id">
        <div class="box-header">
          <div class="cve-list">
              <h4><router-link :to="{name: 'Cves'}"><i class="fa fa-list"><span  style="margin-left:5px">CVE List</span></i></router-link></h4>
          </div>
        </div>
        <div class="box-body">
          <div class="cve-name">
            <h3>{{dt.name}}</h3>
              <span>ID: {{dt.name}}</span><br>
              <span>Published: {{dt.published}}</span> <br>
              <span>Last updated: {{dt.modified}}</span><br>
          
          </div>
          <div class="cve-infomation">
            <div class="cve-description">
              <h3>Description</h3>
              <div>{{dt.description}}</div>
            </div>
            <div class="cvss">
              <h3>CVSS</h3>
              <span> Base Score:  {{dt.severity.String}}</span>
              <div v-if="dt.severity.String != 'N/A'">
                <div><span>Access Vector: {{dt.vector}}</span></div>
                <div><span>Access Complexity: {{dt.complexity}}</span></div>
                <div><span>Authentication: {{dt.authentication}}</span></div>
                <div><span>Confidentiality Impact: {{dt.confidentiality_impact}}</span></div>
                <div><span>Integrity Impact: {{dt.integrity_impact}}</span></div>
                <div><span>Availability Impact: {{dt.availability_impact}}</span></div>
              </div>
            </div>
            <!-- <div v-if='dt.product != ""'>
              <h3>Vulnerable Products</h3>
              <div><router-link v-for="pr in dt.product" :key="pr" :to="{ name: 'Cve Detail', params: {name: cve }}">{{pr}}<br></router-link></div>
            </div> -->
            <div >
              <h3>Vulnerable Products</h3>
              <div v-if='dt.vulnerableProduct != null'>
                <router-link v-for="pr in dt.vulnerableProduct" :key="pr.id" :to="{ name: 'Cpe Detail', params: {id: pr.id }}">{{pr.name}}<br>
                </router-link>
              </div>
              <div v-if='dt.vulnerableProduct == null'>None</div>
            </div>
            <h3>NVTs addressing this CVE</h3>
            <div v-if='dt.nvt != null'>
              <router-link v-for="n in dt.nvt" :key="n.nvt_id" :to="{ name: 'Nvt Detail', params: {id: n.nvt_id}}">{{n.nvt_name}}<br></router-link></div>
            </div>
            <div v-if='dt.nvt == null'>None</div>
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
  props: ['name'],
  data() {
    return {
      detail: []
    }
  },
  mounted() {
    this.getCve(this.name)
  },
  methods: {
    getCve(name) {
      axios.get('http://localhost:8081/cve/' + name)
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
