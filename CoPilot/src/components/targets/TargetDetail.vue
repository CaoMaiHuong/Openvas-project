<template>
  <div class="row">
    <div class="col-md-12">
      <div class="box"  v-for="dt in detail" :key="dt.id">
        <div class="box-header">
          <div class="cve-list">
              <h4><router-link :to="{name: 'Targets'}"><i class="fa fa-list"><span  style="margin-left:5px">{{ $t('targets.targetListMsg') }}</span></i></router-link></h4>
          </div>
        </div>
        <div class="box-body">
          <div class="target-name">
            <h3>{{dt.name}}</h3>
              <span>ID: {{dt.uuid}}</span><br>
              <span>{{ $t('createMsg') }}: {{dt.created}}</span> <br>
              <span>{{ $t('modifyMsg') }}: {{dt.modified}}</span><br>
          </div>
          <div class="target-infomation">
            <div class="host">
              <h3>Host</h3>
              <span> {{ $t('targets.includedMsg') }}:  {{dt.hosts}}<br></span>
              <span> {{ $t('targets.rlOnlyMsg') }}: {{dt.rlonly}}<br></span>
              <span> {{ $t('targets.rlUnifyMsg') }}: {{dt.rlunify}}<br></span>
              <span> {{ $t('targets.aliveTestMsg') }}: {{dt.alivetest}}<br></span>
              <span> {{ $t('targets.portMsg') }}: {{dt.portlist}}</span>
            </div>
            <div class="tasks-using">
              <h3>{{ $t('targets.tasksMsg') }}</h3>
              <div v-if='dt.task != null'>
                <router-link v-for="t in dt.task" :key="t.id" :to="{ name: 'Task Detail', params: {id: t.id }}">{{t.name}}<br>
                </router-link>
              </div>
              <div v-if='dt.task == null'>None</div>
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
    this.getTarget(this.id)
  },
  methods: {
    getTarget(id) {
      axios.get('http://localhost:8081/target/' + id)
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
