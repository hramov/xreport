<script setup lang="ts">
import {ref} from "vue";
const content = ref('')

const handleReset = () => {

}

</script>

<template>
  <v-container>
    <div class="html-wrapper">
      <div class="html">
        <h3 style="text-align: center">Шаблон</h3>
        <v-textarea rows="20" style="padding: 20px" v-model="content"></v-textarea>
      </div>
      <div class="divider"></div>
      <div class="preview">
        <h3 style="text-align: center">Предпросмотр</h3>
        <div style="padding: 20px;" v-html="content"></div>
      </div>
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
  </v-container>
</template>

<style scoped>
.html-wrapper {
  display: flex;
  justify-content: space-between;
}

.html {
  width: 50%;
}

.preview {
  width: 50%;
}

.divider {
  width: 1px;
  background-color: white;
}
</style>