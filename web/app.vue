<template>
  <div class="bg-color min-h-screen">
    <Navigation />
    <ControlBar />
  </div>
</template>

<script>
import Navigation from "~/components/Navigation.vue"
import ControlBar from "~/components/ControlBar.vue";

export default {
  components: {
    Navigation: Navigation,
    ControlBar: ControlBar// Register the Navigation component
  },
};
</script>
