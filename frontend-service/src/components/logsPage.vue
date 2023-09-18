<script setup lang="ts">
import { ref, onMounted } from 'vue';
const logs = ref<any[]>([]);
const getData = async () => {
  const data = await fetch("http://localhost:8080/logs")
  const jsonData = await data.json();
  console.log(jsonData)
  const dataArray = jsonData.logs.map((data: any) => {
    const split = data.split(',').map((item: any) => item.trim());
    const status = split[0].replace(/"/g, '');
    const message = split.slice(1).join(', ').replace(/"/g, '');
    console.log(status, message)
    return {
      status: status,
      message: message
    }
  })

  logs.value = dataArray
}

const clearLog = async () => {
  const options = {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  };
  await fetch("http://localhost:8080/logs", options)

  logs.value = []
}

onMounted(() => {
  getData();
})
</script>

<template>
  <div class="d-flex justify-content-end">
    <button class="btn btn-danger" @click.prevent="clearLog" type="button">Clear Log</button>
  </div>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">LogMessage</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(log, index) in logs" :key="index" :class="log.status === 'success' ? 'text-success' : 'text-danger'">
        <th scope="row">{{ index }}</th>
        <td>{{ log.message }}</td>
      </tr>
    </tbody>
  </table>
</template>