<template>
  <div>
    <div class="row">
      <div class="col-md-12">
        <div class="box"  v-for="dt in detail" :key="dt.uuid">
          <div class="box-header">
            <div class="cve-list">
                <h4><router-link :to="{name: 'Nvts'}"><i class="fa fa-list"><span  style="margin-left:5px">{{ $t('nvts.nvtListMsg') }}</span></i></router-link></h4>
            </div>
          </div>
          <div class="box-body">
            <div class="nvt-name">
              <h3>NVT: {{dt.name}}</h3>
                <span>ID: {{dt.uuid}}</span><br>
                <span>{{ $t('publishMsg') }}: {{dt.created}}</span> <br>
                <span>{{ $t('lastUpdate') }}: {{dt.modified}}</span><br>
            </div>
            <div class="nvt-infomation">
              <div class="nvt-summary" v-if='dt.tag.summary!=""'>
                <h3>{{ $t('nvts.summaryMsg') }}</h3>
                <div>{{dt.tag.summary}}</div>
              </div>
              <div class="scoring">
                  <h3>{{ $t('nvts.scoring.name') }}</h3>
                  <span> {{ $t('nvts.scoring.base') }}:  {{dt.severity}}<br></span>
                  <span v-if='dt.tag.cvss_vector!=""'> {{ $t('nvts.scoring.baseVector') }}: {{dt.tag.cvss_vector}}</span>
              </div>
              <div class="insight" v-if='dt.tag.insight!=""'>
                  <h3>{{ $t('nvts.insightMsg') }}</h3>
                  <span> {{dt.tag.insight}}</span>
              </div>
              <div class="detection">
                  <h3>{{ $t('nvts.detecMsg.name') }}</h3>
                  <div>{{dt.tag.vuldetect}}</div>
                  <span> {{ $t('nvts.detecMsg.qod') }}:  {{dt.tag.qod_type}}&nbsp;({{dt.qod}})</span>
              </div>
              <div class="affected" v-if='dt.tag.affected!=""'>
                  <h3>{{ $t('nvts.affectMsg') }}</h3>
                  <span> {{dt.tag.affected}}</span>
              </div>
              <div class="impact" v-if='dt.tag.impact!=""'>
                  <h3>{{ $t('nvts.impactMsg') }}</h3>
                  <span> {{dt.tag.impact}}</span>
              </div>
              <div class="solution">
                  <h3>{{ $t('nvts.solutionMsg') }}</h3>
                  <span> {{ $t('nvts.solutionType') }}: {{dt.tag.solution_type}}<br></span>
                  <span> {{dt.tag.solution}}</span>
              </div>
              <div class="family">
                  <h3>{{ $t('nvts.familyMsg') }}</h3>
                  <span> {{dt.family}}</span>
              </div>
                  
              <div class="references">
                <h3>{{ $t('nvts.refer') }}</h3>
                <div v-if='dt.cve != ""'>
                    <span>CVE</span><br>
                  <router-link v-for="c in dt.cve" :key="c" :to="{ name: 'Chi tiết CVE', params: {name: c }}">{{c}}<br>
                  </router-link>
                </div>
                <div v-if='dt.xref != null'>
                  <span>{{ $t('nvts.otherMsg') }}</span>
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
