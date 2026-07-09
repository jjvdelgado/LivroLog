<template>
  <div class="info-perfil">
    <!-- Modo Visualização -->
    <div v-if="!editMode" class="profile-view">
      <!-- Foto do Usuário -->
      <!-- Foto do Usuário -->
      <div class="profile-image-container">
        <div class="profile-image">
          <img 
            v-if="userData.avatar" 
            :src="userData.avatar" 
            :alt="userData.name"
            class="profile-image-img"
          />
          <div v-else class="default-avatar">
            <span class="avatar-initials">{{ getInitials(userData.name) }}</span>
          </div>
        </div>
        <button @click="enableEditMode" class="edit-button">
          <svg class="edit-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
          </svg>
          Editar Perfil
        </button>
      </div>

      <!-- Informações do Usuário -->
      <div class="profile-info">
        <h1 class="user-name">{{ userData.name }}</h1>
        <p class="user-email">{{ user.email }}</p>
        <p class="user-bio">{{ userData.bio }}</p>
        
        <!-- Estatísticas -->
        <div class="stats-section">
          <h3 class="stats-title">Estatísticas</h3>
          <div class="stats-grid">
            <div class="stat-item">
              <span class="stat-value">{{ user.reviewsCount }}</span>
              <span class="stat-label">Reviews Feitos</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ user.averageRating.toFixed(1) }}</span>
              <span class="stat-label">Avaliação Média</span>
            </div>
            <div class="stat-item">
              <span class="stat-value">{{ user.memberSince }}</span>
              <span class="stat-label">Membro desde</span>
            </div>
          </div>
        </div>

        <!-- Reviews Recentes -->
        <div class="recent-reviews">
          <h3 class="reviews-title">Reviews Recentes</h3>
          <div class="reviews-list">
            <div v-for="review in user.recentReviews" :key="review.id" class="review-item">
              <div class="review-header">
                <span class="book-title">{{ review.bookTitle }}</span>
                <div class="review-rating">
                  <span 
                    v-for="star in 5" 
                    :key="star"
                    class="star"
                    :class="{ 'filled': star <= review.rating }"
                  >
                    ★
                  </span>
                </div>
              </div>
              <p v-if="review.comment" class="review-comment">{{ review.comment }}</p>
              <span class="review-date">{{ formatDate(review.date) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modo Edição -->
    <div v-else class="profile-edit">
      <form @submit.prevent="saveProfile" class="edit-form">
        <h2 class="edit-title">Editar Perfil</h2>
        
        <!-- Upload de Foto -->
        <div class="photo-upload-section">
          <div class="current-photo">
            <img 
              v-if="editData.avatar" 
              :src="editData.avatar" 
              :alt="editData.name"
              class="preview-image"
            />
            <div v-else class="default-avatar preview-size">
              <svg class="person-icon" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
              </svg>
            </div>
          </div>
          <div class="upload-controls">
            <label for="photo-upload" class="upload-button">
              <svg class="upload-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
              </svg>
              Alterar Foto
            </label>
            <input 
              id="photo-upload" 
              type="file" 
              accept="image/*" 
              @change="handlePhotoUpload"
              class="hidden-input"
            />
          </div>
        </div>

        <!-- Campos do Formulário -->
        <div class="form-group">
          <label for="name" class="form-label">Nome</label>
          <input 
            id="name"
            v-model="editData.name" 
            type="text" 
            class="form-input"
            required
          />
        </div>

        <div class="form-group">
          <label for="bio" class="form-label">Bio (Opcional)</label>
          <textarea 
            id="bio"
            v-model="editData.bio" 
            class="form-textarea"
            rows="3"
            placeholder="Conte um pouco sobre você..."
          ></textarea>
        </div>

        <!-- Botões de Ação -->
        <div class="form-actions">
          <button type="button" @click="cancelEdit" class="cancel-button">
            Cancelar
          </button>
          <button type="submit" class="save-button">
            Salvar Alterações
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'InfoPerfil',
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      editMode: false,
      editData: {}
    }
  },
  computed: {
    // Computed property para dados do usuário com localStorage
    userData() {
      const savedData = localStorage.getItem('userProfile')
      if (savedData) {
        const parsed = JSON.parse(savedData)
        // Mescla os dados salvos com os dados originais
        return { ...this.user, ...parsed }
      }
      return this.user
    }
  },
  methods: {
    enableEditMode() {
      this.editMode = true
      this.editData = { ...this.userData }
    },
    cancelEdit() {
      this.editMode = false
      this.editData = {}
    },
    saveProfile() {
      // Salvar no localStorage
      localStorage.setItem('userProfile', JSON.stringify(this.editData))
      
      // Emitir evento para o componente pai com os dados atualizados
      this.$emit('profile-updated', this.editData)
      this.editMode = false
      console.log('Perfil salvo:', this.editData)
    },
    handlePhotoUpload(event) {
      const file = event.target.files[0]
      if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
          this.editData.avatar = e.target.result
        }
        reader.readAsDataURL(file)
      }
    },
    formatDate(dateString) {
      const date = new Date(dateString)
      return date.toLocaleDateString('pt-BR', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric'
      })
    }
  }
}
</script>

