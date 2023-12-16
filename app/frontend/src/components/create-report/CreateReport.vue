<script setup lang="ts">

import Form from "./Form.vue";

const submit = () => {
  // TODO save service with backend
}
</script>

<template>
  <v-container>
    <h2>Создать новый отчет</h2>
    <br/>
    <Form @submit="submit" />
  </v-container>
</template>