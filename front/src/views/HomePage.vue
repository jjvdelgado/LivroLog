<template>
  <div class="home-page">
    <!-- Header Component -->
    <AppHeader />

    <!-- Main Content -->
    <main class="main-content">
      <!-- Barra de Pesquisa Centralizada -->
      <div class="search-container">
        <div class="search-bar">
          <input type="text" placeholder="Pesquisar livros..." class="search-input" v-model="searchQuery"
            @input="handleSearch" />
          <button class="search-button" @click="performSearch">
            <svg class="search-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="m21 21-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </button>
        </div>
        
        <!-- Botão Adicionar Livro -->
        <div class="add-book-section" v-if="showAddBookButton">
          <button 
            @click="navigateToAddBook" 
            class="add-book-btn"
          >
            <span class="btn-icon">+</span>
            Adicionar Livro
          </button>
        </div>
      </div>

      <!-- Grid de Livros Component -->
      <GridLivros :search-query="searchQuery" @book-selected="handleBookSelection" />
    </main>
  </div>
</template>

<script>
import AppHeader from '@/components/AppHeader.vue'
import GridLivros from '@/components/GridLivros.vue'

import { useRequest } from '@/composables/useRequest'

const { request } = useRequest()

export default {
  name: 'HomePage',
  components: {
    AppHeader,
    GridLivros
  },
  data() {
    return {
      searchQuery: '', // nome do livro
      livros: [],
      erro: null,
      // Futuramente será baseado na condição de admin logado
      showAddBookButton: true // Por enquanto sempre visível
    }
  },
  methods: {
    handleSearch() {
      this.performSearch()
    },
    async performSearch() {
      this.erro = null
      try {
        const url = `http://localhost:8080/livros/pesquisa?titulo=${encodeURIComponent(this.searchQuery)}`
        const resultado = await request(url, 'GET')
        this.livros = resultado
        console.log('Livros encontrados:', resultado)
      } catch (e) {
        this.erro = e.message
        console.error('Erro ao buscar livros:', e)
      }
    },
    async fetchLivrosIniciais() {
      this.erro = null
      try {
        const resultado = await request(`http://localhost:8080/livros?limit=20`, 'GET')
        this.livros = resultado
        console.log('Livros iniciais:', resultado)
      } catch (e) {
        this.erro = e.message
        console.error('Erro ao buscar livros iniciais:', e)
      }
    },
    handleBookSelection(book) {
      console.log('Livro selecionado na HomePage:', book.title)
    },
    navigateToAddBook() {
      this.$router.push('/addnewbook')
    }
  },
  mounted() {
    this.fetchLivrosIniciais()
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.home-page {
  min-height: 100vh;
  background-color: #1a1a1a;
  /* Fundo escuro */
  margin: 0;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.main-content {
  padding-top: 64px;
  /* Altura do header fixo */
  height: calc(100vh - 64px);
  /* Preenche o resto da tela */
  width: 100%;
  background-color: #1a1a1a;
  /* Fundo escuro que contrasta com o header #2c3e50 */
  margin: 0;
  padding-left: 0;
  padding-right: 0;
  position: absolute;
  left: 0;
  right: 0;
  top: 64px;
  overflow-x: hidden;
  overflow-y: auto;
  /* Permite scroll vertical para o grid */
}

/* Container da barra de pesquisa */
.search-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  gap: 20px;
  padding: 40px 20px;
  width: 100%;
}

/* Barra de pesquisa */
.search-bar {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 600px;
  background-color: #2a2a2a;
  border-radius: 25px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease;
}

.search-bar:hover {
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.4);
}

.search-bar:focus-within {
  box-shadow: 0 6px 16px rgba(44, 62, 80, 0.3);
  border: 2px solid #2c3e50;
}

/* Input de pesquisa */
.search-input {
  flex: 1;
  padding: 16px 24px;
  background: transparent;
  border: none;
  outline: none;
  color: #ffffff;
  font-size: 16px;
  font-weight: 400;
  border-radius: 25px;
}

.search-input::placeholder {
  color: #888;
  font-style: italic;
}

/* Botão de pesquisa */
.search-button {
  padding: 12px 16px;
  background-color: #2c3e50;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  margin-right: 8px;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-button:hover {
  background-color: #34495e;
  transform: scale(1.05);
}

.search-button:active {
  transform: scale(0.95);
}

/* Ícone de pesquisa */
.search-icon {
  width: 20px;
  height: 20px;
  color: #ffffff;
}

/* Seção de adicionar livro */
.add-book-section {
  display: flex;
  align-items: center;
}

.add-book-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background: linear-gradient(135deg, #2c3e50, #3498db);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(52, 152, 219, 0.3);
  white-space: nowrap;
}

.add-book-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(52, 152, 219, 0.4);
}

.btn-icon {
  font-size: 1.2rem;
  font-weight: bold;
}

/* Responsividade */
@media (max-width: 768px) {
  .search-container {
    padding: 30px 15px;
    flex-direction: column;
    align-items: center;
    gap: 15px;
  }

  .search-bar {
    max-width: 100%;
  }

  .search-input {
    padding: 14px 20px;
    font-size: 14px;
  }

  .search-button {
    padding: 10px 14px;
  }

  .search-icon {
    width: 18px;
    height: 18px;
  }

  .add-book-btn {
    padding: 10px 20px;
    font-size: 0.9rem;
  }
}

@media (max-width: 480px) {
  .search-container {
    padding: 20px 10px;
    gap: 12px;
  }

  .search-input {
    padding: 12px 18px;
    font-size: 14px;
  }

  .search-button {
    padding: 8px 12px;
  }

  .search-icon {
    width: 16px;
    height: 16px;
  }

  .add-book-btn {
    padding: 8px 16px;
    font-size: 0.8rem;
  }

  .btn-icon {
    font-size: 1rem;
  }
}
</style>