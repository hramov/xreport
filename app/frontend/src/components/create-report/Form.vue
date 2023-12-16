<script setup lang="ts">
import {ref} from "vue";

const emit = defineEmits(['submit'])
const model = ref({
  recipients: [],
});

const handleReset = () => {
  model.value = {
    recipients: [],
  };
}

const addRecipient = () => {
  model.value.recipients.push('')
}

const onSubmit = () => {
  model.value.recipients.filter((r) => r.length > 0)
  emit('submit', model)
}
</script>

<template>
  <form @submit.prevent="onSubmit">

    <v-select
      v-model="model.service"
      label="Сервис"
      :items="['Дашборд']"
    >
    </v-select>

    <v-text-field
        v-model="model.title"
        label="Название"
    ></v-text-field>

    <h3>Получатели</h3>
    <br/>
    <v-btn icon="mdi-plus" @click="addRecipient" style="margin-bottom: 20px"></v-btn>

    <div v-for="(_, index) in model.recipients" :key="index">
      <v-text-field
          v-model="model.recipients[index]"
          label="Email"
      ></v-text-field>
    </div>


    <br/>
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