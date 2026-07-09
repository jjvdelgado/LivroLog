<template>
  <div class="review-page">
    <!-- Header Component -->
    <AppHeader />

    <!-- Main Content -->
    <main class="main-content">
      <!-- Conteúdo da página de review -->
      <div class="content-container">
        <!-- Componente de Informações do Livro -->
        <InfoLivros :book="selectedBook" />

        <!-- Componente de Reviews e Comentários -->
        <ReviewsComments 
          :reviews="bookReviews"
          :isAdmin="currentUser.isAdmin"
          :currentUserId="currentUser.id"
          @review-submitted="handleReviewSubmitted"
          @edit-review="handleEditReview"
          @delete-review="handleDeleteReview"
          @show-message="handleShowMessage"
        />
      </div>
    </main>

    <!-- Modal de Confirmação de Exclusão -->
    <div v-if="showDeleteModal" class="modal-overlay" @click="closeDeleteModal">
      <div class="modal-content delete-modal" @click.stop>
        <div class="modal-header">
          <h3>Confirmar Exclusão</h3>
          <button class="close-btn" @click="closeDeleteModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="warning-icon">⚠️</div>
          <p>Tem certeza que deseja excluir este comentário?</p>
          <p class="warning-text">Esta ação não pode ser desfeita.</p>
        </div>
        <div class="modal-footer">
          <button class="btn btn-cancel" @click="closeDeleteModal">Cancelar</button>
          <button class="btn btn-delete" @click="confirmDelete">Excluir</button>
        </div>
      </div>
    </div>

    <!-- Modal de Edição -->
    <div v-if="showEditModal" class="modal-overlay" @click="closeEditModal">
      <div class="modal-content edit-modal" @click.stop>
        <div class="modal-header">
          <h3>Editar Comentário</h3>
          <button class="close-btn" @click="closeEditModal">&times;</button>
        </div>
        <div class="modal-body">
          <div class="form-group">
            <label for="rating">Nota:</label>
            <div class="rating-container">
              <div class="stars">
                <span 
                  v-for="star in 5" 
                  :key="star"
                  class="star"
                  :class="{ active: star <= editData.rating }"
                  @click="setRating(star)"
                >
                  ★
                </span>
              </div>
              <span class="rating-text">{{ editData.rating }}/5</span>
            </div>
          </div>
          <div class="form-group">
            <label for="comment">Comentário:</label>
            <textarea 
              id="comment"
              v-model="editData.comment"
              placeholder="Escreva seu comentário sobre o livro..."
              rows="4"
            ></textarea>
            <div class="char-count">
              {{ editData.comment ? editData.comment.length : 0 }}/500
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn btn-cancel" @click="closeEditModal">Cancelar</button>
          <button class="btn btn-save" @click="confirmEdit" :disabled="!editData.rating">
            Salvar Alterações
          </button>
        </div>
      </div>
    </div>

    <!-- Modal de Sucesso -->
    <div v-if="showSuccessModal" class="modal-overlay" @click="closeSuccessModal">
      <div class="modal-content success-modal" @click.stop>
        <div class="modal-body">
          <div class="success-icon">✓</div>
          <p class="success-message">{{ successMessage }}</p>
          <button class="btn btn-primary" @click="closeSuccessModal">OK</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppHeader from '@/components/AppHeader.vue'
import InfoLivros from '@/components/InfoLivros.vue'
import ReviewsComments from '@/components/ReviewsComments.vue'

import { useRequest } from '@/composables/useRequest'

const { request } = useRequest()

