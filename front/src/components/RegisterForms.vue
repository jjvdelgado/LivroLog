<template>
  <div class="register-container">
    <div class="register-card">
      <div class="register-header">
        <h1 class="register-title">Criar nova conta</h1>
        <p class="register-subtitle">Junte-se à nossa biblioteca</p>
      </div>

      <form class="register-form" @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="name" class="form-label">Nome completo</label>
          <input type="text" id="name" v-model="name" class="form-input" :class="{ 'error': errors.name }"
            placeholder="Digite seu nome completo" required />
          <span v-if="errors.name" class="error-message">{{ errors.name }}</span>
        </div>

        <div class="form-group">
          <label for="email" class="form-label">E-mail</label>
          <input type="email" id="email" v-model="email" class="form-input" :class="{ 'error': errors.email }"
            placeholder="Digite seu e-mail" required />
          <span v-if="errors.email" class="error-message">{{ errors.email }}</span>
        </div>

        <div class="form-group">
          <label for="password" class="form-label">Senha</label>
          <input type="password" id="password" v-model="password" class="form-input"
            :class="{ 'error': errors.password }" placeholder="Digite sua senha" required />
          <span v-if="errors.password" class="error-message">{{ errors.password }}</span>
          <div class="password-requirements">
            <small class="requirement-text">
              A senha deve ter pelo menos 8 caracteres, incluindo:
            </small>
            <ul class="requirements-list">
              <li :class="{ 'valid': passwordValidation.length }">Mínimo 8 caracteres</li>
              <li :class="{ 'valid': passwordValidation.uppercase }">Pelo menos 1 letra maiúscula</li>
              <li :class="{ 'valid': passwordValidation.lowercase }">Pelo menos 1 letra minúscula</li>
              <li :class="{ 'valid': passwordValidation.number }">Pelo menos 1 número</li>
              <li :class="{ 'valid': passwordValidation.special }">Pelo menos 1 caractere especial</li>
            </ul>
          </div>
        </div>

        <div class="form-group">
          <label for="confirmPassword" class="form-label">Confirmar senha</label>
          <input type="password" id="confirmPassword" v-model="confirmPassword" class="form-input"
            :class="{ 'error': errors.confirmPassword }" placeholder="Confirme sua senha" required />
          <span v-if="errors.confirmPassword" class="error-message">{{ errors.confirmPassword }}</span>
        </div>

        <button type="submit" class="register-button" :disabled="!isFormValid">
          Criar conta
        </button>
      </form>

      <div class="register-actions">
        <div class="login-section">
          <span class="login-text">Já tem uma conta?</span>
          <button class="action-button login-button" @click="handleLogin">
            Entrar
          </button>
        </div>
      </div>
    </div>

    <!-- Pop-up de Sucesso -->
    <div v-if="showSuccessModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="success-icon">
          <svg width="60" height="60" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="12" cy="12" r="10" stroke="#4CAF50" stroke-width="2" />
            <path d="m9 12 2 2 4-4" stroke="#4CAF50" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </div>
        <h2 class="modal-title">Conta criada com sucesso!</h2>
        <p class="modal-message">Bem-vindo à nossa biblioteca! Sua conta foi criada e você já pode começar a explorar
          nossos recursos.</p>
        <button class="modal-button" @click="navigateToHomePage">
          Entrar na aplicação
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { useRequest } from '@/composables/useRequest'

const { request } = useRequest()

