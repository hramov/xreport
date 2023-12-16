<script setup lang="ts">
import {ref} from "vue";

const emit = defineEmits(['submit'])
const model = ref({});

const handleReset = () => {
  model.value = {};
}
</script>

<template>
  <form @submit.prevent="emit('submit', model)">
    <v-text-field
        v-model="model.title"
        label="Название"
    ></v-text-field>

    <h3>Почтовый сервер</h3>
    <br/>

    <v-text-field
        v-model="model.smtp_ip"
        label="IP-адрес"
    ></v-text-field>

    <v-text-field
        v-model="model.smtp_port"
        type="number"
        min="0"
        label="Порт"
    ></v-text-field>

    <v-text-field
        v-model="model.smtp_email"
        label="Email"
        type="email"
    ></v-text-field>

    <v-text-field
        v-model="model.smtp_password"
        label="Пароль"
        type="password"
    ></v-text-field>

    <v-btn
        class="me-4"
        type="submit"
    >
      Сохранить
    </v-btn>

    <v-btn @click="handleReset">
      Очистить
    </v-btn>
  </form>
</template>