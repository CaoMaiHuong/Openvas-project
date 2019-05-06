<template>
  <div>
    <div id="updateTask" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Sửa đổi task
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
            <form v-on:submit.prevent="updateTask(taskData.id)" class="update-task" style="padding: 0px">
              <div class="form-group">
                <label class="control-label" for="name">{{ $t('tasks.nameMsg') }}</label>
                <input class="form-control" v-model="taskData.name" name="name" v-validate="'required'" type="text">
                <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
              </div>
              <!-- {{taskData}} -->
              <div class="form-group">
                <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                <input v-model='taskData.comment' class="form-control" name='comment'>
              </div>
              <div class="form-group">
                <label class="control-label" for='target'>{{ $t('tasks.scanTarget') }}</label>
                <select onfocus='this.size=10;' onblur='this.size=1;' onchange='this.size=1; this.blur();' class="form-control" v-model="taskData.target">
                  <option class= "choose-target" v-for="t in targets" :key="t.id" v-bind:value="t.id">{{t.name}}</option>
                </select>
              </div>
              <div class="form-group">
                <label class="control-label" for="alert">{{ $t('tasks.alert') }}</label>
                <select class="form-control" v-model="alert">
                  <option value>Please select one</option>
                </select>
              </div>
              <div class="form-group">
                <label class="control-label" for='schedule'>{{ $t('tasks.schedule') }}</label>
                <select class="form-control" v-model="taskData.schedule.String">
                  <option value>Scan Config Default</option>
                  <option>ICMP Ping</option>
                  <option>ARP Ping</option>
                </select>
              </div>
              <div class="form-group result">
                <label class="control-label" for='in_assets'>{{ $t('tasks.addResult') }}</label>
                <input type="radio" name="in_assets" v-model="taskData.in_assets" value="yes">Có<br>
                <input type="radio" name="in_assets" v-model="taskData.in_assets" style="margin-left: 30px;" value="no">Không<br>
              </div>
              <div class="form-group apply-override">
                <label class="control-label" for='assets_apply_overrides'>{{ $t('tasks.apply') }}</label>
                <input type="radio" name="assets_apply_overrides" v-model="taskData.assets_apply_overrides" value="yes">Có<br>
                <input type="radio" name="assets_apply_overrides" v-model="taskData.assets_apply_overrides" style="margin-left: 30px;"  value="no">Không<br>
              </div>
              <div class="form-group assets_min_qod">
                <label class="control-label" for='assets_min_qod'>{{ $t('tasks.minQod') }}</label>
                <input type="number" name="assets_min_qod" v-model="taskData.assets_min_qod"><br>
              </div>
              <div class="form-group alterable-task">
                <label class="control-label" for='alterable'>{{ $t('tasks.alterableTask') }}</label>
                <input type="radio" name="alterable" v-model="taskData.alterable" value="1">Có<br>
                <input type="radio" name="alterable" v-model="taskData.alterable" style="margin-left: 30px;" value="0">Không<br>
              </div>
              <div class="form-group auto-delete">
                <label class="control-label" for='auto_delete'>{{ $t('tasks.delReport.name') }}</label><br>
                <input type="radio" onclick="document.getElementById('dele_data').disabled = true;" name="auto_delete" v-model="taskData.auto_delete" value="no" >{{ $t('tasks.delReport.delReport1') }}<br>
                <input type="radio" onclick="document.getElementById('dele_data').disabled = false;" name="auto_delete" v-model="taskData.auto_delete" value="keep">{{ $t('tasks.delReport.delReport2') }}&nbsp;
                <input type="number" id="dele_data" class="delete-data" v-model="taskData.auto_delete_data" disabled>&nbsp;báo cáo mới nhất<br>
              </div>
              <div class="form-group">
                <label class="control-label" for='scanner'>{{ $t('tasks.scanner') }}</label>
                <select class="form-control" v-model=taskData.scanner>
                  <option v-for="s in scanners" :key="s.id" v-bind:value=s.id>{{s.name}}</option>
                </select>
              </div>
              <div class="form-group">
                <label class="control-label" for='config'>{{ $t('tasks.scanConfig') }}</label>
                <select class="form-control" v-model=taskData.config>
                  <option v-for="c in configs" :key="c.id" v-bind:value=c.id>{{c.name}}</option>
                </select>
              </div>
              <div class="form-group">
                <label class="control-label" for='network'>{{ $t('tasks.networkInterface') }}</label>
                <input type="text" v-model="network" name="network">
              </div>
              <div class="form-group">
                <label class="control-label" for='hosts_ordering'>{{ $t('tasks.orderTarget') }}</label>
                <select class="form-control" v-model="taskData.hosts_ordering">
                  <option value="sequential">Sequential</option>
                  <option value="random">Random</option>
                  <option value="reverse">Reverse</option>
                </select>
              </div>
              <div class="form-group max_checks">
                <label class="control-label max_checks" for='max_checks'>{{ $t('tasks.maxExecutedNvt') }}</label>
                <input type="number" v-model="taskData.max_checks" name="max_checks"><br>
              </div>
              <div class="form-group max_hosts">
                <label class="control-label max_hosts" for='max_hosts'>{{ $t('tasks.maxScanned') }}</label>
                <input type="number" v-model="taskData.max_hosts" name="max_hosts"><br>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="submit" @click="updateTask(taskData.id)" class="modal-submit btn btn-primary">Lưu</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import axios from 'axios'

  export default {
    name: 'UpdateTask',
    props: ['taskData'],
//     props: {
//       taskData: {
//         type: Array,
//         default: () => []
//       }
//     },
// //     props: {
//   arr: {
//     type: Array,
//     default: () => ({})
//   }
// }
    data() {
      return {
        targets: [],
        scanners: [],
        configs: []
        // // name: '',
        // comment: '',
        // // target: '',
        // alert: '',
        // schedule: '',
        // in_assets: 'yes',
        // assets_apply_overrides: 'yes',
        // assets_min_qod: '70',
        // alterable: '0',
        // auto_delete: 'no',
        // auto_delete_data: '5',
        // scanner: 1,
        // config: 1,
        // network: '',
        // hosts_ordering: 'sequential',
        // max_checks: '4',
        // max_hosts: '20'
      }
    },
    // computed: {
    //   name: {
    //     get() {
    //       return this.task_name
    //     },
    //     set(newValue) {
    //       this.$emit('update:task_name', newValue)
    //     }
    //   }
    // },
    created() {
      this.getTarget()
    },
    mounted() {
      this.getScanner()
    },
    beforeMount() {
      this.getConfig()
    },
    methods: {
      getTarget() {
        axios.get('http://localhost:8081/task_target')
        .then(response => {
          this.targets = response.data
        })
      },
      // beforeCreated() {
      //   this.name = this.taskData.name
      // },
      // computed: {
      //   namee() {
      //     return 'fgdfgfdg'
      //   }
      // },
      getScanner() {
        axios.get('http://localhost:8081/scanners')
        .then(response => {
          this.scanners = response.data
        })
      },
      getConfig() {
        axios.get('http://localhost:8081/configs')
        .then(response => {
          this.configs = response.data
        })
      },
      updateTask(id) {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'put',
              url: 'http://localhost:8081/task/' + id,
              data: {
                name: this.taskData.name,
                comment: this.taskData.comment,
                target: this.taskData.target,
                alert: this.taskData.alert,
                schedule: this.taskData.schedule,
                in_assets: this.taskData.in_assets,
                assets_apply_overrides: this.taskData.assets_apply_overrides,
                assets_min_qod: this.taskData.assets_min_qod,
                alterable: this.taskData.alterable,
                auto_delete: this.taskData.auto_delete,
                auto_delete_data: this.taskData.auto_delete_data,
                scanner: this.taskData.scanner,
                config: this.taskData.config,
                network: this.taskData.network,
                hosts_ordering: this.taskData.hosts_ordering,
                max_checks: this.taskData.max_checks,
                max_hosts: this.taskData.max_hosts
              }
            })
            .then(response => {
              // // this.$router.push('/targets')
              // this.$router.go(0)
            })
          }
        })
      }
    }
  }
</script>
<style>
  #updateTask .form-control{
    display: unset;
    width: 100%;
  } 
  #updateTask .form-group{
    display: flex;
    align-items: center;
    margin-bottom: 15px;
    flex-wrap: wrap;  
  }
  #updateTask label {
    min-width: 125px;
  } 
  #updateTask .modal-body {
    height: 400px;
    overflow: overlay;
    padding: 15px 30px;
    }
    #updateTask .form-group.auto-delete{
      display: unset;
    }
  .result, .apply-override, .alterable-task {
    display: flex;
  }
  .max_checks, .max_hosts, .result label, .apply-override label, .alterable-task label{
    margin-right: 30px;
  }
  .result:last-child, .apply-override input:last-child, .alterable-task input:last-child {
    margin-left: 30px;
  }
</style>
