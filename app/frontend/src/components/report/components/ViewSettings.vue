<script setup lang="ts">
import {ref} from "vue";

const jsondata = ref([
  {
    user: '1',
    name: 'John',
    phone: '123',
    gender: 'M',
    age: 20,
    birth: '2020-01-01'
  }
])
</script>

<template>
  <vue-excel-editor v-model="jsondata" style="color: black">
    <vue-excel-column field="user"   label="User ID"       type="string" />
    <vue-excel-column field="name"   label="Name"          type="string" />
    <vue-excel-column field="phone"  label="Contact"       type="string" />
    <vue-excel-column field="gender" label="Gender"        type="select" :options="['F','M','U']" />
    <vue-excel-column field="age"    label="Age"           type="number" />
    <vue-excel-column field="birth"  label="Date Of Birth" type="date" />
  </vue-excel-editor>
</template>