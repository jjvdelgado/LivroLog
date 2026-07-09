<template>
  <div class="add-book-form">
    <div class="form-container">
      <h2 class="form-title">Adicionar Novo Livro</h2>
      
      <form @submit.prevent="handleSubmit" class="book-form">
        <!-- Título -->
        <div class="form-group">
          <label for="titulo" class="form-label">Título *</label>
          <input
            type="text"
            id="titulo"
            v-model="bookData.titulo"
            class="form-input"
            placeholder="Digite o título do livro"
            required
          />
          <span v-if="errors.titulo" class="error-message">{{ errors.titulo }}</span>
        </div>

        <!-- Autor -->
        <div class="form-group">
          <label for="autor" class="form-label">Autor *</label>
          <input
            type="text"
            id="autor"
            v-model="bookData.autor"
            class="form-input"
            placeholder="Digite o nome do autor"
            required
          />
          <span v-if="errors.autor" class="error-message">{{ errors.autor }}</span>
        </div>

        <!-- Ano de Publicação e Editora -->
        <div class="form-row">
          <div class="form-group">
            <label for="anoPublicacao" class="form-label">Ano de Publicação *</label>
            <input
              type="number"
              id="anoPublicacao"
              v-model.number="bookData.anoPublicacao"
              class="form-input"
              placeholder="Ex: 2023"
              min="1000"
              max="2024"
              required
            />
            <span v-if="errors.anoPublicacao" class="error-message">{{ errors.anoPublicacao }}</span>
          </div>

          <div class="form-group">
            <label for="editora" class="form-label">Editora *</label>
            <input
              type="text"
              id="editora"
              v-model="bookData.editora"
              class="form-input"
              placeholder="Digite o nome da editora"
              required
            />
            <span v-if="errors.editora" class="error-message">{{ errors.editora }}</span>
          </div>
        </div>

        <!-- Gênero e Número de Páginas -->
        <div class="form-row">
          <div class="form-group">
            <label for="genero" class="form-label">Gênero *</label>
            <select
              id="genero"
              v-model="bookData.genero"
              class="form-select"
              required
            >
              <option value="">Selecione um gênero</option>
              <option value="Ficção">Ficção</option>
              <option value="Não-ficção">Não-ficção</option>
              <option value="Romance">Romance</option>
              <option value="Mistério">Mistério</option>
              <option value="Fantasia">Fantasia</option>
              <option value="Ficção Científica">Ficção Científica</option>
              <option value="Biografia">Biografia</option>
              <option value="História">História</option>
              <option value="Autoajuda">Autoajuda</option>
              <option value="Aventura">Aventura</option>
              <option value="Drama">Drama</option>
              <option value="Terror">Terror</option>
              <option value="Thriller">Thriller</option>
              <option value="Poesia">Poesia</option>
              <option value="Outro">Outro</option>
            </select>
            <span v-if="errors.genero" class="error-message">{{ errors.genero }}</span>
          </div>

          <div class="form-group">
            <label for="numeroPaginas" class="form-label">Número de Páginas *</label>
            <input
              type="number"
              id="numeroPaginas"
              v-model.number="bookData.numeroPaginas"
              class="form-input"
              placeholder="Ex: 350"
              min="1"
              required
            />
            <span v-if="errors.numeroPaginas" class="error-message">{{ errors.numeroPaginas }}</span>
          </div>
        </div>

        <!-- Sinopse -->
        <div class="form-group">
          <label for="sinopse" class="form-label">Sinopse *</label>
          <textarea
            id="sinopse"
            v-model="bookData.sinopse"
            class="form-textarea"
            placeholder="Digite a sinopse do livro"
            rows="4"
            required
          ></textarea>
          <span v-if="errors.sinopse" class="error-message">{{ errors.sinopse }}</span>
        </div>

        <!-- Imagem da Capa -->
        <div class="form-group">
          <label for="imagemCapa" class="form-label">Imagem da Capa *</label>
          <div class="file-upload-container">
            <input
              type="file"
              id="imagemCapa"
              @change="handleFileUpload"
              accept="image/*"
              class="file-input"
              required
            />
            <label for="imagemCapa" class="file-label">
              <span class="file-icon">📁</span>
              {{ bookData.imagemCapa ? bookData.imagemCapa.name : 'Escolher arquivo' }}
            </label>
          </div>
          <span v-if="errors.imagemCapa" class="error-message">{{ errors.imagemCapa }}</span>
          
          <!-- Preview da imagem -->
          <div v-if="imagePreview" class="image-preview">
            <img :src="imagePreview" alt="Preview da capa" class="preview-image" />
          </div>
        </div>

        <!-- Botões -->
        <div class="form-actions">
          <button type="button" @click="resetForm" class="btn btn-secondary">
            Limpar
          </button>
          <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
            {{ isSubmitting ? 'Salvando...' : 'Adicionar Livro' }}
          </button>
        </div>
      </form>
    </div>

    <!-- Modal de Sucesso -->
    <div v-if="showSuccessModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <div class="success-icon">✓</div>
          <h3 class="modal-title">Sucesso!</h3>
        </div>
        <div class="modal-body">
          <p class="success-message">O livro "{{ bookData.titulo }}" foi adicionado com sucesso!</p>
        </div>
        <div class="modal-footer">
          <button @click="goToHomePage" class="btn btn-primary modal-btn">
            Continuar
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FormsAddBook',
  data() {
    return {
      bookData: {
        titulo: '',
        autor: '',
        anoPublicacao: '',
        editora: '',
        genero: '',
        numeroPaginas: '',
        sinopse: '',
        imagemCapa: null
      },
      errors: {},
      imagePreview: null,
      isSubmitting: false,
      showSuccessModal: false
    }
  },
  methods: {
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.bookData.imagemCapa = file;
        this.createImagePreview(file);
        this.clearError('imagemCapa');
      }
    },
    
    createImagePreview(file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.imagePreview = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    
    validateForm() {
      this.errors = {};
      
      if (!this.bookData.titulo.trim()) {
        this.errors.titulo = 'O título é obrigatório';
      }
      
      if (!this.bookData.autor.trim()) {
        this.errors.autor = 'O autor é obrigatório';
      }
      
      if (!this.bookData.anoPublicacao) {
        this.errors.anoPublicacao = 'O ano de publicação é obrigatório';
      } else if (this.bookData.anoPublicacao < 1000 || this.bookData.anoPublicacao > 2024) {
        this.errors.anoPublicacao = 'Digite um ano válido';
      }
      
      if (!this.bookData.editora.trim()) {
        this.errors.editora = 'A editora é obrigatória';
      }
      
      if (!this.bookData.genero) {
        this.errors.genero = 'O gênero é obrigatório';
      }
      
      if (!this.bookData.numeroPaginas) {
        this.errors.numeroPaginas = 'O número de páginas é obrigatório';
      } else if (this.bookData.numeroPaginas < 1) {
        this.errors.numeroPaginas = 'Digite um número válido de páginas';
      }
      
      if (!this.bookData.sinopse.trim()) {
        this.errors.sinopse = 'A sinopse é obrigatória';
      }
      
      if (!this.bookData.imagemCapa) {
        this.errors.imagemCapa = 'A imagem da capa é obrigatória';
      }
      
      return Object.keys(this.errors).length === 0;
    },
    
    clearError(field) {
      if (this.errors[field]) {
        delete this.errors[field];
      }
    },
    
    async handleSubmit() {
      if (!this.validateForm()) {
        return;
      }
      
      this.isSubmitting = true;
      
      try {
        // Aqui você faria a chamada para a API
        const formData = new FormData();
        formData.append('titulo', this.bookData.titulo);
        formData.append('autor', this.bookData.autor);
        formData.append('anoPublicacao', this.bookData.anoPublicacao);
        formData.append('editora', this.bookData.editora);
        formData.append('genero', this.bookData.genero);
        formData.append('numeroPaginas', this.bookData.numeroPaginas);
        formData.append('sinopse', this.bookData.sinopse);
        formData.append('imagemCapa', this.bookData.imagemCapa);
        
        // Exemplo de chamada para API
        // const response = await fetch('/api/books', {
        //   method: 'POST',
        //   body: formData
        // });
        
        // if (response.ok) {
        //   this.showSuccessModal = true;
        // }
        
        // Simulação de sucesso
        setTimeout(() => {
          this.isSubmitting = false;
          this.showSuccessModal = true;
        }, 1500);
        
      } catch (error) {
        console.error('Erro ao adicionar livro:', error);
        alert('Erro ao adicionar livro. Tente novamente.');
        this.isSubmitting = false;
      }
    },

    closeModal() {
      this.showSuccessModal = false;
    },

    goToHomePage() {
      this.showSuccessModal = false;
      this.$router.push('/HomePage');
    },
    
    resetForm() {
      this.bookData = {
        titulo: '',
        autor: '',
        anoPublicacao: '',
        editora: '',
        genero: '',
        numeroPaginas: '',
        sinopse: '',
        imagemCapa: null
      };
      this.errors = {};
      this.imagePreview = null;
      this.isSubmitting = false;
      this.showSuccessModal = false;
      
      // Limpar o input de arquivo
      const fileInput = document.getElementById('imagemCapa');
      if (fileInput) {
        fileInput.value = '';
      }
    }
  }
}
</script>

