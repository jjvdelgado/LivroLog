<template>
  <div class="books-section">
    <!-- Mensagem quando nenhum livro é encontrado -->
    <div v-if="filteredBooks.length === 0 && searchQuery" class="no-results">
      <p>Nenhum livro encontrado para "{{ searchQuery }}"</p>
    </div>
    
    <!-- Grid de livros -->
    <div class="books-grid" v-else>
      <div 
        v-for="book in filteredBooks" 
        :key="book.id" 
        class="book-card"
        @click="selectBook(book)"
      >
        <div class="book-cover">
          <img :src="book.cover" :alt="book.title" />
        </div>
        <div class="book-info">
          <h3 class="book-title">{{ book.title }}</h3>
          <p class="book-author">{{ book.author }}</p>
          <div class="book-rating">
            <div class="stars">
              <span 
                v-for="star in 5" 
                :key="star"
                class="star"
                :class="{ 'filled': star <= book.rating }"
              >
                ★
              </span>
            </div>
            <span class="rating-text">{{ book.rating }}/5</span>
          </div>
          <p class="book-price">{{ book.price }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GridLivros',
  props: {
    searchQuery: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      books: [
        {
          id: 1,
          title: '1984',
          author: 'George Orwell',
          rating: 5,
          price: 'R$ 29,90',
          cover: 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=300&h=400&fit=crop'
        },
        {
          id: 2,
          title: 'Dom Casmurro',
          author: 'Machado de Assis',
          rating: 4,
          price: 'R$ 25,50',
          cover: 'https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=300&h=400&fit=crop'
        },
        {
          id: 3,
          title: 'O Pequeno Príncipe',
          author: 'Antoine de Saint-Exupéry',
          rating: 5,
          price: 'R$ 22,90',
          cover: 'https://images.unsplash.com/photo-1512820790803-83ca734da794?w=300&h=400&fit=crop'
        },
        {
          id: 4,
          title: 'Harry Potter e a Pedra Filosofal',
          author: 'J.K. Rowling',
          rating: 5,
          price: 'R$ 35,00',
          cover: 'https://images.unsplash.com/photo-1621351183012-e2f9972dd9bf?w=300&h=400&fit=crop'
        },
        {
          id: 5,
          title: 'O Cortiço',
          author: 'Aluísio Azevedo',
          rating: 3,
          price: 'R$ 28,00',
          cover: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=300&h=400&fit=crop'
        },
        {
          id: 6,
          title: 'A Revolução dos Bichos',
          author: 'George Orwell',
          rating: 4,
          price: 'R$ 26,90',
          cover: 'https://images.unsplash.com/photo-1589829085413-56de8ae18c73?w=300&h=400&fit=crop'
        },
        {
          id: 7,
          title: 'Cem Anos de Solidão',
          author: 'Gabriel García Márquez',
          rating: 5,
          price: 'R$ 42,00',
          cover: 'https://images.unsplash.com/photo-1592496431122-2349e0fbc666?w=300&h=400&fit=crop'
        },
        {
          id: 8,
          title: 'O Hobbit',
          author: 'J.R.R. Tolkien',
          rating: 4,
          price: 'R$ 38,50',
          cover: 'https://images.unsplash.com/photo-1516979187457-637abb4f9353?w=300&h=400&fit=crop'
        }
      ]
    }
  },
  computed: {
    filteredBooks() {
      if (!this.searchQuery) {
        return this.books
      }
      
      const query = this.searchQuery.toLowerCase()
      return this.books.filter(book => 
        book.title.toLowerCase().includes(query) ||
        book.author.toLowerCase().includes(query)
      )
    }
  },
  methods: {
    selectBook(book) {
      // Emitir evento para o componente pai
      this.$emit('book-selected', book)
      console.log('Livro selecionado:', book.title)
    }
  }
}
</script>

<style scoped>
/* Seção de Livros */
.books-section {
  padding: 40px 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.section-title {
  color: #ffffff;
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 30px;
  text-align: center;
}

/* Mensagem de nenhum resultado */
.no-results {
  text-align: center;
  padding: 40px 20px;
  color: #888;
  font-size: 18px;
}

/* Grid de Livros */
.books-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 25px;
  margin-top: 20px;
}

/* Card do Livro */
.book-card {
  background-color: #2a2a2a;
  border-radius: 15px;
  overflow: hidden;
  transition: all 0.3s ease;
  cursor: pointer;
  border: 2px solid transparent;
}

.book-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.4);
  border-color: #2c3e50;
}

/* Capa do Livro */
.book-cover {
  width: 100%;
  height: 280px;
  overflow: hidden;
  position: relative;
}

.book-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.book-card:hover .book-cover img {
  transform: scale(1.05);
}

/* Informações do Livro */
.book-info {
  padding: 20px;
}

.book-title {
  color: #ffffff;
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
  line-height: 1.3;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.book-author {
  color: #888;
  font-size: 14px;
  margin-bottom: 12px;
  font-style: italic;
}

/* Rating com Estrelas */
.book-rating {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
}

.stars {
  display: flex;
  gap: 2px;
}

.star {
  font-size: 16px;
  color: #444;
  transition: color 0.2s ease;
}

.star.filled {
  color: #ffd700;
}

.rating-text {
  color: #888;
  font-size: 12px;
  font-weight: 500;
}

/* Preço do Livro */
.book-price {
  color: #ffffff;
  font-size: 16px;
  font-weight: 700;
  margin-top: 8px;
}

/* Responsividade */
@media (max-width: 768px) {
  .books-grid {
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 20px;
  }
}

@media (max-width: 480px) {
  .books-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 15px;
  }
}
</style>