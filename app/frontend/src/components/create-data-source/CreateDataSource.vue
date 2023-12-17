<script setup lang="ts">

import Form from "./Form.vue";

const submit = () => {
  // TODO save service with backend
}
</script>

<template>
  <v-container>
    <h2>Создать новый источник данных</h2>
    <br/>

    <v-banner
        lines="one"
        icon="$warning"
        color="warning"
    >
      <v-banner-text style="font-size: 18px">
        Учетная запись, используемая в источниках, должна иметь права ТОЛЬКО на чтение данных!
      </v-banner-text>
      <v-banner-actions></v-banner-actions>
    </v-banner>
    <br/>
    <Form @submit="submit" />
  </v-container>
</template>