<script setup lang="ts">
import {ref} from "vue";

const drawer = ref(false)

const items = [
  {title: 'Сервис', path:'/create-service' },
  {title: 'Отчет', path:'/create-report'},
  {title: 'Источник данных', path:'/create-data-source'},
]
</script>

<template>
  <v-app id="inspire">
    <v-app-bar>
<!--      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>-->

      <v-app-bar-title><router-link to="/" style="text-decoration: none; color: white;">Генератор отчетов</router-link></v-app-bar-title>

      <v-spacer></v-spacer>

      <router-link to="/audit" style="text-decoration: none; color: white;">
        <v-list-item>
          <v-list-item-title>Аудит</v-list-item-title>
        </v-list-item>
      </router-link>

      <v-btn>Создать
        <v-menu activator="parent">
          <v-list>
            <router-link v-for="(item, index) in items" :to="item.path" :key="index" style="text-decoration: none; color: white;">
              <v-list-item
                  :value="index"
              >
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item>
            </router-link>
          </v-list>
        </v-menu>
      </v-btn>

    </v-app-bar>

    <v-navigation-drawer
        v-model="drawer"
        temporary
    >
    </v-navigation-drawer>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<style scoped>
</style>