<style scoped>
.add-book-form {
  min-height: 100vh;
  background-color: #1a1a1a;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 100px 20px 20px 20px; /* Adicionado padding-top para compensar o header */
}

.form-container {
  background-color: #2a2a2a;
  border-radius: 12px;
  padding: 40px;
  width: 100%;
  max-width: 800px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
}

.form-title {
  color: #ecf0f1;
  text-align: center;
  margin-bottom: 30px;
  font-size: 2.2em;
  font-weight: 600;
}

.book-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-row {
  display: flex;
  gap: 20px;
}

.form-row .form-group {
  flex: 1;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  color: #ecf0f1;
  font-weight: 500;
  margin-bottom: 8px;
  font-size: 1.1em;
}

.form-input,
.form-select,
.form-textarea {
  background-color: #3a3a3a;
  border: 2px solid #3a3a3a;
  border-radius: 8px;
  color: #ecf0f1;
  font-size: 1em;
  padding: 12px 16px;
  transition: all 0.3s ease;
}

.form-input:focus,
.form-select:focus,
.form-textarea:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.1);
}

.form-input::placeholder,
.form-textarea::placeholder {
  color: #95a5a6;
}

.form-select {
  cursor: pointer;
}

.form-select option {
  background-color: #3a3a3a;
  color: #ecf0f1;
}

