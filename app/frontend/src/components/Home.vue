<script setup lang="ts">

</script>

<template>
  <v-container>
    <h2>Сервисы</h2>
    <br/>
    <v-row>
      <template v-for="n in 1" :key="n">
        <v-col
            class="mt-2"
            cols="12"
        >
          <strong>Дашборд</strong>
        </v-col>

<!--        <v-col-->
<!--            v-for="j in 6"-->
<!--            :key="`${n}${j}`"-->
<!--            cols="6"-->
<!--            md="2"-->
<!--        >-->
<!--          <router-link :to="'/report/' + j"><v-sheet height="150"></v-sheet></router-link>-->
<!--        </v-col>-->
        <v-col>
          Нет отчетов
        </v-col>
      </template>
    </v-row>
  </v-container>
</template>