<template>
    <div id="addTask" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              Tạo mới task
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="createTask" class="create-user" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">{{ $t('tasks.nameMsg') }}</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">{{ $t('commentMsg') }}</label>
                  <input v-model='comment' class="form-control" name='comment'>
                </div>
                <div class="form-group">
                  <label class="control-label" for='target'>{{ $t('tasks.scanTarget') }}</label>
                  <select onfocus='this.size=10;' onblur='this.size=1;' onchange='this.size=1; this.blur();' name="target" class="form-control"  v-model="target">
                    <option class= "choose-target" v-for="t in targets" :key="t.id" v-bind:value="t.id" >{{t.name}}</option>
                  </select>
                  <span v-if="errors.has('target')">{{ errors.first('target') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="alert">{{ $t('tasks.alert') }}</label>
                  <select class="form-control" v-model="alert">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='schedule'>{{ $t('tasks.schedule') }}</label>
                  <select class="form-control" v-model="schedule">
                    <option></option>
                  </select>
                </div>
                <div class="form-group result">
                  <label class="control-label" for='in_assets'>{{ $t('tasks.addResult') }}</label>
                  <input type="radio" name="in_assets" v-model="in_assets" value="yes">Có<br>
                  <input type="radio" name="in_assets" style="margin-left: 30px;" v-model="in_assets" value="no">Không<br>
                </div>
                <div class="form-group apply-override">
                  <label class="control-label" for='assets_apply_overrides'>{{ $t('tasks.apply') }}</label>
                  <input type="radio" name="assets_apply_overrides" v-model="assets_apply_overrides" value="yes">Có<br>
                  <input type="radio" name="assets_apply_overrides" style="margin-left: 30px;" v-model="assets_apply_overrides" value="no">Không<br>
                </div>
                <div class="form-group assets_min_qod">
                  <label class="control-label" for='assets_min_qod'>{{ $t('tasks.minQod') }}</label>
                  <input type="number" class="form-control" name="assets_min_qod" v-model="assets_min_qod"><br>
                </div>
                <div class="form-group alterable-task">
                  <label class="control-label" for='alterable'>{{ $t('tasks.alterableTask') }}</label>
                  <input type="radio" name="alterable" v-model="alterable" value="1">Có<br>
                  <input type="radio" name="alterable" style="margin-left: 30px;" v-model="alterable" value="0">Không<br>
                </div>
                <div class="form-group auto-delete">
                  <label class="control-label" for='auto_delete'>{{ $t('tasks.delReport.name') }}</label><br>
                  <input type="radio" onclick="document.getElementById('del_data').disabled = true;" name="auto_delete" v-model="auto_delete" value="no" >{{ $t('tasks.delReport.delReport1') }}<br>
                  <input type="radio" onclick="document.getElementById('del_data').disabled = false;" name="auto_delete" v-model="auto_delete" value="keep">{{ $t('tasks.delReport.delReport2') }}&nbsp;
                  <input type="number" class="delete-data" id="del_data" v-model="auto_delete_data" disabled>&nbsp;báo cáo mới nhất<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scanner'>{{ $t('tasks.scanner') }}</label>
                  <select class="form-control" v-model="scanner">
                    <option v-for="s in scanners" :key="s.id" v-bind:value=s.id>{{s.name}}</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='config'>{{ $t('tasks.scanConfig') }}</label>
                  <select class="form-control" v-model="config">
                    <option v-for="c in configs" :key="c.id" v-bind:value=c.id>{{c.name}}</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='network'>{{ $t('tasks.networkInterface') }}</label>
                  <input type="text" class="form-control" v-model="network" name="network">
                </div>
                <div class="form-group">
                  <label class="control-label" for='hosts_ordering'>{{ $t('tasks.orderTarget') }}</label>
                  <select class="form-control" v-model="hosts_ordering">
                    <option value="sequential">Sequential</option>
                    <option value="random">Random</option>
                    <option value="reverse">Reverse</option>
                  </select>
                </div>
                <div class="form-group max_checks">
                  <label class="control-label" for='max_checks'>{{ $t('tasks.maxExecutedNvt') }}</label>
                  <input type="number"  class="form-control" v-model="max_checks" name="max_checks"><br>
                </div>
                <div class="form-group max_hosts">
                  <label class="control-label" for='max_hosts'>{{ $t('tasks.maxScanned') }}</label>
                  <input type="number" class="form-control" v-model="max_hosts" name="max_hosts"><br>
                </div>
                
              </form>
          </div>
          <div class="modal-footer">
            <button type="submit"  @click="createTask()" class="btn btn-primary modal-submit" >Lưu</button>
          </div>
        </div>
      </div>
    </div>
</template>
<script>
  import axios from 'axios'

  export default {
    name: 'CreateTask',
    data(router) {
      return {
        targets: [],
        scanners: [],
        configs: [],
        name: '',
        comment: '',
        target: '',
        alert: '',
        schedule: '',
        in_assets: 'yes',
        assets_apply_overrides: 'yes',
        assets_min_qod: '70',
        alterable: '0',
        auto_delete: 'no',
        auto_delete_data: '5',
        scanner: 1,
        config: 1,
        network: '',
        hosts_ordering: 'sequential',
        max_checks: '4',
        max_hosts: '20'
      }
    },
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
      createTask() {
        this.$validator.validateAll().then(res => {
          if (res) {
            axios({
              method: 'post',
              url: 'http://localhost:8081/task',
              data: {
                name: this.name,
                comment: this.comment,
                target: this.target,
                alert: this.alert,
                schedule: this.schedule,
                in_assets: this.in_assets,
                assets_apply_overrides: this.assets_apply_overrides,
                assets_min_qod: this.assets_min_qod,
                alterable: this.alterable,
                auto_delete: this.auto_delete,
                auto_delete_data: this.auto_delete_data,
                scanner: this.scanner,
                config: this.config,
                network: this.network,
                hosts_ordering: this.hosts_ordering,
                max_checks: this.max_checks,
                max_hosts: this.max_hosts
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
  /* #addTask {
    display: flex;
    align-items: center;
  } */
  #addTask .modal-body {
    height: 400px;
    overflow: overlay;
    padding: 15px 30px;
  }
  .result, .apply-override, .alterable-task {
    display: flex;
  }
  .result label, .apply-override label, .alterable-task label{
    margin-right: 30px;
  }
  .result:last-child, .apply-override input:last-child, .alterable-task input:last-child {
    margin-left: 30px;
  }
  .assets_min_qod {
    display: flex;
    align-items: center;
  }
   .assets_min_qod label{
     width: 90px;
   }
  .delete-data{
    width: 50px;
    text-align: center;
  }
  .max_checks, .max_hosts {
    display: flex;
    align-items: center;
  }
  .max_checks input, .max_hosts input{
    width: 80px;
    text-align: center;
  }
  .max_checks label, .max_hosts label{
    margin-right: 15px;
  }
  .modal-submit {
    margin-right: 20px;
  }
</style>
