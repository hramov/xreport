<script setup lang="ts">
import ShowModal from "./components/ShowModal.vue";
import StatusModal from "./components/StatusModal.vue";

const items = [
  {
    'сервис': 'Дашборд',
    'отчет': 'Еженедельный отчет директора',
    'дата': new Date().toLocaleString('ru-RU'),
    'получатели': 'khramovsi@gvc.rzd,orelikovmg@gvc.rzd',
    'действия': 'показать | скачать | статус',
  },
]
</script>

<template>
  <v-container>
    <v-data-table :items="items">
      <template v-slot:top>
        <v-toolbar
            flat
        >
          <v-toolbar-title>Последние отправленные отчеты</v-toolbar-title>
        </v-toolbar>
      </template>
    </v-data-table>

    <ShowModal />
    <StatusModal />

  </v-container>
</template>