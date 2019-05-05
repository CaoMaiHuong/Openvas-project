<template>
  <div class="info-box">
    <span :class="['info-box-icon', colorClass]">
      <h1 class="info-box-title">{{title}}</h1>
    </span>
    <div class="info-box-content">
      <h1 class="info-box-number">{{number}}</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InfoBox',
  props: {
    text: {
      type: String,
      required: true
    },
    number: {
      type: String,
      default: ''
    },
    title: {
      type: String,
      default: ''
    },
    colorClass: {
      type: String,
      default: 'bg-default'
    }
  }
}
</script>