export default {
  name: 'ReviewPage',
  components: {
    AppHeader,
    InfoLivros,
    ReviewsComments
  },
  data() {
    return {
      // Dados do usuário atual (substitua pela sua lógica de autenticação)
      currentUser: {
        id: 1, // ID do usuário atual - deve vir do sistema de auth
        name: 'Maria Silva',
        avatar: '/images/avatars/maria.jpg',
        isAdmin: true // Mude para true para testar como admin
      },
      // Dados de exemplo - substitua pelos dados reais do livro selecionado
      selectedBook: {
        title: 'O Senhor dos Anéis: A Sociedade do Anel',
        author: 'J.R.R. Tolkien',
        year: 1954,
        publisher: 'HarperCollins',
        genre: 'Fantasia',
        pages: 576,
        rating: 4.8,
        reviewsCount: 2847,
        image: '/images/lotr-fellowship.jpg',
        synopsis: 'Em uma terra fantástica e única, um hobbit recebe de presente de seu tio um anel mágico e maligno que precisa ser destruído antes que caia nas mãos do mal. Para isso, o hobbit Frodo tem a ajuda de magos, anões, elfos e outros hobbits, além de seu fiel companheiro Sam, para realizar sua missão. O livro fala sobre a amizade, a persistência e a vitória do bem sobre o mal.'
      },
      // Dados de exemplo dos reviews (adicionei IDs aos usuários)
      bookReviews: [
        {
          id: 1,
          user: {
            id: 1, // Mesmo ID do usuário atual - poderá editar/excluir
            name: 'Maria Silva',
            avatar: '/images/avatars/maria.jpg'
          },
          rating: 5,
          comment: 'Simplesmente magnífico! Tolkien criou um mundo tão rico e detalhado que é impossível não se apaixonar. A jornada do Frodo é emocionante do início ao fim, e os personagens são inesquecíveis. Recomendo para todos os amantes de fantasia!',
          date: '2024-06-15T10:30:00Z'
        },
        {
          id: 2,
          user: {
            id: 2, // Usuário diferente - não poderá editar, apenas admin pode excluir
            name: 'João Santos',
            avatar: '/images/avatars/joao.jpg'
          },
          rating: 4,
          comment: 'Excelente livro! A narrativa é envolvente e os detalhes do mundo são impressionantes. Algumas partes são um pouco lentas, mas no geral é uma obra-prima da literatura fantástica.',
          date: '2024-06-10T14:22:00Z'
        },
        {
          id: 3,
          user: {
            id: 3,
            name: 'Ana Costa',
            avatar: '/images/avatars/ana.jpg'
          },
          rating: 5,
          comment: null, // Review apenas com nota
          date: '2024-06-08T09:15:00Z'
        },
        {
          id: 4,
          user: {
            id: 4,
            name: 'Carlos Oliveira',
            avatar: '/images/avatars/carlos.jpg'
          },
          rating: 3,
          comment: 'Boa história, mas achei muito descritivo em alguns momentos. Para quem gosta de fantasy é uma boa pedida, mas não é para todos os gostos.',
          date: '2024-06-05T16:45:00Z'
        },
        {
          id: 5,
          user: {
            id: 5,
            name: 'Lucia Ferreira',
            avatar: '/images/avatars/lucia.jpg'
          },
          rating: 4,
          comment: null, // Review apenas com nota
          date: '2024-06-03T11:20:00Z'
        },
        {
          id: 6,
          user: {
            id: 6,
            name: 'Pedro Almeida',
            avatar: '/images/avatars/pedro.jpg'
          },
          rating: 5,
          comment: 'Clássico absoluto! Mesmo depois de tantos anos, continua sendo uma das melhores histórias de fantasia já escritas. A profundidade dos personagens e a riqueza do mundo criado por Tolkien são incomparáveis.',
          date: '2024-05-28T13:10:00Z'
        }
      ],
      // Controle dos modais
      showDeleteModal: false,
      showEditModal: false,
      showSuccessModal: false,
      successMessage: '',
      reviewToDelete: null,
      reviewToEdit: null,
      editData: {
        rating: 0,
        comment: ''
      }
    }
  },
  methods: {
    // Método para lidar com novo review
    handleReviewSubmitted(reviewData) {
      console.log('Novo review recebido:', reviewData)
      
      // Adiciona o novo review à lista
      this.bookReviews.unshift(reviewData)
      
      // Aqui você pode fazer a chamada para o backend
      this.salvarAvaliacao(reviewData)
    },

    // Método para lidar com edição de review - MELHORADO
    handleEditReview(review) {
      console.log('Editando review:', review)
      
      this.reviewToEdit = review
      this.editData = {
        rating: review.rating,
        comment: review.comment || ''
      }
      this.showEditModal = true
    },

    // Método para lidar com exclusão de review - MELHORADO
    handleDeleteReview(review) {
      console.log('Solicitando exclusão de review:', review)
      
      this.reviewToDelete = review
      this.showDeleteModal = true
    },

    // Confirma a exclusão
    async confirmDelete() {
      if (!this.reviewToDelete) return
      
      try {
        // Remove o review da lista
        const index = this.bookReviews.findIndex(r => r.id === this.reviewToDelete.id)
        if (index !== -1) {
          this.bookReviews.splice(index, 1)
          
          // Chama o backend para excluir
          await this.excluirAvaliacaoNoBackend(this.reviewToDelete.id)
          
          this.closeDeleteModal()
          this.showSuccessMessage('Comentário excluído com sucesso!')
        }
      } catch (error) {
        console.error('Erro ao excluir review:', error)
        this.handleShowMessage({
          type: 'error',
          message: 'Erro ao excluir comentário'
        })
      }
    },

    // Confirma a edição
    async confirmEdit() {
      if (!this.reviewToEdit || !this.editData.rating) return
      
      try {
        // Encontra o review na lista e atualiza
        const index = this.bookReviews.findIndex(r => r.id === this.reviewToEdit.id)
        if (index !== -1) {
          this.bookReviews[index].rating = this.editData.rating
          this.bookReviews[index].comment = this.editData.comment
          this.bookReviews[index].date = new Date().toISOString()
          
          // Chama o backend para atualizar
          await this.atualizarAvaliacaoNoBackend(this.bookReviews[index])
          
          this.closeEditModal()
          this.showSuccessMessage('Comentário atualizado com sucesso!')
        }
      } catch (error) {
        console.error('Erro ao editar review:', error)
        this.handleShowMessage({
          type: 'error',
          message: 'Erro ao atualizar comentário'
        })
      }
    },

    // Controle dos modais
    closeDeleteModal() {
      this.showDeleteModal = false
      this.reviewToDelete = null
    },

    closeEditModal() {
      this.showEditModal = false
      this.reviewToEdit = null
      this.editData = { rating: 0, comment: '' }
    },

    closeSuccessModal() {
      this.showSuccessModal = false
      this.successMessage = ''
    },

    showSuccessMessage(message) {
      this.successMessage = message
      this.showSuccessModal = true
    },

    // Controle da avaliação por estrelas
    setRating(rating) {
      this.editData.rating = rating
    },

    // Método para mostrar mensagens
    handleShowMessage(message) {
      console.log('Mensagem:', message)
      
      // Aqui você pode implementar um sistema de notificações
      // Por exemplo, usando toast/snackbar
      if (message.type === 'error') {
        alert(`ERRO: ${message.message}`)
      } else {
        this.showSuccessMessage(message.message)
      }
    },

    // Métodos para chamadas do backend (adapte conforme sua API)
    async salvarAvaliacao(reviewData) {
      try {
        // Adapte conforme sua estrutura de dados
        const avaliacaoParaBackend = {
          livro_id: this.selectedBook.id,
          nota: reviewData.rating,
          comentario: reviewData.comment,
          usuario_id: this.currentUser.id
        }

        const url = `http://localhost:8080/avaliacoes`
        const resultado = await request(url, 'POST', avaliacaoParaBackend)

        console.log('Avaliação salva com sucesso:', resultado)
        
        // Atualiza o review com o ID retornado pelo backend
        const reviewIndex = this.bookReviews.findIndex(r => r.id === reviewData.id)
        if (reviewIndex !== -1 && resultado.id) {
          this.bookReviews[reviewIndex].id = resultado.id
        }

      } catch (e) {
        console.error('Erro ao salvar avaliação:', e.message)
        this.handleShowMessage({
          type: 'error',
          message: 'Erro ao salvar avaliação'
        })
      }
    },

    async atualizarAvaliacaoNoBackend(review) {
      try {
        const url = `http://localhost:8080/avaliacoes/${review.id}`
        const avaliacaoAtualizada = {
          nota: review.rating,
          comentario: review.comment
        }
        
        await request(url, 'PUT', avaliacaoAtualizada)
        console.log('Avaliação atualizada com sucesso')
        
      } catch (e) {
        console.error('Erro ao atualizar avaliação:', e.message)
        throw e
      }
    },

    async excluirAvaliacaoNoBackend(reviewId) {
      try {
        const url = `http://localhost:8080/avaliacoes/${reviewId}`
        await request(url, 'DELETE')
        console.log('Avaliação excluída com sucesso')
        
      } catch (e) {
        console.error('Erro ao excluir avaliação:', e.message)
        throw e
      }
    },

    async buscarDetalhesLivro(livroId) {
      try {
        const url = `http://localhost:8080/livros/${livroId}/detalhes`
        const dados = await request(url, 'GET')
        console.log('Detalhes do livro:', dados)

        this.selectedBook = dados
      } catch (e) {
        console.error('Erro ao buscar detalhes do livro:', e.message)
        this.erro = 'Erro ao buscar detalhes'
      }
    }
  },
  
  mounted() {
    // Busca detalhes do livro quando o componente é montado
    // this.buscarDetalhesLivro(this.$route.params.id) // Se você estiver usando roteamento
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.review-page {
  min-height: 100vh;
  background-color: #1a1a1a;
  margin: 0;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.main-content {
  padding-top: 64px;
  height: calc(100vh - 64px);
  width: 100%;
  background-color: #1a1a1a;
  margin: 0;
  padding-left: 0;
  padding-right: 0;
  position: absolute;
  left: 0;
  right: 0;
  top: 64px;
  overflow-x: hidden;
  overflow-y: auto;
}

.content-container {
  display: flex;
  flex-direction: column;
  gap: 40px;
  justify-content: flex-start;
  align-items: stretch;
  padding: 40px 20px;
  width: 100%;
  min-height: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

/* Estilos dos Modais */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(5px);
}

.modal-content {
  background: #2c3e50;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
  max-width: 500px;
  width: 90%;
  max-height: 90vh;
  overflow-y: auto;
  animation: modalSlideIn 0.3s ease-out;
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #34495e;
}

.modal-header h3 {
  color: #ecf0f1;
  margin: 0;
  font-size: 1.3em;
}

.close-btn {
  background: none;
  border: none;
  color: #bdc3c7;
  font-size: 24px;
  cursor: pointer;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: background-color 0.2s;
}

.close-btn:hover {
  background-color: #34495e;
  color: #ecf0f1;
}

.modal-body {
  padding: 20px;
}

.modal-footer {
  padding: 20px;
  border-top: 1px solid #34495e;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Modal de Exclusão */
.delete-modal .modal-body {
  text-align: center;
}

.warning-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.delete-modal p {
  color: #ecf0f1;
  margin-bottom: 8px;
  font-size: 1.1em;
}

.warning-text {
  color: #e74c3c;
  font-size: 0.9em;
}

/* Modal de Edição */
.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  color: #ecf0f1;
  margin-bottom: 8px;
  font-weight: 500;
}

.rating-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.stars {
  display: flex;
  gap: 4px;
}

.star {
  font-size: 24px;
  color: #7f8c8d;
  cursor: pointer;
  transition: color 0.2s;
}

.star:hover,
.star.active {
  color: #f39c12;
}

.rating-text {
  color: #bdc3c7;
  font-size: 0.9em;
}

.form-group textarea {
  width: 100%;
  min-height: 100px;
  background-color: #34495e;
  border: 1px solid #4a5f7a;
  border-radius: 8px;
  color: #ecf0f1;
  padding: 12px;
  font-family: inherit;
  font-size: 14px;
  resize: vertical;
  transition: border-color 0.2s;
}

.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
}

.form-group textarea::placeholder {
  color: #7f8c8d;
}

.char-count {
  text-align: right;
  color: #7f8c8d;
  font-size: 0.8em;
  margin-top: 4px;
}

/* Modal de Sucesso */
.success-modal .modal-body {
  text-align: center;
  padding: 30px;
}

.success-icon {
  font-size: 48px;
  color: #27ae60;
  margin-bottom: 16px;
}

.success-message {
  color: #ecf0f1;
  font-size: 1.1em;
  margin-bottom: 20px;
}

/* Botões */
.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s;
  min-width: 80px;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-cancel {
  background-color: #7f8c8d;
  color: white;
}

.btn-cancel:hover:not(:disabled) {
  background-color: #95a5a6;
}

.btn-delete {
  background-color: #e74c3c;
  color: white;
}

.btn-delete:hover:not(:disabled) {
  background-color: #c0392b;
}

.btn-save {
  background-color: #27ae60;
  color: white;
}

.btn-save:hover:not(:disabled) {
  background-color: #229954;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #2980b9;
}

/* Responsividade */
@media (max-width: 768px) {
  .content-container {
    padding: 30px 15px;
  }
  
  .modal-content {
    margin: 20px;
    width: calc(100% - 40px);
  }
  
  .modal-footer {
    flex-direction: column;
  }
  
  .modal-footer .btn {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .content-container {
    padding: 20px 10px;
  }
  
  .modal-header h3 {
    font-size: 1.1em;
  }
  
  .stars .star {
    font-size: 20px;
  }
}
</style>