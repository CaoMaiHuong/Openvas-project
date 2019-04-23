<template>
    <div id="addTask" class="modal fade" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <span class="modal-header__title">
              New Task
            </span>
            <button type="button" class="close" data-dismiss="modal">&times;</button>  
          </div>
          <div class="modal-body">   
              <form v-on:submit.prevent="createTarget" class="create-user" style="padding: 0px">
                <div class="form-group">
                  <label class="control-label" for="name">Name</label>
                  <input class="form-control" v-model="name" name="name" v-validate="'required'" placeholder="Enter ..." type="text">
                  <span v-if="errors.has('name')">{{ errors.first('name') }}</span>
                </div>
                <div class="form-group">
                  <label class="control-label" for="comment">Comment</label>
                  <input v-model='comment' class="form-control" name='comment'>
                </div>
                <div class="form-group">
                  <label class="control-label" for='target'>Scan Targets</label>
                  <select class="form-control" v-model="portlist">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for="port-list">Alerts</label>
                  <select class="form-control" v-model="portlist">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='alivetest'>Schedule</label>
                  <select class="form-control" v-model="portlist">
                    <option value>Scan Config Default</option>
                    <option>ICMP Ping</option>
                    <option>ARP Ping</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='result'>Add results to Assets</label>
                  <input type="radio" name="result" value="Yes">Yes<br>
                  <input type="radio" name="result" value="No">No<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='apply-override'>Apply Override</label>
                  <input type="radio" name="apply-override" value="Yes">Yes<br>
                  <input type="radio" name="apply-override" value="No">No<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='min-qod'>Min QoD</label>
                  <input type="number" name="min-qod" value="70"><br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='alterable-task'>Alterable Task</label>
                  <input type="radio" name="alterable-task" value="Yes">Yes<br>
                  <input type="radio" name="alterable-task" value="No">No<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='auto-delete'>Auto Delete Reports</label>
                  <input type="radio" name="auto-delete" value="Yes">Do not automatically delete reports<br>
                  <input type="radio" name="auto-delete" value="No">Automatically delete oldest reports but always keep newest<br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scanner'>Scanner</label>
                  <select class="form-control" v-model="scanner">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scan_config'>Scan Config</label>
                  <select class="form-control" v-model="scan_config">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='network'>Network Source Interface</label>
                  <input type="text" name="network"><br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='scanner'>Order for target hosts</label>
                  <select class="form-control" v-model="order">
                    <option value>Please select one</option>
                  </select>
                </div>
                <div class="form-group">
                  <label class="control-label" for='max_excutednvt'>Maximum concurrently executed NVTs per host</label>
                  <input type="number" name="max_excutednvt"><br>
                </div>
                <div class="form-group">
                  <label class="control-label" for='max_scanhost'>Maximum concurrently scanned hosts</label>
                  <input type="number" name="max_scanhost"><br>
                </div>
              </form>
          </div>
          <div class="modal-footer">
              <button type="submit" class="btn btn-primary" data-dismiss="modal">Save</button>
          </div>
        </div>
      </div>
    </div>
</template>
