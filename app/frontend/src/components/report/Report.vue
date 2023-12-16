<script setup lang="ts">
import {ref} from "vue";
import DataSettings from "./components/DataSettings.vue";
import ViewSettings from "./components/ViewSettings.vue";

const tab = ref('data')
</script>

<template>
  <v-container>
    <div class="report-title">
      <h2>Длинное название отчета</h2>
      <v-tabs
          v-model="tab"
          fixed-tabs
      >
        <v-tab value="report">
          Отчет
        </v-tab>
        <v-tab value="data">
          Данные
        </v-tab>
        <v-tab value="view">
          Внешний вид
        </v-tab>
      </v-tabs>
    </div>

    <v-window v-model="tab">
      <v-window-item value="data">
        <DataSettings />
      </v-window-item>
      <v-window-item value="view">
        <ViewSettings />
      </v-window-item>
    </v-window>

  </v-container>
</template>

<style scoped>
.report-title {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
</style>