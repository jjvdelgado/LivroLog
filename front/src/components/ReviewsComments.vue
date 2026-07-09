<template>
  <div class="reviews-comments">
    <div class="reviews-header">
      <h2 class="reviews-title">Reviews dos Usuários</h2>
      <span class="reviews-count">{{ reviews.length }} review{{ reviews.length !== 1 ? 's' : '' }}</span>
    </div>

    <!-- Botão para adicionar review -->
    <div class="add-review-section">
      <button 
        @click="toggleReviewForm" 
        class="add-review-btn"
        :class="{ 'active': showReviewForm }"
      >
        <span class="btn-icon">{{ showReviewForm ? '×' : '+' }}</span>
        {{ showReviewForm ? 'Cancelar' : 'Adicionar Review' }}
      </button>
    </div>

    <!-- Formulário de review -->
    <transition name="form-slide">
      <div v-if="showReviewForm" class="review-form-container">
        <form @submit.prevent="submitReview" class="review-form">
          <h3 class="form-title">Deixe sua avaliação</h3>
          
          <!-- Seleção de estrelas -->
          <div class="rating-input">
            <label class="input-label">Nota *</label>
            <div class="stars-input">
              <button
                type="button"
                v-for="star in 5"
                :key="star"
                @click="setRating(star)"
                @mouseover="hoverRating = star"
                @mouseleave="hoverRating = 0"
                class="star-btn"
                :class="{ 'filled': star <= (hoverRating || newReview.rating) }"
              >
                ★
              </button>
            </div>
            <span class="rating-text">{{ getRatingText(newReview.rating) }}</span>
          </div>

          <!-- Campo de comentário -->
          <div class="comment-input">
            <label class="input-label">Comentário (opcional)</label>
            <textarea
              v-model="newReview.comment"
              placeholder="Compartilhe sua opinião sobre o livro..."
              class="comment-textarea"
              rows="4"
              maxlength="500"
            ></textarea>
            <div class="char-count">
              {{ newReview.comment.length }}/500
            </div>
          </div>

          <!-- Botões do formulário -->
          <div class="form-actions">
            <button
              type="button"
              @click="toggleReviewForm"
              class="btn-secondary"
            >
              Cancelar
            </button>
            <button
              type="submit"
              class="btn-primary"
              :disabled="!newReview.rating || isSubmitting"
            >
              {{ isSubmitting ? 'Publicando...' : 'Publicar Review' }}
            </button>
          </div>
        </form>
      </div>
    </transition>

    <div class="reviews-list">
      <div 
        v-for="review in reviews" 
        :key="review.id"
        class="review-item"
      >
        <!-- Foto e Nome do Usuário -->
        <div class="user-info">
          <img 
            :src="review.user.avatar || '/images/default-avatar.jpg'" 
            :alt="review.user.name"
            class="user-avatar"
          />
          <div class="user-details">
            <h4 class="user-name">{{ review.user.name }}</h4>
            <span class="review-date">{{ formatDate(review.date) }}</span>
          </div>
          
          <!-- Botões de Ação -->
          <div class="review-actions">
            <!-- Botão de Editar - Apenas para comentários próprios -->
            <button 
              v-if="canEditReview(review)" 
              @click="editReview(review)"
              class="action-btn edit-btn"
              title="Editar review"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                <path d="m18.5 2.5 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
              </svg>
            </button>

            <!-- Botão de Excluir - Para admins ou comentários próprios -->
            <button 
              v-if="canDeleteReview(review)" 
              @click="deleteReview(review)"
              class="action-btn delete-btn"
              title="Excluir review"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M3 6h18"></path>
                <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"></path>
                <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"></path>
                <line x1="10" y1="11" x2="10" y2="17"></line>
                <line x1="14" y1="11" x2="14" y2="17"></line>
              </svg>
            </button>
          </div>
        </div>

        <!-- Avaliação com Estrelas -->
        <div class="rating-section">
          <div class="stars">
            <span 
              v-for="star in 5" 
              :key="star"
              class="star"
              :class="{ 'filled': star <= review.rating }"
            >
              ★
            </span>
          </div>
          <span class="rating-value">{{ review.rating }}</span>
        </div>

        <!-- Comentário (se existir) -->
        <div v-if="review.comment" class="comment-section">
          <p class="comment-text">{{ review.comment }}</p>
        </div>

        <!-- Indicador de review apenas com nota -->
        <div v-else class="rating-only">
          <span class="rating-only-text">Avaliação sem comentário</span>
        </div>
      </div>
    </div>

    <!-- Mensagem quando não há reviews -->
    <div v-if="reviews.length === 0" class="no-reviews">
      <p class="no-reviews-text">Ainda não há reviews para este livro.</p>
      <p class="no-reviews-subtitle">Seja o primeiro a avaliar!</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ReviewsComments',
  props: {
    reviews: {
      type: Array,
      required: true,
      default: () => []
    },
    // Props para controlar visibilidade dos botões
    isAdmin: {
      type: Boolean,
      default: false
    },
    currentUserId: {
      type: [String, Number],
      default: null
    }
  },
  data() {
    return {
      showReviewForm: false,
      isSubmitting: false,
      hoverRating: 0,
      newReview: {
        rating: 0,
        comment: ''
      }
    }
  },
  methods: {
    toggleReviewForm() {
      this.showReviewForm = !this.showReviewForm
      if (!this.showReviewForm) {
        this.resetForm()
      }
    },
    
    setRating(rating) {
      this.newReview.rating = rating
    },
    
    getRatingText(rating) {
      const texts = {
        1: 'Muito ruim',
        2: 'Ruim',
        3: 'Regular',
        4: 'Bom',
        5: 'Excelente'
      }
      return rating > 0 ? texts[rating] : 'Selecione uma nota'
    },
    
    async submitReview() {
      if (!this.newReview.rating) return
      
      this.isSubmitting = true
      
      try {
        // Simular dados do usuário (você deve pegar do sistema de autenticação)
        const reviewData = {
          id: Date.now(), // Em produção, isso viria do backend
          rating: this.newReview.rating,
          comment: this.newReview.comment.trim() || null,
          date: new Date().toISOString(),
          user: {
            id: this.currentUserId, // ID do usuário atual
            name: 'Usuário Atual', // Pegar do sistema de auth
            avatar: '/images/default-avatar.jpg' // Pegar do sistema de auth
          }
        }
        
        // Emitir evento para o componente pai
        this.$emit('review-submitted', reviewData)
        
        // Resetar formulário
        this.resetForm()
        this.showReviewForm = false
        
        // Feedback visual (opcional)
        this.$emit('show-message', {
          type: 'success',
          message: 'Review publicado com sucesso!'
        })
        
      } catch (error) {
        console.error('Erro ao enviar review:', error)
        this.$emit('show-message', {
          type: 'error',
          message: 'Erro ao publicar review. Tente novamente.'
        })
      } finally {
        this.isSubmitting = false
      }
    },
    
    resetForm() {
      this.newReview = {
        rating: 0,
        comment: ''
      }
      this.hoverRating = 0
    },
    
    formatDate(dateString) {
      const date = new Date(dateString)
      const now = new Date()
      const diffTime = Math.abs(now - date)
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
      
      if (diffDays === 1) {
        return 'Ontem'
      } else if (diffDays <= 7) {
        return `${diffDays} dias atrás`
      } else if (diffDays <= 30) {
        const weeks = Math.floor(diffDays / 7)
        return `${weeks} semana${weeks > 1 ? 's' : ''} atrás`
      } else {
        return date.toLocaleDateString('pt-BR', {
          day: '2-digit',
          month: '2-digit',
          year: 'numeric'
        })
      }
    },

    // Métodos para verificar permissões
    canEditReview(review) {
      // Só pode editar se for o próprio usuário
      return this.currentUserId && review.user.id === this.currentUserId
    },

    canDeleteReview(review) {
      // Pode excluir se for admin OU se for o próprio usuário
      return this.isAdmin || (this.currentUserId && review.user.id === this.currentUserId)
    },

    // Métodos para ações dos botões
    editReview(review) {
      // Emitir evento para o componente pai lidar com a edição
      this.$emit('edit-review', review)
      console.log('Editando review:', review.id)
    },

    deleteReview(review) {
      // Confirmar antes de excluir
      if (confirm('Tem certeza que deseja excluir este review?')) {
        // Emitir evento para o componente pai lidar com a exclusão
        this.$emit('delete-review', review)
        console.log('Excluindo review:', review.id)
      }
    }
  }
}
</script>

