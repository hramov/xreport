<script setup lang="ts">
const props = defineProps(['dialog'])
</script>

<template>
  <div class="text-center">
    <v-dialog
        v-model="props.dialog"
        activator="parent"
        width="auto"
    >
      <v-card>
        <v-card-text>
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </v-card-text>
        <v-card-actions>
          <v-btn color="primary" block @click="props.dialog = false">Close Dialog</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>