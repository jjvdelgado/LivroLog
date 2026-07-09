<template>
  <div class="add-new-book-page">
    <!-- Header Component -->
    <AppHeader />
    <!-- Main Content -->
    <main class="main-content">
      <!-- Add Book Form Component -->
      <FormsAddBook />
    </main>
  </div>
</template>

<script>
import AppHeader from '@/components/AppHeader.vue'
import FormsAddBook from '@/components/FormsAddBook.vue'

export default {
  name: 'AddNewBookPage',
  components: {
    AppHeader,
    FormsAddBook
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.add-new-book-page {
  min-height: 100vh;
  background-color: #1a1a1a; /* Fundo escuro */
  margin: 0;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.main-content {
  min-height: 100vh; /* Mudado de height para min-height */
  width: 100%;
  background-color: #1a1a1a; /* Fundo escuro que contrasta com o header #2c3e50 */
  margin: 0;
  padding-left: 0;
  padding-right: 0;
  overflow-x: hidden;
  overflow-y: auto;
  /* Removido position: absolute e propriedades de posicionamento */
}
</style>