<style scoped>
.reviews-comments {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* Header dos Reviews */
.reviews-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 2px solid #2c3e50;
}

.reviews-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #ffffff;
  margin: 0;
}

.reviews-count {
  font-size: 1rem;
  color: #888;
  background-color: #2a2a2a;
  padding: 6px 12px;
  border-radius: 16px;
}

/* Seção de adicionar review */
.add-review-section {
  margin-bottom: 24px;
}

.add-review-btn {
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
}

.add-review-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(52, 152, 219, 0.4);
}

.add-review-btn.active {
  background: linear-gradient(135deg, #e74c3c, #c0392b);
  box-shadow: 0 2px 8px rgba(231, 76, 60, 0.3);
}

.btn-icon {
  font-size: 1.2rem;
  font-weight: bold;
}

/* Formulário de review */
.review-form-container {
  background-color: #2a2a2a;
  border-radius: 12px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.3);
}

.review-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-title {
  font-size: 1.4rem;
  color: #ffffff;
  margin: 0;
  text-align: center;
}

.input-label {
  color: #cccccc;
  font-weight: 600;
  margin-bottom: 8px;
  display: block;
}

/* Input de rating */
.rating-input {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.stars-input {
  display: flex;
  gap: 4px;
}

.star-btn {
  background: none;
  border: none;
  font-size: 2rem;
  color: #444;
  cursor: pointer;
  transition: all 0.2s ease;
  padding: 4px;
  border-radius: 4px;
}

.star-btn:hover {
  transform: scale(1.1);
}

.star-btn.filled {
  color: #ffd700;
}

.rating-text {
  color: #888;
  font-size: 0.9rem;
  margin-top: 4px;
}

/* Input de comentário */
.comment-input {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.comment-textarea {
  background-color: #1a1a1a;
  border: 2px solid #444;
  border-radius: 8px;
  color: #ffffff;
  padding: 12px;
  font-size: 1rem;
  resize: vertical;
  min-height: 100px;
  transition: border-color 0.2s ease;
}

.comment-textarea:focus {
  outline: none;
  border-color: #3498db;
}

.comment-textarea::placeholder {
  color: #666;
}

.char-count {
  color: #888;
  font-size: 0.8rem;
  text-align: right;
}

/* Botões do formulário */
.form-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.btn-secondary {
  background-color: #444;
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-secondary:hover {
  background-color: #555;
}

.btn-primary {
  background: linear-gradient(135deg, #27ae60, #2ecc71);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 600;
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(39, 174, 96, 0.4);
}

.btn-primary:disabled {
  background: #666;
  cursor: not-allowed;
}

/* Animação do formulário */
.form-slide-enter-active,
.form-slide-leave-active {
  transition: all 0.3s ease;
}

.form-slide-enter-from {
  opacity: 0;
  transform: translateY(-20px);
}

.form-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

/* Lista de Reviews */
.reviews-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.review-item {
  background-color: #2a2a2a;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
  transition: transform 0.2s ease;
}

.review-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

/* Informações do Usuário */
.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.user-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #2c3e50;
}

.user-details {
  flex: 1;
}

.user-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #ffffff;
  margin: 0 0 4px 0;
}

.review-date {
  font-size: 0.9rem;
  color: #888;
}

/* Botões de Ação */
.review-actions {
  display: flex;
  gap: 8px;
  margin-left: auto;
}

.action-btn {
  background: none;
  border: none;
  padding: 8px;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0.7;
}

.action-btn:hover {
  opacity: 1;
  transform: scale(1.05);
}

.edit-btn {
  color: #3498db;
  background-color: rgba(52, 152, 219, 0.1);
}

.edit-btn:hover {
  background-color: rgba(52, 152, 219, 0.2);
}

.delete-btn {
  color: #e74c3c;
  background-color: rgba(231, 76, 60, 0.1);
}

.delete-btn:hover {
  background-color: rgba(231, 76, 60, 0.2);
}

/* Seção de Avaliação */
.rating-section {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.stars {
  display: flex;
  gap: 2px;
}

.star {
  font-size: 1.2rem;
  color: #444;
  transition: color 0.2s ease;
}

.star.filled {
  color: #ffd700;
}

.rating-value {
  font-size: 1rem;
  font-weight: 600;
  color: #ffffff;
  background-color: #2c3e50;
  padding: 4px 8px;
  border-radius: 8px;
}

/* Seção de Comentário */
.comment-section {
  margin-top: 12px;
}

.comment-text {
  color: #cccccc;
  line-height: 1.5;
  font-size: 1rem;
  margin: 0;
  text-align: justify;
}

/* Review apenas com nota */
.rating-only {
  margin-top: 8px;
}

.rating-only-text {
  color: #888;
  font-style: italic;
  font-size: 0.9rem;
}

/* Sem Reviews */
.no-reviews {
  text-align: center;
  padding: 60px 20px;
  background-color: #2a2a2a;
  border-radius: 12px;
  border: 2px dashed #444;
}

.no-reviews-text {
  font-size: 1.2rem;
  color: #888;
  margin: 0 0 8px 0;
}

.no-reviews-subtitle {
  font-size: 1rem;
  color: #666;
  margin: 0;
}

/* Responsividade */
@media (max-width: 768px) {
  .reviews-comments {
    padding: 16px;
  }

  .reviews-header {
    flex-direction: column;
    gap: 12px;
    align-items: flex-start;
  }

  .reviews-title {
    font-size: 1.5rem;
  }

  .review-form-container {
    padding: 20px;
  }

  .form-actions {
    flex-direction: column;
  }

  .stars-input {
    justify-content: center;
  }

  .review-item {
    padding: 16px;
  }

  .user-info {
    gap: 10px;
  }

  .user-avatar {
    width: 40px;
    height: 40px;
  }

  .user-name {
    font-size: 1rem;
  }

  .rating-section {
    flex-wrap: wrap;
    gap: 8px;
  }

  .review-actions {
    margin-left: 0;
    margin-top: 8px;
  }
}

@media (max-width: 480px) {
  .reviews-comments {
    padding: 12px;
  }

  .reviews-title {
    font-size: 1.3rem;
  }

  .review-form-container {
    padding: 16px;
  }

  .add-review-btn {
    width: 100%;
    justify-content: center;
  }

  .review-item {
    padding: 12px;
  }

  .user-avatar {
    width: 36px;
    height: 36px;
  }

  .stars {
    gap: 1px;
  }

  .star {
    font-size: 1rem;
  }

  .star-btn {
    font-size: 1.5rem;
  }

  .user-info {
    flex-wrap: wrap;
  }

  .review-actions {
    width: 100%;
    justify-content: flex-end;
  }
}
</style>