.form-textarea {
  resize: vertical;
  min-height: 100px;
}

.file-upload-container {
  position: relative;
}

.file-input {
  position: absolute;
  opacity: 0;
  width: 100%;
  height: 100%;
  cursor: pointer;
}

.file-label {
  display: flex;
  align-items: center;
  gap: 12px;
  background-color: #3a3a3a;
  border: 2px dashed #7f8c8d;
  border-radius: 8px;
  padding: 16px;
  color: #ecf0f1;
  cursor: pointer;
  transition: all 0.3s ease;
}

.file-label:hover {
  border-color: #3498db;
  background-color: #4a4a4a;
}

.file-icon {
  font-size: 1.2em;
}

.image-preview {
  margin-top: 15px;
  text-align: center;
}

.preview-image {
  max-width: 200px;
  max-height: 300px;
  border-radius: 8px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.3);
}

.error-message {
  color: #e74c3c;
  font-size: 0.9em;
  margin-top: 5px;
  font-weight: 500;
}

.form-actions {
  display: flex;
  gap: 15px;
  justify-content: flex-end;
  margin-top: 30px;
}

.btn {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-size: 1em;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 120px;
}

.btn-primary {
  background-color: #3498db;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #2980b9;
  transform: translateY(-2px);
}

.btn-primary:disabled {
  background-color: #7f8c8d;
  cursor: not-allowed;
}

.btn-secondary {
  background-color: #95a5a6;
  color: white;
}

.btn-secondary:hover {
  background-color: #7f8c8d;
}

@media (max-width: 768px) {
  .form-container {
    padding: 20px;
  }
  
  .form-row {
    flex-direction: column;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .form-title {
    font-size: 1.8em;
  }
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

.modal-content {
  background-color: #2a2a2a;
  border-radius: 16px;
  padding: 0;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
  animation: modalSlideIn 0.3s ease-out;
}

.modal-header {
  text-align: center;
  padding: 30px 30px 20px;
  border-bottom: 1px solid #3a3a3a;
}

.success-icon {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: #27ae60;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2em;
  font-weight: bold;
  margin: 0 auto 16px;
  animation: successPulse 0.6s ease-out;
}

.modal-title {
  color: #ecf0f1;
  font-size: 1.5em;
  font-weight: 600;
  margin: 0;
}

.modal-body {
  padding: 20px 30px;
}

.success-message {
  color: #ecf0f1;
  font-size: 1.1em;
  text-align: center;
  line-height: 1.5;
  margin: 0;
}

.modal-footer {
  padding: 20px 30px 30px;
  text-align: center;
}

.modal-btn {
  min-width: 120px;
  padding: 12px 30px;
  font-size: 1.1em;
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: translateY(-20px) scale(0.9);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes successPulse {
  0% {
    transform: scale(0.8);
    opacity: 0;
  }
  50% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

@media (max-width: 480px) {
  .modal-content {
    margin: 20px;
  }
  
  .modal-header,
  .modal-body,
  .modal-footer {
    padding-left: 20px;
    padding-right: 20px;
  }
  
  .success-icon {
    width: 50px;
    height: 50px;
    font-size: 1.5em;
  }
  
  .modal-title {
    font-size: 1.3em;
  }
  
  .success-message {
    font-size: 1em;
  }
}
</style>