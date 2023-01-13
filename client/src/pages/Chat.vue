<script setup lang="ts">
import * as socketIo from 'socket.io-client'

const socket = socketIo.connect("ws://localhost:9090")
let massage = ""


const handelSend = (e: Event) => {
  e.preventDefault();
    if (massage) {
      console.log("I am Sending Stuff")
      socket.emit('chat message', massage);
      massage = '';
    }
    console.log(" I Dose Not Do anyThing")
}

socket.on('chat message', (msg) => {
  console.log("Massage from server ", msg)
  });

</script>

<template>

<main>
  <form v-on:submit="handelSend" class="flex flex-col gap-2">
    <label for="send-massage-input">Send Massage</label>
    <input type="text" id="send-massage-input" v-model="massage" placeholder="...">
    <button type="submit">Send</button>
  </form>
</main>

</template>

<style scoped>

</style>
