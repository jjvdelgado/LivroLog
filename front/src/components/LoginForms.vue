<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h1 class="login-title">Entrar na sua conta</h1>
        <p class="login-subtitle">Acesse sua biblioteca pessoal</p>
      </div>

      <form class="login-form" @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="email" class="form-label">E-mail</label>
          <input type="email" id="email" v-model="email" class="form-input" placeholder="Digite seu e-mail" required />
        </div>

        <div class="form-group">
          <label for="password" class="form-label">Senha</label>
          <input type="password" id="password" v-model="password" class="form-input" placeholder="Digite sua senha"
            required />
        </div>

        <button type="submit" class="login-button">
          Entrar
        </button>
      </form>

      <div class="login-actions">
        <button class="action-button forgot-password" @click="handleForgotPassword">
          Esqueci minha senha
        </button>

        <div class="register-section">
          <span class="register-text">Não tem uma conta?</span>
          <button class="action-button register-button" @click="handleRegister">
            Criar conta
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'
import { useRequest } from '@/composables/useRequest'

const { request } = useRequest()

export default {
  name: 'LoginForms',
  setup() {
    const router = useRouter()

    return {
      router
    }
  },
  data() {
    return {
      email: '',
      password: '',
      isLoading: false
    }
  },
  methods: {
    async handleLogin() {
      this.isLoading = true

      try {
        // Lógica de login será implementada aqui
        console.log('Login attempt:', { email: this.email, password: this.password })

        // Simulando uma requisição (remova depois)
        await new Promise(resolve => setTimeout(resolve, 1000))

        // Aqui você faria a chamada para sua API de login
        // const response = await loginAPI(this.email, this.password)

        request('http://localhost:8080/login', 'POST', {
          email: this.email,
          senha: this.password
        })
          .then(() => {
            this.router.push('/homepage')
          })
          .catch(err => {
            console.error('Erro na requisição:', err)
          })

        // Redirecionar após login bem-sucedido
        //  // ou para onde quiser após o login

        alert('Login realizado com sucesso!')

      } catch (error) {
        console.error('Erro no login:', error)
        alert('Erro ao fazer login. Tente novamente.')
      } finally {
        this.isLoading = false
      }
    },

    handleForgotPassword() {
      // Lógica para esqueci minha senha será implementada aqui
      console.log('Forgot password clicked')
      // Você pode navegar para uma página de recuperação de senha
      // this.router.push('/forgot-password')
    },

    handleRegister() {
      // Navegação para cadastro
      console.log('Register clicked')
      this.router.push('/register')
    },

    goToHome() {
      // Método para voltar à página inicial
      this.router.push('/')
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
  width: 100%;
}

.login-card {
  background-color: #2a2a2a;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  width: 100%;
  max-width: 420px;
  transition: all 0.3s ease;
}

.login-card:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.5);
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.login-title {
  color: #ffffff;
  font-size: 28px;
  font-weight: 700;
  margin-bottom: 8px;
}

.login-subtitle {
  color: #888;
  font-size: 16px;
  font-weight: 400;
}

.login-form {
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

.form-input::placeholder {
  color: #666;
  font-style: italic;
}

.login-button {
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

.login-button:hover {
  background-color: #34495e;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(44, 62, 80, 0.3);
}

.login-button:active {
  transform: translateY(0);
}

.login-actions {
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

.forgot-password {
  margin-bottom: 24px;
  display: block;
}

.register-section {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.register-text {
  color: #888;
  font-size: 14px;
}

.register-button {
  font-weight: 600;
}

/* Responsividade */
@media (max-width: 768px) {
  .login-card {
    padding: 32px 24px;
    margin: 20px;
  }

  .login-title {
    font-size: 24px;
  }

  .login-subtitle {
    font-size: 14px;
  }

  .form-input {
    padding: 14px 16px;
  }

  .login-button {
    padding: 14px;
  }
}

@media (max-width: 480px) {
  .login-card {
    padding: 24px 20px;
    margin: 10px;
  }

  .login-title {
    font-size: 22px;
  }

  .form-input {
    padding: 12px 14px;
    font-size: 14px;
  }

  .login-button {
    padding: 12px;
    font-size: 14px;
  }

  .register-section {
    flex-direction: column;
    gap: 12px;
  }
}
</style>