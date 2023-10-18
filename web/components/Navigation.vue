<!-- components/Navigation.vue -->
<template>
  <nav class="base-color h-[80px] place-content-center justify-center flex absolute w-full shadow-lg p-2">
    <img class="mr-auto mx-1" src="../assets/livealLogoWithText.svg" width="160">
    <ul class="mx-1 gap-3 place-items-center flex">
      <li class="">
        <a href="/">
          <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48">
            <path
                d="M11 39h7.5V26.5h11V39H37V19.5L24 9.75 11 19.5Zm-3 3V18L24 6l16 12v24H26.5V29.5h-5V42Zm16-17.65Z"></path>
          </svg>
        </a>
      </li>
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48">
          <path
              d="M39.8 41.95 26.65 28.8q-1.5 1.3-3.5 2.025-2 .725-4.25.725-5.4 0-9.15-3.75T6 18.75q0-5.3 3.75-9.05 3.75-3.75 9.1-3.75 5.3 0 9.025 3.75 3.725 3.75 3.725 9.05 0 2.15-.7 4.15-.7 2-2.1 3.75L42 39.75Zm-20.95-13.4q4.05 0 6.9-2.875Q28.6 22.8 28.6 18.75t-2.85-6.925Q22.9 8.95 18.85 8.95q-4.1 0-6.975 2.875T9 18.75q0 4.05 2.875 6.925t6.975 2.875Z"></path>
        </svg>
      </li>
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 96 960 960">
          <path
              d="M180 936q-24 0-42-18t-18-42V276q0-24 18-42t42-18h291v60H180v600h291v60H180Zm486-185-43-43 102-102H375v-60h348L621 444l43-43 176 176-174 174Z"></path>
        </svg>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  // Component logic if needed
};
</script>

<style scoped>
/* CSS styles specific to this component */
</style>