<template>
  <!-- Main content -->
  <section class="content">
    <!-- GitHub hint -->
    <div class="row" v-for="dash in dashboards" :key="dash.nvtNumber">
      <!-- Info boxes -->
      <router-link to="/nvts">
      <div class="col-md-4 col-sm-6 col-xs-12">
        <info-box color-class="bg-aqua"
                  title="NVT"
                  text="NVTs"
                  :number="dash.nvtNumber"></info-box>
      </div>
      </router-link>
      <!-- /.col -->
      <router-link to="/cves">
      <div class="col-md-4 col-sm-6 col-xs-12">
        <info-box color-class="bg-red"
                  title="CVE"
                  text="CVES"
                  :number="dash.cveNumber"></info-box>
      </div>
      </router-link>
      <!-- /.col -->
      <!-- fix for small devices only -->
      <div class="clearfix visible-sm-block"></div>
      <router-link to="/cpes">
      <div class="col-md-4 col-sm-6 col-xs-12">
        <info-box color-class="bg-green"
                  title="CPE"
                  text="CPES"
                  :number="dash.cpeNumber"></info-box>
      </div>
      </router-link>
      <!-- /.col -->
      <!-- <router-link to="/users">
      <div class="col-md-3 col-sm-6 col-xs-12">
        <info-box color-class="bg-yellow"
                  :icon-classes="['ion', 'ion-ios-people-outline']"
                  text="New Members"
                  :number="userNumbers"></info-box>
      </div>
      </router-link> -->
      <!-- /.col -->
    </div>
    <!-- /.row -->

    <div class="col-xs-12">
      <div class="box">
        <div class="box-header with-border">
          <h3 class="box-title"></h3>
          <div class="box-body">
            <div class="col-sm-6 col-xs-12">
              <p class="text-center">
                <strong>NVTs by Severity Class (Total: {{nvttotal}})</strong>
              </p>
              <canvas id="nvtbyseverity"></canvas>
            </div>
            <hr class="visible-xs-block">
            <div class="col-sm-6 col-xs-12">
              <p class="text-center">
                <strong>CVEs by Severity Class (Total: {{cvetotal}})</strong>
              </p>
              <canvas id="cvebyseverity"></canvas>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- /.row -->

    <!-- Main row -->
    <!-- /.row -->
  </section>
  <!-- /.content -->
</template>

<script>
import Chart from 'chart.js'
import Alert from '../widgets/Alert'
import InfoBox from '../widgets/InfoBox'
import ProcessInfoBox from '../widgets/ProcessInfoBox'
import axios from 'axios'
export default {
  name: 'Dashboard',
  components: {
    Alert,
    InfoBox,
    ProcessInfoBox
  },
  data () {
    return {
      users: [],
      dashboards: [],
      cvetotal: 0,
      nvttotal: 0,
      // high: '',
      // medium: '',
      // low: '',
      generateRandomNumbers (numbers, max, min) {
        var a = []
        for (var i = 0; i < numbers; i++) {
          a.push(Math.floor(Math.random() * (max - min + 1)) + max)
        }
        return a
      }
    }
  },
  created() {
    axios
    .get('http://localhost:8081/dashboard')
    .then(response => {
      // JSON responses are automatically parsed.
      this.dashboards = response.data
      // return this.dashboards
    })
  },
  beforeMount() {
    axios
    .get('http://localhost:8081/nvtbyseverity')
    .then(response => {
      this.nvttotal = response.data.total
      this.$nextTick(() => {
        var pieChartCanvas = document.getElementById('nvtbyseverity').getContext('2d')
        var pieConfig = {
          type: 'pie',
          data: {
            labels: ['High', 'Medium', 'Low', 'Log', 'N/A'],
            datasets: [{
              data: [response.data.high.Int64, response.data.medium.Int64, response.data.low.Int64, response.data.log.Int64, response.data.na.Int64],
              backgroundColor: ['rgb(216, 0, 0)', 'rgb(255, 165, 0)', 'rgb(135, 206, 235)', 'rgb(221, 221, 221)', 'rgb(192, 192, 192)'],
              hoverBackgroundColor: ['rgb(216, 0, 0)', 'rgb(255, 165, 0)', 'rgb(135, 206, 235)', 'rgb(221, 221, 221)', 'rgb(192, 192, 192)']
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: !this.isMobile,
            legend: {
              position: 'bottom',
              display: true
            }
          }
        }
        new Chart(pieChartCanvas, pieConfig) // eslint-disable-line no-new
      })
    })
  },
  computed: {
    coPilotNumbers () {
      return this.generateRandomNumbers(12, 1000000, 10000)
    },
    personalNumbers () {
      return this.generateRandomNumbers(12, 1000000, 10000)
    },
    isMobile () {
      return (window.innerWidth <= 800 && window.innerHeight <= 600)
    },
    userNumbers () {
      return this.users
    }
  },
  mounted () {
    axios
    .get('http://localhost:8081/cvebyseverity')
    .then(response => {
      this.cvetotal = response.data.total
      this.$nextTick(() => {
        // // var ctx = document.getElementById('trafficBar').getContext('2d')
        // // var config = {
        // //   type: 'line',
        // //   data: {
        // //     labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
        // //     datasets: [{
        // //       label: 'CoPilot',
        // //       fill: false,
        // //       borderColor: '#284184',
        // //       pointBackgroundColor: '#284184',
        // //       backgroundColor: 'rgba(0, 0, 0, 0)',
        // //       data: this.coPilotNumbers
        // //     }, {
        // //       label: 'Personal Site',
        // //       borderColor: '#4BC0C0',
        // //       pointBackgroundColor: '#4BC0C0',
        // //       backgroundColor: 'rgba(0, 0, 0, 0)',
        // //       data: this.personalNumbers
        // //     }]
        // //   },
        // //   options: {
        // //     responsive: true,
        // //     maintainAspectRatio: !this.isMobile,
        // //     legend: {
        // //       position: 'bottom',
        // //       display: true
        // //     },
        // //     tooltips: {
        // //       mode: 'label',
        // //       xPadding: 10,
        // //       yPadding: 10,
        // //       bodySpacing: 10
        // //     }
        // //   }
        // // }

        // new Chart(ctx, config) // eslint-disable-line no-new

        var pieChartCanvas = document.getElementById('cvebyseverity').getContext('2d')
        var pieConfig = {
          type: 'pie',
          data: {
            labels: ['High', 'Medium', 'Low', 'Log', 'N/A'],
            datasets: [{
              data: [response.data.high.Int64, response.data.medium.Int64, response.data.low.Int64, response.data.log.Int64, response.data.na.Int64],
              backgroundColor: ['rgb(216, 0, 0)', 'rgb(255, 165, 0)', 'rgb(135, 206, 235)', 'rgb(221, 221, 221)', 'rgb(192, 192, 192)'],
              hoverBackgroundColor: ['rgb(216, 0, 0)', 'rgb(255, 165, 0)', 'rgb(135, 206, 235)', 'rgb(221, 221, 221)', 'rgb(192, 192, 192)']
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: !this.isMobile,
            legend: {
              position: 'bottom',
              display: true
            }
          }
        }

        new Chart(pieChartCanvas, pieConfig) // eslint-disable-line no-new
      })
    })
  }
}
</script>
<style>
.info-box {
  cursor: pointer;
}
.info-box-content {
  text-align: center;
  vertical-align: middle;
  display: inherit;
}
.info-box-number {
  font-size: 30px;
}
.info-box-title{
  margin-top: .67em;
  font-weight: 600;
}
.fullCanvas {
  width: 100%;
}
</style>
