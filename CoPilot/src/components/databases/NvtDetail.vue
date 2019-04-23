<template>
  <div>
    <div class="row">
      <div class="col-md-12">
        <div class="box"  v-for="dt in detail" :key="dt.uuid">
          <div class="box-header">
            <div class="cve-list">
                <h4><router-link :to="{name: 'Nvts'}"><i class="fa fa-list"><span  style="margin-left:5px">NVT List</span></i></router-link></h4>
            </div>
          </div>
          <div class="box-body">
            <div class="nvt-name">
              <h3>NVT: {{dt.name}}</h3>
                <span>ID: {{dt.uuid}}</span><br>
                <span>Published: {{dt.created}}</span> <br>
                <span>Last updated: {{dt.modified}}</span><br>
            </div>
            <div class="nvt-infomation">
              <div class="nvt-summary" v-if='dt.tag.summary!=""'>
                <h3>Summary</h3>
                <div>{{dt.tag.summary}}</div>
              </div>
              <div class="scoring">
                  <h3>Scoring</h3>
                  <span> CVSS Base:  {{dt.severity}}<br></span>
                  <span v-if='dt.tag.cvss_vector!=""'> CVSS Base Vector: {{dt.tag.cvss_vector}}</span>
              </div>
              <div class="insight" v-if='dt.tag.insight!=""'>
                  <h3>Insight</h3>
                  <span> {{dt.tag.insight}}</span>
              </div>
              <div class="detection">
                  <h3>Detection Method</h3>
                  <span> Quality of Detection:  {{dt.tag.qod_type}}({{dt.qod}})</span>
              </div>
              <div class="affected" v-if='dt.tag.affected!=""'>
                  <h3>Affected Software/OS</h3>
                  <span> {{dt.tag.affected}}</span>
              </div>
              <div class="impact" v-if='dt.tag.impact!=""'>
                  <h3>Impact</h3>
                  <span> {{dt.tag.impact}}</span>
              </div>
              <div class="solution">
                  <h3>Solution</h3>
                  <span> Solution Type: {{dt.tag.solution_type}}<br></span>
                  <span> {{dt.tag.solution}}</span>
              </div>
              <div class="family">
                  <h3>Family</h3>
                  <span> {{dt.family}}</span>
              </div>
                  
              <div class="references">
                <h3>References</h3>
                <div v-if='dt.cve != ""'>
                    <span>CVE</span><br>
                  <router-link v-for="c in dt.cve" :key="c" :to="{ name: 'Cve Detail', params: {name: c }}">{{c}}<br>
                  </router-link>
                </div>
                <div v-if='dt.xref != null'>
                  <span>Other</span>
                  <div v-for="x in dt.xref" :key="x"> 
                    <a v-bind:href=x target="_blank">{{x}}<br></a>
                  </div>
                </div>
              </div>
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
    this.getNvt(this.id)
  },
  methods: {
    getNvt(id) {
      axios.get('http://localhost:8081/nvt/' + id)
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