<style scoped>
.info-perfil {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* Modo Visualização */
.profile-view {
  display: flex;
  gap: 40px;
  background-color: #2a2a2a;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.profile-image-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  flex-shrink: 0;
}

.profile-image {
  width: 200px;
  height: 200px;
  border-radius: 50%;
  border: 4px solid #2c3e50;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.profile-image-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.default-avatar {
  width: 100%;
  height: 100%;
  background-color: #1a1a1a;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
}

.default-avatar.preview-size {
  width: 120px;
  height: 120px;
  border: 3px solid #2c3e50;
}

.person-icon {
  width: 60%;
  height: 60%;
}

.edit-button {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  background-color: #2c3e50;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.edit-button:hover {
  background-color: #34495e;
  transform: translateY(-2px);
}

.edit-icon {
  width: 18px;
  height: 18px;
}

.profile-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.user-name {
  font-size: 2.5rem;
  font-weight: bold;
  color: #ffffff;
  margin: 0;
}

.user-email {
  font-size: 1.1rem;
  color: #888;
  margin: 0;
}

.user-bio {
  font-size: 1.1rem;
  color: #888;
  margin: 0;
}

/* Estatísticas */
.stats-section {
  background-color: #1a1a1a;
  padding: 20px;
  border-radius: 8px;
}

.stats-title {
  font-size: 1.3rem;
  color: #2c3e50;
  margin: 0 0 16px 0;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
}

.stat-item {
  text-align: center;
}

.stat-value {
  display: block;
  font-size: 2rem;
  font-weight: bold;
  color: #ffffff;
}

.stat-label {
  font-size: 0.9rem;
  color: #888;
}

/* Reviews Recentes */
.recent-reviews {
  background-color: #1a1a1a;
  padding: 20px;
  border-radius: 8px;
}

.reviews-title {
  font-size: 1.3rem;
  color: #2c3e50;
  margin: 0 0 16px 0;
}

.reviews-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.review-item {
  background-color: #2a2a2a;
  padding: 16px;
  border-radius: 8px;
  border-left: 4px solid #2c3e50;
}

.review-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.book-title {
  font-weight: 600;
  color: #ffffff;
}

.review-rating {
  display: flex;
  gap: 2px;
}

.star {
  color: #444;
  font-size: 1rem;
}

.star.filled {
  color: #ffd700;
}

.review-comment {
  color: #cccccc;
  margin: 8px 0;
  font-size: 0.9rem;
  line-height: 1.4;
}

.review-date {
  font-size: 0.8rem;
  color: #888;
}

/* Modo Edição - ALTERAÇÕES AQUI */
.profile-edit {
  background-color: #2a2a2a;
  border-radius: 12px;
  padding: 30px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  /* Definindo largura fixa para evitar mudanças dinâmicas */
  width: 100%;
  max-width: 1200px; /* Mesma largura máxima do container pai */
  box-sizing: border-box;
}

.edit-form {
  width: 600px; /* Largura fixa em vez de max-width */
  margin: 0 auto;
  /* Garantindo que o formulário não mude de tamanho */
  box-sizing: border-box;
}

.edit-title {
  font-size: 2rem;
  color: #ffffff;
  margin: 0 0 30px 0;
  text-align: center;
}

/* Upload de Foto */
.photo-upload-section {
  display: flex;
  align-items: center;
  gap: 30px;
  margin-bottom: 30px;
  padding: 20px;
  background-color: #1a1a1a;
  border-radius: 8px;
  /* Largura fixa para evitar mudanças */
  width: 100%;
  box-sizing: border-box;
}

.current-photo {
  flex-shrink: 0;
}

.preview-image {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #2c3e50;
}

.upload-controls {
  flex: 1;
}

.upload-button {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  background-color: #2c3e50;
  color: white;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  width: fit-content;
}

.upload-button:hover {
  background-color: #34495e;
}

.upload-icon {
  width: 18px;
  height: 18px;
}

.hidden-input {
  display: none;
}

/* Formulário */
.form-group {
  margin-bottom: 20px;
  /* Largura fixa para os grupos do formulário */
  width: 100%;
  box-sizing: border-box;
}

.form-label {
  display: block;
  color: #2c3e50;
  font-weight: 600;
  margin-bottom: 8px;
}

.form-input,
.form-textarea {
  width: 100%;
  padding: 12px 16px;
  background-color: #1a1a1a;
  border: 2px solid #444;
  border-radius: 8px;
  color: #ffffff;
  font-size: 1rem;
  transition: border-color 0.3s ease;
  /* Garantindo largura fixa */
  box-sizing: border-box;
  /* Evitando mudanças de tamanho no resize */
  resize: none;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: #2c3e50;
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

/* Botões de Ação */
.form-actions {
  display: flex;
  gap: 16px;
  justify-content: flex-end;
  margin-top: 30px;
  /* Largura fixa */
  width: 100%;
  box-sizing: border-box;
}

.cancel-button,
.save-button {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  /* Largura fixa dos botões */
  min-width: 120px;
}

.cancel-button {
  background-color: #666;
  color: white;
}

.cancel-button:hover {
  background-color: #777;
}

.save-button {
  background-color: #2c3e50;
  color: white;
}

.save-button:hover {
  background-color: #34495e;
}

/* Responsividade */
@media (max-width: 768px) {
  .profile-view {
    flex-direction: column;
    gap: 24px;
    padding: 20px;
  }

  .profile-image {
    width: 150px;
    height: 150px;
  }

  .user-name {
    font-size: 2rem;
  }

  .stats-grid {
    grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  }

  .photo-upload-section {
    flex-direction: column;
    gap: 20px;
    text-align: center;
  }

  .form-actions {
    flex-direction: column;
  }

  .review-header {
    flex-direction: column;
    gap: 8px;
    align-items: flex-start;
  }

  /* Ajustando o formulário para mobile */
  .edit-form {
    width: 100%;
    max-width: 500px;
  }
}

@media (max-width: 480px) {
  .info-perfil {
    padding: 12px;
  }

  .profile-view,
  .profile-edit {
    padding: 16px;
  }

  .profile-image {
    width: 120px;
    height: 120px;
  }

  .preview-image {
    width: 100px;
    height: 100px;
  }

  .user-name {
    font-size: 1.8rem;
  }

  .stats-grid {
    grid-template-columns: 1fr;
  }

  /* Formulário ainda menor para telas muito pequenas */
  .edit-form {
    width: 100%;
    max-width: 400px;
  }
}
</style>