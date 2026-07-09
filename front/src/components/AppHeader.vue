<template>
  <header class="header">
    <div class="header-container">
      <!-- Botão do Menu Lateral (Esquerda) -->
      <button class="menu-button" @click="toggleSidebar" aria-label="Abrir menu">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <line x1="3" y1="6" x2="21" y2="6"></line>
          <line x1="3" y1="12" x2="21" y2="12"></line>
          <line x1="3" y1="18" x2="21" y2="18"></line>
        </svg>
      </button>

      <!-- Logo/Nome da Aplicação (Centro) -->
      <h1 class="logo">LivroLog</h1>

      <!-- Botão do Perfil (Direita) -->
      <button class="profile-button" @click="goToProfile" aria-label="Acessar perfil">
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
          <circle cx="12" cy="7" r="4"></circle>
        </svg>
      </button>
    </div>

    <!-- Menu Lateral -->
    <div class="sidebar-overlay" :class="{ active: isSidebarOpen }" @click="closeSidebar">
      <nav class="sidebar" :class="{ open: isSidebarOpen }" @click.stop>
        <div class="sidebar-header">
          <h2>Menu</h2>
          <button class="close-button" @click="closeSidebar" aria-label="Fechar menu">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
        </div>

        <ul class="sidebar-menu">
          <li>
            <a href="#" @click="closeSidebar">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
                <polyline points="9,22 9,12 15,12 15,22"></polyline>
              </svg>
              Início
            </a>
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script>
// Importar o router
import { useRouter } from 'vue-router'

export default {
  name: 'AppHeader',
  data() {
    return {
      isSidebarOpen: false
    }
  },
  setup() {
    // Para Vue 3 com Composition API
    const router = useRouter()
    return { router }
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen
    },
    closeSidebar() {
      this.$router.push('/homepage')
      this.isSidebarOpen = false
    },
    goToProfile() {
      // Usando o router importado
      this.$router.push('/profile')
    }
  },
  mounted() {
    // Fecha o sidebar ao pressionar ESC
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && this.isSidebarOpen) {
        this.closeSidebar()
      }
    })
  }
}
</script>

<style scoped>
.header {
  background-color: #2c3e50;
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  height: 64px;
}

.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 16px;
  height: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.menu-button,
.profile-button {
  background: none;
  border: none;
  color: white;
  cursor: pointer;
  padding: 8px;
  border-radius: 6px;
  transition: background-color 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.menu-button:hover,
.profile-button:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.logo {
  font-size: 24px;
  font-weight: bold;
  margin: 0;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  color: white;
}

/* Sidebar Styles */
.sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 999;
  opacity: 0;
  visibility: hidden;
  transition: opacity 0.3s ease, visibility 0.3s ease;
}

.sidebar-overlay.active {
  opacity: 1;
  visibility: visible;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  width: 280px;
  height: 100%;
  background-color: #2c3e50;
  color: white;
  transform: translateX(-100%);
  transition: transform 0.3s ease;
  box-shadow: 2px 0 8px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
}

.sidebar.open {
  transform: translateX(0);
}

.sidebar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background-color: #34495e;
}

.sidebar-header h2 {
  margin: 0;
  font-size: 18px;
  color: white;
}

.close-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  border-radius: 4px;
  color: white;
  transition: background-color 0.2s ease;
}

.close-button:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.sidebar-menu {
  list-style: none;
  padding: 0;
  margin: 0;
}

.sidebar-menu li {
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.sidebar-menu li:last-child {
  border-bottom: none;
}

.sidebar-menu a {
  display: flex;
  align-items: center;
  padding: 16px 20px;
  color: white;
  text-decoration: none;
  transition: background-color 0.2s ease;
  gap: 12px;
}

.sidebar-menu a:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

/* Responsive */
@media (max-width: 768px) {
  .header-container {
    padding: 0 12px;
  }

  .logo {
    font-size: 20px;
  }

  .sidebar {
    width: 100%;
    max-width: 320px;
  }
}
</style>