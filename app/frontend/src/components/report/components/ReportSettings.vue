<script setup lang="ts">
import {ref} from "vue";

const model = ref({
  recipients: [],
})

const handleReset = () => {

}

const addRecipient = () => {

}

const onSubmit = () => {

}
</script>

<template>
  <v-form @submit.prevent="onSubmit">
    <v-text-field
        v-model="model.title"
        label="Название"
    ></v-text-field>

    <v-text-field
        v-model="model.time"
        label="Периодичность отправки"
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

    <v-checkbox label="Приклыдывать файл Excel" v-model="model.is_excel"></v-checkbox>

    <v-btn
        class="me-4"
        type="submit"
    >
      Сохранить
    </v-btn>

    <v-btn @click="handleReset">
      Очистить
    </v-btn>
  </v-form>

</template>