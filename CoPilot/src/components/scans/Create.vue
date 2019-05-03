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
              <form v-on:submit.prevent="createTarget" class="create-user" style="padding: 0px">
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
                  <select class="form-control" v-model="portlist">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for="port-list">{{ $t('tasks.alert') }}</label>
                  <select class="form-control" v-model="portlist">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='alivetest'>{{ $t('tasks.schedule') }}</label>
                  <select class="form-control" v-model="alivetest">
                    <option value>Scan Config Default</option>
                    <option>ICMP Ping</option>
                    <option>ARP Ping</option>
                  </select>
                </div>
                <div class="form-group result">
                  <label class="control-label" for='result'>{{ $t('tasks.addResult') }}</label>
                  <input type="radio" name="result" value="Yes">Có<br>
                  <input type="radio" name="result" value="No">Không<br>
                </div>
                <div class="form-group apply-override">
                  <label class="control-label" for='apply-override'>{{ $t('tasks.apply') }}</label>
                  <input type="radio" name="apply-override" value="Yes">Có<br>
                  <input type="radio" name="apply-override" value="No">Không<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='min-qod'>{{ $t('tasks.minQod') }}</label>
                  <input type="number" name="min-qod" value="70"><br>
                </div>
                <div class="form-group alterable-task">
                  <label class="control-label" for='alterable-task'>{{ $t('tasks.alterableTask') }}</label>
                  <input type="radio" name="alterable-task" value="Yes">Có<br>
                  <input type="radio" name="alterable-task" value="No">Không<br>
                </div>
                <div class="form-group auto-delete">
                  <label class="control-label" for='auto-delete'>{{ $t('tasks.delReport.name') }}</label><br>
                  <input type="radio" name="auto-delete" value="Yes">{{ $t('tasks.delReport.delReport1') }}<br>
                  <input type="radio" name="auto-delete" value="No">{{ $t('tasks.delReport.delReport2') }}&nbsp;<input type="number">&nbsp;báo cáo mới nhất<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scanner'>{{ $t('tasks.scanner') }}</label>
                  <select class="form-control" v-model="scanner">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scan_config'>{{ $t('tasks.scanConfig') }}</label>
                  <select class="form-control" v-model="scan_config">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='network'>{{ $t('tasks.networkInterface') }}</label>
                  <input type="text" name="network"><br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scanner'>{{ $t('tasks.orderTarget') }}</label>
                  <select class="form-control" v-model="order">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='max_excutednvt'>{{ $t('tasks.maxExecutedNvt') }}</label>
                  <input type="number" name="max_excutednvt"><br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='max_scanhost'>{{ $t('tasks.maxScanned') }}</label>
                  <input type="number" name="max_scanhost"><br>
                </div>
              </form>
          </div>
          <div class="modal-footer">
              <button type="submit" class="btn btn-primary" data-dismiss="modal">Lưu</button>
          </div>
        </div>
      </div>
    </div>
</template>
<style>
  #addTask {
    display: flex;
    align-items: center;
  }
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
</style>
