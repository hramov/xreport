<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

const emit = defineEmits(["submit"]);

const model = ref({});

const drivers = ref([]);

const loading = ref(false);
const canCheckConnection = ref(true);

const loadDrivers = async () => {
  const response = await axios.get("http://localhost:3000/driver/");
  drivers.value = response.data.data;
};

const checkConnection = async () => {
  loading.value = true;
  const response = await axios.post(
    "http://localhost:3000/source/check",
    model.value
  );
  loading.value = false;
  console.log(response);
};

const submit = () => {
  emit("submit", model);
};

const handleReset = () => {
  model.value = {};
};

onMounted(() => {
  loadDrivers();
});
</script>

<template>
  <v-form style="width: 700px; margin: 0 auto">
    <v-select
      label="Драйвер *"
      v-model="model.driver"
      :items="drivers"
      item-title="title"
      item-value="code"
    />

    <v-text-field label="Название *" v-model="model.title" />

    <v-text-field label="Хост *" v-model="model.host" />

    <v-text-field label="Порт *" placeholder="5432" v-model="model.port" />

    <v-text-field label="Пользователь *" v-model="model.user" />

    <v-text-field label="Пароль *" v-model="model.password" />

    <v-text-field label="Название БД" v-model="model.database" />

    <v-btn
      variant="text"
      :loading="loading"
      :disabled="loading"
      @click="checkConnection"
      >Проверить подключение</v-btn
    >
    <span style="margin-left: 20px">Успешно</span>
    <br />
    <br />
    <v-btn class="me-4" type="submit"> Сохранить </v-btn>

    <v-btn @click="handleReset"> Очистить </v-btn>
  </v-form>
</template>
