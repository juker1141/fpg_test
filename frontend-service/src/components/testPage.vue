<script setup lang="ts">
import { ref } from "vue";

const response = ref();
const apiUrl = ref("http://192.168.88.90:8081/test/get?a=11&b=22");
const username = ref("");
const isInit = ref(true);
const isSuccess = ref(false);

const loggerData = async (status: string, data: string) => {
  const jsonData = `"${status}", "${data}"`
  const requestData = {
    log: jsonData,
  };
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestData)
  };

  const logData = await fetch("http://localhost:8080/log", options);
  const logJsonData = await logData.json()
  console.log(logJsonData)


  const getLog = await fetch("http://localhost:8080/logs");
  const getLogJson = await getLog.json();
  console.log(getLogJson)
}

const fetchData = async () => {
  try {
    if (apiUrl.value === "") {
      const err = new Error("API Address can not be empty.")
      throw (err)
    }

    const result = await fetch(apiUrl.value);

    const logMessage = `HTTP log, User = ${username.value}, URL = ${result.url}, Status = ${result.status}, StatusText = ${result.statusText}`;
    if (!result.ok) {
      const errorData = await result.json();
      const err = JSON.stringify(errorData);
      throw new Error(`${logMessage}, Err = ${err}`);
    }

    const data = await result.json();

    response.value = data;
    isSuccess.value = true;

    loggerData("success", logMessage)
  } catch (err: any) {
    loggerData("error", err)
    response.value = err;
    isSuccess.value = false;
  } finally {
    isInit.value = false;
  }
}

</script>

<template>
  <form @submit.prevent="fetchData">
    <div class="mb-3">
      <label class="form-label">Username</label>
      <input type="text" class="form-control mb-3" placeholder="User Nasme" aria-label="User name" v-model="username">
      <label class="form-label">API address</label>
      <input type="text" class="form-control" placeholder="Test API Address" aria-label="API Address" v-model="apiUrl">
    </div>
    <div class="d-flex justify-content-center mb-3">
      <button class="btn btn-primary" type="submit" id="submit-button">Submit</button>
    </div>
    <div class="border border-2 px-3 py-4 rounded-3 position-relative" :class="isInit ? '' : isSuccess
      ? 'border-success text-success' : 'border-danger text-danger'">
      <div class="p-2 rounded-circle position-absolute top-0 start-0 translate-middle shadow-sm" :class="isInit ? '' : isSuccess
        ? 'bg-success' : 'bg-danger'"></div>
      {{ response }}
    </div>
  </form>
</template>