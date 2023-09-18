<script setup lang="ts">
import { ref } from "vue";
import testPage from "./components/testPage.vue"
import logsPage from "./components/logsPage.vue"

type PageStatus = "test" | "log";
const pageStatus = ref<PageStatus>("test");

function switchPage(target: PageStatus) {
  pageStatus.value = target
}
</script>

<template>
  <div class="container mt-24 ">
    <div class="row align-items-center">
      <div class="col col-lg-6 offset-lg-3">
        <h2 class="text-center mb-3">iFPG 台塑測試 API</h2>
        <ul class="nav nav-tabs mb-3">
          <li class="nav-item">
            <a @click.prevent="switchPage('test')" :class="pageStatus === 'test' && 'active'" class="nav-link"
              href="#">測試頁</a>
          </li>
          <li class="nav-item">
            <a @click.prevent="switchPage('log')" :class="pageStatus === 'log' && 'active'" class="nav-link"
              href="#">Log</a>
          </li>
        </ul>
        <testPage v-if="pageStatus === 'test'" />
        <logsPage v-if="pageStatus === 'log'" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.mt-24 {
  margin-top: 10rem;
}
</style>