export default {
  name: 'RegisterForms',
  data() {
    return {
      name: '',
      email: '',
      password: '',
      confirmPassword: '',
      showSuccessModal: false,
      errors: {
        name: '',
        email: '',
        password: '',
        confirmPassword: ''
      }
    }
  },
  computed: {
    passwordValidation() {
      return {
        length: this.password.length >= 8,
        uppercase: /[A-Z]/.test(this.password),
        lowercase: /[a-z]/.test(this.password),
        number: /\d/.test(this.password),
        special: /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(this.password)
      }
    },
    isPasswordStrong() {
      return Object.values(this.passwordValidation).every(valid => valid)
    },
    isFormValid() {
      return this.name.trim().length >= 2 &&
        this.isValidEmail(this.email) &&
        this.isPasswordStrong &&
        this.password === this.confirmPassword &&
        this.password.length > 0
    }
  },
  watch: {
    name() {
      this.validateName()
    },
    email() {
      this.validateEmail()
    },
    password() {
      this.validatePassword()
      if (this.confirmPassword) {
        this.validateConfirmPassword()
      }
    },
    confirmPassword() {
      this.validateConfirmPassword()
    }
  },
  methods: {
    validateName() {
      this.errors.name = ''
      if (this.name.trim().length === 0) {
        this.errors.name = 'Nome é obrigatório'
      } else if (this.name.trim().length < 2) {
        this.errors.name = 'Nome deve ter pelo menos 2 caracteres'
      } else if (!/^[a-zA-ZÀ-ÿ\s]+$/.test(this.name.trim())) {
        this.errors.name = 'Nome deve conter apenas letras e espaços'
      }
    },
    validateEmail() {
      this.errors.email = ''
      if (this.email.trim().length === 0) {
        this.errors.email = 'E-mail é obrigatório'
      } else if (!this.isValidEmail(this.email)) {
        this.errors.email = 'E-mail deve ter um formato válido'
      }
    },
    validatePassword() {
      this.errors.password = ''
      if (this.password.length === 0) {
        this.errors.password = 'Senha é obrigatória'
      } else if (!this.isPasswordStrong) {
        this.errors.password = 'Senha não atende aos requisitos de segurança'
      }
    },
    validateConfirmPassword() {
      this.errors.confirmPassword = ''
      if (this.confirmPassword.length === 0) {
        this.errors.confirmPassword = 'Confirmação de senha é obrigatória'
      } else if (this.password !== this.confirmPassword) {
        this.errors.confirmPassword = 'As senhas não coincidem'
      }
    },
    isValidEmail(email) {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      return emailRegex.test(email)
    },
    async handleRegister() {
      // Valida todos os campos antes de prosseguir
      this.validateName()
      this.validateEmail()
      this.validatePassword()
      this.validateConfirmPassword()

      // Verifica se há erros
      const hasErrors = Object.values(this.errors).some(error => error !== '')

      if (hasErrors || !this.isFormValid) {
        return
      }

      // Simula o cadastro (aqui você faria a chamada para o backend)
      console.log('Register attempt:', {
        name: this.name.trim(),
        email: this.email.trim(),
        password: this.password
      })

      try {
        await request('http://localhost:8080/usuarios', 'POST', {
          nome: this.name,
          email: this.email,
          senha: this.password,
          descricao: ""
        }).then(() => {
          this.handleLogin()
        })
      } catch (e) {
        console.error('Erro ao cadastrar:', e.message)
      }

      // Mostra o modal de sucesso
      this.showSuccessModal = true
    },
    handleLogin() {
      this.$router.push('/login')
    },
    closeModal() {
      this.showSuccessModal = false
    },
    navigateToHomePage() {
      this.$router.push('/HomePage')
    }
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 40px 20px;
  width: 100%;
}

.register-card {
  background-color: #2a2a2a;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  width: 100%;
  max-width: 500px;
  transition: all 0.3s ease;
}

.register-card:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
}

.register-header {
  text-align: center;
  margin-bottom: 32px;
}

.register-title {
  color: #ffffff;
  font-size: 28px;
  font-weight: 700;
  margin-bottom: 8px;
}

.register-subtitle {
  color: #888;
  font-size: 16px;
  font-weight: 400;
}

.register-form {
  margin-bottom: 32px;
}

.form-group {
  margin-bottom: 24px;
}

.form-label {
  display: block;
  color: #ffffff;
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 8px;
}

