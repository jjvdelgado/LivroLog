<template>
  <div class="login-page">
    <!-- Main Content -->
    <main class="main-content">
      <!-- Login Form Component -->
      <LoginForms />
    </main>
  </div>
</template>

<script>
import LoginForms from '@/components/LoginForms.vue'

export default {
  name: 'LoginPage',
  components: {
    LoginForms
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.login-page {
  min-height: 100vh;
  background-color: #1a1a1a; /* Fundo escuro */
  margin: 0;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.main-content {
  height: 100vh; /* Preenche toda a tela */
  width: 100%;
  background-color: #1a1a1a; /* Fundo escuro que contrasta com o header #2c3e50 */
  margin: 0;
  padding-left: 0;
  padding-right: 0;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  overflow-x: hidden;
  overflow-y: auto;
}
</style>