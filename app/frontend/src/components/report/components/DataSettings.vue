<script setup lang="ts">

const getData = () => {

}

const onSubmit = () => {

}

const handleReset = () => {

}
</script>

<template>
  <div style="width: 50%; margin: 0 auto">
    <v-form @submit.prevent="onSubmit">
      <v-select label="Источник" :items="[]" item-title="title" item-value="driver"/>
      <v-textarea label="Запрос" rows="8" ></v-textarea>
      <v-btn @click="getData" class="mb-6">Выполнить</v-btn>

      <v-table
          height="250px"
          fixed-header
      >
        <thead>
        <tr>
          <th class="text-center" v-for="h in 4" :key="h">
            {{ h }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="item in 4"
            :key="item"
        >
          <td class="text-center" v-for="col in 4" :key="col as string">{{ col }}</td>
        </tr>
        </tbody>
      </v-table>
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
    </v-form>
  </div>
</template>