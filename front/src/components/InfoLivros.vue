<template>
  <div class="info-livros">
    <!-- Imagem do Livro -->
    <div class="book-image-container">
      <img 
        :src="book.image || '/images/book-placeholder.jpg'" 
        :alt="book.title"
        class="book-image"
      />
    </div>

    <!-- Informações do Livro -->
    <div class="book-info">
      <h1 class="book-title">{{ book.title }}</h1>
      
      <!-- Avaliação com Estrelas -->
      <div class="rating-section">
        <div class="stars-container">
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
          <span class="rating-text">{{ book.rating.toFixed(1) }}</span>
          <span class="reviews-count">({{ book.reviewsCount }} reviews)</span>
        </div>
      </div>

      <!-- Informações Básicas -->
      <div class="basic-info">
        <div class="info-item">
          <span class="info-label">Autor:</span>
          <span class="info-value">{{ book.author }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">Ano de Lançamento:</span>
          <span class="info-value">{{ book.year }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">Editora:</span>
          <span class="info-value">{{ book.publisher }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">Gênero:</span>
          <span class="info-value">{{ book.genre }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">Páginas:</span>
          <span class="info-value">{{ book.pages }}</span>
        </div>
      </div>

      <!-- Sinopse -->
      <div class="synopsis-section">
        <h3 class="synopsis-title">Sinopse</h3>
        <p class="synopsis-text">{{ book.synopsis }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InfoLivros',
  props: {
    book: {
      type: Object,
      required: true,
      default: () => ({
        title: 'Título do Livro',
        author: 'Nome do Autor',
        year: 2024,
        publisher: 'Editora',
        genre: 'Ficção',
        pages: 350,
        rating: 4.5,
        reviewsCount: 128,
        image: '/images/book-placeholder.jpg',
        synopsis: 'Esta é uma sinopse de exemplo do livro. Aqui você encontrará um resumo envolvente da história, personagens principais e temas abordados na obra.'
      })
    }
  }
}
</script>

<style scoped>
.info-livros {
  display: flex;
  gap: 40px;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  background-color: #2a2a2a;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

/* Container da Imagem */
.book-image-container {
  flex-shrink: 0;
}

.book-image {
  width: 200px;
  height: 300px;
  object-fit: cover;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.4);
}

/* Informações do Livro */
.book-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.book-title {
  font-size: 2.5rem;
  font-weight: bold;
  color: #ffffff;
  margin: 0;
  line-height: 1.2;
}

/* Seção de Avaliação */
.rating-section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.stars-container {
  display: flex;
  align-items: center;
  gap: 12px;
}

.stars {
  display: flex;
  gap: 2px;
}

.star {
  font-size: 1.5rem;
  color: #444;
  transition: color 0.2s ease;
}

.star.filled {
  color: #ffd700;
}

.rating-text {
  font-size: 1.1rem;
  font-weight: bold;
  color: #ffffff;
}

.reviews-count {
  font-size: 0.9rem;
  color: #888;
}

/* Informações Básicas */
.basic-info {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.info-label {
  font-weight: 600;
  color: #ffffff;
  min-width: 140px;
}

.info-value {
  color: #ffffff;
  font-weight: 400;
}

/* Seção da Sinopse */
.synopsis-section {
  margin-top: 8px;
}

.synopsis-title {
  font-size: 1.3rem;
  font-weight: bold;
  color: #ffffff;
  margin-bottom: 12px;
}

.synopsis-text {
  color: #cccccc;
  line-height: 1.6;
  font-size: 1rem;
  text-align: justify;
}

/* Responsividade */
@media (max-width: 768px) {
  .info-livros {
    flex-direction: column;
    gap: 24px;
    padding: 16px;
  }

  .book-image-container {
    align-self: center;
  }

  .book-image {
    width: 160px;
    height: 240px;
  }

  .book-title {
    font-size: 2rem;
    text-align: center;
  }

  .rating-section {
    align-items: center;
  }

  .info-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }

  .info-label {
    min-width: auto;
  }
}

@media (max-width: 480px) {
  .info-livros {
    padding: 12px;
  }

  .book-image {
    width: 140px;
    height: 210px;
  }

  .book-title {
    font-size: 1.8rem;
  }

  .stars-container {
    flex-direction: column;
    gap: 8px;
  }
}
</style>