.form-input {
  width: 100%;
  padding: 16px 20px;
  background-color: #1a1a1a;
  border: 2px solid #3a3a3a;
  border-radius: 12px;
  color: #ffffff;
  font-size: 16px;
  transition: all 0.3s ease;
  box-sizing: border-box;
}

.form-input:focus {
  outline: none;
  border-color: #2c3e50;
  box-shadow: 0 0 0 3px rgba(44, 62, 80, 0.1);
}

.form-input.error {
  border-color: #e74c3c;
  box-shadow: 0 0 0 3px rgba(231, 76, 60, 0.1);
}

.form-input::placeholder {
  color: #666;
  font-style: italic;
}

.error-message {
  color: #e74c3c;
  font-size: 12px;
  margin-top: 4px;
  display: block;
}

.password-requirements {
  margin-top: 8px;
}

.requirement-text {
  color: #888;
  font-size: 12px;
  margin-bottom: 4px;
}

.requirements-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.requirements-list li {
  color: #666;
  font-size: 11px;
  margin-bottom: 2px;
  position: relative;
  padding-left: 16px;
}

.requirements-list li:before {
  content: '✗';
  color: #e74c3c;
  position: absolute;
  left: 0;
  font-weight: bold;
}

.requirements-list li.valid:before {
  content: '✓';
  color: #27ae60;
}

.requirements-list li.valid {
  color: #27ae60;
}

.register-button {
  width: 100%;
  padding: 16px;
  background-color: #2c3e50;
  border: none;
  border-radius: 12px;
  color: #ffffff;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 8px;
}

.register-button:hover:not(:disabled) {
  background-color: #34495e;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(44, 62, 80, 0.3);
}

.register-button:active:not(:disabled) {
  transform: translateY(0);
}

.register-button:disabled {
  background-color: #666;
  cursor: not-allowed;
  opacity: 0.6;
}

.register-actions {
  text-align: center;
}

.action-button {
  background: none;
  border: none;
  color: #2c3e50;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
}

.action-button:hover {
  color: #34495e;
  text-decoration: underline;
}

.login-section {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.login-text {
  color: #888;
  font-size: 14px;
}

.login-button {
  font-weight: 600;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: #2a2a2a;
  border-radius: 16px;
  padding: 40px;
  text-align: center;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

.success-icon {
  margin-bottom: 20px;
}

.modal-title {
  color: #ffffff;
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 16px;
}

.modal-message {
  color: #888;
  font-size: 16px;
  margin-bottom: 32px;
  line-height: 1.5;
}

.modal-button {
  background-color: #27ae60;
  color: #ffffff;
  border: none;
  border-radius: 12px;
  padding: 16px 32px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.modal-button:hover {
  background-color: #2ecc71;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(39, 174, 96, 0.3);
}

/* Responsividade */
@media (max-width: 768px) {
  .register-container {
    padding: 30px 15px;
  }

  .register-card {
    padding: 32px 24px;
    max-width: 100%;
  }

  .register-title {
    font-size: 24px;
  }

  .register-subtitle {
    font-size: 14px;
  }

  .form-input {
    padding: 14px 16px;
  }

  .register-button {
    padding: 14px;
  }

  .modal-content {
    padding: 32px 24px;
  }

  .modal-title {
    font-size: 20px;
  }

  .modal-message {
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  .register-container {
    padding: 20px 10px;
  }

  .register-card {
    padding: 24px 20px;
  }

  .register-title {
    font-size: 22px;
  }

  .form-input {
    padding: 12px 14px;
    font-size: 14px;
  }

  .register-button {
    padding: 12px;
    font-size: 14px;
  }

  .login-section {
    flex-direction: column;
    gap: 12px;
  }

  .modal-content {
    padding: 24px 20px;
  }

  .modal-title {
    font-size: 18px;
  }

  .modal-button {
    padding: 12px 24px;
    font-size: 14px;
  }
}
</style>